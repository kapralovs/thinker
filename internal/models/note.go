package models

type Note struct {
	ID     int64    `json:"id"`
	UserID int64    `json:"user_id"`
	Title  string   `json:"title"`
	Text   string   `json:"text"`
	Tags   []string `json:"tags"`
}
