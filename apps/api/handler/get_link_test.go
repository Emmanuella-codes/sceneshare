package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/service"
)

func TestGetLink(t *testing.T) {
	tests := []struct {
		name           string
		code           string
		mockSetup      func(*testing.T, *mockLinkService)
		expectedStatus int
		expectedCode   string
	}{
		{
			name: "success",
			code: "abc1234",
			mockSetup: func(t *testing.T, m *mockLinkService) {
				t.Helper()
				m.getLink = func(_ context.Context, code string) (*dtos.LinkResponse, error) {
					if code != "abc1234" {
						t.Fatalf("expected code %q, got %q", "abc1234", code)
					}
					return &dtos.LinkResponse{
						ID:           "link-1",
						ShortCode:    code,
						ShortURL:     "http://localhost:3001/r/abc1234",
						Platform:     "youtube",
						ContentID:    "dQw4w9WgXcQ",
						TimestampS:   83,
						TimestampFmt: "1:23",
						ClickCount:   2,
						CreatedAt:    "2025-01-01T00:00:00Z",
					}, nil
				}
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "not found",
			code: "missing1",
			mockSetup: func(_ *testing.T, m *mockLinkService) {
				m.getLink = func(_ context.Context, _ string) (*dtos.LinkResponse, error) {
					return nil, service.ErrNotFound
				}
			},
			expectedStatus: http.StatusNotFound,
			expectedCode:   "LINK_NOT_FOUND",
		},
		{
			name: "expired",
			code: "expired1",
			mockSetup: func(_ *testing.T, m *mockLinkService) {
				m.getLink = func(_ context.Context, _ string) (*dtos.LinkResponse, error) {
					return nil, service.ErrExpired
				}
			},
			expectedStatus: http.StatusGone,
			expectedCode:   "LINK_EXPIRED",
		},
		{
			name: "internal error",
			code: "broken12",
			mockSetup: func(_ *testing.T, m *mockLinkService) {
				m.getLink = func(_ context.Context, _ string) (*dtos.LinkResponse, error) {
					return nil, errors.New("boom")
				}
			},
			expectedStatus: http.StatusInternalServerError,
			expectedCode:   "INTERNAL_ERROR",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			svc := &mockLinkService{}
			tc.mockSetup(t, svc)
			h := New(svc)

			r := withCode(httptest.NewRequest(http.MethodGet, "/api/v1/links/"+tc.code, nil), tc.code)
			w := httptest.NewRecorder()

			h.GetLink(w, r)

			res := w.Result()
			if res.StatusCode != tc.expectedStatus {
				t.Fatalf("expected status %d, got %d", tc.expectedStatus, res.StatusCode)
			}

			if tc.expectedCode != "" {
				body := decodeJSON[errorResponse](t, res)
				if body.Code != tc.expectedCode {
					t.Fatalf("expected code %q, got %q", tc.expectedCode, body.Code)
				}
				return
			}

			body := decodeJSON[dtos.LinkResponse](t, res)
			if body.ShortCode != tc.code {
				t.Fatalf("expected short_code %q, got %q", tc.code, body.ShortCode)
			}
			if body.ShortURL != "http://localhost:3001/r/abc1234" {
				t.Fatalf("expected short_url %q, got %q", "http://localhost:3001/r/abc1234", body.ShortURL)
			}
			if body.TimestampFmt != "1:23" {
				t.Fatalf("expected timestamp_fmt %q, got %q", "1:23", body.TimestampFmt)
			}
		})
	}
}
