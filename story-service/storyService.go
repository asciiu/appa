package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	commonResp "github.com/asciiu/appa/common/constants/response"
	protoStory "github.com/asciiu/appa/story-service/proto/story"
	"gopkg.in/libgit2/git2go.v27"
)

type StoryService struct {
	DB *sql.DB
}

func (service *StoryService) NewStory(ctx context.Context, req *protoStory.NewStoryRequest, res *protoStory.StoryResponse) error {
	path := fmt.Sprintf("%s/%s", req.UserID, req.Title)
	_, err := git.InitRepository(path, false)
	if err != nil {
		log.Println("init repo error: ", err)
	}
	//fmt.Println(repo)

	res.Status = commonResp.Success
	res.Data = &protoStory.StoryData{
		Story: &protoStory.Story{
			StoryID: path,
			UserID:  req.UserID,
		},
	}

	return nil
}
