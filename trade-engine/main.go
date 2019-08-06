package main

import (
	"fmt"
	"log"
	"os"

	"github.com/asciiu/appa/common/db"
	"github.com/asciiu/appa/trade-engine/models"
	"github.com/asciiu/appa/trade-engine/proto/trade"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

func main() {
	dbURL := fmt.Sprintf("%s", os.Getenv("DB_URL"))

	// Create a new service. Include some options here.
	srv := k8s.NewService(
		// This name must match the package name given in your protobuf definition
		micro.Name("trade-engine"),
		micro.Version("latest"),
	)

	// Init will parse the command line flags.
	srv.Init()

	appaDB, err := db.NewDB(dbURL)

	if err != nil {
		log.Fatalf(err.Error())
	}

	engine := TradeEngine{
		DB:         appaDB,
		OrderBooks: make(map[string]*models.OrderBook),
	}
	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	trade.RegisterTradeEngineHandler(srv.Server(), &engine)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
