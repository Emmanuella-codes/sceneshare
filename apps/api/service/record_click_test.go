package service

import (
	"context"
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/models"
)

func TestRecordClick(t *testing.T) {
	called := make(chan models.ClickEvent, 1)
	svc := NewLinkService(&mockStore{
		incrementClickCount: func(_ context.Context, event models.ClickEvent) error {
			called <- event
			return nil
		},
	}, "http://localhost:3001")

	svc.RecordClick("link-1", "test-agent", "https://example.com")

	event := <-called
	if event.LinkID != "link-1" {
		t.Fatalf("expected link id %q, got %q", "link-1", event.LinkID)
	}
	if event.UserAgent != "test-agent" {
		t.Fatalf("expected user agent %q, got %q", "test-agent", event.UserAgent)
	}
	if event.Referrer != "https://example.com" {
		t.Fatalf("expected referrer %q, got %q", "https://example.com", event.Referrer)
	}
}
