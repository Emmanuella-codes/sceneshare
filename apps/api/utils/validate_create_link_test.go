package utils

import (
	"testing"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
)

func TestValidateCreateLinkInput(t *testing.T) {
	validInput := func() *dtos.CreateLinkInput {
		return &dtos.CreateLinkInput{
			Platform:   models.PlatformYoutube,
			ContentID:  "dQw4w9WgXcQ",
			TimestampS: 83,
		}
	}

	tests := []struct {
		name          string
		mutate        func(*dtos.CreateLinkInput)
		expectedField string
		expectedError string
	}{
		{
			name:   "valid payload",
			mutate: func(_ *dtos.CreateLinkInput) {},
		},
		{
			name: "missing platform",
			mutate: func(input *dtos.CreateLinkInput) {
				input.Platform = ""
			},
			expectedField: "platform",
			expectedError: "platform: is required",
		},
		{
			name: "unsupported platform",
			mutate: func(input *dtos.CreateLinkInput) {
				input.Platform = models.Platform("netflix")
			},
			expectedField: "platform",
			expectedError: "platform: unsupported platform",
		},
		{
			name: "missing content id",
			mutate: func(input *dtos.CreateLinkInput) {
				input.ContentID = ""
			},
			expectedField: "content_id",
			expectedError: "content_id: is required",
		},
		{
			name: "invalid youtube content id",
			mutate: func(input *dtos.CreateLinkInput) {
				input.ContentID = "bad-id"
			},
			expectedField: "content_id",
			expectedError: "content_id: YouTube video ID must be 11 alphanumeric characters",
		},
		{
			name: "negative timestamp",
			mutate: func(input *dtos.CreateLinkInput) {
				input.TimestampS = -1
			},
			expectedField: "timestamp_s",
			expectedError: "timestamp_s: must be at least 0",
		},
		{
			name: "expires_in zero",
			mutate: func(input *dtos.CreateLinkInput) {
				v := 0
				input.ExpiresIn = &v
			},
			expectedField: "expires_in",
			expectedError: "expires_in: must be greater than 0",
		},
		{
			name: "expires_in negative",
			mutate: func(input *dtos.CreateLinkInput) {
				v := -10
				input.ExpiresIn = &v
			},
			expectedField: "expires_in",
			expectedError: "expires_in: must be greater than 0",
		},
		{
			name: "valid http thumbnail",
			mutate: func(input *dtos.CreateLinkInput) {
				v := "http://example.com/thumb.jpg"
				input.Thumbnail = &v
			},
		},
		{
			name: "valid https thumbnail",
			mutate: func(input *dtos.CreateLinkInput) {
				v := "https://example.com/thumb.jpg"
				input.Thumbnail = &v
			},
		},
		{
			name: "invalid thumbnail url",
			mutate: func(input *dtos.CreateLinkInput) {
				v := "not-a-url"
				input.Thumbnail = &v
			},
			expectedField: "thumbnail",
			expectedError: "thumbnail: must be a valid http or https URL",
		},
		{
			name: "relative thumbnail url",
			mutate: func(input *dtos.CreateLinkInput) {
				v := "/thumb.jpg"
				input.Thumbnail = &v
			},
			expectedField: "thumbnail",
			expectedError: "thumbnail: must be a valid http or https URL",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			input := validInput()
			tc.mutate(input)

			err := ValidateCreateLinkInput(input)
			if tc.expectedError == "" {
				if err != nil {
					t.Fatalf("expected no error, got %v", err)
				}
				return
			}

			if err == nil {
				t.Fatal("expected validation error, got nil")
			}

			ve, ok := err.(*ValidationError)
			if !ok {
				t.Fatalf("expected *ValidationError, got %T", err)
			}
			if ve.Field != tc.expectedField {
				t.Fatalf("expected field %q, got %q", tc.expectedField, ve.Field)
			}
			if ve.Error() != tc.expectedError {
				t.Fatalf("expected error %q, got %q", tc.expectedError, ve.Error())
			}
		})
	}
}
