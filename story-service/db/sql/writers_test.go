package sql_test

import (
	"fmt"
	"log"
	"testing"

	repoUser "github.com/asciiu/appa/api-graphql/db/sql"
	user "github.com/asciiu/appa/api-graphql/models"
	"github.com/asciiu/appa/common/db"
	repoAuthor "github.com/asciiu/appa/story-service/db/sql"
	"github.com/stretchr/testify/assert"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func TestInsertWriter(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := user.NewUser("tastytest", "test@email", "password")
	err = repoUser.InsertUser(db, user)
	assert.Nil(t, err, fmt.Sprintf("could not insert new user %s", err))

	err = repoAuthor.InsertWriter(db, user.ID, "A Test of Will")
	assert.Equal(t, nil, err, "err should be nil")

	repoUser.DeleteUserHard(db, user.ID)
}
