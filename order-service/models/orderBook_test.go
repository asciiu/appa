package models

import (
	"fmt"
	"testing"
	"time"

	constOrder "github.com/asciiu/appa/order-service/constants"
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestOrderBook(t *testing.T) {
	book := NewOrderBook("test-btc")
	order := protoOrder.Order{
		OrderID:    "#1",
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	book.AddBuyOrder(&order)

	assert.Equal(t, 1, len(book.BuyOrders), "should be 1 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
}

func TestOrderBookWrongOrder(t *testing.T) {
	book := NewOrderBook("test-btc")
	order := protoOrder.Order{
		OrderID:    uuid.New().String(),
		MarketName: "bch-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	book.AddBuyOrder(&order)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
}

func TestOrderBookWrongSideOrder(t *testing.T) {
	book := NewOrderBook("test-btc")
	order := protoOrder.Order{
		OrderID:    uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	book.AddSellOrder(&order)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
}

func TestOrderBookMatchSellOrder(t *testing.T) {
	now := time.Now().UTC()
	book := NewOrderBook("test-btc")
	order0 := &protoOrder.Order{
		OrderID:    uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	order1 := &protoOrder.Order{
		OrderID:    "#1",
		MarketName: "test-btc",
		Price:      0.01,
		Size:       1.2,
		Side:       "sell",
		CreatedOn:  now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:    "#2",
		MarketName: "test-btc",
		Price:      0.007,
		Size:       0.2,
		Side:       "sell",
		CreatedOn:  now.Add(time.Second * 1).String(),
	}

	orders := []*protoOrder.Order{order0, order1, order2}
	for _, o := range orders {
		book.AddSellOrder(o)
	}

	buyOrder := &protoOrder.Order{
		OrderID:    "#buyer",
		MarketName: "test-btc",
		Price:      0.007,
		Size:       0.9,
		Side:       "buy",
	}
	match := book.MatchSellOrders(buyOrder)

	for _, o := range match {
		fmt.Printf("%+v\n", o)
	}

	assert.Equal(t, 1, len(match), "should be 1 match sell order")
}

func TestOrderBookMatchBuyOrder(t *testing.T) {
	now := time.Now().UTC()
	book := NewOrderBook("test-btc")
	order0 := &protoOrder.Order{
		OrderID:    uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	order1 := &protoOrder.Order{
		OrderID:    "#1",
		MarketName: "test-btc",
		Price:      0.01,
		Size:       1.2,
		Side:       "buy",
		CreatedOn:  now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:    "#2",
		MarketName: "test-btc",
		Price:      0.007,
		Size:       0.2,
		Side:       "buy",
		CreatedOn:  now.Add(time.Second * 1).String(),
	}
	order3 := &protoOrder.Order{
		OrderID:    "#4",
		MarketName: "test-btc",
		Price:      0.007,
		Size:       2.7,
		Side:       "buy",
		CreatedOn:  now.Add(time.Second * 20).String(),
	}
	order4 := &protoOrder.Order{
		OrderID:    "#3",
		MarketName: "test-btc",
		Price:      0.007,
		Size:       0.9,
		Side:       "buy",
		CreatedOn:  now.Add(time.Second * 2).String(),
	}
	order5 := &protoOrder.Order{
		OrderID:    "#0",
		MarketName: "test-btc",
		Price:      0.00034,
		Size:       0.9,
		Side:       "buy",
		CreatedOn:  now.Add(time.Second * 100).String(),
	}

	orders := []*protoOrder.Order{order0, order1, order2, order3, order4, order5}
	for _, o := range orders {
		book.AddBuyOrder(o)
	}

	sellOrder := &protoOrder.Order{
		OrderID:    "#sell",
		MarketName: "test-btc",
		Price:      0.007,
		Size:       1.9,
		Side:       "sell",
	}
	match := book.MatchBuyOrders(sellOrder)

	//for _, o := range match {
	//	fmt.Printf("%+v\n", o)
	//}

	assert.Equal(t, 3, len(match), "should be 3 matched buy orders")
}
