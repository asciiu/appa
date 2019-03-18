package models

type Story struct {
	ID        string `json:"id"`
	AuthorID  string `json:"author_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	CreatedOn string `json:"createdOn"`
}
