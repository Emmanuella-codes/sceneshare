package utils

import (
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/models"
)

func TestValidateContentID(t *testing.T) {
	tests := []struct {
		name          string
		platform      models.Platform
		id            string
		expectedError string
	}{
		{
			name:     "valid youtube id",
			platform: models.PlatformYoutube,
			id:       "dQw4w9WgXcQ",
		},
		{
			name:          "invalid youtube id",
			platform:      models.PlatformYoutube,
			id:            "short",
			expectedError: "content_id: YouTube video ID must be 11 alphanumeric characters",
		},
		{
			name:          "unsupported platform",
			platform:      models.Platform("netflix"),
			id:            "whatever1234",
			expectedError: "platform: unsupported platform",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ValidateContentID(tc.platform, tc.id)
			if tc.expectedError == "" {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				return
			}

			if err == nil {
				t.Fatal("expected validation error, got nil")
			}
			if err.Error() != tc.expectedError {
				t.Fatalf("expected error %q, got %q", tc.expectedError, err.Error())
			}
		})
	}
}
