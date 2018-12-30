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

func (book *OrderBook) AddOrder(order *protoOrder.Order) {
	switch {
	case order.Side == constOrder.Buy:
		book.AddBuyOrder(order)
	case order.Side == constOrder.Sell:
		book.AddSellOrder(order)
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

func (book *OrderBook) AddSellOrder(order *protoOrder.Order) {
	if order.Side != constOrder.Sell || order.MarketName != book.MarketName {
		return
	}
	book.SellOrders = append(book.SellOrders, order)

	if len(book.SellOrders) > 1 {
		book.SellOrders = MergeSort(book.SellOrders)
	}
}

// func (book *OrderBook) MatchBuyOrders(sellOrder *protoOrder.Order) []int {
// 	if sellOrder.Side != constOrder.Sell {
// 		return nil
// 	}

// 	return MatchIndices(book.BuyOrders, sellOrder.Price, sellOrder.Size)
// }

// func (book *OrderBook) MatchSellOrders(buyOrder *protoOrder.Order) []int {
// 	if buyOrder.Side != constOrder.Buy {
// 		return nil
// 	}

// 	return MatchIndices(book.SellOrders, buyOrder.Price, buyOrder.Size)
// }

func (book *OrderBook) RemoveBuyOrders(from, to int, upToSize float64) {
	sum := upToSize
	for i := from; i < to; i++ {
		order := book.BuyOrders[i]
		sum -= order.Size
		if sum < 0 {
			order.Size -= sum
			break
		}
		book.BuyOrders = append(book.BuyOrders[:i], book.BuyOrders[i+1:]...)
	}
}

func (book *OrderBook) RemoveSellOrders(from, to int, upToSize float64) {
	sum := upToSize
	for i := from; i < to; i++ {
		order := book.SellOrders[i]
		sum -= order.Size
		if sum < 0 {
			order.Size -= sum
			break
		}
		book.SellOrders = append(book.SellOrders[:i], book.SellOrders[i+1:]...)
	}
}
