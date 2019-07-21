package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

func NewUser(username, email, password string) *User {
	newID := uuid.New()
	now := string(pq.FormatTimestamp(time.Now().UTC()))

	user := User{
		ID:            newID.String(),
		Username:      username,
		Email:         email,
		EmailVerified: false,
		PasswordHash:  HashAndSalt([]byte(password)),
		CreatedOn:     now,
		UpdatedOn:     now,
	}
	return &user
}

func HashAndSalt(pwd []byte) string {

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

type User struct {
	ID            string
	Username      string
	Email         string
	EmailVerified bool
	PasswordHash  string
	CreatedOn     string
	UpdatedOn     string
}

//type UserInfo struct {
//	UserID   string `json:"userID"`
//	Username string `json:"username"`
//	Email    string `json:"email"`
//}
//
//func (user *User) Info() *UserInfo {
//	return &UserInfo{
//		UserID:   user.ID,
//		Username: user.Username,
//		Email:    user.Email,
//	}
//}
