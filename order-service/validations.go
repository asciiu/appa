package main

import (
	constOrder "github.com/asciiu/oldiez/order-service/constants"
)

// plans must have a title
func ValidateSide(side string) bool {
	switch side {
	case constOrder.Buy:
		return true
	case constOrder.Sell:
		return true
	default:
		return false
	}
}
