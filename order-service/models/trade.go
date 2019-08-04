package models

type Trade struct {
	TakerOrderID string `json:"taker_order_id"`
	MakerOrderID string `json:"maker_order_id"`
	Amount       uint64 `json:"amount"`
	Price        uint64 `json:"price"`
	Side         string `json:"side"`
}
