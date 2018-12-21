package models

import (
	"testing"

	constOrder "github.com/asciiu/appa/order-service/constants"
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestOrderBook(t *testing.T) {
	book := NewOrderBook("test-btc")
	order := protoOrder.Order{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MarketName: "test-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	book.AddOrder(&order)

	assert.Equal(t, 1, len(book.BuyQ), "should be 1 order in buys")
	assert.Equal(t, 0, len(book.SellQ), "should be 0 order in sells")
}

func TestOrderBookWrongOrder(t *testing.T) {
	book := NewOrderBook("test-btc")
	order := protoOrder.Order{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MarketName: "bch-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	book.AddOrder(&order)

	assert.Equal(t, 0, len(book.BuyQ), "should be 0 order in buys")
	assert.Equal(t, 0, len(book.SellQ), "should be 0 order in sells")
}
