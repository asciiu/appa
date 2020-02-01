package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asciiu/appa/common/db"
	"github.com/asciiu/appa/order-service/models"
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

func NewOrderService(name, dbUrl string) micro.Service {
	// Create a new service. Include some options here.
	srv := k8s.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name(name),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	appaDB, err := db.NewDB(dbUrl)

	if err != nil {
		log.Fatalf(err.Error())
	}

	service := OrderService{
		DB:         appaDB,
		OrderBooks: make(map[string]*models.OrderBook),
	}
	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	protoOrder.RegisterOrderServiceHandler(srv.Server(), &service)

	return srv
}

func main() {
	dbUrl := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	srv := NewOrderService("appa.orders", dbUrl)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
