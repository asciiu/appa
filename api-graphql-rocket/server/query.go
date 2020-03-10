//go:generate go run github.com/99designs/gqlgen
package server

import (
	"context"
	"fmt"

	"github.com/asciiu/appa/api-graphql-rocket/graph/generated"
	"github.com/asciiu/appa/api-graphql/auth"
	balanceRepo "github.com/asciiu/appa/lib/balance/db/sql"
	balance "github.com/asciiu/appa/lib/balance/models"
	user "github.com/asciiu/appa/lib/user/models"
)

func (r *graphQLServer) Balances(ctx context.Context) ([]*balance.Balance, error) {
	loginUser := auth.ForContext(ctx)
	if loginUser == nil {
		return []*balance.Balance{}, fmt.Errorf("unauthorized")
	}

	balances, err := balanceRepo.FindUserBalances(r.DB, loginUser.ID)
	return balances, err
}

func (r *graphQLServer) Users(ctx context.Context) ([]*user.User, error) {
	if loginUser := auth.ForContext(ctx); loginUser == nil {
		return []*user.User{}, fmt.Errorf("unauthorized")
	}

	return []*user.User{}, nil
}

func (s *graphQLServer) Query() generated.QueryResolver {
	return s
}
