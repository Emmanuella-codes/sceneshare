package service

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
	"github.com/Emmanuella-codes/sceneshare/api/store"
	"github.com/Emmanuella-codes/sceneshare/api/utils"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

const alphabet = "abcdefghjkmnpqrstuvwxyzABCDEFGHJKMNPQRSTUVWXYZ23456789"
const codeLength = 7

type LinkService struct {
	store   *store.Store
	baseURL string
}

func NewLinkService(store *store.Store, baseURL string) *LinkService {
	return &LinkService{store: store, baseURL: baseURL}
}

func (s *LinkService) CreateLink(ctx context.Context, input *dtos.CreateLinkInput) (*dtos.LinkResponse, error) {
	if !input.Platform.IsValid() {
		return nil, &utils.ValidationError{Field: "platform", Message: "must be one of: youtube"} // extend when new platforms are added
	}

	if err := utils.ValidateContentID(input.Platform, input.ContentID); err != nil {
		return nil, err
	}

	code, err := gonanoid.Generate(alphabet, codeLength)
	if err != nil {
		return nil, fmt.Errorf("generating code: %w", err)
	}

	ownerToken, err := gonanoid.Generate(alphabet, 24)
	if err != nil {
		return nil, fmt.Errorf("generating owner token: %w", err)
	}

	params := dtos.CreateLinkParams{
		ShortCode:  code,
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

	created, err := s.store.CreateLink(ctx, params)
	if err != nil {
		return nil, err
	}
	resp := toResponse(created, s.baseURL)
	resp.OwnerToken = &created.OwnerToken
	return resp, nil
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
	return s.store.DeleteLink(ctx, code, token)
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

func BuildDeepLink(link *models.Link) string {
	switch link.Platform {
	case models.PlatformYoutube:
		return fmt.Sprintf("https://www.youtube.com/watch?v=%s", link.ContentID)
	default:
		return ""
	}
}

func (s *LinkService) getLink(ctx context.Context, code string) (*models.Link, error) {
	return s.store.GetLinkByCode(ctx, code)
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
