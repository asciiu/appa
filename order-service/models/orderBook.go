package models

import (
	constOrder "github.com/asciiu/appa/order-service/constants"
)

type OrderBook struct {
	MarketName string
	//LastPrice uint64
	//Side string
	BuyOrders  []*Order
	SellOrders []*Order
}

func NewOrderBook(marketName string) *OrderBook {
	return &OrderBook{
		MarketName: marketName,
		BuyOrders:  make([]*Order, 0),
		SellOrders: make([]*Order, 0),
	}
}

func (book *OrderBook) AddOrder(order *Order) {
	if order.MarketName != book.MarketName {
		return
	}

	switch {
	case order.Side == constOrder.Buy:
		book.addBuyOrder(order)
	case order.Side == constOrder.Sell:
		book.addSellOrder(order)
	}
}

// buy orders will be kept sorted in acending price order
// the last order should be the highest priced order
func (book *OrderBook) addBuyOrder(order *Order) {
	n := len(book.BuyOrders)

	if n == 0 {
		book.BuyOrders = append(book.BuyOrders, order)
		return
	}

	var i int
	for i := n - 1; i >= 0; i-- {
		buyOrder := book.BuyOrders[i]
		if buyOrder.Price < order.Price {
			break
		}
	}
	if i == n-1 {
		book.BuyOrders = append(book.BuyOrders, order)
	} else {
		copy(book.BuyOrders[i+1:], book.BuyOrders[i:])
		book.BuyOrders[i] = order
	}
}

// sell orders will be kept sorted in descending price order
// the last order should the lowest priced order
func (book *OrderBook) addSellOrder(order *Order) {
	n := len(book.SellOrders)
	var i int
	for i := n - 1; i >= 0; i-- {
		sellOrder := book.SellOrders[i]
		if sellOrder.Price > order.Price {
			break
		}
	}
	if i == n-1 {
		book.SellOrders = append(book.SellOrders, order)
	} else {
		copy(book.SellOrders[i+1:], book.SellOrders[i:])
		book.SellOrders[i] = order
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
// func (book *OrderBook) ProcessSellOrder(sellOrder *protoOrder.Order) (buyOrders []*protoOrder.Order) {
// 	if sellOrder.Side != constOrder.Sell {
// 		return
// 	}
// 	if sellOrder.Amount <= 0 {
// 		// sell order amount should always be > 0
// 		return
// 	}

// 	filledOrders := make([]*protoOrder.Order, 0)
// 	sellAmount := sellOrder.Amount

// 	// buy orders are presorted by price from low to high
// 	// start with buy orders with the highest price
// 	for i := len(book.BuyOrders) - 1; i >= 0; i-- {
// 		buyOrder := book.BuyOrders[i]

// 		// lower buy prices should be be filled by a higher price
// 		if buyOrder.Price < sellOrder.Price {
// 			break
// 		}

// 		// fill buy orders with a higher price than the sell
// 		// the sell price is at the seller's price since it is lower
// 		//trade := Trade{
// 		//	MakerOrderID: buyOrder.OrderID,
// 		//	TakerOrderID: sellOrder.OrderID,
// 		//	Price:        sellOrder.Price,
// 		//	Side:         constants.Sell,
// 		//}

// 		// if the sell amount is less than the buy order amount
// 		// the entire sell order is filled
// 		if sellAmount >= buyOrder.Amount {
// 			// fill the entire buy order
// 			buyOrder.Fill = buyOrder.Amount
// 			// subtract the bought amount from the running sellAmount
// 			sellOrder.Amount -= buyOrder.Amount
// 			// remove filled orders
// 			book.BuyOrders = book.BuyOrders[:i]
// 		} else {
// 			// Trade{
// 			// 	MakerOrderID: buyOrder.OrderID,
// 			// 	TakerOrderID: sellOrder.OrderID,
// 			// 	Amount:       sellAmount,
// 			// 	Price:        sellOrder.Price,
// 			// 	Side:         constants.Sell,
// 			// }

// 			// TODO update buy order fill column
// 			// set order status to filled
// 			buyOrder.Fill = sellAmount

// 			// all sold out
// 			sellOrder.Amount = 0
// 			break
// 		}
// 		filledOrders = append(filledOrders, buyOrder)
// 	}

// 	// TODO
// 	// update filled orders

// 	// add the remaining sell order amount
// 	if sellOrder.Amount > 0 {
// 		book.AddSellOrder(sellOrder)
// 	}

// 	return
// }

// // get all sell orders that a buy order can fill
// func (book *OrderBook) ProcessBuyOrder(buyOrder *protoOrder.Order) (sellOrders []*protoOrder.Order) {
// 	if buyOrder.Side != constOrder.Buy {
// 		return
// 	}

// 	sellOrders = make([]*protoOrder.Order, 0)
// 	buySize := buyOrder.Amount
// 	for i, sell := range book.SellOrders {
// 		if sell.Price <= buyOrder.Price && buySize > 0 {
// 			sellSize := sell.Amount

// 			if buySize < sellSize {
// 				sell.Fill = buySize
// 				sell.Amount -= buySize
// 				buySize = 0
// 			} else {
// 				sell.Fill = sellSize
// 				buySize -= sellSize
// 				book.SellOrders = book.SellOrders[i+1:]
// 			}

// 			sellOrders = append(sellOrders, sell)
// 		}
// 	}

// 	return
// }

// func (book *OrderBook) FindBuyOrder(buyOrder *protoOrder.Order) (index int) {
// 	index = -1
// 	if len(book.BuyOrders) > 0 {
// 		index = FindOrder(book.BuyOrders, buyOrder)
// 	}
// 	return
// }

// func (book *OrderBook) FindSellOrder(sellOrder *protoOrder.Order) (index int) {
// 	index = -1
// 	if len(book.SellOrders) > 0 {
// 		index = FindOrder(book.SellOrders, sellOrder)
// 	}
// 	return
// }

// func (book *OrderBook) CancelOrder(order *protoOrder.Order) {
// 	switch {
// 	case order.Side == constOrder.Buy:
// 		book.CancelBuyOrder(order)
// 	case order.Side == constOrder.Sell:
// 		book.CancelSellOrder(order)
// 	}
// }

// func (book *OrderBook) CancelBuyOrder(buyOrder *protoOrder.Order) {
// 	if i := book.FindBuyOrder(buyOrder); i >= 0 {
// 		book.BuyOrders = append(book.BuyOrders[:i], book.BuyOrders[i+1:]...)
// 	}
// }

// func (book *OrderBook) CancelSellOrder(sellOrder *protoOrder.Order) {
// 	if i := book.FindSellOrder(sellOrder); i >= 0 {
// 		book.BuyOrders = append(book.SellOrders[:i], book.SellOrders[i+1:]...)
// 	}
// }
