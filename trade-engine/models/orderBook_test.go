package models

import (
	"database/sql"
	"testing"

	"github.com/asciiu/appa/common/db"
	constants "github.com/asciiu/appa/trade-engine/constants"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func setup() *sql.DB {
	dbURL := "postgres://postgres@localhost:5432/appa_test?&sslmode=disable"
	db, _ := db.NewDB(dbURL)
	return db
}
func TestProcessBuyOrder(t *testing.T) {

	book := NewOrderBook("test-btc")
	order := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     1024,
		Price:      1000,
	}
	filledOrders, trades := book.Process(&order)

	assert.Equal(t, 1, len(book.BuyOrders), "should be 1 order in buys")
	assert.Equal(t, 0, len(filledOrders), "should be 0 filled orders")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 0, len(trades), "should be no trades")
}

func TestProcessSellOrder(t *testing.T) {

	book := NewOrderBook("test-btc")
	order := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     1024,
		Price:      1000,
	}
	filledOrders, trades := book.Process(&order)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 0, len(filledOrders), "should be 0 filled orders")
	assert.Equal(t, 1, len(book.SellOrders), "should be 1 order in sells")
	assert.Equal(t, 0, len(trades), "should be no trades")
}

func TestWrongMarketName(t *testing.T) {

	book := NewOrderBook("test-btc")
	order := Order{
		ID:         uuid.New().String(),
		MarketName: "test-bch",
		Side:       constants.Sell,
		Amount:     1024,
		Price:      1000,
	}
	filledOrders, trades := book.Process(&order)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 0, len(filledOrders), "should be 0 filled orders")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 0, len(trades), "should be no trades")
}

func TestSellFill(t *testing.T) {

	book := NewOrderBook("test-btc")
	sell := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     1024,
		Price:      1000,
	}
	book.Process(&sell)

	buy := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     700,
		Price:      1010,
	}

	remainingSellAmount := sell.Amount - buy.Amount
	boughtAmount := buy.Amount
	filledOrders, trades := book.Process(&buy)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 1, len(book.SellOrders), "should be 1 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, 1, len(filledOrders), "should be 1 order filled")

	assert.Equal(t, filledOrders[0].Amount, remainingSellAmount, "remaining amount of filled order is wrong")
	assert.Equal(t, filledOrders[0].Filled, boughtAmount, "filled amount is wrong")
	assert.Equal(t, filledOrders[0].Side, sell.Side, "filled order side incorrect")

	assert.Equal(t, uint64(0), buy.Amount, "buy amount should be 0")
	assert.Equal(t, boughtAmount, buy.Filled, "fill amount should be original buy amount")
	assert.Equal(t, constants.Completed, buy.Status, "buy should be complete")

	assert.Equal(t, boughtAmount, trades[0].Amount, "trade amount shold be full buy amount")
	assert.Equal(t, constants.Buy, trades[0].Side, "side should be buy")
	assert.Equal(t, buy.ID, trades[0].TakerOrderID, "taker is buy order")
	assert.Equal(t, sell.ID, trades[0].MakerOrderID, "maker is sell order")
	assert.Equal(t, sell.Price, trades[0].Price, "price not at sell order price")
	assert.Equal(t, book.SellOrders[0].Amount, remainingSellAmount, "remaining sell amount is incorrect")
}

func TestPartialSell(t *testing.T) {

	book := NewOrderBook("test-btc")
	buy := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     400,
		Price:      9000,
	}
	book.Process(&buy)

	sell := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     600,
		Price:      8999,
	}

	remainingSellAmount := sell.Amount - buy.Amount
	boughtAmount := buy.Amount
	filledOrders, trades := book.Process(&sell)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 1, len(book.SellOrders), "should be 1 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, 1, len(filledOrders), "should be 1 order filled")

	assert.Equal(t, filledOrders[0].Amount, uint64(0), "remaining amount of filled order is wrong")
	assert.Equal(t, filledOrders[0].Filled, boughtAmount, "filled amount is wrong")
	assert.Equal(t, filledOrders[0].Side, buy.Side, "filled order side incorrect")

	assert.Equal(t, remainingSellAmount, sell.Amount, "sell amount is wrong")
	assert.Equal(t, boughtAmount, sell.Filled, "fill amount should be original buy amount")
	assert.Equal(t, constants.Pending, buy.Status, "sell should be pending still")

	assert.Equal(t, boughtAmount, trades[0].Amount, "trade amount shold be full buy amount")
	assert.Equal(t, constants.Sell, trades[0].Side, "side should be sell")
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
		Side:       constants.Buy,
		Amount:     400,
		Price:      9000,
	}
	book.Process(&buy)

	sell := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     300,
		Price:      8999,
	}

	remainingBuyAmount := buy.Amount - sell.Amount
	boughtAmount := sell.Amount
	filledOrders, trades := book.Process(&sell)

	assert.Equal(t, 1, len(book.BuyOrders), "should be 1 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, 1, len(filledOrders), "should be 1 order filled")

	assert.Equal(t, remainingBuyAmount, filledOrders[0].Amount, "remaining amount of filled order is wrong")
	assert.Equal(t, boughtAmount, filledOrders[0].Filled, "filled amount is wrong")
	assert.Equal(t, constants.Buy, filledOrders[0].Side, "filled order side incorrect")
	assert.Equal(t, constants.Pending, filledOrders[0].Status, "flled order should still be pending")

	assert.Equal(t, boughtAmount, trades[0].Amount, "trade amount shold be full sell amount")
	assert.Equal(t, constants.Sell, trades[0].Side, "side should be sell")
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
		Side:       constants.Sell,
		Amount:     1024,
		Price:      1000,
	}
	book.Process(&sell)

	buy := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     1050,
		Price:      1010,
	}

	// remaining buy amount
	remainingBuyAmount := buy.Amount - sell.Amount
	boughtAmount := sell.Amount
	filledOrders, trades := book.Process(&buy)

	assert.Equal(t, 1, len(book.BuyOrders), "should be 1 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, 1, len(filledOrders), "should be 1 order filled")

	assert.Equal(t, uint64(0), filledOrders[0].Amount, "remaining amount of filled order is wrong")
	assert.Equal(t, uint64(boughtAmount), filledOrders[0].Filled, "filled amount is wrong")
	assert.Equal(t, filledOrders[0].Side, sell.Side, "filled order side incorrect")
	assert.Equal(t, constants.Completed, filledOrders[0].Status, "filled order should be completed")

	assert.Equal(t, boughtAmount, trades[0].Amount, "trade amount shold be full sell amount")
	assert.Equal(t, constants.Buy, trades[0].Side, "side should be buy")
	assert.Equal(t, buy.ID, trades[0].TakerOrderID, "taker is buy order")
	assert.Equal(t, sell.ID, trades[0].MakerOrderID, "maker is sell order")
	assert.Equal(t, sell.Price, trades[0].Price, "price not at sell order price")
	assert.Equal(t, book.BuyOrders[0].Amount, remainingBuyAmount, "remaining buy amount is incorrect")
}

func TestBuySortOrder(t *testing.T) {

	buyOrders := []*Order{
		&Order{
			ID:         "1",
			MarketName: "test-btc",
			Side:       constants.Buy,
			Amount:     100,
			Price:      1000,
		},
		&Order{
			ID:         "2",
			MarketName: "test-btc",
			Side:       constants.Buy,
			Amount:     204,
			Price:      2000,
		},
		&Order{
			ID:         "0",
			MarketName: "test-btc",
			Side:       constants.Buy,
			Amount:     400,
			Price:      1000,
		},
	}
	book := NewOrderBook("test-btc")
	for _, order := range buyOrders {
		book.Process(order)
	}

	// for _, order := range book.BuyOrders {
	// 	fmt.Printf("%+v\n", order)
	// }

	assert.Equal(t, 3, len(book.BuyOrders), "should be 3 orders")
	assert.Equal(t, buyOrders[1].ID, book.BuyOrders[2].ID, "highest price buy should be last")
	assert.Equal(t, buyOrders[0].ID, book.BuyOrders[1].ID, "first added order should be second")
	assert.Equal(t, buyOrders[2].ID, book.BuyOrders[0].ID, "last added order should be first")
}

func TestSellSortOrder(t *testing.T) {

	sellOrders := []*Order{
		&Order{
			ID:         "1",
			MarketName: "test-btc",
			Side:       constants.Sell,
			Amount:     100,
			Price:      1000,
		},
		&Order{
			ID:         "2",
			MarketName: "test-btc",
			Side:       constants.Sell,
			Amount:     400,
			Price:      200,
		},
		&Order{
			ID:         "0",
			MarketName: "test-btc",
			Side:       constants.Sell,
			Amount:     204,
			Price:      1000,
		},
	}

	book := NewOrderBook("test-btc")
	for _, order := range sellOrders {
		book.Process(order)
	}

	assert.Equal(t, 3, len(book.SellOrders), "should be 3 orders")
	assert.Equal(t, sellOrders[1].ID, book.SellOrders[2].ID, "lowest price sell should be last")
	assert.Equal(t, sellOrders[0].ID, book.SellOrders[1].ID, "first added order should be second")
	assert.Equal(t, sellOrders[2].ID, book.SellOrders[0].ID, "last added order should be first")
}

// Test first in first out buy. Buy orders at the same price should be FIFO
func TestFIFOBuyOrders(t *testing.T) {

	book := NewOrderBook("test-btc")
	buy1 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     1024,
		Price:      1000,
	}
	_, trades := book.Process(&buy1)
	assert.Equal(t, 0, len(trades), "should be 0 trades")

	buy2 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     200,
		Price:      1000,
	}
	_, trades = book.Process(&buy2)
	assert.Equal(t, 0, len(trades), "should be 0 trades")

	sell := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     800,
		Price:      900,
	}
	// remaining buy1 amount after sell
	remainingBuy1Amount := buy1.Amount - sell.Amount
	boughtAmount := sell.Amount
	filledOrders, trades := book.Process(&sell)

	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, 2, len(book.BuyOrders), "should be 2 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 1, len(filledOrders), "should be 1 order filled")

	assert.Equal(t, remainingBuy1Amount, filledOrders[0].Amount, "remaining amount of filled order is wrong")
	assert.Equal(t, boughtAmount, filledOrders[0].Filled, "filled amount is wrong")
	assert.Equal(t, buy1.Side, filledOrders[0].Side, "filled order side incorrect")
	assert.Equal(t, constants.Pending, filledOrders[0].Status, "partial filled order should still be pending")

	assert.Equal(t, boughtAmount, trades[0].Amount, "trade amount shold be full sell amount")
	assert.Equal(t, constants.Sell, trades[0].Side, "side should be sell")
	assert.Equal(t, sell.ID, trades[0].TakerOrderID, "taker should be sell order ID")
	assert.Equal(t, buy1.ID, trades[0].MakerOrderID, "maker should be buy order I")
	assert.Equal(t, buy1.Price, trades[0].Price, "price not at sell order price")

	assert.Equal(t, remainingBuy1Amount, book.BuyOrders[1].Amount, "remaining buy amount is incorrect")
	assert.Equal(t, boughtAmount, book.BuyOrders[1].Filled, "buy filled amount is incorrect")
	assert.Equal(t, buy1.ID, book.BuyOrders[1].ID, "first buy order is incorrect")

	assert.Equal(t, uint64(0), sell.Amount, "sell amount should be 0")
	assert.Equal(t, boughtAmount, sell.Filled, "sell fill is wrong")
	assert.Equal(t, constants.Completed, sell.Status, "sell should be complete")
}

func TestFIFOSellOrders(t *testing.T) {

	book := NewOrderBook("test-btc")
	sell1 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     1024,
		Price:      850,
	}
	filledOrders, trades := book.Process(&sell1)
	assert.Equal(t, 0, len(trades), "should be 0 trades")
	assert.Equal(t, 0, len(filledOrders), "should be 0 filled orders")

	sell2 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     200,
		Price:      850,
	}
	filledOrders, trades = book.Process(&sell2)
	assert.Equal(t, 0, len(trades), "should be 0 trades")
	assert.Equal(t, 0, len(filledOrders), "should be 0 filled orders")

	buy := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     800,
		Price:      900,
	}
	// remaining sell1 amount after buy
	remainingSell1Amount := sell1.Amount - buy.Amount
	boughtAmount := buy.Amount
	filledOrders, trades = book.Process(&buy)

	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, 1, len(filledOrders), "should be 1 order filled")

	assert.Equal(t, remainingSell1Amount, filledOrders[0].Amount, "remaining amount of filled order is wrong")
	assert.Equal(t, boughtAmount, filledOrders[0].Filled, "filled amount is wrong")
	assert.Equal(t, filledOrders[0].Side, sell1.Side, "filled order side incorrect")

	assert.Equal(t, 2, len(book.SellOrders), "should be 2 order in sells")
	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, boughtAmount, trades[0].Amount, "trade amount should be full buy amount")

	assert.Equal(t, constants.Buy, trades[0].Side, "side should be buy")
	assert.Equal(t, buy.ID, trades[0].TakerOrderID, "taker should be buy order ID")
	assert.Equal(t, sell1.ID, trades[0].MakerOrderID, "maker should be sell order I")

	assert.Equal(t, uint64(0), buy.Amount, "buy amount should be 0")
	assert.Equal(t, boughtAmount, buy.Filled, "buy filled amount wrong")
	assert.Equal(t, constants.Completed, buy.Status, "buy status should be complete")

	assert.Equal(t, sell1.Price, trades[0].Price, "price not at sell order price")
	assert.Equal(t, remainingSell1Amount, book.SellOrders[1].Amount, "remaining sell amount of first order is incorrect")
	assert.Equal(t, sell1.ID, book.SellOrders[1].ID, "first selll order is incorrect")
}

func TestMultiSellOrderFill(t *testing.T) {

	book := NewOrderBook("test-btc")
	sell1 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     1024,
		Price:      850,
	}
	filledOrders, trades := book.Process(&sell1)
	assert.Equal(t, 0, len(trades), "should be 0 trades")
	assert.Equal(t, 0, len(filledOrders), "should be 0 filled orders")

	sell2 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     200,
		Price:      850,
	}
	filledOrders, trades = book.Process(&sell2)
	assert.Equal(t, 0, len(trades), "should be 0 trades")
	assert.Equal(t, 0, len(filledOrders), "should be 0 filled orders")

	sell3 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     200,
		Price:      910,
	}
	filledOrders, trades = book.Process(&sell3)
	assert.Equal(t, 0, len(trades), "should be 0 trades")
	assert.Equal(t, 0, len(filledOrders), "should be 0 filled orders")

	buy := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     1500,
		Price:      900,
	}

	sell1Amount := sell1.Amount
	sell2Amount := sell2.Amount
	remainingBuyAmount := buy.Amount - sell1.Amount - sell2.Amount
	boughtAmount := sell1.Amount + sell2.Amount

	filledOrders, trades = book.Process(&buy)

	assert.Equal(t, 2, len(trades), "should be 2 trades")
	assert.Equal(t, 2, len(filledOrders), "should be 2 orders filled")

	// filled orders tests
	assert.Equal(t, uint64(0), filledOrders[0].Amount, "remaining amount of filled order is wrong")
	assert.Equal(t, sell1Amount, filledOrders[0].Filled, "filled amount is wrong")
	assert.Equal(t, sell1.Side, filledOrders[0].Side, "filled order side incorrect")
	assert.Equal(t, uint64(0), filledOrders[1].Amount, "remaining amount of filled order is wrong")
	assert.Equal(t, sell2Amount, filledOrders[1].Filled, "filled amount is wrong")
	assert.Equal(t, sell2.Side, filledOrders[1].Side, "filled order side incorrect")

	assert.Equal(t, 1, len(book.SellOrders), "should be 2 order in sells")
	assert.Equal(t, 1, len(book.BuyOrders), "should be 0 order in buys")

	// trade tests
	assert.Equal(t, sell1Amount, trades[0].Amount, "trade amount should be full sell 1 amount")
	assert.Equal(t, constants.Buy, trades[0].Side, "side should be buy")
	assert.Equal(t, buy.ID, trades[0].TakerOrderID, "taker should be buy order ID")
	assert.Equal(t, sell1.ID, trades[0].MakerOrderID, "maker should be sell order I")
	assert.Equal(t, sell2Amount, trades[1].Amount, "trade amount should be full sell 2 amount")
	assert.Equal(t, constants.Buy, trades[1].Side, "side should be buy")
	assert.Equal(t, buy.ID, trades[1].TakerOrderID, "taker should be buy order ID")
	assert.Equal(t, sell2.ID, trades[1].MakerOrderID, "maker should be sell order I")

	// buy order
	assert.Equal(t, remainingBuyAmount, buy.Amount, "buy amount wrong")
	assert.Equal(t, boughtAmount, buy.Filled, "buy filled amount wrong")
	assert.Equal(t, constants.Pending, buy.Status, "buy status should be pending")
	assert.Equal(t, remainingBuyAmount, book.BuyOrders[0].Amount, "buy amount wrong")
	assert.Equal(t, boughtAmount, book.BuyOrders[0].Filled, "buy filled wrong")
	assert.Equal(t, constants.Pending, book.BuyOrders[0].Status, "buy status should be pending")
}

func TestCancelSellOrder(t *testing.T) {

	sellOrders := []*Order{
		&Order{
			ID:         "1",
			MarketName: "test-btc",
			Side:       constants.Sell,
			Amount:     100,
			Price:      1000,
		},
		&Order{
			ID:         "2",
			MarketName: "test-btc",
			Side:       constants.Sell,
			Amount:     400,
			Price:      200,
		},
		&Order{
			ID:         "0",
			MarketName: "test-btc",
			Side:       constants.Sell,
			Amount:     204,
			Price:      1000,
		},
	}

	book := NewOrderBook("test-btc")
	for _, order := range sellOrders {
		book.Process(order)
	}
	err := book.Cancel(sellOrders[2])

	assert.Nil(t, err, "error from cancel should be nil")
	assert.Equal(t, 2, len(book.SellOrders), "should be 2 orders")

	assert.Equal(t, sellOrders[0].ID, book.SellOrders[0].ID, "first sell order should have highest price")
	assert.Equal(t, sellOrders[1].ID, book.SellOrders[1].ID, "last sell order should have lowest price")
}

func TestCancelBuyOrder(t *testing.T) {

	// these buy orders will be sorted by acending price
	buyOrders := []*Order{
		&Order{
			ID:         "1",
			MarketName: "test-btc",
			Side:       constants.Buy,
			Amount:     100,
			Price:      1000,
		},
		&Order{
			ID:         "2",
			MarketName: "test-btc",
			Side:       constants.Buy,
			Amount:     204,
			Price:      2000,
		},
		&Order{
			ID:         "0",
			MarketName: "test-btc",
			Side:       constants.Buy,
			Amount:     400,
			Price:      1000,
		},
	}

	book := NewOrderBook("test-btc")
	for _, order := range buyOrders {
		book.Process(order)
	}

	// remove the highest priced order at with ID #2
	err := book.Cancel(buyOrders[1])
	assert.Nil(t, err, "error was not nil for cancel order")

	assert.Equal(t, 2, len(book.BuyOrders), "should be 2 orders")

	assert.Equal(t, buyOrders[2].ID, book.BuyOrders[0].ID, "first buy order is incorrect")
	assert.Equal(t, buyOrders[0].ID, book.BuyOrders[1].ID, "last buy order is incorrect")
}
