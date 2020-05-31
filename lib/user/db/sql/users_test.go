package sql_test

import (
	"strings"
	"testing"

	db "github.com/asciiu/appa/lib/db/sql"
	"github.com/asciiu/appa/lib/user/db/sql"
	userRepo "github.com/asciiu/appa/lib/user/db/sql"
	user "github.com/asciiu/appa/lib/user/models"
	"github.com/asciiu/appa/lib/util"
	"github.com/stretchr/testify/assert"
)

var testDB = "postgres://postgres@localhost/appa_test?&sslmode=disable"

func TestInsertUser(t *testing.T) {
	db, err := db.NewDB(testDB)
	util.CheckErr(err)
	defer db.Close()

	newUser := user.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = userRepo.InsertUser(db, newUser)
	assert.Nil(t, err, "insert should have nil error")

	userRepo.DeleteUserHard(db, newUser.ID)
}

func TestInsertUserDupeUsername(t *testing.T) {
	db, err := db.NewDB(testDB)
	util.CheckErr(err)
	defer db.Close()

	newUser := user.NewUser("flowtester", "test@email", "Yo yo yo!!")
	err = userRepo.InsertUser(db, newUser)
	err = userRepo.InsertUser(db, newUser)

	if !strings.Contains(err.Error(), "duplicate key") {
		t.Errorf("should have failed on duplicate key")
	}
	userRepo.DeleteUserHard(db, newUser.ID)
}

func TestFindUserByEmail(t *testing.T) {
	db, err := db.NewDB(testDB)
	util.CheckErr(err)
	defer db.Close()

	newUser := user.NewUser("flowtester", "test@email", "Yo yo yo!!")
	sql.InsertUser(db, newUser)

	email := "test@email"
	foundUser, err := userRepo.FindUserByEmail(db, email)
	if err != nil {
		t.Errorf("%s", err)
	}

	assert.Equal(t, newUser.ID, foundUser.ID, "ids do not match")
	assert.Equal(t, "test@email", foundUser.Email, "email does not match")
	assert.Equal(t, false, foundUser.EmailVerified, "a new user will default email verified to false")
	assert.Equal(t, "flowtester", foundUser.Username, "email does not match")

	userRepo.DeleteUserHard(db, newUser.ID)
}

func TestFindUserByID(t *testing.T) {
	db, err := db.NewDB(testDB)
	util.CheckErr(err)
	defer db.Close()

	newUser := user.NewUser("flowtester", "test@email", "Yo yo yo!!")
	userRepo.InsertUser(db, newUser)

	email := "test@email"
	foundUser, err := sql.FindUserByEmail(db, email)
	if err != nil {
		t.Errorf("%s", err)
	}

	assert.Equal(t, newUser.ID, foundUser.ID, "ids do not match")
	assert.Equal(t, email, foundUser.Email, "email does not match")
	assert.Equal(t, false, foundUser.EmailVerified, "a new user will default email verified to false")
	assert.Equal(t, "flowtester", foundUser.Username, "email does not match")

	userRepo.DeleteUserHard(db, newUser.ID)
}

func TestFindUserID(t *testing.T) {
	db, err := db.NewDB(testDB)
	util.CheckErr(err)
	defer db.Close()

	newUser := user.NewUser("flowtester", "test@email", "Yo yo yo!!")
	userRepo.InsertUser(db, newUser)

	foundUser, err := userRepo.FindUserByID(db, newUser.ID)
	if err != nil {
		t.Errorf("%s", err)
	}

	assert.Equal(t, newUser.ID, foundUser.ID, "ids do not match")
	assert.Equal(t, "test@email", foundUser.Email, "email does not match")
	assert.Equal(t, false, foundUser.EmailVerified, "a new user will default email verified to false")
	assert.Equal(t, "flowtester", foundUser.Username, "email does not match")

	userRepo.DeleteUserHard(db, newUser.ID)
}
