package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asciiu/gomo/common/db"
	msg "github.com/asciiu/gomo/common/messages"
	kp "github.com/asciiu/gomo/key-service/proto/key"
	micro "github.com/micro/go-micro"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.srv.key-service"),
	)

	// Init will parse the command line flags.
	srv.Init()

	dbUrl := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	gomoDB, err := db.NewDB(dbUrl)
	if err != nil {
		log.Fatalf(err.Error())
	}

	publisher := micro.NewPublisher(msg.TopicNewKey, srv.Client())
	keyService := KeyService{gomoDB, publisher}

	kp.RegisterKeyServiceHandler(srv.Server(), &keyService)

	listener1 := KeyVerifiedListener{gomoDB}
	// handles key verified events
	micro.RegisterSubscriber(msg.TopicKeyVerified, srv.Server(), &listener1)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}