package models

type OrderBook struct {
	MarketName string
	BuyOrders  []*Order
	SellOrders []*Order
}

type Order struct {
	OrderID    string
	UserID     string
	MarketName string
	Side       string
	Size       float64
	CreatedOn  string
}
