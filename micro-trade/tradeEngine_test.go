package main

import (
	"context"
	"log"
	"testing"

	repo "github.com/asciiu/appa/api-graphql/db/sql"
	user "github.com/asciiu/appa/api-graphql/models"
	"github.com/asciiu/appa/lib/constants/response"
	"github.com/asciiu/appa/lib/db"
	"github.com/asciiu/appa/micro-trade/constants"
	tradeRepo "github.com/asciiu/appa/micro-trade/db/sql"
	"github.com/asciiu/appa/micro-trade/proto/trade"
	"github.com/stretchr/testify/assert"
)

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		panic(err)
	}
}

func setupEngine() (*TradeEngine, *user.User) {
	dbUrl := "postgres://postgres@localhost:5432/appa_test?&sslmode=disable"
	db, _ := db.NewDB(dbUrl)
	engine := NewTradeEngine(db)

	user := user.NewUser("testy", "test@email", "hash")
	err := repo.InsertUser(db, user)
	checkErr(err)

	return engine, user
}

func TestProcessOrder(t *testing.T) {
	engine, user := setupEngine()

	defer engine.DB.Close()

	req := trade.NewOrderRequest{
		UserID:     user.ID,
		MarketName: "btc-usd",
		Side:       constants.Sell,
		Amount:     1.0,
		Price:      200,
	}

	res := trade.OrderResponse{}
	engine.Process(context.Background(), &req, &res)

	// assert response
	assert.Equal(t, response.Success, res.Status, "expected success got: "+res.Message)
	assert.Equal(t, req.Side, res.Data.Order.Side, "side incorrect")
	assert.Equal(t, req.Amount, res.Data.Order.Amount, "amounts should be the same")
	assert.Equal(t, uint64(0), res.Data.Order.Filled, "fill incorrect")
	assert.Equal(t, constants.Pending, res.Data.Order.Status, "status incorrect")

	// assert db data
	order, err := tradeRepo.FindOrderByID(engine.DB, res.Data.Order.OrderID)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, order.ID, res.Data.Order.OrderID, "id mismatch")
	assert.Equal(t, req.Amount, order.Amount, "amounts should be the same")
	assert.Equal(t, uint64(0), order.Filled, "fill incorrect")
	assert.Equal(t, constants.Pending, order.Status, "status incorrect")

	repo.DeleteUserHard(engine.DB, user.ID)
}

func TestProcessTrade(t *testing.T) {
	engine, user := setupEngine()

	defer engine.DB.Close()

	or1 := &trade.NewOrderRequest{
		UserID:     user.ID,
		MarketName: "btc-usd",
		Side:       constants.Sell,
		Amount:     12,
		Price:      10000000,
	}
	or2 := &trade.NewOrderRequest{
		UserID:     user.ID,
		MarketName: "btc-usd",
		Side:       constants.Sell,
		Amount:     20,
		Price:      1000000,
	}
	or3 := &trade.NewOrderRequest{
		UserID:     user.ID,
		MarketName: "btc-usd",
		Side:       constants.Sell,
		Amount:     7,
		Price:      100000,
	}
	or4 := &trade.NewOrderRequest{
		UserID:     user.ID,
		MarketName: "btc-usd",
		Side:       constants.Sell,
		Amount:     2,
		Price:      100000,
	}
	buy := &trade.NewOrderRequest{
		UserID:     user.ID,
		MarketName: "btc-usd",
		Side:       constants.Buy,
		Amount:     8,
		Price:      100010,
	}

	sellReqs := []*trade.NewOrderRequest{or1, or2, or3, or4}
	res := &trade.OrderResponse{}
	for _, req := range sellReqs {
		engine.Process(context.Background(), req, res)
	}

	engine.Process(context.Background(), buy, res)

	// assert response
	assert.Equal(t, response.Success, res.Status, "expected success got: "+res.Message)
	assert.Equal(t, buy.Side, res.Data.Order.Side, "side incorrect")
	assert.Equal(t, uint64(0), res.Data.Order.Amount, "amounts should be 0")
	assert.Equal(t, buy.Amount, res.Data.Order.Filled, "fill should be entire amount")
	assert.Equal(t, constants.Completed, res.Data.Order.Status, "status incorrect")

	// assert order data for buy order
	order, err := tradeRepo.FindOrderByID(engine.DB, res.Data.Order.OrderID)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, uint64(0), order.Amount, "amount should be 0")
	assert.Equal(t, buy.Amount, order.Filled, "fill incorrect")
	assert.Equal(t, constants.Completed, order.Status, "status incorrect")

	// assert trade data
	tp, err := tradeRepo.FindUserTrades(engine.DB, user.ID, 0, 100)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, uint32(2), tp.Total, "should be 2 trades")
	assert.Equal(t, res.Data.Order.OrderID, tp.Trades[0].TakerOrderID, "maker not correct")
	assert.Equal(t, res.Data.Order.OrderID, tp.Trades[1].TakerOrderID, "maker not correct")
	assert.Equal(t, or3.Amount, tp.Trades[0].Amount, "amount not correct")
	assert.Equal(t, or3.Price, tp.Trades[0].Price, "price not correct")
	assert.Equal(t, uint64(7), tp.Trades[1].Amount, "amount not correct")
	assert.Equal(t, or4.Price, tp.Trades[1].Price, "price not correct")

	// assert order data for sell orders
	// trade 1 - maker sell order
	order, err = tradeRepo.FindOrderByID(engine.DB, tp.Trades[0].MakerOrderID)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, uint64(0), order.Amount, "amount should be 0")
	assert.Equal(t, uint64(7), order.Filled, "fill incorrect")
	assert.Equal(t, constants.Completed, order.Status, "status incorrect")

	// trade 2 - maker sell order
	order, err = tradeRepo.FindOrderByID(engine.DB, tp.Trades[1].MakerOrderID)
	assert.Nil(t, err, "error should be nil")
	assert.Equal(t, uint64(0), order.Amount, "amount should be 0")
	assert.Equal(t, uint64(7), order.Filled, "fill incorrect")
	assert.Equal(t, constants.Completed, order.Status, "status incorrect")

	repo.DeleteUserHard(engine.DB, user.ID)
}
