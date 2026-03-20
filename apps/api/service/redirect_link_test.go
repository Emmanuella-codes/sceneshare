package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/Emmanuella-codes/sceneshare/api/models"
	"github.com/Emmanuella-codes/sceneshare/api/store"
)

func TestGetLinkForRedirect(t *testing.T) {
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
		{
			name: "internal error",
			code: "broken12",
			store: &mockStore{
				getLinkByCode: func(context.Context, string) (*models.Link, error) {
					return nil, errors.New("boom")
				},
			},
			expectedErr: errors.New("boom"),
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewLinkService(tc.store, "http://localhost:3001")
			link, err := svc.GetLinkForRedirect(context.Background(), tc.code)

			if tc.expectedErr != nil {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if err.Error() != tc.expectedErr.Error() {
					t.Fatalf("expected error %q, got %q", tc.expectedErr.Error(), err.Error())
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if link.ShortCode != tc.code {
				t.Fatalf("expected short_code %q, got %q", tc.code, link.ShortCode)
			}
		})
	}
}
