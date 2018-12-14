package models

type OrderBook struct {
	MarketName string
	BuyOrders  []*Order
	SellOrders []*Order
}

type Order struct {
	UserID string
	Side   string
	Size   float64
}
