package models

import (
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
		Amount:     1,
		Price:      0.01,
	}
	book.AddOrder(&order)

	assert.Equal(t, 1, len(book.BuyOrders), "should be 1 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
}

func TestOrderBookWrongMarketName(t *testing.T) {
	book := NewOrderBook("test-btc")
	order := protoOrder.Order{
		OrderID:    uuid.New().String(),
		MarketName: "bch-btc",
		Side:       constOrder.Buy,
		Amount:     1,
		Price:      0.01,
	}
	book.AddBuyOrder(&order)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
}

// Order book should reject adding a sell order when the order side is buy
func TestOrderBookRejectBuyOrder(t *testing.T) {
	book := NewOrderBook("test-btc")
	order := protoOrder.Order{
		OrderID:    uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Amount:     1,
		Price:      0.01,
	}
	book.AddSellOrder(&order)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
}

func TestOrderBookRejectSellOrder(t *testing.T) {
	book := NewOrderBook("test-btc")
	order := protoOrder.Order{
		OrderID:    uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Amount:     1,
		Price:      0.01,
	}
	book.AddBuyOrder(&order)

	assert.Equal(t, 0, len(book.BuyOrders), "should be 0 order in buys")
	assert.Equal(t, 0, len(book.SellOrders), "should be 0 order in sells")
}

func TestOrderBookMatchSellOrder(t *testing.T) {
	now := time.Now().UTC()
	book := NewOrderBook("test-btc")
	order0 := &protoOrder.Order{
		OrderID:    "#0",
		MarketName: "test-btc",
		Side:       constOrder.Sell,
		Amount:     1,
		Price:      0.01,
		CreatedOn:  now.String(),
	}
	order1 := &protoOrder.Order{
		OrderID:    "#1",
		MarketName: "test-btc",
		Price:      0.01,
		Amount:     1.2,
		Side:       constOrder.Sell,
		CreatedOn:  now.Add(time.Second * 10).String(),
	}
	order2 := &protoOrder.Order{
		OrderID:    "#2",
		MarketName: "test-btc",
		Price:      0.007,
		Amount:     0.2,
		Side:       constOrder.Sell,
		CreatedOn:  now.Add(time.Second * 1).String(),
	}

	orders := []*protoOrder.Order{order0, order1, order2}
	for _, o := range orders {
		book.AddOrder(o)
	}

	//for _, o := range book.SellOrders {
	//	fmt.Printf("%+v\n", o)
	//}

	buyOrder := &protoOrder.Order{
		OrderID:    "#buyer",
		MarketName: "test-btc",
		Price:      0.01,
		Amount:     0.9,
		Side:       "buy",
	}
	//fmt.Printf("buy order %+v\n", buyOrder)

	matches := book.FillSellOrders(buyOrder)
	//for _, o := range matches {
	//	fmt.Printf("%+v\n", o)
	//}
	//for _, o := range book.SellOrders {
	//	fmt.Printf("%+v\n", o)
	//}

	assert.Equal(t, 2, len(matches), "should be 2 matched sell orders")
	assert.Equal(t, 2, len(book.SellOrders), "should be 2 sell orders")
	assert.Equal(t, 0.2, matches[0].Fill, "fill for 1st match should be 0.2")
	assert.Equal(t, 0.7, matches[1].Fill, "fill for 2nd match should be 0.7")

	book.CancelSellOrder(order2)
	i := book.FindSellOrder(order2)

	assert.Equal(t, -1, i, "cancelled order should not be found")
}

func TestOrderBookMatchBuyOrder(t *testing.T) {
	now := time.Now().UTC()
	book := NewOrderBook("test-btc")
	order0 := &protoOrder.Order{
		OrderID:    "#0",
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Amount:     1.0,
		Price:      0.01,
	}
	order1 := &protoOrder.Order{
		OrderID:    "#1",
		MarketName: "test-btc",
		Price:      0.01,
		Amount:     1.2,
		Side:       "buy",
		CreatedOn:  now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:    "#2",
		MarketName: "test-btc",
		Price:      0.007,
		Amount:     0.2,
		Side:       "buy",
		CreatedOn:  now.Add(time.Second * 1).String(),
	}
	order3 := &protoOrder.Order{
		OrderID:    "#4",
		MarketName: "test-btc",
		Price:      0.007,
		Amount:     2.7,
		Side:       "buy",
		CreatedOn:  now.Add(time.Second * 20).String(),
	}
	order4 := &protoOrder.Order{
		OrderID:    "#3",
		MarketName: "test-btc",
		Price:      0.007,
		Amount:     0.9,
		Side:       "buy",
		CreatedOn:  now.Add(time.Second * 2).String(),
	}
	order5 := &protoOrder.Order{
		OrderID:    "#5",
		MarketName: "test-btc",
		Price:      0.00034,
		Amount:     0.9,
		Side:       "buy",
		CreatedOn:  now.Add(time.Second * 100).String(),
	}

	orders := []*protoOrder.Order{order0, order1, order2, order3, order4, order5}
	for _, o := range orders {
		book.AddOrder(o)
	}
	//for _, o := range book.BuyOrders {
	//	fmt.Printf("%+v\n", o)
	//}

	sellOrder := &protoOrder.Order{
		OrderID:    "#sell",
		MarketName: "test-btc",
		Price:      0.007,
		Amount:     1.9,
		Side:       "sell",
	}
	//fmt.Printf("sell order %+v\n", sellOrder)
	matches := book.FillBuyOrders(sellOrder)

	//for _, o := range matches {
	//	fmt.Printf("%+v\n", o)
	//}

	assert.Equal(t, 2, len(matches), "should be 2 matched buy orders")
	assert.Equal(t, 1.2, matches[0].Fill, "fill for 1st match should be 1.2")
	assert.Equal(t, 0.7, matches[1].Fill, "fill for 2nd match should be 0.7")
	assert.Equal(t, 5, len(book.BuyOrders), "should be 5 buy orders")

	book.CancelBuyOrder(order1)
	i := book.FindBuyOrder(order1)

	assert.Equal(t, -1, i, "cancelled order should not be found")
}
