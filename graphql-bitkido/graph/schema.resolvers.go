package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	generated1 "github.com/asciiu/appa/api-graphql-rocket/graph/generated"
	"github.com/asciiu/appa/api-graphql-rocket/graph/model"
	"github.com/asciiu/appa/lib/user/models"
)

func (r *mutationResolver) Signup(ctx context.Context, email string, username string, password string) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Signin(ctx context.Context, email string, password string, remember bool) (*model.TokenUser, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) Signout(ctx context.Context, selector string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PostBet(ctx context.Context, description string, amount int) (*model.Bet, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) PostMessage(ctx context.Context, input *model.MessageInput) (*model.Message, error) {
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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) Login(ctx context.Context, email string, password string, remember bool) (*model.Token, error) {
	panic(fmt.Errorf("not implemented"))
}
