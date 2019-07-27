package sql_test

import (
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

	story := models.NewStory(user.ID, "test", "{\"some\":\"json\"}")
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

	story1 := models.NewStory(user.ID, "one", "{\"some\":\"json\"}")
	err = sql.InsertStory(db, story1)
	assert.Nil(t, err, "insert story failed")

	story2 := models.NewStory(user.ID, "two", "{\"some\":\"json\"}")
	err = sql.InsertStory(db, story2)
	assert.Nil(t, err, "insert story failed")

	story3 := models.NewStory(user.ID, "three", "{\"some\":\"json\"}")
	err = sql.InsertStory(db, story3)
	assert.Nil(t, err, "insert story failed")

	story4 := models.NewStory(user.ID, "four", "{\"some\":\"json\"}")
	err = sql.InsertStory(db, story4)
	assert.Nil(t, err, "insert story failed")

	page := uint32(0)
	pageSize := uint32(10)
	pagedStories, err := sql.StoryTitles(db, constants.Unpublished, page, pageSize)

	assert.Nil(t, err, "insert story failed")

	assert.Equal(t, page, pagedStories.Page, "page should be 0")
	assert.Equal(t, pageSize, pagedStories.PageSize, "page size should be 10")
	assert.Equal(t, 4, len(pagedStories.Stories), "should be 4 stories")

	sql.DeleteUserHard(db, user.ID)
}

func TestFindStoryByID(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := models.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = sql.InsertUser(db, user)
	assert.Nil(t, err, "insert new user failed")

	story1 := models.NewStory(user.ID, "one", "{\"some\":\"json\"}")
	err = sql.InsertStory(db, story1)
	assert.Nil(t, err, "insert story failed")

	story2 := models.NewStory(user.ID, "two", "{\"some\": \"json\"}")
	err = sql.InsertStory(db, story2)
	assert.Nil(t, err, "insert story failed")

	foundStory, err := sql.FindStoryByID(db, story2.ID)
	assert.Nil(t, err, "find story failed")

	assert.Equal(t, story2.ID, foundStory.ID, "id must match")
	assert.Equal(t, story2.Title, foundStory.Title, "titles do not match")
	assert.Equal(t, story2.Title, foundStory.Title, "titles do not match")
	assert.Equal(t, story2.Content, foundStory.Content, "content do not match")

	sql.DeleteUserHard(db, user.ID)
}
