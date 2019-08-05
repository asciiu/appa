package models

type Order struct {
	ID         string `json:"id"`
	UserID     string `json:"userID"`
	MarketName string `json:"marketName"`
	Amount     uint64 `json:"amount"`
	Price      uint64 `json:"price"`
	Side       string `json:"side"`
	Status     string `json:"status"`
	filled     string `json:"filled"`
	CreateOn   string `json:"createdOn"`
	UpdatedOn  string `json:"createdOn"`
}

type Trade struct {
	TakerOrderID string `json:"taker_order_id"`
	MakerOrderID string `json:"maker_order_id"`
	Amount       uint64 `json:"amount"`
	Price        uint64 `json:"price"`
	Side         string `json:"side"`
}
