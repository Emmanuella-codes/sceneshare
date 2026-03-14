package utils

import (
	"fmt"
	"regexp"

	"github.com/Emmanuella-codes/sceneshare/api/models"
)

// youtubeIDRegex validates YouTube video IDs (11 alphanumeric + - _)
var youtubeIDRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{11}$`)

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func ValidateContentID(platform models.Platform, id string) error {
	switch platform {
	case models.PlatformYoutube: 
		if !youtubeIDRegex.MatchString(id) {
			return &ValidationError{Field: "content_id", Message: "YouTube video ID must be 11 alphanumeric characters"}
		}
	default:
		if len(id) == 0 || len(id) > 200 {
			return &ValidationError{Field: "content_id", Message: "content ID must be between 1 and 200 characters"}
		}
	}
	return nil
}
