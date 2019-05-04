package main

import (
	"context"
	"database/sql"

	protoStory "github.com/asciiu/appa/story-service/proto/story"
)

type StoryService struct {
	DB *sql.DB
}

func (service *StoryService) NewStory(ctx context.Context, req *protoStory.NewStoryRequest, res *protoStory.StoryResponse) error {
	return nil
}
