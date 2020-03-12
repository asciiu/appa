package server

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/asciiu/appa/api-graphql-rocket/auth"
	"github.com/asciiu/appa/api-graphql-rocket/graph/generated"
	roken "github.com/asciiu/appa/api-graphql-rocket/graph/model"
	tokenRepo "github.com/asciiu/appa/lib/refreshToken/db/sql"
	token "github.com/asciiu/appa/lib/refreshToken/models"
	userRepo "github.com/asciiu/appa/lib/user/db/sql"
	user "github.com/asciiu/appa/lib/user/models"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

func (r *graphQLServer) Signup(ctx context.Context, email, username, password string) (*user.User, error) {
	newUser := user.NewUser(username, email, password)
	if err := userRepo.InsertUser(r.DB, newUser); err != nil {
		return nil, err
	}
	return newUser, nil
}

func (r *graphQLServer) Login(ctx context.Context, email, password string, remember bool) (*roken.Token, error) {
	log.Info(fmt.Sprintf("login: %s", email))

	loginUser, err := userRepo.FindUserByEmail(r.DB, email)
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
			jwt, err := auth.CreateJwtToken(loginUser.ID, auth.JwtDuration)
			if err != nil {
				return nil, err
			}

			// issue a refresh token if remember is true
			if remember {
				refreshToken := token.NewRefreshToken(loginUser.ID)
				expiresOn := time.Now().Add(auth.RefreshDuration)
				selectAuth := refreshToken.Renew(expiresOn)

				// this needs to be checked
				if _, err := tokenRepo.InsertRefreshToken(r.DB, refreshToken); err != nil {
					log.Error(fmt.Sprintf("failed to insert refresh token: %s", err.Error()))
				}

				return &roken.Token{
					Jwt:     &jwt,
					Refresh: &selectAuth,
				}, nil
			}

			return &roken.Token{Jwt: &jwt}, nil
		}

		return nil, fmt.Errorf("incorrect password/email")
	}
}

func (s *graphQLServer) Mutation() generated.MutationResolver {
	return s
}
