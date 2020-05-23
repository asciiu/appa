package server

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	roken "github.com/asciiu/appa/api-graphql-rocket/graph/model"
	tokenRepo "github.com/asciiu/appa/lib/refreshToken/db/sql"
	token "github.com/asciiu/appa/lib/refreshToken/models"
	userRepo "github.com/asciiu/appa/lib/user/db/sql"
	user "github.com/asciiu/appa/lib/user/models"
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
