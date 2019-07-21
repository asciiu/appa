package sql_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/asciiu/appa/common/db"
	"github.com/asciiu/appa/user-service/db/sql"
	"github.com/asciiu/appa/user-service/models"
	"github.com/stretchr/testify/assert"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func TestInsertUser(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := models.NewUser("jonni5", "test@email", "password")
	newUser, err := sql.InsertUser(db, user)
	assert.Nil(t, err, fmt.Sprintf("could not insert user %s", err))

	if err != nil {
		t.Errorf("%s", err)
	}
	foundUser, err := sql.FindUserByID(db, user.ID)
	assert.Nil(t, err, fmt.Sprintf("could not find user by id %s", err))

	assert.Equal(t, newUser.ID, foundUser.ID, "user ids do not match")
	assert.Equal(t, newUser.Username, foundUser.Username, "usernames do not match")
	assert.Equal(t, newUser.Email, foundUser.Email, "emails do not match")

	err = sql.DeleteUserHard(db, user.ID)
	assert.Nil(t, err, fmt.Sprintf("could not delete user %s", err))
}

// func TestFindUser(t *testing.T) {
// 	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
// 	checkErr(err)
// 	defer db.Close()

// 	email := "test@email"
// 	user, err := sql.FindUser(db, email)
// 	if err != nil {
// 		t.Errorf("%s", err)
// 	}
// 	if user == nil {
// 		t.Errorf("user not found %s", email)
// 	}

// 	sqlStatement := `delete from users where email = $1`
// 	db.Exec(sqlStatement, email)
// }
