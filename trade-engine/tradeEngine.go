package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/asciiu/appa/common/constants/response"
	tradeRepo "github.com/asciiu/appa/trade-engine/db/sql"
	"github.com/asciiu/appa/trade-engine/models"
	"github.com/asciiu/appa/trade-engine/proto/trade"
)

type TradeEngine struct {
	DB         *sql.DB
	OrderBooks map[string]*models.OrderBook
}

func (service *TradeEngine) AddOrder(ctx context.Context, req *trade.NewOrderRequest, res *trade.OrderResponse) error {
	newOrder := models.NewOrder(req.UserID, req.MarketName, req.Side, req.Amount, req.Price)
	//now := string(pq.FormatTimestamp(time.Now().UTC()))

	// newOrder := trade.Order{
	// 	OrderID:    uuid.New().String(),
	// 	UserID:     req.UserID,
	// 	MarketName: req.MarketName,
	// 	Side:       req.Side,
	// 	Amount:     req.Amount,
	// 	Price:      req.Price,
	// 	Status:     constants.Pending,
	// 	CreatedOn:  now,
	// 	UpdatedOn:  now,
	// }

	if err := tradeRepo.InsertOrder(service.DB, newOrder); err != nil {
		msg := fmt.Sprintf("insert order failed %s", err.Error())
		log.Println(msg)

		res.Status = response.Error
		res.Message = msg
		return nil
	}

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

	//if book, ok := service.OrderBooks[newOrder.MarketName]; ok {
	//	filledOrders := book.FillOrders(&newOrder)
	//	sumFilled := 0.0
	//	for _, f := range filledOrders {
	//		sumFilled += f.Fill
	//	}
	//	newOrder.Fill = sumFilled

	//	if sumFilled < newOrder.Size {
	//		newOrder.Size -= sumFilled
	//	} else if sumFilled == newOrder.Size {
	//		newOrder.Size = 0.0
	//		newOrder.Status = constants.Filled
	//	}
	//	book.AddOrder(&newOrder)
	//} else {
	//	newOrderBook := models.NewOrderBook(newOrder.MarketName)
	//	newOrderBook.AddOrder(&newOrder)
	//	service.OrderBooks[newOrder.MarketName] = newOrderBook
	//}

	res.Status = response.Success
	res.Data = &trade.OrderData{
		Order: &trade.Order{
			OrderID:    newOrder.ID,
			UserID:     newOrder.UserID,
			MarketName: newOrder.MarketName,
			Side:       newOrder.Side,
			Amount:     newOrder.Amount,
			Price:      newOrder.Price,
			Status:     newOrder.Status,
			CreatedOn:  newOrder.CreatedOn,
			UpdatedOn:  newOrder.UpdatedOn,
		},
	}
	return nil
}

func (service *TradeEngine) CancelOrder(ctx context.Context, req *trade.OrderRequest, res *trade.StatusResponse) error {
	_, err := tradeRepo.FindOrderByID(service.DB, req.OrderID)
	switch {
	case err == sql.ErrNoRows:
		res.Status = response.Nonentity
		res.Message = fmt.Sprintf("OrderID not found %s", req.OrderID)
	case err != nil:
		res.Status = response.Error
		res.Message = err.Error()
		//case err == nil:
		//	if err := tradeRepo.DeleteOrder(service.DB, req.OrderID, req.UserID); err == nil {
		//		book := service.OrderBooks[order.MarketName]
		//		book.CancelOrder(order)
		//		res.Status = response.Success

		//	} else {
		//		res.Status = response.Error
		//		res.Message = err.Error()
		//	}
	}

	return nil
}

func (service *TradeEngine) FindOrder(ctx context.Context, req *trade.OrderRequest, res *trade.OrderResponse) error {
	_, err := tradeRepo.FindOrderByID(service.DB, req.OrderID)

	switch {
	case err == sql.ErrNoRows:
		res.Status = response.Nonentity
		res.Message = fmt.Sprintf("OrderID not found %s", req.OrderID)
	case err != nil:
		res.Status = response.Error
		res.Message = err.Error()
	case err == nil:
		res.Status = response.Success
		//	res.Data = &trade.OrderData{
		//		Order: order,
		//	}
	}

	return nil
}

func (service *TradeEngine) FindUserOrders(ctx context.Context, req *trade.UserOrdersRequest, res *trade.OrdersPageResponse) error {
	_, err := tradeRepo.FindUserOrders(service.DB, req.UserID, req.Status, req.Page, req.PageSize)

	if err == nil {
		res.Status = response.Success
		//res.Data = trades.OrderPage{
		//}
		//	res.Data = ordersPage
	} else {
		res.Status = response.Error
		res.Message = err.Error()
	}

	return nil
}
