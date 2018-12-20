package models

import (
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

type OrderBook struct {
	MarketName string
	BuyOrders  []*protoOrder.Order
	SellOrders []*protoOrder.Order
}
