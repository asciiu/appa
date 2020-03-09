//go:generate go run github.com/99designs/gqlgen
package server

import (
	"context"
	"encoding/json"
	"log"

	"github.com/asciiu/appa/api-graphql-chat/graph/generated"
	graph "github.com/asciiu/appa/api-graphql-chat/graph/model"
)

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

func (s *graphQLServer) Query() generated.QueryResolver {
	return s
}
