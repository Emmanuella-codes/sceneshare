package store

import (
	"context"
	"errors"
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
)

func TestCreateAndGetLink(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()
	title := "Test title"
	thumbnail := "https://example.com/thumb.jpg"

	created, err := s.CreateLink(ctx, dtos.CreateLinkParams{
		ShortCode:  "abc1234",
		Platform:   models.PlatformYoutube,
		ContentID:  "dQw4w9WgXcQ",
		TimestampS: 83,
		Title:      &title,
		Thumbnail:  &thumbnail,
		OwnerToken: "owner-token",
	})
	if err != nil {
		t.Fatalf("create link: %v", err)
	}

	got, err := s.GetLinkByCode(ctx, "abc1234")
	if err != nil {
		t.Fatalf("get link by code: %v", err)
	}

	if got.ID != created.ID {
		t.Fatalf("expected id %q, got %q", created.ID, got.ID)
	}
	if got.ShortCode != "abc1234" {
		t.Fatalf("expected short code %q, got %q", "abc1234", got.ShortCode)
	}
	if got.Platform != models.PlatformYoutube {
		t.Fatalf("expected platform %q, got %q", models.PlatformYoutube, got.Platform)
	}
	if got.ContentID != "dQw4w9WgXcQ" {
		t.Fatalf("expected content id %q, got %q", "dQw4w9WgXcQ", got.ContentID)
	}
	if got.TimestampS != 83 {
		t.Fatalf("expected timestamp %d, got %d", 83, got.TimestampS)
	}
	if got.Title == nil || *got.Title != title {
		t.Fatalf("expected title %q, got %#v", title, got.Title)
	}
	if got.Thumbnail == nil || *got.Thumbnail != thumbnail {
		t.Fatalf("expected thumbnail %q, got %#v", thumbnail, got.Thumbnail)
	}
	if got.OwnerToken != "owner-token" {
		t.Fatalf("expected owner token %q, got %q", "owner-token", got.OwnerToken)
	}
}

func TestCreateLink_CodeConflict(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()

	params := dtos.CreateLinkParams{
		ShortCode:  "dup1234",
		Platform:   models.PlatformYoutube,
		ContentID:  "dQw4w9WgXcQ",
		TimestampS: 83,
		OwnerToken: "owner-token",
	}

	if _, err := s.CreateLink(ctx, params); err != nil {
		t.Fatalf("create original link: %v", err)
	}

	_, err := s.CreateLink(ctx, params)
	if !errors.Is(err, ErrCodeConflict) {
		t.Fatalf("expected ErrCodeConflict, got %v", err)
	}
}
