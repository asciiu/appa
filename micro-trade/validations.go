package main

import (
	"github.com/asciiu/appa/micro-trade/constants"
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
