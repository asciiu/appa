package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/asciiu/appa/api-graphql-rocket/graph/generated"
	"github.com/asciiu/appa/api-graphql-rocket/graph/model"
)

func (r *mutationResolver) PostMessage(ctx context.Context, user string, text string) (*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Messages(ctx context.Context) ([]*model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) MessagePosted(ctx context.Context, user string) (<-chan *model.Message, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) UserJoined(ctx context.Context, user string) (<-chan string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated1.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated1.QueryResolver { return &queryResolver{r} }

// Subscription returns generated1.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated1.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
