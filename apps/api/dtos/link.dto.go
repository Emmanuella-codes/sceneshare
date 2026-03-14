package dtos

import (
	"time"

	"github.com/Emmanuella-codes/sceneshare/api/models"
)

// raw JSON payload decoded from the HTTP request
type CreateLinkInput struct {
	Platform   models.Platform `json:"platform" validate:"required"`
	ContentID  string          `json:"content_id" validate:"required,min=1,max=200"`
	TimestampS int             `json:"timestamp_s" validate:"min=0"`
	Title      *string         `json:"title" validate:"omitempty,max=500"`
	Thumbnail  *string         `json:"thumbnail" validate:"omitempty,url,max=1000"`
	ExpiresIn  *int            `json:"expires_in" validate:"omitempty,duration"`
}

// short code generation parameters
type CreateLinkParams struct {
	ShortCode  string
	Platform   models.Platform
	ContentID  string
	TimestampS int
	Title      *string
	Thumbnail  *string
	OwnerToken string
	ExpiresAt  *time.Time
}

type LinkResponse struct {
	ID           string  `json:"id"`
	ShortCode    string  `json:"short_code"`
	ShortURL     string  `json:"short_url"`
	Platform     string  `json:"platform"`
	ContentID    string  `json:"content_id"`
	TimestampS   int     `json:"timestamp_s"`
	TimestampFmt string  `json:"timestamp_fmt"`
	Title        *string `json:"title"`
	Thumbnail    *string `json:"thumbnail"`
	OwnerToken   *string `json:"owner_token,omitempty"`
	ClickCount   int     `json:"click_count"`
	CreatedAt    string  `json:"created_at"`
	ExpiresAt    *string `json:"expires_at,omitempty"`
}
