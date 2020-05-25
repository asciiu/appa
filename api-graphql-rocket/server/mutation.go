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
	tokenRepo "github.com/asciiu/appa/lib/refreshToken/db/sql"
	token "github.com/asciiu/appa/lib/refreshToken/models"
	userRepo "github.com/asciiu/appa/lib/user/db/sql"
	user "github.com/asciiu/appa/lib/user/models"
	"github.com/segmentio/ksuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (srv *graphQLServer) Signup(ctx context.Context, email, username, password string) (*user.User, error) {
	log.Info(fmt.Sprintf("Signup: %s", email))
	newUser := user.NewUser(username, email, password)
	if err := userRepo.InsertUser(srv.DB, newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (srv *graphQLServer) Signin(ctx context.Context, email, password string, remember bool) (*roken.TokenUser, error) {
	log.Info(fmt.Sprintf("Signin: %s", email))

	loginUser, err := userRepo.FindUserByEmail(srv.DB, email)
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
				refreshToken := token.NewRefreshToken(loginUser.ID)
				expiresOn := time.Now().Add(refreshDuration)
				selectAuth := refreshToken.Renew(expiresOn)

				// this needs to be checked
				if _, err := tokenRepo.InsertRefreshToken(srv.DB, refreshToken); err != nil {
					log.Error(fmt.Sprintf("failed to insert refresh token: %s", err.Error()))
				}
				tok.Refresh = &selectAuth
			}

			return &roken.TokenUser{Token: &tok, User: loginUser}, nil
		}

		return nil, fmt.Errorf("incorrect password/email")
	}
}

func (s *graphQLServer) PostMessage(ctx context.Context, userID string, username string, text string, avatarURL string) (*model.Message, error) {
	//func (s *graphQLServer) PostMessage(ctx context.Context, user string, text string) (*graph.Message, error) {
	log.Info(fmt.Sprintf("PostMessage: %s: %s", username, text))

	err := s.createUser(username)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	// Create message
	m := &graph.Message{
		ID:        ksuid.New().String(),
		UserID:    userID,
		Username:  username,
		Text:      text,
		Type:      "comment",
		AvatarURL: avatarURL,
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
