package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/utils"
)

func TestCreateLink(t *testing.T) {
	tests := []struct {
		name           string
		body           string
		mockSetup      func(*testing.T, *mockLinkService)
		expectedStatus int
		expectedCode   string
	}{
		{
			name: "valid request",
			body: `{"platform":"youtube","content_id":"dQw4w9WgXcQ","timestamp_s":83}`,
			mockSetup: func(t *testing.T, m *mockLinkService) {
				t.Helper()
				m.createLink = func(_ context.Context, input *dtos.CreateLinkInput) (*dtos.LinkResponse, error) {
					if input.Platform != "youtube" {
						t.Fatalf("expected platform youtube, got %q", input.Platform)
					}
					if input.ContentID != "dQw4w9WgXcQ" {
						t.Fatalf("expected content_id dQw4w9WgXcQ, got %q", input.ContentID)
					}
					if input.TimestampS != 83 {
						t.Fatalf("expected timestamp_s 83, got %d", input.TimestampS)
					}
					return &dtos.LinkResponse{ShortCode: "123456"}, nil
				}
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:           "invalid json",
			body:           `{bad}`,
			mockSetup:      func(_ *testing.T, _ *mockLinkService) {},
			expectedStatus: http.StatusBadRequest,
			expectedCode:   "INVALID_REQUEST",
		},
		{
			name: "validation error",
			body: `{"platform":"youtube","content_id":"dQw4w9WgXcQ","timestamp_s":83}`,
			mockSetup: func(_ *testing.T, m *mockLinkService) {
				m.createLink = func(_ context.Context, _ *dtos.CreateLinkInput) (*dtos.LinkResponse, error) {
					return nil, &utils.ValidationError{Field: "platform", Message: "is required"}
				}
			},
			expectedStatus: http.StatusBadRequest,
			expectedCode:   "VALIDATION_ERROR",
		},
		{
			name: "internal error",
			body: `{"platform":"youtube","content_id":"dQw4w9WgXcQ","timestamp_s":83}`,
			mockSetup: func(_ *testing.T, m *mockLinkService) {
				m.createLink = func(_ context.Context, _ *dtos.CreateLinkInput) (*dtos.LinkResponse, error) {
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

			r := httptest.NewRequest(http.MethodPost, "/api/v1/links", strings.NewReader(tc.body))
			w := httptest.NewRecorder()
			h.CreateLink(w, r)

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

			body := decodeJSON[dtos.LinkResponse](t, res)
			if body.ShortCode != "123456" {
				t.Errorf("wanted short_code %q, got %q", "123456", body.ShortCode)
			}
		})
	}
}
