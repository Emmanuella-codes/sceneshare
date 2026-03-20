package store

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
)

func TestGetLinkByCode_NotFound(t *testing.T) {
	s := newTestStore(t)

	_, err := s.GetLinkByCode(context.Background(), "missing1")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("expected ErrNotFound, got %v", err)
	}
}

func TestGetLinkByCode_Expired(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()
	expiredAt := time.Now().Add(-1 * time.Minute)

	_, err := s.CreateLink(ctx, dtos.CreateLinkParams{
		ShortCode:  "expired1",
		Platform:   models.PlatformYoutube,
		ContentID:  "dQw4w9WgXcQ",
		TimestampS: 83,
		OwnerToken: "owner-token",
		ExpiresAt:  &expiredAt,
	})
	if err != nil {
		t.Fatalf("create expired link: %v", err)
	}

	_, err = s.GetLinkByCode(ctx, "expired1")
	if !errors.Is(err, ErrExpired) {
		t.Fatalf("expected ErrExpired, got %v", err)
	}
}
