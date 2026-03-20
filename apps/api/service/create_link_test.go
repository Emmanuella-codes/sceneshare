package service

import (
	"context"
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
	"github.com/Emmanuella-codes/sceneshare/api/store"
	"github.com/Emmanuella-codes/sceneshare/api/utils"
)

func TestCreateLink(t *testing.T) {
	validInput := func() *dtos.CreateLinkInput {
		return &dtos.CreateLinkInput{
			Platform:   models.PlatformYoutube,
			ContentID:  "dQw4w9WgXcQ",
			TimestampS: 83,
		}
	}

	tests := []struct {
		name        string
		input       *dtos.CreateLinkInput
		store       *mockStore
		expectedErr error
		check       func(*testing.T, *dtos.LinkResponse)
	}{
		{
			name:  "success",
			input: validInput(),
			store: &mockStore{
				createLink: func(_ context.Context, params dtos.CreateLinkParams) (*models.Link, error) {
					if params.ContentID != "dQw4w9WgXcQ" {
						t.Fatalf("expected content id to be forwarded")
					}
					if params.TimestampS != 83 {
						t.Fatalf("expected timestamp to be forwarded")
					}
					if params.ShortCode == "" {
						t.Fatal("expected generated short code")
					}
					if len(params.OwnerToken) != 24 {
						t.Fatalf("expected owner token length 24, got %d", len(params.OwnerToken))
					}
					return &models.Link{
						ID:         "link-1",
						ShortCode:  params.ShortCode,
						Platform:   params.Platform,
						ContentID:  params.ContentID,
						TimestampS: params.TimestampS,
						OwnerToken: params.OwnerToken,
						CreatedAt:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
					}, nil
				},
			},
			check: func(t *testing.T, resp *dtos.LinkResponse) {
				t.Helper()
				if resp.ShortCode == "" {
					t.Fatal("expected short code in response")
				}
				if !strings.HasPrefix(resp.ShortURL, "http://localhost:3001/r/") {
					t.Fatalf("expected short_url to use /r/, got %q", resp.ShortURL)
				}
				if resp.OwnerToken == nil || len(*resp.OwnerToken) != 24 {
					t.Fatal("expected owner token in response")
				}
				if resp.TimestampFmt != "1:23" {
					t.Fatalf("expected timestamp_fmt %q, got %q", "1:23", resp.TimestampFmt)
				}
			},
		},
		{
			name: "validation error returns early",
			input: &dtos.CreateLinkInput{
				Platform:   models.PlatformYoutube,
				ContentID:  "bad-id",
				TimestampS: 83,
			},
			store: &mockStore{
				createLink: func(context.Context, dtos.CreateLinkParams) (*models.Link, error) {
					t.Fatal("store should not be called when input validation fails")
					return nil, nil
				},
			},
			expectedErr: &utils.ValidationError{},
		},
		{
			name:  "code conflict retries and succeeds",
			input: validInput(),
			store: func() *mockStore {
				attempts := 0
				return &mockStore{
					createLink: func(_ context.Context, params dtos.CreateLinkParams) (*models.Link, error) {
						attempts++
						if attempts == 1 {
							return nil, store.ErrCodeConflict
						}
						return &models.Link{
							ID:         "link-1",
							ShortCode:  params.ShortCode,
							Platform:   params.Platform,
							ContentID:  params.ContentID,
							TimestampS: params.TimestampS,
							OwnerToken: params.OwnerToken,
							CreatedAt:  time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC),
						}, nil
					},
				}
			}(),
			check: func(t *testing.T, resp *dtos.LinkResponse) {
				t.Helper()
				if resp.ShortCode == "" {
					t.Fatal("expected short code after retry")
				}
			},
		},
		{
			name:  "conflict after max attempts",
			input: validInput(),
			store: &mockStore{
				createLink: func(_ context.Context, _ dtos.CreateLinkParams) (*models.Link, error) {
					return nil, store.ErrCodeConflict
				},
			},
			expectedErr: store.ErrCodeConflict,
		},
		{
			name:  "store error is mapped",
			input: validInput(),
			store: &mockStore{
				createLink: func(_ context.Context, _ dtos.CreateLinkParams) (*models.Link, error) {
					return nil, store.ErrExpired
				},
			},
			expectedErr: ErrExpired,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := NewLinkService(tc.store, "http://localhost:3001")
			resp, err := svc.CreateLink(context.Background(), tc.input)

			if tc.expectedErr != nil {
				if err == nil {
					t.Fatalf("expected error %v, got nil", tc.expectedErr)
				}
				var validationErr *utils.ValidationError
				if _, ok := tc.expectedErr.(*utils.ValidationError); ok {
					if !errors.As(err, &validationErr) {
						t.Fatalf("expected validation error, got %v", err)
					}
					return
				}
				if !errors.Is(err, tc.expectedErr) {
					t.Fatalf("expected error %v, got %v", tc.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if tc.check != nil {
				tc.check(t, resp)
			}
		})
	}
}
