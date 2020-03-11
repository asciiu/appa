package server_test

import (
	"context"
	"log"
	"testing"

	"github.com/asciiu/appa/api-graphql-rocket/server"
	userQuery "github.com/asciiu/appa/lib/user/db/sql"
	"github.com/stretchr/testify/assert"
)

func TestSignup(t *testing.T) {
	cfg := server.Config{
		RedisURL: "localhost:6379",
		DBURL:    "postgres://postgres@localhost:5432/appa_test?sslmode=disable",
	}

	s, err := server.NewGraphQLServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	email := "test.email"
	username := "jerry"

	user, err := s.Signup(context.Background(), email, username, "password")
	if err != nil {
		log.Fatal(err)
	}

	assert.Equal(t, email, user.Email, "email does not match")
	assert.Equal(t, username, user.Username, "username does not match")
	assert.Equal(t, false, user.EmailVerified, "email verified should be false")

	userQuery.DeleteUserHard(s.DB, user.ID)
}

func TestLogin(t *testing.T) {
	cfg := server.Config{
		RedisURL: "localhost:6379",
		DBURL:    "postgres://postgres@localhost:5432/appa_test?sslmode=disable",
	}

	s, err := server.NewGraphQLServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	email := "test.email"
	username := "jerry"

	user, err := s.Signup(context.Background(), email, username, "password")
	if err != nil {
		log.Fatal(err)
	}
	userQuery.UpdateEmailVerified(s.DB, user.ID, true)

	token, err := s.Login(context.Background(), email, "password", true)
	if err != nil {
		log.Fatal(err)
	}

	assert.NotEmpty(t, token.Jwt, "must contain jwt")
	assert.NotEmpty(t, token.Refresh, "must contain refresh")

	userQuery.DeleteUserHard(s.DB, user.ID)
}
