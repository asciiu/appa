package graphql

import (
	"context"
	"fmt"

	"github.com/asciiu/appa/api-graphql/auth"
	"github.com/asciiu/appa/api-graphql/models"
	balanceRepo "github.com/asciiu/appa/lib/balance/db/sql"
	balance "github.com/asciiu/appa/lib/balance/models"
	user "github.com/asciiu/appa/lib/user/models"
	log "github.com/sirupsen/logrus"
)

type queryResolver struct{ *Resolver }

func (r *queryResolver) Balances(ctx context.Context) ([]*balance.Balance, error) {
	loginUser := auth.ForContext(ctx)
	if loginUser == nil {
		return []*balance.Balance{}, fmt.Errorf("unauthorized")
	}

	balances, err := balanceRepo.FindUserBalances(r.DB, loginUser.ID)
	return balances, err
}

func (r *queryResolver) Users(ctx context.Context) ([]*user.User, error) {
	if loginUser := auth.ForContext(ctx); loginUser == nil {
		return []*user.User{}, fmt.Errorf("unauthorized")
	}

	return []*user.User{}, nil
}

func (r *queryResolver) Info(ctx context.Context) (*user.User, error) {
	loginUser := auth.ForContext(ctx)
	if loginUser == nil {
		return &user.User{}, fmt.Errorf("unauthorized")
	}

	return loginUser, nil
}

func (r *queryResolver) GetUser(ctx context.Context) (*user.User, error) {
	loginUser := auth.ForContext(ctx)
	if loginUser == nil {
		return nil, fmt.Errorf("unauthorized")
	}

	return loginUser, nil
}

func (r *queryResolver) UserSummary(ctx context.Context) (*user.UserSummary, error) {
	loginUser := auth.ForContext(ctx)
	if loginUser == nil {
		return nil, fmt.Errorf("unauthorized")
	}

	log.Info("UserSummary", loginUser)

	balances, err := balanceRepo.FindUserBalances(r.DB, loginUser.ID)
	if err != nil {
		log.Error("encountered error when pulling balances: ", err)
	}

	for _, balance := range balances {
		if balance.Symbol == "BTC" {
			return &user.UserSummary{
				User:    loginUser,
				Balance: balance,
			}, nil
		}
	}

	// TODO when adding more currencies you'll need to determine
	// total balance in BTC

	return nil, nil
}

func (r *queryResolver) FindOrder(ctx context.Context, id string) (*user.Order, error) {
	loginUser := auth.ForContext(ctx)
	if loginUser == nil {
		return nil, fmt.Errorf("unauthorized")
	}

	return &user.Order{ID: id, Txt: "yeah!"}, nil
}

func (r *queryResolver) ListStories(ctx context.Context) ([]*models.Story, error) {
	//user := auth.ForContext(ctx)
	//if user == nil {
	//	return []models.Story{}, fmt.Errorf("unauthorized")
	//}

	stories := []*models.Story{}
	return stories, nil
}
