package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asciiu/appa/common/db"
	protoStory "github.com/asciiu/appa/story-service/proto/story"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dbURL := fmt.Sprintf("%s", os.Getenv("DB_URL"))

	srv := k8s.NewService(
		micro.Name("stories"),
		micro.Version("latest"),
	)
	srv.Init()

	appaDB, err := db.NewDB(dbURL)
	check(err)

	service := StoryService{
		DB: appaDB,
	}
	protoStory.RegisterStoryServiceHandler(srv.Server(), &service)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
