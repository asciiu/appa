package models

import (
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

type OrderBook struct {
	MarketName string
	Side       string
	Orders     []*protoOrder.Order
}

func NewOrderBook(marketName, side string) *OrderBook {
	return &OrderBook{
		MarketName: marketName,
		Orders:     make([]*protoOrder.Order, 0),
		Side:       side,
	}
}
