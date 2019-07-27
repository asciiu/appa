package apiql

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/asciiu/appa/api-graphql/auth"
	repo "github.com/asciiu/appa/api-graphql/db/sql"
	"github.com/asciiu/appa/api-graphql/models"
	"github.com/vektah/gqlparser/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) Signup(ctx context.Context, email, username, password string) (*models.User, error) {
	user := models.NewUser(username, email, password)
	if err := repo.InsertUser(r.DB, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, email, password string, remember bool) (*Token, error) {
	user, err := repo.FindUserByEmail(r.DB, email)
	switch {
	case err != nil && strings.Contains(err.Error(), "no rows"):
		return nil, gqlerror.Errorf("incorrect password/email")
	case err != nil:
		return nil, err
	case !user.EmailVerified:
		// only verified accounts should be able to login
		return nil, gqlerror.Errorf("email account not verified")
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
				if _, err := repo.InsertRefreshToken(r.DB, refreshToken); err != nil {
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

func (r *mutationResolver) CreateStory(ctx context.Context, title, jsonData string) (string, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return "", fmt.Errorf("unauthorized")
	}

	story := models.NewStory(user.ID, title, jsonData)
	if err := repo.InsertStory(r.DB, story); err != nil {
		return "", err
	}

	return story.ID, nil
}

func (r *mutationResolver) UpdateStory(ctx context.Context, storyID, title, jsonData, status string) (bool, error) {
	user := auth.ForContext(ctx)
	if user == nil {
		return false, fmt.Errorf("unauthorized")
	}

	story := models.Story{
		ID:       storyID,
		AuthorID: user.ID,
		Title:    title,
		Content:  jsonData,
		Status:   status,
	}

	if err := repo.UpdateStory(r.DB, &story); err != nil {
		return false, err
	}

	return true, nil
}
