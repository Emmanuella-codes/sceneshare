package utils

import (
	"fmt"
	"net/url"
	"regexp"

	"github.com/Emmanuella-codes/sceneshare/api/dtos"
	"github.com/Emmanuella-codes/sceneshare/api/models"
	"github.com/go-playground/validator/v10"
)

// youtubeIDRegex validates YouTube video IDs (11 alphanumeric + - _)
var youtubeIDRegex = regexp.MustCompile(`^[a-zA-Z0-9_-]{11}$`)
var validate = newValidator()

type ValidationError struct {
	Field   string
	Message string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// enforces API-level constraints before persistence.
func ValidateCreateLinkInput(input *dtos.CreateLinkInput) error {
	if err := validate.Struct(input); err != nil {
		if verrs, ok := err.(validator.ValidationErrors); ok && len(verrs) > 0 {
			verr := verrs[0]
			return &ValidationError{Field: jsonFieldName(verr.Field()), Message: validationMessage(verr)}
		}
		return err
	}

	if err := ValidateContentID(input.Platform, input.ContentID); err != nil {
		return err
	}

	return nil
}

// registers custom rules used by the API payloads.
func newValidator() *validator.Validate {
	v := validator.New()
	_ = v.RegisterValidation("http_url", func(fl validator.FieldLevel) bool {
		value := fl.Field().String()
		if value == "" {
			return true
		}

		u, err := url.ParseRequestURI(value)
		if err != nil {
			return false
		}

		return (u.Scheme == "http" || u.Scheme == "https") && u.Host != ""
	})

	return v
}

func jsonFieldName(field string) string {
	switch field {
	case "Platform":
		return "platform"
	case "ContentID":
		return "content_id"
	case "TimestampS":
		return "timestamp_s"
	case "Title":
		return "title"
	case "Thumbnail":
		return "thumbnail"
	case "ExpiresIn":
		return "expires_in"
	default:
		return field
	}
}

func validationMessage(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return "is required"
	case "oneof":
		return "must be one of: youtube"
	case "min":
		return fmt.Sprintf("must be at least %s", err.Param())
	case "max":
		return fmt.Sprintf("must be %s characters or fewer", err.Param())
	case "gt":
		return fmt.Sprintf("must be greater than %s", err.Param())
	case "http_url":
		return "must be a valid http or https URL"
	default:
		return "is invalid"
	}
}

func ValidateContentID(platform models.Platform, id string) error {
	switch platform {
	case models.PlatformYoutube:
		if !youtubeIDRegex.MatchString(id) {
			return &ValidationError{Field: "content_id", Message: "YouTube video ID must be 11 alphanumeric characters"}
		}
	default:
		return &ValidationError{Field: "platform", Message: "unsupported platform"}
	}
	return nil
}
