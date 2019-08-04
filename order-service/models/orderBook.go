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

// func (book *OrderBook) FillOrders(order *protoOrder.Order) (filledOrders []*protoOrder.Order) {
// 	filledOrders = make([]*protoOrder.Order, 0)
// 	switch {
// 	case order.Side == constOrder.Buy:
// 		filledOrders = book.FillSellOrders(order)
// 	case order.Side == constOrder.Sell:
// 		filledOrders = book.FillBuyOrders(order)
// 	}
// 	return
// }

// get all buy orders that can fill a sell order
func (book *OrderBook) ProcessSellOrder(sellOrder *protoOrder.Order) (buyOrders []*protoOrder.Order) {
	if sellOrder.Side != constOrder.Sell {
		return
	}
	if sellOrder.Amount <= 0 {
		// sell order amount should always be > 0
		return
	}

	filledOrders := make([]*protoOrder.Order, 0)
	sellAmount := sellOrder.Amount

	// buy orders are presorted by price from low to high
	// start with buy orders with the highest price
	for i := len(book.BuyOrders) - 1; i >= 0; i-- {
		buyOrder := book.BuyOrders[i]

		// lower buy prices should be be filled by a higher price
		if buyOrder.Price < sellOrder.Price {
			break
		}

		// fill buy orders with a higher price than the sell
		// the sell price is at the seller's price since it is lower
		//trade := Trade{
		//	MakerOrderID: buyOrder.OrderID,
		//	TakerOrderID: sellOrder.OrderID,
		//	Price:        sellOrder.Price,
		//	Side:         constants.Sell,
		//}

		// if the sell amount is less than the buy order amount
		if sellAmount >= buyOrder.Amount {
			// fill the entire buy order
			buyOrder.Fill = buyOrder.Amount
			// subtract the bought amount from the running sellAmount
			sellOrder.Amount -= buyOrder.Amount
			// remove filled orders
			book.BuyOrders = book.BuyOrders[:i]
		} else {
			// Trade{
			// 	MakerOrderID: buyOrder.OrderID,
			// 	TakerOrderID: sellOrder.OrderID,
			// 	Amount:       sellAmount,
			// 	Price:        sellOrder.Price,
			// 	Side:         constants.Sell,
			// }

			// TODO update buy order fill column
			// set order status to filled
			buyOrder.Fill = sellAmount

			// all sold out
			sellOrder.Amount = 0
			break
		}
		filledOrders = append(filledOrders, buyOrder)
	}

	// TODO
	// update filled orders

	// add the remaining sell order amount
	if sellOrder.Amount > 0 {
		book.AddSellOrder(sellOrder)
	}

	return
}

// get all sell orders that a buy order can fill
func (book *OrderBook) ProcessBuyOrder(buyOrder *protoOrder.Order) (sellOrders []*protoOrder.Order) {
	if buyOrder.Side != constOrder.Buy {
		return
	}

	sellOrders = make([]*protoOrder.Order, 0)
	buySize := buyOrder.Amount
	for i, sell := range book.SellOrders {
		if sell.Price <= buyOrder.Price && buySize > 0 {
			sellSize := sell.Amount

			if buySize < sellSize {
				sell.Fill = buySize
				sell.Amount -= buySize
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

func (book *OrderBook) CancelOrder(order *protoOrder.Order) {
	switch {
	case order.Side == constOrder.Buy:
		book.CancelBuyOrder(order)
	case order.Side == constOrder.Sell:
		book.CancelSellOrder(order)
	}
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
