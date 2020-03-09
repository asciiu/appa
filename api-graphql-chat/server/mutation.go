//go:generate go run github.com/99designs/gqlgen
package server

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/asciiu/appa/api-graphql-chat/graph/generated"
	graph "github.com/asciiu/appa/api-graphql-chat/graph/model"
	"github.com/segmentio/ksuid"
)

func (s *graphQLServer) PostMessage(ctx context.Context, user string, text string) (*graph.Message, error) {
	err := s.createUser(user)
	if err != nil {
		return nil, err
	}

	// Create message
	m := &graph.Message{
		ID:        ksuid.New().String(),
		CreatedAt: time.Now().UTC(),
		Text:      text,
		User:      user,
	}
	mj, _ := json.Marshal(m)
	if err := s.redisClient.LPush("messages", mj).Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	// Notify new message
	s.mutex.Lock()
	for _, ch := range s.messageChannels {
		ch <- m
	}
	s.mutex.Unlock()
	return m, nil
}

func (s *graphQLServer) Mutation() generated.MutationResolver {
	return s
}
