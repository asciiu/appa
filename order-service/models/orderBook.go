package models

import (
	constOrder "github.com/asciiu/appa/order-service/constants"
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

type OrderBook struct {
	MarketName string
	BuyQ       map[float64]*OrderQueue
	SellQ      map[float64]*OrderQueue
}

func NewOrderBook(marketName string) *OrderBook {
	return &OrderBook{
		MarketName: marketName,
		BuyQ:       make(map[float64]*OrderQueue, 0),
		SellQ:      make(map[float64]*OrderQueue, 0),
	}
}

func (book *OrderBook) AddOrder(order *protoOrder.Order) {
	switch {
	case order.MarketName != book.MarketName:
		return
	case order.Side == constOrder.Buy:
		if queue, ok := book.BuyQ[order.Price]; ok {
			queue.AddOrder(order)
		} else {
			queue = NewOrderQueue(order.Price)
			queue.AddOrder(order)
			book.BuyQ[order.Price] = queue
		}
	case order.Side == constOrder.Sell:
		if queue, ok := book.SellQ[order.Price]; ok {
			queue.AddOrder(order)
		} else {
			queue = NewOrderQueue(order.Price)
			queue.AddOrder(order)
			book.BuyQ[order.Price] = queue
		}
	}
}

func (book *OrderBook) CancelOrder(order *protoOrder.Order) {
	switch {
	case order.MarketName != book.MarketName:
		return
	case order.Side == constOrder.Buy:
		if queue, ok := book.BuyQ[order.Price]; ok {
			queue.RemoveOrder(order.OrderID)
		}
	case order.Side == constOrder.Sell:
		if queue, ok := book.SellQ[order.Price]; ok {
			queue.RemoveOrder(order.OrderID)
		}
	}
}
