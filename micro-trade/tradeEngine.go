package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/asciiu/appa/lib/constants/response"
	"github.com/asciiu/appa/micro-trade/constants"
	repo "github.com/asciiu/appa/micro-trade/db/sql"
	"github.com/asciiu/appa/micro-trade/models"
	"github.com/asciiu/appa/micro-trade/proto/trade"
)

type TradeEngine struct {
	DB         *sql.DB
	OrderBooks map[string]*models.OrderBook
}

func NewTradeEngine(db *sql.DB) *TradeEngine {
	return &TradeEngine{
		DB:         db,
		OrderBooks: make(map[string]*models.OrderBook),
	}
}

func (service *TradeEngine) Process(ctx context.Context, req *trade.NewOrderRequest, res *trade.OrderResponse) error {
	newOrder := models.NewOrder(
		req.UserID,
		req.MarketName,
		req.Side,
		req.Amount,
		req.Price,
	)

	// validate the order
	switch {
	case !ValidateSide(newOrder.Side):
		res.Status = response.Fail
		res.Message = "side must be buy or sell"
		return nil
	case !ValidateType(newOrder.Type):
		res.Status = response.Fail
		res.Message = "type must be limit"
		return nil
	}

	book, ok := service.OrderBooks[newOrder.MarketName]
	if !ok {
		// new order book for market name
		book = models.NewOrderBook(newOrder.MarketName)
		service.OrderBooks[newOrder.MarketName] = book
	}
	// process the new order
	filledOrders, trades := book.Process(newOrder)

	if err := repo.InsertOrder(service.DB, newOrder); err != nil {
		msg := fmt.Sprintf("insert order failed %s", err.Error())
		log.Println(msg)

		res.Status = response.Error
		res.Message = msg
		return nil
	}

	for _, trade := range trades {
		if err := repo.InsertTrade(service.DB, trade); err != nil {
			msg := fmt.Sprintf("insert order failed %s", err.Error())
			log.Println(msg)

			res.Status = response.Error
			res.Message = msg
			return nil
		}
	}

	for _, filled := range filledOrders {
		if err := repo.UpdateOrder(service.DB, filled); err != nil {
			msg := fmt.Sprintf("update order failed %s", err.Error())
			log.Println(msg)

			res.Status = response.Error
			res.Message = msg
			return nil
		}
	}

	res.Status = response.Success
	res.Data = &trade.OrderData{
		Order: &trade.Order{
			OrderID:    newOrder.ID,
			UserID:     newOrder.UserID,
			MarketName: newOrder.MarketName,
			Side:       newOrder.Side,
			Amount:     newOrder.Amount,
			Filled:     newOrder.Filled,
			Price:      newOrder.Price,
			Status:     newOrder.Status,
			CreatedOn:  newOrder.CreatedOn,
			UpdatedOn:  newOrder.UpdatedOn,
		},
	}
	return nil
}

func (service *TradeEngine) Cancel(ctx context.Context, req *trade.OrderRequest, res *trade.StatusResponse) error {
	order, err := repo.FindOrderByID(service.DB, req.OrderID)
	switch {
	case err == sql.ErrNoRows:
		res.Status = response.Nonentity
		res.Message = fmt.Sprintf("OrderID not found %s", req.OrderID)
	case err != nil:
		res.Status = response.Error
		res.Message = err.Error()
	case err == nil:
		order.Status = constants.Cancelled
		if err := repo.UpdateOrder(service.DB, order); err != nil {
			msg := fmt.Sprintf("update order failed %s", err.Error())
			log.Println(msg)

			res.Status = response.Error
			res.Message = msg
			return nil
		}

		book, ok := service.OrderBooks[order.MarketName]
		if !ok {
			res.Status = response.Error
			res.Message = fmt.Sprintf("order book not found for %s", order.MarketName)
			return nil
		}

		if err = book.Cancel(order); err != nil {
			res.Status = response.Error
			res.Message = fmt.Sprintf("could not cancel order %s", err.Error())
			return nil
		}

		res.Status = response.Success
	}

	return nil
}

func (service *TradeEngine) FindOrder(ctx context.Context, req *trade.OrderRequest, res *trade.OrderResponse) error {
	order, err := repo.FindOrderByID(service.DB, req.OrderID)

	switch {
	case err == sql.ErrNoRows:
		res.Status = response.Nonentity
		res.Message = fmt.Sprintf("OrderID not found %s", req.OrderID)
	case err != nil:
		res.Status = response.Error
		res.Message = err.Error()
	case err == nil:
		res.Status = response.Success
		res.Data = &trade.OrderData{
			Order: &trade.Order{
				OrderID:    order.ID,
				UserID:     order.UserID,
				MarketName: order.MarketName,
				Side:       order.Side,
				Amount:     order.Amount,
				Price:      order.Price,
				Status:     order.Status,
				CreatedOn:  order.CreatedOn,
				UpdatedOn:  order.UpdatedOn,
			},
		}
	}

	return nil
}

func (service *TradeEngine) FindUserOrders(ctx context.Context, req *trade.UserOrdersRequest, res *trade.OrdersPageResponse) error {
	page, err := repo.FindUserOrders(service.DB, req.UserID, req.Status, req.Page, req.PageSize)

	if err == nil {
		orders := make([]*trade.Order, len(page.Orders))
		for _, o := range page.Orders {
			orders = append(orders, &trade.Order{
				OrderID:    o.ID,
				UserID:     o.UserID,
				MarketName: o.MarketName,
				Side:       o.Side,
				Amount:     o.Amount,
				Price:      o.Price,
				Status:     o.Status,
				CreatedOn:  o.CreatedOn,
				UpdatedOn:  o.UpdatedOn,
			})
		}

		res.Status = response.Success
		res.Data = &trade.OrdersPage{
			Page:     page.Page,
			PageSize: page.PageSize,
			Total:    page.Total,
			Orders:   orders,
		}
	} else {
		res.Status = response.Error
		res.Message = err.Error()
	}

	return nil
}
