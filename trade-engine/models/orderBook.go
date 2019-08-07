package models

import (
	"fmt"

	constants "github.com/asciiu/appa/trade-engine/constants"
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
// the last order should be the highest priced order.
// When the price is equal the older order should come
// after the newer order.
func (book *OrderBook) addBuyOrder(order *Order) {
	order.Status = constants.Pending
	n := len(book.BuyOrders)

	if n == 0 || book.BuyOrders[n-1].Price < order.Price {
		book.BuyOrders = append(book.BuyOrders, order)
		return
	}

	var i int
	for i := n - 1; i >= 0; i-- {
		buyOrder := book.BuyOrders[i]
		if buyOrder.Price < order.Price {
			i++
			break
		}
	}

	book.BuyOrders = append(book.BuyOrders, order)
	copy(book.BuyOrders[i+1:], book.BuyOrders[i:])
	book.BuyOrders[i] = order
}

// sell orders will be kept sorted in descending price order
// the last order should the lowest priced order
func (book *OrderBook) addSellOrder(order *Order) {
	order.Status = constants.Pending
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

// Cancel a buy order.
func (book *OrderBook) cancelBuyOrder(orderID string) error {
	for i, o := range book.BuyOrders {
		if o.ID == orderID {
			// TODO update status of persisted order to "cancelled"
			// need to do this after the orders are persisted
			// persisted orders will eventually load at startup using
			// status == pending
			book.removeBuyOrder(i)
			return nil
		}
	}
	return fmt.Errorf("orderID: %s not found", orderID)
}

// Cancel a sell order.
func (book *OrderBook) cancelSellOrder(orderID string) error {
	for i, o := range book.SellOrders {
		if o.ID == orderID {
			// TODO update status of persisted order to "cancelled"
			// need to do this after the orders are persisted
			// persisted orders will eventually load at startup using
			// status == pending
			book.removeSellOrder(i)
			return nil
		}
	}
	return fmt.Errorf("orderID: %s not found", orderID)
}

// Cancel a limit order
func (book *OrderBook) Cancel(order *Order) error {
	if order.MarketName != book.MarketName {
		return fmt.Errorf("market name for order should be %s got %s", book.MarketName, order.MarketName)
	}

	if order.Side == constants.Sell {
		return book.cancelSellOrder(order.ID)
	}
	return book.cancelBuyOrder(order.ID)
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
// returns the orders that were filled and the trades associated with the orders
func (book *OrderBook) Process(order *Order) ([]*Order, []*Trade) {
	if order.MarketName != book.MarketName {
		return []*Order{}, []*Trade{}
	}

	switch {
	case order.Side == constants.Buy:
		return book.processLimitBuy(order)
	case order.Side == constants.Sell:
		return book.processLimitSell(order)
	}

	// side no recognized will result in no filled orders and trades
	return []*Order{}, []*Trade{}
}

// Process a limit buy order
func (book *OrderBook) processLimitBuy(buyOrder *Order) ([]*Order, []*Trade) {
	trades := make([]*Trade, 0, 1)
	orders := make([]*Order, 0, 1)

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

				trade := NewTrade(
					buyOrder.ID,
					sellOrder.ID,
					buyOrder.Side,
					0,
					sellOrder.Price,
				)

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
					sellOrder.Filled += buyOrder.Amount

					buyOrder.Filled += buyOrder.Amount
					buyOrder.Amount = 0
					buyOrder.Status = constants.Completed

					// if the sell order amount == 0 then the sell
					// order has been filled - remove it
					if sellOrder.Amount == 0 {
						sellOrder.Status = constants.Completed
						book.removeSellOrder(j)
					}
					orders = append(orders, sellOrder)

					return orders, trades
				}

				// the entire limit sell order will be filled by this buy
				// order because the sell amount < buy amount.
				trade.Amount = sellOrder.Amount
				trades = append(trades, trade)

				// update buy order amount
				buyOrder.Amount -= sellOrder.Amount
				buyOrder.Filled += sellOrder.Amount
				if buyOrder.Amount == 0 {
					buyOrder.Status = constants.Completed
				}

				sellOrder.Filled = sellOrder.Amount
				sellOrder.Amount = 0
				sellOrder.Status = constants.Completed
				orders = append(orders, sellOrder)

				// buy order should be removed
				book.removeSellOrder(j)
			}
		}
	}

	if buyOrder.Amount > 0 {
		book.addBuyOrder(buyOrder)
	}

	return orders, trades
}

// Process a limit sell order
func (book *OrderBook) processLimitSell(sellOrder *Order) ([]*Order, []*Trade) {
	orders := make([]*Order, 0, 1)
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
		// order and should be filled from i2 -> i1
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
			for j := i2; sellOrder.Amount > 0 && j >= i1; j-- {
				buyOrder := book.BuyOrders[j]

				trade := NewTrade(
					sellOrder.ID,
					buyOrder.ID,
					sellOrder.Side,
					0,
					buyOrder.Price,
				)

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
					buyOrder.Filled += sellOrder.Amount

					sellOrder.Filled += sellOrder.Amount
					sellOrder.Amount = 0
					sellOrder.Status = constants.Completed

					// if the buy order amount == 0 then the buy
					// order has been filled - remove it
					if buyOrder.Amount == 0 {
						buyOrder.Status = constants.Completed
						book.removeBuyOrder(j)
					}
					orders = append(orders, buyOrder)

					return orders, trades
				}

				// the entire buy order will be filled by this sell
				// order because the sell amount > buy amount.
				trade.Amount = buyOrder.Amount
				trades = append(trades, trade)

				// update new sell order amount
				sellOrder.Amount -= buyOrder.Amount
				sellOrder.Filled += buyOrder.Amount
				if sellOrder.Amount == 0 {
					sellOrder.Status = constants.Completed
				}

				buyOrder.Filled = buyOrder.Amount
				buyOrder.Amount = 0
				orders = append(orders, buyOrder)

				// buy order should be removed
				book.removeBuyOrder(j)
			}
		}
	}

	// add remaining sell amount to our sell orders
	if sellOrder.Amount > 0 {
		book.addSellOrder(sellOrder)
	}

	return orders, trades
}
