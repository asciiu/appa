package apiql

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/asciiu/appa/apiql/auth"
	repoUser "github.com/asciiu/appa/apiql/db/sql"
	"github.com/asciiu/appa/apiql/models"
	"golang.org/x/crypto/bcrypt"
)

type Resolver struct {
	DB *sql.DB
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SignUp(ctx context.Context, input NewUser) (*models.User, error) {
	user := models.NewUser(input.Username, input.Email, input.Password)
	if err := repoUser.InsertUser(r.DB, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, input NewLogin) (*Token, error) {
	user, err := repoUser.FindUserByEmail(r.DB, input.Email)
	switch {
	case err != nil && strings.Contains(err.Error(), "no rows"):
		return nil, fmt.Errorf("incorrect password/email")
	case err != nil:
		return nil, err
	default:
		if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)) == nil {
			jwt, err := auth.CreateJwtToken(user.ID, auth.JwtDuration)
			if err != nil {
				return nil, err
			}

			// issue a refresh token if remember is true
			if input.Remember {
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

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	if user := auth.ForContext(ctx); user == nil {
		return []models.User{}, fmt.Errorf("Access denied")
	}

	return []models.User{}, nil
}
