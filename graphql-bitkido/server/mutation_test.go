package server_test

import (
	"context"
	"log"
	"testing"

	"github.com/asciiu/appa/api-graphql-rocket/server"
	"github.com/asciiu/appa/lib/db/gopg"
	userpg "github.com/asciiu/appa/lib/user/db/gopg"
	userModels "github.com/asciiu/appa/lib/user/models"
	"github.com/stretchr/testify/assert"
)

func TestMutations(t *testing.T) {
	cfg := server.Config{
		RedisURL: "localhost:6379",
		DBURL:    "postgres://postgres@localhost:5432/appa_test?sslmode=disable",
	}

	s, err := server.NewGraphQLServer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db, err := gopg.NewDB(cfg.DBURL)
	if err != nil {
		log.Fatal(err)
	}

	email := "test.email"
	username := "jerry"
	user := &userModels.User{}
	userRepo := userpg.NewUserRepo(db)

	t.Run("Signup", func(t *testing.T) {

		user, err = s.Signup(context.Background(), email, username, "password")
		if err != nil {
			log.Fatal(err)
		}

		user.EmailVerified = true
		userRepo.UpdateEmailVerified(user.ID, true)

		assert.Equal(t, email, user.Email, "email does not match")
		assert.Equal(t, username, user.Username, "username does not match")
		//assert.Equal(t, false, user.EmailVerified, "email verified should be false")
	})

	t.Run("Signin", func(t *testing.T) {
		token, err := s.Signin(context.Background(), email, "password", true)
		if err != nil {
			log.Fatal(err)
		}

		assert.NotEmpty(t, token.Token.Jwt, "must contain jwt")
		assert.NotEmpty(t, token.Token.Refresh, "must contain refresh")
	})

	t.Run("Incorrect password", func(t *testing.T) {
		token, err := s.Signin(context.Background(), email, "passwo", true)
		assert.Equal(t, "incorrect password/email", err.Error())
		assert.Nil(t, token, "token should be nil")
	})

	userRepo.DeleteUserHard(user.ID)
}
