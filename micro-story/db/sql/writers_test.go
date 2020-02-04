package sql_test

import (
	"fmt"
	"testing"

	repoUser "github.com/asciiu/appa/api-graphql/db/sql"
	user "github.com/asciiu/appa/api-graphql/models"
	"github.com/asciiu/appa/lib/db"
	util "github.com/asciiu/appa/lib/util"
	repoWriter "github.com/asciiu/appa/micro-story/db/sql"
	"github.com/stretchr/testify/assert"
)

func TestInsertWriter(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	util.CheckErr(err)
	defer db.Close()

	user := user.NewUser("tastytest", "test@email", "password")
	err = repoUser.InsertUser(db, user)
	assert.Nil(t, err, fmt.Sprintf("could not insert new user %s", err))

	err = repoWriter.InsertWriter(db, user.ID, "A Test of Will")
	assert.Equal(t, nil, err, "err should be nil")

	repoUser.DeleteUserHard(db, user.ID)
}
