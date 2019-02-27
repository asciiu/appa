package apiql

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func (r *mutationResolver) RegisterUser(ctx context.Context, input NewUser) (*models.User, error) {
	user := models.NewUser(input.Username, input.Email, input.Password)
	if err := repoUser.InsertUser(r.DB, user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *mutationResolver) Login(ctx context.Context, input NewLogin) (*Token, error) {
	user, err := repoUser.FindUserByEmail(r.DB, input.Email)
	if err != nil {
		return nil, err
	}

	if bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)) == nil {
		jwt := "jwt tokie"
		refresh := fmt.Sprintf("refresh tokie %b", input.Remember)

		return &Token{
			Jwt:     &jwt,
			Refresh: &refresh,
		}, nil
	}

	return nil, errors.New("incorrect password/email")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	panic("not implemented")
}
