package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Emmanuella-codes/sceneshare/api/models"
	"github.com/Emmanuella-codes/sceneshare/api/store"
)

func TestGetLink(t *testing.T) {
	tests := []struct {
		name        string
		code        string
		store       *mockStore
		expectedErr error
	}{
		{
			name: "success",
			code: "abc1234",
			store: &mockStore{
				getLinkByCode: func(_ context.Context, code string) (*models.Link, error) {
					if code != "abc1234" {
						t.Fatalf("expected code %q, got %q", "abc1234", code)
					}
					return &models.Link{
						ID:         "link-1",
						ShortCode:  code,
						Platform:   models.PlatformYoutube,
						ContentID:  "dQw4w9WgXcQ",
						TimestampS: 83,
						CreatedAt:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
					}, nil
				},
			},
		},
		{
			name: "not found",
			code: "missing1",
			store: &mockStore{
				getLinkByCode: func(context.Context, string) (*models.Link, error) {
					return nil, store.ErrNotFound
				},
			},
			expectedErr: ErrNotFound,
		},
		{
			name: "expired",
			code: "expired1",
			store: &mockStore{
				getLinkByCode: func(context.Context, string) (*models.Link, error) {
					return nil, store.ErrExpired
				},
			},
			expectedErr: ErrExpired,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewLinkService(tc.store, "http://localhost:3001")
			resp, err := svc.GetLink(context.Background(), tc.code)

			if tc.expectedErr != nil {
				if !errors.Is(err, tc.expectedErr) {
					t.Fatalf("expected error %v, got %v", tc.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if resp.ShortCode != tc.code {
				t.Fatalf("expected short_code %q, got %q", tc.code, resp.ShortCode)
			}
			if resp.ShortURL != "http://localhost:3001/r/abc1234" {
				t.Fatalf("expected short_url %q, got %q", "http://localhost:3001/r/abc1234", resp.ShortURL)
			}
			if resp.TimestampFmt != "1:23" {
				t.Fatalf("expected timestamp_fmt %q, got %q", "1:23", resp.TimestampFmt)
			}
		})
	}
}
