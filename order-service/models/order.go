package models

import (
	"time"

	"github.com/asciiu/appa/order-service/constants"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Order struct {
	ID         string `json:"id"`
	UserID     string `json:"userID"`
	MarketName string `json:"marketName"`
	Side       string `json:"side"`
	Amount     uint64 `json:"amount"`
	Filled     uint64 `json:"filled"`
	Price      uint64 `json:"price"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	CreatedOn  string `json:"createdOn"`
	UpdatedOn  string `json:"updatedOn"`
}

type OrdersPage struct {
	Page     uint32
	PageSize uint32
	Total    uint32
	Orders   []*Order
}

func NewOrder(userID, marketName, side string, amount, price uint64) *Order {
	now := string(pq.FormatTimestamp(time.Now().UTC()))
	return &Order{
		ID:         uuid.New().String(),
		UserID:     userID,
		MarketName: marketName,
		Side:       side,
		Amount:     amount,
		Price:      price,
		Status:     constants.Pending,
		Type:       constants.LimitOrder,
		CreatedOn:  now,
		UpdatedOn:  now,
	}
}

type Trade struct {
	TakerOrderID string `json:"taker_order_id"`
	MakerOrderID string `json:"maker_order_id"`
	Amount       uint64 `json:"amount"`
	Price        uint64 `json:"price"`
	Side         string `json:"side"`
}
