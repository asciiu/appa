package models

import (
	"testing"
)

func TestOrderBook(t *testing.T) {
	// book := NewOrderBook("test-btc", "buy")
	// order := protoOrder.Order{
	// 	OrderID:    uuid.New().String(),
	// 	UserID:     uuid.New().String(),
	// 	MarketName: "test-btc",
	// 	Side:       constOrder.Buy,
	// 	Size:       1,
	// 	Price:      0.01,
	// 	Type:       constOrder.LimitOrder,
	// }
	//book.AddOrder(&order)

	//assert.Equal(t, 1, len(book.Buys), "should be 1 order in buys")
	//assert.Equal(t, 0, len(book.Sells), "should be 0 order in sells")
}

func TestOrderBookWrongOrder(t *testing.T) {
	// book := NewOrderBook("test-btc")
	// order := protoOrder.Order{
	// 	OrderID:    uuid.New().String(),
	// 	UserID:     uuid.New().String(),
	// 	MarketName: "bch-btc",
	// 	Side:       constOrder.Buy,
	// 	Size:       1,
	// 	Price:      0.01,
	// 	Type:       constOrder.LimitOrder,
	// }
	// book.AddOrder(&order)

	// assert.Equal(t, 0, len(book.Buys), "should be 0 order in buys")
	// assert.Equal(t, 0, len(book.Sells), "should be 0 order in sells")
}
