package apiql

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/asciiu/appa/api-graphql/auth"
	repoUser "github.com/asciiu/appa/api-graphql/db/sql"
	"github.com/asciiu/appa/api-graphql/models"
	"github.com/vektah/gqlparser/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Signup(ctx context.Context, email, username, password string) (*models.User, error) {
	user := models.NewUser(username, email, password)
	if err := repoUser.InsertUser(r.DB, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, email, password string, remember bool) (*Token, error) {
	user, err := repoUser.FindUserByEmail(r.DB, email)
	switch {
	case err != nil && strings.Contains(err.Error(), "no rows"):
		return nil, gqlerror.Errorf("incorrect password/email")
	case err != nil:
		return nil, err
	default:
		if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password)) == nil {
			jwt, err := auth.CreateJwtToken(user.ID, auth.JwtDuration)
			if err != nil {
				return nil, err
			}

			// issue a refresh token if remember is true
			if remember {
				refreshToken := models.NewRefreshToken(user.ID)
				expiresOn := time.Now().Add(auth.RefreshDuration)
				selectAuth := refreshToken.Renew(expiresOn)

				// this needs to be checked
				if _, err := repoUser.InsertRefreshToken(r.DB, refreshToken); err != nil {
					log.Println("failed to insert refresh token: ", err)
				}

				return &Token{
					Jwt:     &jwt,
					Refresh: &selectAuth,
				}, nil
			}

			return &Token{Jwt: &jwt}, nil
		}

		return nil, fmt.Errorf("incorrect password/email")
	}
}
