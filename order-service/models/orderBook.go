package models

import (
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

type Order struct {
	OrderID string
	UserID  string
	Side    string
	Price   float64
	Size    float64
}

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
