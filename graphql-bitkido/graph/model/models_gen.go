// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"time"

	"github.com/asciiu/appa/lib/user/models"
)

type Bet struct {
	ID        string `json:"id"`
	UserID    string `json:"userID"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatarURL"`
}

type Message struct {
	ID        string    `json:"id"`
	UserID    string    `json:"userID"`
	Username  string    `json:"username"`
	Text      string    `json:"text"`
	Type      string    `json:"type"`
	AvatarURL string    `json:"avatarURL"`
	CreatedAt time.Time `json:"createdAt"`
}

type MessageInput struct {
	UserID    string `json:"userID"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	AvatarURL string `json:"avatarURL"`
}

type Token struct {
	Jwt     *string `json:"jwt"`
	Refresh *string `json:"refresh"`
}

type TokenUser struct {
	User  *models.User `json:"user"`
	Token *Token       `json:"token"`
}
