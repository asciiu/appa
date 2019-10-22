package sql_test

import (
	"log"
	"strings"
	"testing"

	"github.com/asciiu/appa/api-graphql/db/sql"
	"github.com/asciiu/appa/api-graphql/models"
	"github.com/asciiu/appa/lib/db"
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

	user := models.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = sql.InsertUser(db, user)
	assert.Nil(t, err, "insert should have nil error")

	sql.DeleteUserHard(db, user.ID)
}

func TestInsertUserDupeUsername(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	user := models.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = sql.InsertUser(db, user)
	err = sql.InsertUser(db, user)

	if !strings.Contains(err.Error(), "duplicate key") {
		t.Errorf("should have failed on duplicate key")
	}
	sql.DeleteUserHard(db, user.ID)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	newUser := models.NewUser("flowtester", "test@email", "Yo yo yo!!")
	sql.InsertUser(db, newUser)

	email := "test@email"
	foundUser, err := sql.FindUserByEmail(db, email)
	if err != nil {
		t.Errorf("%s", err)
	}

	assert.Equal(t, newUser.ID, foundUser.ID, "ids do not match")
	assert.Equal(t, "test@email", foundUser.Email, "email does not match")
	assert.Equal(t, false, foundUser.EmailVerified, "a new user will default email verified to false")
	assert.Equal(t, "flowtester", foundUser.Username, "email does not match")

	sql.DeleteUserHard(db, newUser.ID)
}

func TestFindUserByID(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	newUser := models.NewUser("flowtester", "test@email", "Yo yo yo!!")
	sql.InsertUser(db, newUser)

	email := "test@email"
	foundUser, err := sql.FindUserByEmail(db, email)
	if err != nil {
		t.Errorf("%s", err)
	}

	assert.Equal(t, newUser.ID, foundUser.ID, "ids do not match")
	assert.Equal(t, email, foundUser.Email, "email does not match")
	assert.Equal(t, false, foundUser.EmailVerified, "a new user will default email verified to false")
	assert.Equal(t, "flowtester", foundUser.Username, "email does not match")

	sql.DeleteUserHard(db, newUser.ID)
}

func TestFindUserID(t *testing.T) {
	db, err := db.NewDB("postgres://postgres@localhost/appa_test?&sslmode=disable")
	checkErr(err)
	defer db.Close()

	newUser := models.NewUser("flowtester", "test@email", "Yo yo yo!!")
	sql.InsertUser(db, newUser)

	foundUser, err := sql.FindUserByID(db, newUser.ID)
	if err != nil {
		t.Errorf("%s", err)
	}

	assert.Equal(t, newUser.ID, foundUser.ID, "ids do not match")
	assert.Equal(t, "test@email", foundUser.Email, "email does not match")
	assert.Equal(t, false, foundUser.EmailVerified, "a new user will default email verified to false")
	assert.Equal(t, "flowtester", foundUser.Username, "email does not match")

	sql.DeleteUserHard(db, newUser.ID)
}
