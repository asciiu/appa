package models

import (
	"testing"

	constOrder "github.com/asciiu/appa/order-service/constants"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProcessBuyOrder(t *testing.T) {
	book := NewOrderBook("test-btc")
	order := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Amount:     1024,
		Price:      1000,
	}
	trades := book.Process(&order)

	assert.Equal(t, 1, len(book.BuyOrders), "should be 1 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 0, len(trades), "should be no trades")
}

func TestProcessSellOrder(t *testing.T) {
	book := NewOrderBook("test-btc")
	order := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Amount:     1024,
		Price:      1000,
	}
	trades := book.Process(&order)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 1, len(book.SellOrders), "should be 1 order in sells")
	assert.Equal(t, 0, len(trades), "should be no trades")
}

func TestWrongMarketName(t *testing.T) {
	book := NewOrderBook("test-btc")
	order := Order{
		ID:         uuid.New().String(),
		MarketName: "test-bch",
		Side:       constOrder.Sell,
		Amount:     1024,
		Price:      1000,
	}
	trades := book.Process(&order)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 0, len(trades), "should be no trades")
}

func TestSellFill(t *testing.T) {
	book := NewOrderBook("test-btc")
	sell := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Amount:     1024,
		Price:      1000,
	}
	book.Process(&sell)

	buy := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Amount:     700,
		Price:      1010,
	}

	amount := sell.Amount - buy.Amount
	trades := book.Process(&buy)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 1, len(book.SellOrders), "should be 1 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, buy.Amount, trades[0].Amount, "trade amount shold be full buy amount")
	assert.Equal(t, constOrder.Buy, trades[0].Side, "side should be buy")
	assert.Equal(t, buy.ID, trades[0].TakerOrderID, "taker is buy order")
	assert.Equal(t, sell.ID, trades[0].MakerOrderID, "maker is sell order")
	assert.Equal(t, sell.Price, trades[0].Price, "price not at sell order price")
	assert.Equal(t, book.SellOrders[0].Amount, amount, "remaining sell amount is incorrect")
}

func TestPartialSell(t *testing.T) {
	book := NewOrderBook("test-btc")
	buy := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Amount:     400,
		Price:      9000,
	}
	book.Process(&buy)

	sell := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Amount:     600,
		Price:      8999,
	}

	remainingSellAmount := sell.Amount - buy.Amount
	trades := book.Process(&sell)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 1, len(book.SellOrders), "should be 1 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, buy.Amount, trades[0].Amount, "trade amount shold be full buy amount")
	assert.Equal(t, constOrder.Sell, trades[0].Side, "side should be sell")
	assert.Equal(t, sell.ID, trades[0].TakerOrderID, "taker is sell order")
	assert.Equal(t, buy.ID, trades[0].MakerOrderID, "maker is buy order")
	assert.Equal(t, buy.Price, trades[0].Price, "price not at buy order price")
	assert.Equal(t, book.SellOrders[0].Amount, remainingSellAmount, "remaining sell amount is incorrect")
}

func TestBuyFill(t *testing.T) {
	book := NewOrderBook("test-btc")
	buy := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Amount:     400,
		Price:      9000,
	}
	book.Process(&buy)

	sell := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Amount:     300,
		Price:      8999,
	}

	remainingBuyAmount := buy.Amount - sell.Amount
	trades := book.Process(&sell)

	assert.Equal(t, 1, len(book.BuyOrders), "should be 1 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, sell.Amount, trades[0].Amount, "trade amount shold be full sell amount")
	assert.Equal(t, constOrder.Sell, trades[0].Side, "side should be sell")
	assert.Equal(t, sell.ID, trades[0].TakerOrderID, "taker is sell order")
	assert.Equal(t, buy.ID, trades[0].MakerOrderID, "maker is buy order")
	assert.Equal(t, buy.Price, trades[0].Price, "price not at buy order price")
	assert.Equal(t, book.BuyOrders[0].Amount, remainingBuyAmount, "remaining buy amount is incorrect")
}

func TestPartialBuy(t *testing.T) {
	book := NewOrderBook("test-btc")
	sell := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Amount:     1024,
		Price:      1000,
	}
	book.Process(&sell)

	buy := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Amount:     1050,
		Price:      1010,
	}

	// remaining buy amount
	amount := buy.Amount - sell.Amount
	trades := book.Process(&buy)

	assert.Equal(t, 1, len(book.BuyOrders), "should be 1 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, sell.Amount, trades[0].Amount, "trade amount shold be full sell amount")
	assert.Equal(t, constOrder.Buy, trades[0].Side, "side should be buy")
	assert.Equal(t, buy.ID, trades[0].TakerOrderID, "taker is buy order")
	assert.Equal(t, sell.ID, trades[0].MakerOrderID, "maker is sell order")
	assert.Equal(t, sell.Price, trades[0].Price, "price not at sell order price")
	assert.Equal(t, book.BuyOrders[0].Amount, amount, "remaining buy amount is incorrect")
}

// Test first in first out buy. Buy orders at the same price should be FIFO
func TestFIFOBuy(t *testing.T) {
	book := NewOrderBook("test-btc")
	buy1 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Amount:     1024,
		Price:      1000,
	}
	trades := book.Process(&buy1)
	assert.Equal(t, 0, len(trades), "should be 0 trades")

	buy2 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Amount:     200,
		Price:      1000,
	}
	trades = book.Process(&buy2)
	assert.Equal(t, 0, len(trades), "should be 0 trades")

	sell := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Amount:     800,
		Price:      900,
	}
	// remaining buy1 amount after sell
	amount := buy1.Amount - sell.Amount
	trades = book.Process(&sell)
	assert.Equal(t, 1, len(trades), "should be 1 trade")

	assert.Equal(t, 2, len(book.BuyOrders), "should be 2 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, sell.Amount, trades[0].Amount, "trade amount shold be full sell amount")
	assert.Equal(t, constOrder.Sell, trades[0].Side, "side should be sell")
	assert.Equal(t, sell.ID, trades[0].TakerOrderID, "taker should be sell order ID")
	assert.Equal(t, buy1.ID, trades[0].MakerOrderID, "maker should be buy order I")
	assert.Equal(t, buy1.Price, trades[0].Price, "price not at sell order price")
	assert.Equal(t, book.BuyOrders[0].Amount, amount, "remaining buy amount is incorrect")
	assert.Equal(t, buy1.ID, book.BuyOrders[0].ID, "first buy order is incorrect")
}
