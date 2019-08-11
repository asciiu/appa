package main

import (
	"fmt"
	"log"

	"github.com/asciiu/appa/ether-service/examples"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//dbURL := fmt.Sprintf("%s", os.Getenv("DB_URL"))

	//// Create a new service. Include some options here.
	//srv := k8s.NewService(
	//	// This name must match the package name given in your protobuf definition
	//	micro.Name("trade-engine"),
	//	micro.Version("latest"),
	//)

	//// Init will parse the command line flags.
	//srv.Init()

	//appaDB, err := db.NewDB(dbURL)

	//if err != nil {
	//	log.Fatalf(err.Error())
	//}

	// Register our service with the gRPC server, this will tie our
	// implementation into the auto-generated interface code for our
	// protobuf definition.
	//trade.RegisterTradeEngineHandler(srv.Server(), engine)

	//if err := srv.Run(); err != nil {
	//	log.Fatal(err)
	//}

	// connects to local ganache-cli
	//client, err := ethclient.Dial("http://localhost:8545")
	//if err != nil {
	//	log.Fatal(err)
	//}

	client, err := ethclient.Dial("https://mainnet.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("connected!")

	examples.Balance(client)
}
