package models

import (
	"testing"

	constOrder "github.com/asciiu/appa/order-service/constants"
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestQueue(t *testing.T) {
	queue := NewOrderQueue(0.01)
	order := protoOrder.Order{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MarketName: "bch-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	queue.AddOrder(&order)

	assert.Equal(t, 1, len(queue.Orders), "should be 1 order in queue")
}

func TestPopQueue(t *testing.T) {
	queue := NewOrderQueue(0.01)
	order := protoOrder.Order{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MarketName: "bch-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	queue.AddOrder(&order)
	order2 := queue.Pop()

	assert.Equal(t, order.OrderID, order2.OrderID, "order IDs do not match")
	assert.Equal(t, order.UserID, order2.UserID, "user IDs do not match")
	assert.Equal(t, order.Size, order2.Size, "sizes do not match")
	assert.Equal(t, 0, len(queue.Orders), "queue should be empty")
}

func TestUniqueQueue(t *testing.T) {
	queue := NewOrderQueue(0.01)
	order := protoOrder.Order{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MarketName: "bch-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	// cannot add same order twice
	queue.AddOrder(&order)
	queue.AddOrder(&order)

	assert.Equal(t, 1, len(queue.Orders), "queue should be empty")
}

func TestRemoveQueue(t *testing.T) {
	queue := NewOrderQueue(0.01)
	order1 := protoOrder.Order{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MarketName: "bch-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	order2 := protoOrder.Order{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MarketName: "bch-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}
	order3 := protoOrder.Order{
		OrderID:    uuid.New().String(),
		UserID:     uuid.New().String(),
		MarketName: "bch-btc",
		Side:       constOrder.Buy,
		Size:       1,
		Price:      0.01,
		Type:       constOrder.LimitOrder,
	}

	// cannot add same order twice
	queue.AddOrder(&order1)
	queue.AddOrder(&order2)
	queue.AddOrder(&order3)

	assert.Equal(t, 3, len(queue.Orders), "queue should be empty")

	queue.RemoveOrder(order2.OrderID)

	assert.Equal(t, 2, len(queue.Orders), "queue should be empty")
}
