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

// get all buy orders that can fill a sell order
func (book *OrderBook) FillBuyOrders(sellOrder *protoOrder.Order) (buyOrders []*protoOrder.Order) {
	if sellOrder.Side != constOrder.Sell {
		return
	}

	buyOrders = make([]*protoOrder.Order, 0)
	sellSize := sellOrder.Size
	for i := len(book.BuyOrders) - 1; i >= 0; i-- {
		buy := book.BuyOrders[i]
		if buy.Price >= sellOrder.Price && sellSize > 0 {
			buySize := buy.Size

			if sellSize < buySize {
				buy.Fill = sellSize
				buy.Size -= sellSize
				sellSize = 0
			} else {
				buy.Fill = buySize
				sellSize -= buySize
				// remove filled orders
				book.BuyOrders = book.BuyOrders[:i]
			}

			buyOrders = append(buyOrders, buy)
		}
	}

	return
}

// get all sell orders that a buy order can fill
func (book *OrderBook) FillSellOrders(buyOrder *protoOrder.Order) (sellOrders []*protoOrder.Order) {
	if buyOrder.Side != constOrder.Buy {
		return
	}

	sellOrders = make([]*protoOrder.Order, 0)
	buySize := buyOrder.Size
	for i, sell := range book.SellOrders {
		if sell.Price <= buyOrder.Price && buySize > 0 {
			sellSize := sell.Size

			if buySize < sellSize {
				sell.Fill = buySize
				sell.Size -= buySize
				buySize = 0
			} else {
				sell.Fill = sellSize
				buySize -= sellSize
				book.SellOrders = book.SellOrders[i+1:]
			}

			sellOrders = append(sellOrders, sell)
		}
	}

	return
}

func (book *OrderBook) FindBuyOrder(buyOrder *protoOrder.Order) (index int) {
	index = -1
	if len(book.BuyOrders) > 0 {
		index = FindOrder(book.BuyOrders, buyOrder)
	}
	return
}

func (book *OrderBook) FindSellOrder(sellOrder *protoOrder.Order) (index int) {
	index = -1
	if len(book.SellOrders) > 0 {
		index = FindOrder(book.SellOrders, sellOrder)
	}
	return
}

func (book *OrderBook) CancelBuyOrder(buyOrder *protoOrder.Order) {
	if i := book.FindBuyOrder(buyOrder); i >= 0 {
		book.BuyOrders = append(book.BuyOrders[:i], book.BuyOrders[i+1:]...)
	}
}

func (book *OrderBook) CancelSellOrder(sellOrder *protoOrder.Order) {
	if i := book.FindSellOrder(sellOrder); i >= 0 {
		book.BuyOrders = append(book.SellOrders[:i], book.SellOrders[i+1:]...)
	}
}
