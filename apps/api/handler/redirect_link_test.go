package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/models"
	"github.com/Emmanuella-codes/sceneshare/api/service"
)

func TestRedirectLink(t *testing.T) {
	tests := []struct {
		name           string
		code           string
		mockSetup      func(*testing.T, *mockLinkService)
		expectedStatus int
		expectedCode   string
		expectedLocation string
	}{
		{
			name: "valid link",
			code: "abc1234",
			mockSetup: func(t *testing.T, m *mockLinkService) {
				t.Helper()
				m.getLinkForRedirect = func(_ context.Context, code string) (*models.Link, error) {
					if code != "abc1234" {
						t.Fatalf("expected code %q, got %q", "abc1234", code)
					}
					return &models.Link{
						ID:         "123456",
						ShortCode:  code,
						Platform:   models.PlatformYoutube,
						ContentID:  "dQw4w9WgXcQ",
						TimestampS: 83,
					}, nil
				}
			},
			expectedStatus: http.StatusFound,
			expectedLocation: "https://www.youtube.com/watch?t=83s&v=dQw4w9WgXcQ",
		},
		{
			name: "not found",
			code: "missing1",
			mockSetup: func(t *testing.T, m *mockLinkService) {
				t.Helper()
				m.getLinkForRedirect = func(_ context.Context, code string) (*models.Link, error) {
					return nil, service.ErrNotFound
				}
			},
			expectedStatus: http.StatusNotFound,
			expectedCode:   "LINK_NOT_FOUND",
		},
		{
			name: "internal error",
			code: "abc1234",
			mockSetup: func(t *testing.T, m *mockLinkService) {
				t.Helper()
				m.getLinkForRedirect = func(_ context.Context, code string) (*models.Link, error) {
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

			r := withCode(httptest.NewRequest(http.MethodGet, "/r/"+tc.code, nil), tc.code)
			w := httptest.NewRecorder()
			h.Redirect(w, r)

			res := w.Result()
			if res.StatusCode != tc.expectedStatus {
				t.Errorf("wanted %d, got %d", tc.expectedStatus, res.StatusCode)
			}
			if tc.expectedCode != "" {
				body := decodeJSON[errorResponse](t, res)
				if body.Code != tc.expectedCode {
					t.Errorf("wanted code %q, got %q", tc.expectedCode, body.Code)
				}
				return
			}
			if got := res.Header.Get("Location"); got != tc.expectedLocation {
				t.Errorf("wanted location %q, got %q", tc.expectedLocation, got)
			}
		})
	}
}
