package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/asciiu/appa/api-graphql/graph/generated"
	models1 "github.com/asciiu/appa/api-graphql/graph/models"
	models3 "github.com/asciiu/appa/api-graphql/models"
	models2 "github.com/asciiu/appa/lib/balance/models"
	"github.com/asciiu/appa/lib/user/models"
)

func (r *mutationResolver) Signup(ctx context.Context, email string, username string, password string) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Login(ctx context.Context, email string, password string, remember bool) (*models1.Token, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateBet(ctx context.Context, userID string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateStory(ctx context.Context, title string, json string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateStory(ctx context.Context, id string, title string, json string, status string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Balances(ctx context.Context) ([]*models2.Balance, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Info(ctx context.Context) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUser(ctx context.Context) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) UserSummary(ctx context.Context) (*models.UserSummary, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) FindOrder(ctx context.Context, id string) (*models.Order, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListStories(ctx context.Context) ([]*models3.Story, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
