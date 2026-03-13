package model

import "time"

type Platform string

const (
	PlatformYoutube Platform = "youtube"
	PlatformNetflix Platform = "netflix"
	PlatformPrime   Platform = "prime"
)

func (p Platform) IsValid() bool {
	switch p {
	case PlatformYoutube, PlatformNetflix, PlatformPrime:
		return true
	default:
		return false
	}
}

type Link struct {
	ID         string     `db:"id"`
	ShortCode  string     `db:"short_code"`
	Platform   Platform   `db:"platform"`
	ContentID  string     `db:"content_id"`
	TimestampS int        `db:"timestamp_s"`
	Title      string     `db:"title"`
	Thumbnail  string     `db:"thumbnail"`
	CreatedBy  string     `db:"created_by"`
	CreatedAt  time.Time  `db:"created_at"`
	ExpiresAt  *time.Time `db:"expires_at"`
	ClickCount int        `db:"click_count"`
}
