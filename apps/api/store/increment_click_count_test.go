package store

import (
	"context"
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
)

func TestIncrementClickCount(t *testing.T) {
	s := newTestStore(t)
	ctx := context.Background()

	created, err := s.CreateLink(ctx, dtos.CreateLinkParams{
		ShortCode:  "click12",
		Platform:   models.PlatformYoutube,
		ContentID:  "dQw4w9WgXcQ",
		TimestampS: 83,
		OwnerToken: "owner-token",
	})
	if err != nil {
		t.Fatalf("create link: %v", err)
	}

	err = s.IncrementClickCount(ctx, models.ClickEvent{
		LinkID:    created.ID,
		UserAgent: "test-agent",
		Referrer:  "https://example.com",
	})
	if err != nil {
		t.Fatalf("increment click count: %v", err)
	}

	got, err := s.GetLinkByCode(ctx, "click12")
	if err != nil {
		t.Fatalf("get link after click: %v", err)
	}
	if got.ClickCount != 1 {
		t.Fatalf("expected click count %d, got %d", 1, got.ClickCount)
	}

	var eventCount int
	if err := s.db.QueryRow(ctx, `SELECT COUNT(*) FROM click_events WHERE link_id = $1`, created.ID).Scan(&eventCount); err != nil {
		t.Fatalf("count click events: %v", err)
	}
	if eventCount != 1 {
		t.Fatalf("expected click event count %d, got %d", 1, eventCount)
	}
}
