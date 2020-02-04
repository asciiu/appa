package main

import (
	"context"
	"fmt"
	"os/exec"
	"testing"

	"github.com/asciiu/appa/lib/db"
	repoUser "github.com/asciiu/appa/lib/user/db/sql"
	user "github.com/asciiu/appa/lib/user/models"
	util "github.com/asciiu/appa/lib/util"
	protoStory "github.com/asciiu/appa/micro-story/proto/story"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var TEST_DIR = "trash"

func setupService() (*StoryService, *user.User) {
	dbURL := "postgres://postgres@localhost:5432/appa_test?&sslmode=disable"
	testDir := TEST_DIR
	db, _ := db.NewDB(dbURL)

	storyService := StoryService{
		DB:            db,
		DataDirectory: testDir,
	}

	user := user.NewUser("chester", "test@email", "hash")
	err := repoUser.InsertUser(db, user)
	util.CheckErr(err)

	return &storyService, user
}

func TestInitStory(t *testing.T) {
	service, user := setupService()
	title := "Sling Blade"

	defer service.DB.Close()

	storyID := uuid.New()

	req := protoStory.InitStoryRequest{
		StoryID:   storyID.String(),
		UserID:    user.ID,
		Username:  user.Username,
		UserEmail: user.Email,
		Title:     title,
		JsonData:  "Shouldn't have done that to my brother. He was just a boy.",
	}

	res := protoStory.StoryResponse{}
	service.InitStory(context.Background(), &req, &res)
	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)

	path := fmt.Sprintf("%s/%s", TEST_DIR, req.StoryID)
	cmd := exec.Command("rm", "-rf", path)
	err := cmd.Run()
	assert.Nil(t, err, "error for delete should be nil")

	repoUser.DeleteUserHard(service.DB, user.ID)
}

func TestDeleteRepo(t *testing.T) {
	service, user := setupService()

	defer service.DB.Close()

	storyID := uuid.New()
	title := "Forest Gump"
	req1 := protoStory.InitStoryRequest{
		StoryID:   storyID.String(),
		UserID:    user.ID,
		Username:  user.Username,
		UserEmail: user.Email,
		Title:     title,
		JsonData:  "Run forest!",
	}

	res1 := protoStory.StoryResponse{}
	service.InitStory(context.Background(), &req1, &res1)
	assert.Equal(t, "success", res1.Status, "expected success got: "+res1.Message)

	req2 := protoStory.DeleteStoryRequest{
		StoryID: res1.Data.Story.StoryID,
		UserID:  user.ID,
	}
	res2 := protoStory.StoryResponse{}
	service.DeleteStory(context.Background(), &req2, &res2)
	assert.Equal(t, "success", res2.Status, "expected success got: "+res2.Message)

	repoUser.DeleteUserHard(service.DB, user.ID)
}
