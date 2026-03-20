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

func TestGetStats(t *testing.T) {
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
				m.getStats = func(_ context.Context, code string) (*dtos.StatsResponse, error) {
					if code != "abc1234" {
						t.Fatalf("expected code %q, got %q", "abc1234", code)
					}
					return &dtos.StatsResponse{
						ShortCode:  code,
						ClickCount: 7,
						CreatedAt:  "2025-01-01T00:00:00Z",
					}, nil
				}
			},
			expectedStatus: http.StatusOK,
		},
		{
			name: "not found",
			code: "missing1",
			mockSetup: func(_ *testing.T, m *mockLinkService) {
				m.getStats = func(_ context.Context, _ string) (*dtos.StatsResponse, error) {
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
				m.getStats = func(_ context.Context, _ string) (*dtos.StatsResponse, error) {
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
				m.getStats = func(_ context.Context, _ string) (*dtos.StatsResponse, error) {
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

			r := withCode(httptest.NewRequest(http.MethodGet, "/api/v1/links/"+tc.code+"/stats", nil), tc.code)
			w := httptest.NewRecorder()

			h.GetStats(w, r)

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

			body := decodeJSON[dtos.StatsResponse](t, res)
			if body.ShortCode != tc.code {
				t.Fatalf("expected short_code %q, got %q", tc.code, body.ShortCode)
			}
			if body.ClickCount != 7 {
				t.Fatalf("expected click_count %d, got %d", 7, body.ClickCount)
			}
			if body.CreatedAt != "2025-01-01T00:00:00Z" {
				t.Fatalf("expected created_at %q, got %q", "2025-01-01T00:00:00Z", body.CreatedAt)
			}
		})
	}
}
