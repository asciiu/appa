package sql_test

import (
	"fmt"
	"testing"

	"github.com/asciiu/appa/api-graphql/constants"
	"github.com/asciiu/appa/api-graphql/db/sql"
	"github.com/asciiu/appa/api-graphql/models"
	"github.com/asciiu/appa/common/db"
	"github.com/stretchr/testify/assert"
)

func TestInsertStory(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := models.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = sql.InsertUser(db, user)
	assert.Nil(t, err, "insert new user failed")

	story := models.NewStory(user.ID, "test", "this is only a test")
	err = sql.InsertStory(db, story)

	assert.Nil(t, err, "insert story failed")

	sql.DeleteUserHard(db, user.ID)
}

func TestListStories(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := models.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = sql.InsertUser(db, user)
	assert.Nil(t, err, "insert new user failed")

	story1 := models.NewStory(user.ID, "one", "this is only a test")
	err = sql.InsertStory(db, story1)
	assert.Nil(t, err, "insert story failed")

	story2 := models.NewStory(user.ID, "two", "this is only a test")
	err = sql.InsertStory(db, story2)
	assert.Nil(t, err, "insert story failed")

	story3 := models.NewStory(user.ID, "three", "this is only a test")
	err = sql.InsertStory(db, story3)
	assert.Nil(t, err, "insert story failed")

	story4 := models.NewStory(user.ID, "four", "this is only a test")
	err = sql.InsertStory(db, story4)
	assert.Nil(t, err, "insert story failed")

	stories, err := sql.StoryTitles(db, constants.Unpublished, 0, 10)

	assert.Nil(t, err, "insert story failed")

	fmt.Println(stories)

	sql.DeleteUserHard(db, user.ID)
}
