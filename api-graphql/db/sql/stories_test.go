package sql_test

import (
	"testing"

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
