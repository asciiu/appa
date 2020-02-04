package main

import (
	"github.com/asciiu/appa/trade-engine/constants"
)

func ValidateSide(side string) bool {
	switch side {
	case constants.Buy, constants.Sell:
		return true
	default:
		return false
	}
}

func ValidateType(typ string) bool {
	switch typ {
	case constants.LimitOrder:
		return true
	default:
		return false
	}
}
