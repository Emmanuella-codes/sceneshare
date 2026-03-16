package models

type ClickEvent struct {
	LinkID    string `db:"link_id"`
	UserAgent string `db:"user_agent"`
	Referrer  string `db:"referrer"`
}
