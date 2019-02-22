package apiql

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/asciiu/appa/apiql/models"
)

type Resolver struct {
	users []models.User
	todos []models.Todo
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) Todo() TodoResolver {
	return &todoResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input NewTodo) (models.Todo, error) {
	todo := models.Todo{
		Text:   input.Text,
		ID:     fmt.Sprintf("T%d", rand.Int()),
		UserID: input.UserID,
	}
	r.todos = append(r.todos, todo)
	return todo, nil
}

func (r *mutationResolver) RegisterUser(ctx context.Context, input NewUser) (models.User, error) {
	user := *models.NewUser(input.Username, input.Email, input.Password)
	r.users = append(r.users, user)
	return user, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]models.Todo, error) {
	return r.todos, nil
}
func (r *queryResolver) Users(ctx context.Context) ([]models.User, error) {
	return r.users, nil
}

type todoResolver struct{ *Resolver }

func (r *todoResolver) User(ctx context.Context, obj *models.Todo) (models.User, error) {
	return models.User{ID: obj.UserID}, nil
}
