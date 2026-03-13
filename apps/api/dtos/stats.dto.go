package dtos

type StatsResponse struct {
	ShortCode  string `json:"short_code"`
	ClickCount int    `json:"click_count"`
	CreatedAt  string `json:"created_at"`
}
