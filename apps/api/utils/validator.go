package utils

import (
	"fmt"
	"regexp"

	// "github.com/Emmanuella-codes/sceneshare/api/model"
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

// func ValidateContentID(platform model.Platform, id string) error {
// 	switch platform {
// 	case model.PlatformYoutube: 
// 		if !youtubeIDRegex.
// 	}
// }
