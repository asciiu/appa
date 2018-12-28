package models

import (
	constOrder "github.com/asciiu/appa/order-service/constants"
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

type OrderBook struct {
	MarketName string
	BuyOrders  []*protoOrder.Order
	SellOrders []*protoOrder.Order
}

func NewOrderBook(marketName string) *OrderBook {
	return &OrderBook{
		MarketName: marketName,
		BuyOrders:  make([]*protoOrder.Order, 0),
		SellOrders: make([]*protoOrder.Order, 0),
	}
}

func (book *OrderBook) AddBuyOrder(order *protoOrder.Order) {
	if order.Side != constOrder.Buy || order.MarketName != book.MarketName {
		return
	}
	book.BuyOrders = append(book.BuyOrders, order)

	if len(book.BuyOrders) > 1 {
		book.BuyOrders = MergeSort(book.BuyOrders)
	}
}

func (book *OrderBook) MatchBuyOrders(sellOrder *protoOrder.Order) []int {
	if sellOrder.Side != constOrder.Sell {
		return nil
	}

	return MatchIndices(book.BuyOrders, sellOrder.Price, sellOrder.Size)
}

func (book *OrderBook) MatchSellOrders(buyOrder *protoOrder.Order) []int {
	if buyOrder.Side != constOrder.Buy {
		return nil
	}

	return MatchIndices(book.SellOrders, buyOrder.Price, buyOrder.Size)
}

//func (book *OrderBook) RemoveBuyOrders(to, from index) {
//
//}

func (book *OrderBook) AddSellOrder(order *protoOrder.Order) {
	if order.Side != constOrder.Sell || order.MarketName != book.MarketName {
		return
	}
	book.SellOrders = append(book.SellOrders, order)

	if len(book.SellOrders) > 1 {
		book.SellOrders = MergeSort(book.SellOrders)
	}
}
