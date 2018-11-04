package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asciiu/oldiez/common/db"
	protoUser "github.com/asciiu/oldiez/user-service/proto/user"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

func NewUserService(name, dbUrl string) micro.Service {
	// Create a new service. Include some options here.
	srv := k8s.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name(name),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	oldiezDB, err := db.NewDB(dbUrl)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	protoUser.RegisterUserServiceHandler(srv.Server(), &UserService{oldiezDB})

	return srv
}

func main() {
	dbUrl := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	srv := NewUserService("oldiez.users", dbUrl)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
