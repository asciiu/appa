package models

type OrderBook struct {
	MarketName string

	buys  []Stake
	sells []Stake
}

func NewOrderBook(eventName string) *OrderBook {
	return &OrderBook{
		MarketName: eventName,
		buys:       make([]Stake, 0),
		sells:      make([]Stake, 0),
	}
}

func (book *OrderBook) AddStake(stake Stake) {
	switch {
	case stake.Side == "buy":
		book.addBuyOrder(stake)
	case stake.Side == "sell":
		book.addSellOrder(stake)
	}
}

func (book *OrderBook) addBuyOrder(stake Stake) {
	if stake.Side != "buy" {
		return
	}
	book.buys = append(book.buys, stake)
	book.buys = MergeSort(book.buys)
}

func (book *OrderBook) addSellOrder(stake Stake) {
	if stake.Side != "sell" {
		return
	}
	book.sells = append(book.sells, stake)
	book.sells = MergeSort(book.sells)
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

// // get all buy orders that can fill a sell order
// func (book *OrderBook) FillBuyOrders(sellOrder *protoOrder.Order) (buyOrders []*protoOrder.Order) {
// 	if sellOrder.Side != constOrder.Sell {
// 		return
// 	}

// 	buyOrders = make([]*protoOrder.Order, 0)
// 	sellSize := sellOrder.Size
// 	for i := len(book.BuyOrders) - 1; i >= 0; i-- {
// 		buy := book.BuyOrders[i]
// 		if buy.Price >= sellOrder.Price && sellSize > 0 {
// 			buySize := buy.Size

// 			if sellSize < buySize {
// 				buy.Fill = sellSize
// 				buy.Size -= sellSize
// 				sellSize = 0
// 			} else {
// 				buy.Fill = buySize
// 				sellSize -= buySize
// 				// remove filled orders
// 				book.BuyOrders = book.BuyOrders[:i]
// 			}

// 			buyOrders = append(buyOrders, buy)
// 		}
// 	}

// 	return
// }

// // get all sell orders that a buy order can fill
// func (book *OrderBook) FillSellOrders(buyOrder *protoOrder.Order) (sellOrders []*protoOrder.Order) {
// 	if buyOrder.Side != constOrder.Buy {
// 		return
// 	}

// 	sellOrders = make([]*protoOrder.Order, 0)
// 	buySize := buyOrder.Size
// 	for i, sell := range book.SellOrders {
// 		if sell.Price <= buyOrder.Price && buySize > 0 {
// 			sellSize := sell.Size

// 			if buySize < sellSize {
// 				sell.Fill = buySize
// 				sell.Size -= buySize
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
