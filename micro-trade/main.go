package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asciiu/appa/lib/db"
	proto "github.com/asciiu/appa/trade-engine/proto/trade"
	micro "github.com/micro/go-micro/v2"
)

func main() {
	dbURL := fmt.Sprintf("%s", os.Getenv("DB_URL"))

	// Create a new service. Include some options here.
	service := micro.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("trade-engine"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	service.Init()

	appaDB, err := db.NewDB(dbURL)

	if err != nil {
		log.Fatalf(err.Error())
	}

	engine := NewTradeEngine(appaDB)
	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	proto.RegisterTradeEngineHandler(service.Server(), new(engine))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
