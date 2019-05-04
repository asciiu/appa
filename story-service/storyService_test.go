package main

import (
	"log"

	"github.com/asciiu/appa/common/db"
	repoUser "github.com/asciiu/appa/user-service/db/sql"
	user "github.com/asciiu/appa/user-service/models"
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

	user := user.NewUser("first", "last", "test@email", "hash")
	_, err := repoUser.InsertUser(db, user)
	checkErr(err)

	return &storyService, user
}
