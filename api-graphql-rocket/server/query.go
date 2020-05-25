//go:generate go run github.com/99designs/gqlgen
package server

import (
	"context"
	"encoding/json"
	"fmt"

	graph "github.com/asciiu/appa/api-graphql-rocket/graph/model"
	balanceRepo "github.com/asciiu/appa/lib/balance/db/sql"
	balance "github.com/asciiu/appa/lib/balance/models"
	log "github.com/sirupsen/logrus"
)

func (srv *graphQLServer) Balances(ctx context.Context) ([]*balance.Balance, error) {
	loginUser := ForContext(ctx)
	if loginUser == nil {
		return []*balance.Balance{}, fmt.Errorf("unauthorized")
	}

	balances, err := balanceRepo.FindUserBalances(srv.DB, loginUser.ID)
	return balances, err
}

//func (srv *graphQLServer) Users(ctx context.Context) ([]*user.User, error) {
//	if loginUser := ForContext(ctx); loginUser == nil {
//		return []*user.User{}, fmt.Errorf("unauthorized")
//	}
//
//	return []*user.User{}, nil
//}

func (s *graphQLServer) Messages(ctx context.Context) ([]*graph.Message, error) {
	cmd := s.redisClient.LRange("messages", 0, -1)
	if cmd.Err() != nil {
		log.Println(cmd.Err())
		return nil, cmd.Err()
	}
	res, err := cmd.Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	messages := []*graph.Message{}
	for _, mj := range res {
		m := &graph.Message{}
		err = json.Unmarshal([]byte(mj), &m)
		messages = append(messages, m)
	}
	return messages, nil
}

func (s *graphQLServer) Users(ctx context.Context) ([]string, error) {
	cmd := s.redisClient.SMembers("users")
	if cmd.Err() != nil {
		log.Println(cmd.Err())
		return nil, cmd.Err()
	}
	res, err := cmd.Result()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return res, nil
}
