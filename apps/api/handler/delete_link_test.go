package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/service"
)

func TestDeleteLink(t *testing.T) {
	tests := []struct {
		name           string
		code           string
		token          string
		mockSetup      func(*testing.T, *mockLinkService)
		expectedStatus int
		expectedCode   string
	}{
		{
			name:           "missing token",
			code:           "abc1234",
			token:          "",
			mockSetup:      func(_ *testing.T, _ *mockLinkService) {},
			expectedStatus: http.StatusUnauthorized,
			expectedCode:   "MISSING_TOKEN",
		},
		{
			name:  "success",
			code:  "abc1234",
			token: "owner-token",
			mockSetup: func(t *testing.T, m *mockLinkService) {
				t.Helper()
				m.deleteLink = func(_ context.Context, code, token string) error {
					if code != "abc1234" {
						t.Fatalf("expected code %q, got %q", "abc1234", code)
					}
					if token != "owner-token" {
						t.Fatalf("expected token %q, got %q", "owner-token", token)
					}
					return nil
				}
			},
			expectedStatus: http.StatusNoContent,
		},
		{
			name:  "forbidden",
			code:  "abc1234",
			token: "wrong-token",
			mockSetup: func(_ *testing.T, m *mockLinkService) {
				m.deleteLink = func(_ context.Context, _, _ string) error {
					return service.ErrForbidden
				}
			},
			expectedStatus: http.StatusForbidden,
			expectedCode:   "FORBIDDEN",
		},
		{
			name:  "not found",
			code:  "missing1",
			token: "owner-token",
			mockSetup: func(_ *testing.T, m *mockLinkService) {
				m.deleteLink = func(_ context.Context, _, _ string) error {
					return service.ErrNotFound
				}
			},
			expectedStatus: http.StatusNotFound,
			expectedCode:   "LINK_NOT_FOUND",
		},
		{
			name:  "internal error",
			code:  "broken12",
			token: "owner-token",
			mockSetup: func(_ *testing.T, m *mockLinkService) {
				m.deleteLink = func(_ context.Context, _, _ string) error {
					return errors.New("boom")
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

			r := withCode(httptest.NewRequest(http.MethodDelete, "/api/v1/links/"+tc.code, nil), tc.code)
			if tc.token != "" {
				r.Header.Set("X-Owner-Token", tc.token)
			}
			w := httptest.NewRecorder()

			h.DeleteLink(w, r)

			res := w.Result()
			if res.StatusCode != tc.expectedStatus {
				t.Fatalf("expected status %d, got %d", tc.expectedStatus, res.StatusCode)
			}

			if tc.expectedCode == "" {
				return
			}

			body := decodeJSON[errorResponse](t, res)
			if body.Code != tc.expectedCode {
				t.Fatalf("expected code %q, got %q", tc.expectedCode, body.Code)
			}
		})
	}
}
