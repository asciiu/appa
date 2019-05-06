package main

import (
	"context"
	"log"
	"os/exec"
	"testing"

	repoUser "github.com/asciiu/appa/api-graphql/db/sql"
	user "github.com/asciiu/appa/api-graphql/models"
	"github.com/asciiu/appa/common/db"
	protoStory "github.com/asciiu/appa/story-service/proto/story"
	"github.com/stretchr/testify/assert"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func setupService() (*StoryService, *user.User) {
	dbUrl := "postgres://postgres@localhost:5432/appa_test?&sslmode=disable"
	db, _ := db.NewDB(dbUrl)

	storyService := StoryService{
		DB: db,
	}

	user := user.NewUser("chester", "test@email", "hash")
	err := repoUser.InsertUser(db, user)
	checkErr(err)

	return &storyService, user
}

func TestNewRepo(t *testing.T) {
	service, user := setupService()

	defer service.DB.Close()

	req := protoStory.NewStoryRequest{
		UserID:  user.ID,
		Title:   "test story",
		Content: "he said something",
		Rated:   "everyone",
		Status:  "draft",
	}

	res := protoStory.StoryResponse{}
	service.NewStory(context.Background(), &req, &res)
	assert.Equal(t, "success", res.Status, "expected success got: "+res.Message)

	cmd := exec.Command("rm", "-rf", res.Data.Story.UserID)
	err := cmd.Run()
	assert.Nil(t, err, "error for delete should be nil")

	repoUser.DeleteUserHard(service.DB, user.ID)
}

func TestDeleteRepo(t *testing.T) {
	service, user := setupService()

	defer service.DB.Close()

	title := "testing 123"
	req1 := protoStory.NewStoryRequest{
		UserID:  user.ID,
		Title:   title,
		Content: "he said something",
		Rated:   "everyone",
		Status:  "draft",
	}

	res1 := protoStory.StoryResponse{}
	service.NewStory(context.Background(), &req1, &res1)
	assert.Equal(t, "success", res1.Status, "expected success got: "+res1.Message)

	req2 := protoStory.DeleteStoryRequest{
		UserID: user.ID,
		Title:  title,
	}
	res2 := protoStory.StoryResponse{}
	service.DeleteStory(context.Background(), &req2, &res2)
	assert.Equal(t, "success", res2.Status, "expected success got: "+res2.Message)

	cmd := exec.Command("rm", "-rf", res1.Data.Story.UserID)
	err := cmd.Run()
	assert.Nil(t, err, "error for delete should be nil")

	repoUser.DeleteUserHard(service.DB, user.ID)
}
