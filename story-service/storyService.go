package main

import (
	"context"
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"time"

	commonResp "github.com/asciiu/appa/common/constants/response"
	protoStory "github.com/asciiu/appa/story-service/proto/story"
	"gopkg.in/libgit2/git2go.v27"
)

type StoryService struct {
	DB *sql.DB
}

// InitStory - Init new story
func (service *StoryService) InitStory(ctx context.Context, req *protoStory.InitStoryRequest, res *protoStory.StoryResponse) error {
	path := fmt.Sprintf("%s/%s", "database", req.Title)
	repo, err := git.InitRepository(path, false)
	if err != nil {
		msg := fmt.Sprintf("init repo error for %s: %s", req.Title, err)
		log.Println(msg)
		res.Status = commonResp.Fail
		return nil
	}

	filePath := fmt.Sprintf("%s/%s.txt", path, req.Title)

	data := []byte(req.Content)
	err = ioutil.WriteFile(filePath, data, 0644)
	if err != nil {
		msg := fmt.Sprintf("write error for %s.txt: %s", req.Title, err)
		log.Println(msg)
		res.Status = commonResp.Fail
		return nil
	}

	index, err := repo.Index()
	if err != nil {
		msg := fmt.Sprintf("could not obtain repo index for %s: %s", req.Title, err)
		log.Println(msg)
		res.Status = commonResp.Fail
		return nil
	}

	index.AddByPath(filePath)
	index.Write()

	sig := &git.Signature{
		Name:  req.Username,
		Email: req.UserEmail,
		When:  time.Now(),
	}

	treeID, err := index.WriteTreeTo(repo)
	if err != nil {
		log.Println("error on write tree: ", err)
		res.Status = commonResp.Fail
		return nil
	}

	tree, err := repo.LookupTree(treeID)
	if err != nil {
		log.Println("error on lookup tree: ", err)
		res.Status = commonResp.Fail
		return nil
	}

	_, err = repo.CreateCommit("HEAD", sig, sig, "Initial commit.", tree)
	if err != nil {
		log.Println("error on lookup tree: ", err)
		res.Status = commonResp.Fail
		return nil
	} else {
		res.Status = commonResp.Success
		res.Data = &protoStory.StoryData{
			Story: &protoStory.Story{
				StoryID: filePath,
				UserID:  req.UserID,
				Title:   req.Title,
				Content: req.Content,
				Rated:   req.Rated,
				Status:  req.Status,
			},
		}
	}

	return nil
}

func (service *StoryService) DeleteStory(ctx context.Context, req *protoStory.DeleteStoryRequest, res *protoStory.StoryResponse) error {
	path := fmt.Sprintf("%s/%s", req.UserID, req.Title)

	cmd := exec.Command("rm", "-rf", path)
	err := cmd.Run()
	if err != nil {
		res.Status = commonResp.Error
		res.Message = err.Error()
	} else {
		res.Status = commonResp.Success
		res.Data = &protoStory.StoryData{
			Story: &protoStory.Story{
				StoryID: path,
				UserID:  req.UserID,
			},
		}
	}

	return nil
}
