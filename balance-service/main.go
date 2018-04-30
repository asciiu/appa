package main

import (
	"fmt"
	"log"
	"os"

	bp "github.com/asciiu/gomo/balance-service/proto/balance"
	"github.com/asciiu/gomo/common/db"
	msg "github.com/asciiu/gomo/common/messages"
	micro "github.com/micro/go-micro"
)

func main() {
	srv := micro.NewService(
		micro.Name("go.micro.srv.balance"),
	)

	srv.Init()

	dbUrl := fmt.Sprintf("%s", os.Getenv("DB_URL"))
	gomoDB, err := db.NewDB(dbUrl)

	if err != nil {
		log.Fatalf(err.Error())
	}

	balanceService := BalanceService{gomoDB}
	balanceListener := BalanceUpdateListener{gomoDB, srv}
	bp.RegisterBalanceServiceHandler(srv.Server(), &balanceService)

	// subscribe to new key topic with a key validator
	micro.RegisterSubscriber(msg.TopicBalanceUpdate, srv.Server(), &balanceListener)
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
