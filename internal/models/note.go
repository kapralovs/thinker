package models

type Note struct {
	ID        int64    `json:"id"`
	CreatedBy string   `json:"created_by"`
	Title     string   `json:"title"`
	Text      string   `json:"text"`
	Tags      []string `json:"tags"`
}
