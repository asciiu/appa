package main

import (
	"log"

	msg "github.com/asciiu/gomo/common/constants/messages"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"
)

func main() {
	srv := k8s.NewService(
		micro.Name("fomo.binance"),
	)

	srv.Init()

	verifiedPub := micro.NewPublisher(msg.TopicKeyVerified, srv.Client())
	balancePub := micro.NewPublisher(msg.TopicBalanceUpdate, srv.Client())
	filledPub := micro.NewPublisher(msg.TopicOrderFilled, srv.Client())
	keyValidator := KeyValidator{
		KeyVerifiedPub: verifiedPub,
		BalancePub:     balancePub,
	}
	// subscribe to new key topic with a key validator
	micro.RegisterSubscriber(msg.TopicNewKey, srv.Server(), &keyValidator)

	orderFiller := OrderFiller{
		FilledPub: filledPub,
		FailedPub: micro.NewPublisher(msg.TopicNotification, srv.Client()),
	}
	// subscribe to new key topic with a key validator
	micro.RegisterSubscriber(msg.TopicFillOrder, srv.Server(), &orderFiller)

	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
