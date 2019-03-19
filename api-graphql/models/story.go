package models

import (
	"github.com/asciiu/appa/api-graphql/constants"
	"github.com/google/uuid"
)

type Story struct {
	ID        string `json:"id"`
	AuthorID  string `json:"author_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Status    string `json:"status"`
	CreatedOn string `json:"createdOn"`
}

func NewStory(userID, title, content string) *Story {
	newID := uuid.New()

	user := Story{
		ID:       newID.String(),
		AuthorID: userID,
		Title:    title,
		Content:  content,
		Status:   constants.Unpublished,
	}
	return &user
}
