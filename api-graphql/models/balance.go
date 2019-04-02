package models

import (
	"github.com/google/uuid"
)

func NewBalance(userID, symbol, name, address string, amount, locked int64, precision int) *Balance {
	newID := uuid.New()

	balance := Balance{
		ID:        newID.String(),
		UserID:    userID,
		Symbol:    symbol,
		Name:      name,
		Amount:    amount,
		Locked:    locked,
		Precision: precision,
		Address:   address,
	}
	return &balance
}

type Balance struct {
	ID        string
	UserID    string
	Symbol    string
	Name      string
	Amount    int64
	Locked    int64
	Precision int
	Address   string
}

type Currency struct {
	Symbol string
	Name   string
}
