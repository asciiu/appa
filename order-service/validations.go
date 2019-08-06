package main

import (
	constOrder "github.com/asciiu/appa/order-service/constants"
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
	case constOrder.LimitOrder:
		return true
	default:
		return false
	}
}
