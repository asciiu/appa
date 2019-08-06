package models

import (
	constOrder "github.com/asciiu/appa/order-service/constants"
)

// OrderBook has a market
// The only public method should be Process.
type OrderBook struct {
	MarketName string
	LastPrice  uint64
	LastSide   string
	BuyOrders  []*Order
	SellOrders []*Order
}

// NewOrderBook will create a new instance of an order
// book for a market.
func NewOrderBook(marketName string) *OrderBook {
	return &OrderBook{
		MarketName: marketName,
		BuyOrders:  make([]*Order, 0),
		SellOrders: make([]*Order, 0),
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

	if n == 0 || book.SellOrders[n-1].Price > order.Price {
		// we have a new lower seller append it to the end of all sells
		book.SellOrders = append(book.SellOrders, order)
		return
	}

	var i int
	for i := n - 1; i >= 0; i-- {
		sellOrder := book.SellOrders[i]
		if sellOrder.Price > order.Price {
			break
		}
	}

	book.SellOrders = append(book.SellOrders, order)
	copy(book.SellOrders[i+1:], book.SellOrders[i:])
	book.SellOrders[i] = order
}

// Remove a buy order from the order book at a given index
func (book *OrderBook) removeBuyOrder(index int) {
	book.BuyOrders = append(book.BuyOrders[:index], book.BuyOrders[index+1:]...)
}

// Remove a sell order from the order book at a given index
func (book *OrderBook) removeSellOrder(index int) {
	book.SellOrders = append(book.SellOrders[:index], book.SellOrders[index+1:]...)
}

// Process an order and return the trades generated before adding the remaining amount to the market
func (book *OrderBook) Process(order *Order) []*Trade {
	if order.MarketName != book.MarketName {
		return []*Trade{}
	}

	trades := []*Trade{}
	switch {
	case order.Side == constOrder.Buy:
		trades = book.processLimitBuy(order)
	case order.Side == constOrder.Sell:
		trades = book.processLimitSell(order)
	}

	return trades
}

// Process a limit buy order
func (book *OrderBook) processLimitBuy(buyOrder *Order) []*Trade {
	trades := make([]*Trade, 0, 1)
	numSellOrders := len(book.SellOrders)

	if numSellOrders > 0 {
		i1 := 0
		i2 := 0
		count := 0

		// loop through sell orders beginning with
		// orders with the lowest price. Find index
		// range (i1, i2) within sorted (i.e. descending price) sell orders.
		// These are the sell orders that can fill the buy
		// order and should be filled from i2 -> i1
		for i := numSellOrders - 1; i >= 0; i-- {
			sellOrder := book.SellOrders[i]

			// a higher asking sell price cannot fill a lower buy price
			if sellOrder.Price > buyOrder.Price {
				// the first buy order index should be the next order
				// after this one
				i1 = i + 1
				break
			}

			// a buy order with a higher asking price should
			// fill a sell order with a lower asking price
			if sellOrder.Price <= buyOrder.Price {
				count++
				if i2 == 0 {
					i2 = i
				}
			}
		}

		if count > 0 {
			// fills orders from i2 -> i1.
			for j := i2; buyOrder.Amount > 0 && j >= i1; j-- {
				sellOrder := book.SellOrders[j]

				trade := &Trade{
					TakerOrderID: buyOrder.ID,
					MakerOrderID: sellOrder.ID,
					Price:        sellOrder.Price,
					Side:         buyOrder.Side,
				}
				book.LastPrice = sellOrder.Price
				book.LastSide = buyOrder.Side

				// fill the entire buy order
				if sellOrder.Amount >= buyOrder.Amount {
					// sell order amount >= buy amount therefore,
					// the entire buy order will be filled with
					// a single limit sell order. The amount traded will
					// be the buyer's amount.
					trade.Amount = buyOrder.Amount
					trades = append(trades, trade)

					// update new sell order amount
					sellOrder.Amount -= buyOrder.Amount

					// if the sell order amount == 0 then the sell
					// order has been filled - remove it
					if sellOrder.Amount == 0 {
						book.removeSellOrder(j)
					}
					return trades
				}

				// the entire limit sell order will be filled by this buy
				// order because the sell amount < buy amount.
				trade.Amount = sellOrder.Amount
				trades = append(trades, trade)

				// update buy order amount
				buyOrder.Amount -= sellOrder.Amount

				// buy order should be removed
				book.removeSellOrder(j)
			}
		}
	}

	if buyOrder.Amount > 0 {
		book.addBuyOrder(buyOrder)
	}

	return trades
}

// Process a limit sell order
func (book *OrderBook) processLimitSell(sellOrder *Order) []*Trade {
	trades := make([]*Trade, 0, 1)
	numBuyOrders := len(book.BuyOrders)

	if numBuyOrders > 0 {
		i1 := 0
		i2 := 0
		count := 0

		// loop through buy orders beginning with
		// orders with the highest price. Find index
		// range (i1, i2) within sorted buy orders.
		// These are the orders that can fill the sell
		// order and should be filled from i1 -> i2
		for i := numBuyOrders - 1; i >= 0; i-- {
			buyOrder := book.BuyOrders[i]

			// a lower buy price cannot fill a higher sell price
			if buyOrder.Price < sellOrder.Price {
				// the first buy order index should be the next order
				// after this one
				i1 = i + 1
				break
			}

			// any buy order with a higher asking price should
			// fill a sell order with a lower asking price
			if buyOrder.Price >= sellOrder.Price {
				count++
				if i2 == 0 {
					i2 = i
				}
			}
		}

		if count > 0 {

			// fills orders from i1 -> i2.
			for j := i1; sellOrder.Amount > 0 && j <= i2; j++ {
				buyOrder := book.BuyOrders[j]

				trade := &Trade{
					TakerOrderID: sellOrder.ID,
					MakerOrderID: buyOrder.ID,
					Price:        buyOrder.Price,
					Side:         sellOrder.Side,
				}
				book.LastPrice = buyOrder.Price
				book.LastSide = sellOrder.Side

				// fill the entire sell order
				if buyOrder.Amount >= sellOrder.Amount {
					// buy order amount >= sell amount therefore,
					// the entire sell order will be filled with
					// a single buy order. The amount traded will
					// be the seller's amount.
					trade.Amount = sellOrder.Amount
					trades = append(trades, trade)

					// update new buy order amount
					buyOrder.Amount -= sellOrder.Amount

					// if the buy order amount == 0 then the buy
					// order has been filled - remove it
					if buyOrder.Amount == 0 {
						book.removeBuyOrder(j)
					}
					return trades
				}

				// the entire buy order will be filled by this sell
				// order because the sell amount > buy amount.
				trade.Amount = buyOrder.Amount
				trades = append(trades, trade)

				// update new sell order amount
				sellOrder.Amount -= buyOrder.Amount

				// buy order should be removed
				book.removeBuyOrder(j)
			}
		}
	}

	// add remaining sell amount to our sell orders
	if sellOrder.Amount > 0 {
		book.addSellOrder(sellOrder)
	}

	return trades
}
