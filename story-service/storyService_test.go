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
