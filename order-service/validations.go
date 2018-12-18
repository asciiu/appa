package main

import (
	constOrder "github.com/asciiu/oldiez/order-service/constants"
)

func ValidateSide(side string) bool {
	switch side {
	case constOrder.Buy, constOrder.Sell:
		return true
	default:
		return false
	}
}

func ValidateType(typ string) bool {
	switch typ {
	case constOrder.LimitOrder, constOrder.MarketOrder:
		return true
	default:
		return false
	}
}
