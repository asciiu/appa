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
	db := setup()

	book := NewOrderBook("test-btc", db)
	order := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     1024,
		Price:      1000,
	}
	trades := book.Process(&order)

	assert.Equal(t, 1, len(book.BuyOrders), "should be 1 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 0, len(trades), "should be no trades")
}

func TestProcessSellOrder(t *testing.T) {
	db := setup()
	book := NewOrderBook("test-btc", db)
	order := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     1024,
		Price:      1000,
	}
	trades := book.Process(&order)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 1, len(book.SellOrders), "should be 1 order in sells")
	assert.Equal(t, 0, len(trades), "should be no trades")
}

func TestWrongMarketName(t *testing.T) {
	db := setup()
	book := NewOrderBook("test-btc", db)
	order := Order{
		ID:         uuid.New().String(),
		MarketName: "test-bch",
		Side:       constants.Sell,
		Amount:     1024,
		Price:      1000,
	}
	trades := book.Process(&order)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 0, len(trades), "should be no trades")
}

func TestSellFill(t *testing.T) {
	db := setup()
	book := NewOrderBook("test-btc", db)
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

	amount := sell.Amount - buy.Amount
	trades := book.Process(&buy)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 1, len(book.SellOrders), "should be 1 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, buy.Amount, trades[0].Amount, "trade amount shold be full buy amount")
	assert.Equal(t, constants.Buy, trades[0].Side, "side should be buy")
	assert.Equal(t, buy.ID, trades[0].TakerOrderID, "taker is buy order")
	assert.Equal(t, sell.ID, trades[0].MakerOrderID, "maker is sell order")
	assert.Equal(t, sell.Price, trades[0].Price, "price not at sell order price")
	assert.Equal(t, book.SellOrders[0].Amount, amount, "remaining sell amount is incorrect")
}

func TestPartialSell(t *testing.T) {
	db := setup()
	book := NewOrderBook("test-btc", db)
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
	trades := book.Process(&sell)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 1, len(book.SellOrders), "should be 1 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, buy.Amount, trades[0].Amount, "trade amount shold be full buy amount")
	assert.Equal(t, constants.Sell, trades[0].Side, "side should be sell")
	assert.Equal(t, sell.ID, trades[0].TakerOrderID, "taker is sell order")
	assert.Equal(t, buy.ID, trades[0].MakerOrderID, "maker is buy order")
	assert.Equal(t, buy.Price, trades[0].Price, "price not at buy order price")
	assert.Equal(t, book.SellOrders[0].Amount, remainingSellAmount, "remaining sell amount is incorrect")
}

func TestBuyFill(t *testing.T) {
	db := setup()
	book := NewOrderBook("test-btc", db)
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
	trades := book.Process(&sell)

	assert.Equal(t, 1, len(book.BuyOrders), "should be 1 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, sell.Amount, trades[0].Amount, "trade amount shold be full sell amount")
	assert.Equal(t, constants.Sell, trades[0].Side, "side should be sell")
	assert.Equal(t, sell.ID, trades[0].TakerOrderID, "taker is sell order")
	assert.Equal(t, buy.ID, trades[0].MakerOrderID, "maker is buy order")
	assert.Equal(t, buy.Price, trades[0].Price, "price not at buy order price")
	assert.Equal(t, book.BuyOrders[0].Amount, remainingBuyAmount, "remaining buy amount is incorrect")
}

func TestPartialBuy(t *testing.T) {
	db := setup()
	book := NewOrderBook("test-btc", db)
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
	amount := buy.Amount - sell.Amount
	trades := book.Process(&buy)

	assert.Equal(t, 1, len(book.BuyOrders), "should be 1 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
	assert.Equal(t, 1, len(trades), "should be 1 trade")
	assert.Equal(t, sell.Amount, trades[0].Amount, "trade amount shold be full sell amount")
	assert.Equal(t, constants.Buy, trades[0].Side, "side should be buy")
	assert.Equal(t, buy.ID, trades[0].TakerOrderID, "taker is buy order")
	assert.Equal(t, sell.ID, trades[0].MakerOrderID, "maker is sell order")
	assert.Equal(t, sell.Price, trades[0].Price, "price not at sell order price")
	assert.Equal(t, book.BuyOrders[0].Amount, amount, "remaining buy amount is incorrect")
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
	db := setup()
	book := NewOrderBook("test-btc", db)
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
	db := setup()
	book := NewOrderBook("test-btc", db)
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
	db := setup()
	book := NewOrderBook("test-btc", db)
	buy1 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     1024,
		Price:      1000,
	}
	trades := book.Process(&buy1)
	assert.Equal(t, 0, len(trades), "should be 0 trades")

	buy2 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     200,
		Price:      1000,
	}
	trades = book.Process(&buy2)
	assert.Equal(t, 0, len(trades), "should be 0 trades")

	sell := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
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
	assert.Equal(t, constants.Sell, trades[0].Side, "side should be sell")
	assert.Equal(t, sell.ID, trades[0].TakerOrderID, "taker should be sell order ID")
	assert.Equal(t, buy1.ID, trades[0].MakerOrderID, "maker should be buy order I")
	assert.Equal(t, buy1.Price, trades[0].Price, "price not at sell order price")
	assert.Equal(t, book.BuyOrders[1].Amount, amount, "remaining buy amount is incorrect")
	assert.Equal(t, buy1.ID, book.BuyOrders[1].ID, "first buy order is incorrect")
}

func TestFIFOSellOrders(t *testing.T) {
	db := setup()
	book := NewOrderBook("test-btc", db)
	sell1 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     1024,
		Price:      850,
	}
	trades := book.Process(&sell1)
	assert.Equal(t, 0, len(trades), "should be 0 trades")

	sell2 := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Sell,
		Amount:     200,
		Price:      850,
	}
	trades = book.Process(&sell2)
	assert.Equal(t, 0, len(trades), "should be 0 trades")

	buy := Order{
		ID:         uuid.New().String(),
		MarketName: "test-btc",
		Side:       constants.Buy,
		Amount:     800,
		Price:      900,
	}
	// remaining sell1 amount after buy
	amount := sell1.Amount - buy.Amount
	trades = book.Process(&buy)
	assert.Equal(t, 1, len(trades), "should be 1 trade")

	assert.Equal(t, 2, len(book.SellOrders), "should be 2 order in sells")
	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, buy.Amount, trades[0].Amount, "trade amount shold be full buy amount")

	assert.Equal(t, constants.Buy, trades[0].Side, "side should be buy")
	assert.Equal(t, buy.ID, trades[0].TakerOrderID, "taker should be buy order ID")
	assert.Equal(t, sell1.ID, trades[0].MakerOrderID, "maker should be sell order I")

	assert.Equal(t, sell1.Price, trades[0].Price, "price not at sell order price")
	assert.Equal(t, book.SellOrders[1].Amount, amount, "remaining sell amount of first order is incorrect")
	assert.Equal(t, sell1.ID, book.SellOrders[1].ID, "first selll order is incorrect")
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
	db := setup()
	book := NewOrderBook("test-btc", db)
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

	db := setup()
	book := NewOrderBook("test-btc", db)
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
