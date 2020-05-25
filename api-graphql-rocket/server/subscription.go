//go:generate go run github.com/99designs/gqlgen
package server

import (
	"context"
	"fmt"

	graph "github.com/asciiu/appa/api-graphql-rocket/graph/model"
	log "github.com/sirupsen/logrus"
)

func (s *graphQLServer) createUser(user string) error {
	// Upsert user
	if err := s.redisClient.SAdd("users", user).Err(); err != nil {
		return err
	}
	// Notify new user joined
	s.mutex.Lock()
	for _, ch := range s.userChannels {
		ch <- user
	}
	s.mutex.Unlock()
	return nil
}

func (s *graphQLServer) MessagePosted(ctx context.Context, user string) (<-chan *graph.Message, error) {
	log.Info(fmt.Sprintf("MessagePosted: %s\n", user))

	err := s.createUser(user)
	if err != nil {
		return nil, err
	}

	// Create new channel for request
	messages := make(chan *graph.Message, 1)
	s.mutex.Lock()
	s.messageChannels[user] = messages
	s.mutex.Unlock()

	// Delete channel when done
	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.messageChannels, user)
		s.mutex.Unlock()
	}()

	return messages, nil
}

func (s *graphQLServer) UserJoined(ctx context.Context, user string) (<-chan string, error) {
	log.Info(fmt.Sprintf("UserJoined: %s\n", user))

	err := s.createUser(user)
	if err != nil {
		return nil, err
	}

	// Create new channel for request
	users := make(chan string, 1)
	s.mutex.Lock()
	s.userChannels[user] = users
	s.mutex.Unlock()

	// Delete channel when done
	go func() {
		<-ctx.Done()
		s.mutex.Lock()
		delete(s.userChannels, user)
		s.mutex.Unlock()
	}()

	return users, nil
}
