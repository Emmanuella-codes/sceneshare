package service

import (
	"errors"
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/models"
)

func TestBuildDeepLink(t *testing.T) {
	tests := []struct {
		name        string
		link        *models.Link
		expectedURL string
		expectedErr error
	}{
		{
			name: "youtube with timestamp",
			link: &models.Link{
				Platform:   models.PlatformYoutube,
				ContentID:  "dQw4w9WgXcQ",
				TimestampS: 83,
			},
			expectedURL: "https://www.youtube.com/watch?t=83s&v=dQw4w9WgXcQ",
		},
		{
			name: "youtube without timestamp",
			link: &models.Link{
				Platform:   models.PlatformYoutube,
				ContentID:  "dQw4w9WgXcQ",
				TimestampS: 0,
			},
			expectedURL: "https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		},
		{
			name: "unsupported platform",
			link: &models.Link{
				Platform:   "netflix",
				ContentID:  "some-id",
				TimestampS: 83,
			},
			expectedErr: ErrUnsupportedPlatform,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := BuildDeepLink(tc.link)
			if tc.expectedErr != nil {
				if !errors.Is(err, tc.expectedErr) {
					t.Fatalf("expected error %v, got %v", tc.expectedErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("expected no error, got %v", err)
			}
			if got != tc.expectedURL {
				t.Fatalf("expected URL %q, got %q", tc.expectedURL, got)
			}
		})
	}
}
