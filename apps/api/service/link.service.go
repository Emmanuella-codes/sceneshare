package service

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/url"
	"strconv"
	"time"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
	"github.com/Emmanuella-codes/sceneshare/api/store"
	"github.com/Emmanuella-codes/sceneshare/api/utils"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const alphabet = "abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ23456789"
const codeLength = 7
const maxCreateAttempts = 5

var ErrNotFound = errors.New("link not found")
var ErrExpired = errors.New("link expired")
var ErrForbidden = errors.New("forbidden")
var ErrUnsupportedPlatform = errors.New("unsupported platform")

type LinkService struct {
	store   *store.Store
	baseURL string
}

func NewLinkService(store *store.Store, baseURL string) *LinkService {
	return &LinkService{store: store, baseURL: baseURL}
}

func (s *LinkService) CreateLink(ctx context.Context, input *dtos.CreateLinkInput) (*dtos.LinkResponse, error) {
	if err := utils.ValidateCreateLinkInput(input); err != nil {
		return nil, err
	}

	ownerToken, err := gonanoid.Generate(alphabet, 24)
	if err != nil {
		return nil, fmt.Errorf("generating owner token: %w", err)
	}

	params := dtos.CreateLinkParams{
		Platform:   input.Platform,
		ContentID:  input.ContentID,
		TimestampS: input.TimestampS,
		Title:      input.Title,
		Thumbnail:  input.Thumbnail,
		OwnerToken: ownerToken,
	}

	if input.ExpiresIn != nil {
		t := time.Now().Add(time.Duration(*input.ExpiresIn) * time.Second)
		params.ExpiresAt = &t
	}

	// Retry on rare code collisions instead of surfacing them as 500s.
	var created *models.Link
	for attempt := 0; attempt < maxCreateAttempts; attempt++ {
		code, err := gonanoid.Generate(alphabet, codeLength)
		if err != nil {
			return nil, fmt.Errorf("generating code: %w", err)
		}

		params.ShortCode = code
		created, err = s.store.CreateLink(ctx, params)
		if err == nil {
			resp := toResponse(created, s.baseURL)
			resp.OwnerToken = &created.OwnerToken
			return resp, nil
		}
		if !errors.Is(err, store.ErrCodeConflict) {
			return nil, mapStoreError(err)
		}
	}

	return nil, fmt.Errorf("creating short code: %w", store.ErrCodeConflict)
}

func (s *LinkService) GetLink(ctx context.Context, code string) (*dtos.LinkResponse, error) {
	link, err := s.getLink(ctx, code)
	if err != nil {
		return nil, err
	}
	return toResponse(link, s.baseURL), nil
}

func (s *LinkService) GetLinkForRedirect(ctx context.Context, code string) (*models.Link, error) {
	return s.getLink(ctx, code)
}

func (s *LinkService) DeleteLink(ctx context.Context, code, token string) error {
	return mapStoreError(s.store.DeleteLink(ctx, code, token))
}

func (s *LinkService) GetStats(ctx context.Context, code string) (*dtos.StatsResponse, error) {
	link, err := s.getLink(ctx, code)
	if err != nil {
		return nil, err
	}
	return &dtos.StatsResponse{
		ShortCode:  link.ShortCode,
		ClickCount: link.ClickCount,
		CreatedAt:  link.CreatedAt.Format(time.RFC3339),
	}, nil
}

func (s *LinkService) RecordClick(linkID, userAgent, referrer string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.store.IncrementClickCount(ctx, models.ClickEvent{
		LinkID:    linkID,
		UserAgent: userAgent,
		Referrer:  referrer,
	}); err != nil {
		slog.Error("failed to record click", "link_id", linkID, "error", err)
	}
}

// BuildDeepLink preserves the saved playback offset in the platform URL.
func BuildDeepLink(link *models.Link) (string, error) {
	switch link.Platform {
	case models.PlatformYoutube:
		target := &url.URL{
			Scheme: "https",
			Host:   "www.youtube.com",
			Path:   "/watch",
		}
		query := target.Query()
		query.Set("v", link.ContentID)
		if link.TimestampS > 0 {
			query.Set("t", strconv.Itoa(link.TimestampS)+"s")
		}
		target.RawQuery = query.Encode()
		return target.String(), nil
	default:
		return "", ErrUnsupportedPlatform
	}
}

func (s *LinkService) getLink(ctx context.Context, code string) (*models.Link, error) {
	link, err := s.store.GetLinkByCode(ctx, code)
	if err != nil {
		return nil, mapStoreError(err)
	}
	return link, nil
}

// translates persistence failures into service-level errors.
func mapStoreError(err error) error {
	switch {
	case errors.Is(err, store.ErrNotFound):
		return ErrNotFound
	case errors.Is(err, store.ErrExpired):
		return ErrExpired
	case errors.Is(err, store.ErrForbidden):
		return ErrForbidden
	default:
		return err
	}
}

func toResponse(l *models.Link, baseURL string) *dtos.LinkResponse {
	r := &dtos.LinkResponse{
		ID:           l.ID,
		ShortCode:    l.ShortCode,
		ShortURL:     fmt.Sprintf("%s/%s", baseURL, l.ShortCode),
		Platform:     string(l.Platform),
		ContentID:    l.ContentID,
		TimestampS:   l.TimestampS,
		TimestampFmt: utils.FormatTimestamp(l.TimestampS),
		Title:        l.Title,
		Thumbnail:    l.Thumbnail,
		ClickCount:   l.ClickCount,
		CreatedAt:    l.CreatedAt.Format(time.RFC3339),
	}

	if l.ExpiresAt != nil {
		t := l.ExpiresAt.Format(time.RFC3339)
		r.ExpiresAt = &t
	}
	return r
}
