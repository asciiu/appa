package models

import (
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

type OrderQueue struct {
	Price float64
	// fifo
	Orders []*protoOrder.Order
}

func NewOrderQueue(price float64) *OrderQueue {
	return &OrderQueue{
		Price:  price,
		Orders: make([]*protoOrder.Order, 0),
	}
}

func (queue *OrderQueue) AddOrder(order *protoOrder.Order) {
	if order.Price != queue.Price || queue.IsQueued(order.OrderID) {
		return
	}
	queue.Orders = append(queue.Orders, order)
}

func (queue *OrderQueue) IsQueued(orderID string) bool {
	// cannot add order if already existing
	for _, o := range queue.Orders {
		if o.OrderID == orderID {
			return true
		}
	}
	return false
}

func (queue *OrderQueue) RemoveOrder(orderID string) *protoOrder.Order {
	orders := queue.Orders
	for i, order := range orders {
		if order.OrderID == orderID {
			queue.Orders = append(orders[:i], orders[i+1:]...)
			return order
		}
	}
	return nil
}

func (queue *OrderQueue) Pop() *protoOrder.Order {
	orders := queue.Orders
	order := orders[0]
	queue.Orders = orders[1:]
	return order
}
