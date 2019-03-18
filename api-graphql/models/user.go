package models

import (
	"log"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(username, email, password string) *User {
	newID := uuid.New()

	user := User{
		ID:            newID.String(),
		Username:      username,
		Email:         email,
		EmailVerified: false,
		PasswordHash:  GenHash([]byte(password)),
	}
	return &user
}

func GenHash(pwd []byte) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	// hash this using a server secret key
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

type Order struct {
	ID  string
	Txt string
}

type User struct {
	ID            string
	Username      string
	Email         string
	EmailVerified bool
	PasswordHash  string
}

type UserInfo struct {
	UserID   string `json:"userID"`
	Username string `json:"username"`
}

func (user *User) Info() *UserInfo {
	return &UserInfo{
		UserID:   user.ID,
		Username: user.Username,
	}
}
