package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	constRes "github.com/asciiu/appa/common/constants/response"
	"github.com/asciiu/appa/order-service/constants"
	repoOrder "github.com/asciiu/appa/order-service/db/sql"
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type OrderService struct {
	DB *sql.DB
}

func (service *OrderService) AddOrder(ctx context.Context, req *protoOrder.NewOrderRequest, res *protoOrder.OrderResponse) error {
	now := string(pq.FormatTimestamp(time.Now().UTC()))
	newOrder := protoOrder.Order{
		OrderID:    uuid.New().String(),
		UserID:     req.UserID,
		MarketName: req.MarketName,
		Side:       req.Side,
		Size:       req.Size,
		Price:      req.Price,
		Type:       req.Type,
		Status:     constants.Pending,
		CreatedOn:  now,
		UpdatedOn:  now,
	}

	switch {
	case !ValidateSide(newOrder.Side):
		res.Status = constRes.Fail
		res.Message = "side must be buy or sell"
		return nil
	case !ValidateType(newOrder.Type):
		res.Status = constRes.Fail
		res.Message = "type must be limit or market"
		return nil
	}

	if err := repoOrder.InsertOrder(service.DB, &newOrder); err != nil {
		msg := fmt.Sprintf("insert order failed %s", err.Error())
		log.Println(msg)

		res.Status = constRes.Error
		res.Message = msg
		return nil
	}

	res.Status = constRes.Success
	res.Data = &protoOrder.OrderData{
		Order: &newOrder,
	}
	return nil
}

func (service *OrderService) CancelOrder(ctx context.Context, req *protoOrder.OrderRequest, res *protoOrder.StatusResponse) error {
	err := repoOrder.DeleteOrder(service.DB, req.OrderID, req.UserID)
	switch {
	case err == sql.ErrNoRows:
		res.Status = constRes.Nonentity
		res.Message = fmt.Sprintf("OrderID not found %s", req.OrderID)
	case err != nil:
		res.Status = constRes.Error
		res.Message = err.Error()
	case err == nil:
		res.Status = constRes.Success
	}

	return nil
}

func (service *OrderService) FindOrder(ctx context.Context, req *protoOrder.OrderRequest, res *protoOrder.OrderResponse) error {
	order, err := repoOrder.FindOrder(service.DB, req.OrderID, req.UserID)

	switch {
	case err == sql.ErrNoRows:
		res.Status = constRes.Nonentity
		res.Message = fmt.Sprintf("OrderID not found %s", req.OrderID)
	case err != nil:
		res.Status = constRes.Error
		res.Message = err.Error()
	case err == nil:
		res.Status = constRes.Success
		res.Data = &protoOrder.OrderData{
			Order: order,
		}
	}

	return nil
}

func (service *OrderService) FindUserOrders(ctx context.Context, req *protoOrder.UserOrdersRequest, res *protoOrder.OrdersPageResponse) error {
	ordersPage, err := repoOrder.FindUserOrders(service.DB, req.UserID, req.Status, req.Page, req.PageSize)

	if err == nil {
		res.Status = constRes.Success
		res.Data = ordersPage
	} else {
		res.Status = constRes.Error
		res.Message = err.Error()
	}

	return nil
}
