package store

import (
	"context"
	"errors"
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
)

func TestDeleteLink(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := newTestStore(t)
		ctx := context.Background()

		if _, err := s.CreateLink(ctx, dtos.CreateLinkParams{
			ShortCode:  "del1234",
			Platform:   models.PlatformYoutube,
			ContentID:  "dQw4w9WgXcQ",
			TimestampS: 83,
			OwnerToken: "owner-token",
		}); err != nil {
			t.Fatalf("create link: %v", err)
		}

		if err := s.DeleteLink(ctx, "del1234", "owner-token"); err != nil {
			t.Fatalf("delete link: %v", err)
		}

		_, err := s.GetLinkByCode(ctx, "del1234")
		if !errors.Is(err, ErrNotFound) {
			t.Fatalf("expected ErrNotFound after delete, got %v", err)
		}
	})

	t.Run("forbidden", func(t *testing.T) {
		s := newTestStore(t)
		ctx := context.Background()

		if _, err := s.CreateLink(ctx, dtos.CreateLinkParams{
			ShortCode:  "forb123",
			Platform:   models.PlatformYoutube,
			ContentID:  "dQw4w9WgXcQ",
			TimestampS: 83,
			OwnerToken: "owner-token",
		}); err != nil {
			t.Fatalf("create link: %v", err)
		}

		err := s.DeleteLink(ctx, "forb123", "wrong-token")
		if !errors.Is(err, ErrForbidden) {
			t.Fatalf("expected ErrForbidden, got %v", err)
		}
	})

	t.Run("not found", func(t *testing.T) {
		s := newTestStore(t)

		err := s.DeleteLink(context.Background(), "missing1", "owner-token")
		if !errors.Is(err, ErrNotFound) {
			t.Fatalf("expected ErrNotFound, got %v", err)
		}
	})
}
