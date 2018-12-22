package models

import (
	constOrder "github.com/asciiu/appa/order-service/constants"
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

type OrderBook struct {
	MarketName string
	Buys       []*OrderQueue
	Sells      []*OrderQueue
}

func NewOrderBook(marketName string) *OrderBook {
	return &OrderBook{
		MarketName: marketName,
		Buys:       make([]*OrderQueue, 0),
		Sells:      make([]*OrderQueue, 0),
	}
}

func (book *OrderBook) AddOrder(order *protoOrder.Order) {
	switch {
	case order.MarketName != book.MarketName:
		return
	case order.Side == constOrder.Buy:
	case order.Side == constOrder.Sell:
	}
}

func (book *OrderBook) AddBuyOrder(order *protoOrder.Order) {

	for _, queue := range book.Buys {
		if queue.Price == order.Price {
			queue.AddOrder(order)
			break
		}
	}

}

func (book *OrderBook) AddSellOrder(order *protoOrder.Order) {

}

func (book *OrderBook) CancelOrder(order *protoOrder.Order) {
	switch {
	case order.MarketName != book.MarketName:
		return
	case order.Side == constOrder.Buy:
	case order.Side == constOrder.Sell:
	}
}
