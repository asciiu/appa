package main

import (
	"context"
	"database/sql"

	protoOrder "github.com/asciiu/oldiez/order-service/proto/order"
)

type OrderService struct {
	DB *sql.DB
}

func (service *OrderService) AddOrder(ctx context.Context, req *protoOrder.NewOrderRequest, res *protoOrder.OrderResponse) error {
	return nil
}

func (service *OrderService) CancelOrder(ctx context.Context, req *protoOrder.OrderRequest, res *protoOrder.StatusResponse) error {
	return nil
}

func (service *OrderService) FindOrder(ctx context.Context, req *protoOrder.OrderRequest, res *protoOrder.OrderResponse) error {
	return nil
}
