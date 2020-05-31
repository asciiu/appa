package server

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/asciiu/appa/api-graphql-rocket/graph/model"
	graph "github.com/asciiu/appa/api-graphql-rocket/graph/model"
	roken "github.com/asciiu/appa/api-graphql-rocket/graph/model"
	user "github.com/asciiu/appa/lib/user/models"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (srv *graphQLServer) Signup(ctx context.Context, email, username, password string) (*user.User, error) {
	log.Info(fmt.Sprintf("Signup: %s", email))
	//newUser := user.NewUser(username, email, password)
	newUser, err := srv.userController.CreateUser(username, email, password)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}

func (srv *graphQLServer) Signin(ctx context.Context, email, password string, remember bool) (*roken.TokenUser, error) {
	log.Info(fmt.Sprintf("Signin: %s", email))

	loginUser, err := srv.userController.FindUserByEmail(email)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("incorrect password/email")
	case err != nil:
		return nil, err
	case !loginUser.EmailVerified:
		// only verified accounts should be able to login
		return nil, fmt.Errorf("email account not verified")
	default:
		if bcrypt.CompareHashAndPassword([]byte(loginUser.PasswordHash), []byte(password)) == nil {
			jwt, err := createJwtToken(loginUser.ID, jwtDuration)
			if err != nil {
				return nil, err
			}

			tok := roken.Token{Jwt: &jwt}

			// issue a refresh token if remember is true
			if remember {
				if _, selectAuth, err := srv.refreshController.CreateRefreshToken(loginUser.ID, time.Now().Add(refreshDuration)); err != nil {
					log.Error(fmt.Sprintf("failed to insert refresh token: %s", err.Error()))
				} else {
					tok.Refresh = &selectAuth
				}
			}

			return &roken.TokenUser{Token: &tok, User: loginUser}, nil
		}

		return nil, fmt.Errorf("incorrect password/email")
	}
}

func (s *graphQLServer) PostMessage(ctx context.Context, input *model.MessageInput) (*model.Message, error) {
	log.Info(fmt.Sprintf("PostMessage: %s: %s", input.Username, input.Text))

	err := s.createUser(input.Username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create message
	m := &graph.Message{
		ID:        uuid.New().String(),
		UserID:    input.UserID,
		Username:  input.Username,
		Text:      input.Text,
		Type:      "comment",
		AvatarURL: input.AvatarURL,
		CreatedAt: time.Now().UTC(),
	}
	mj, _ := json.Marshal(m)
	if err := s.redisClient.LPush("messages", mj).Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	// Notify new message
	s.mutex.Lock()
	for _, ch := range s.messageChannels {
		ch <- m
	}
	s.mutex.Unlock()
	return m, nil
}
