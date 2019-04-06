package models

import (
	"fmt"
	"io"
	"strconv"

	"github.com/google/uuid"
)

func NewBalance(userID, symbol, name, address string, amount, locked int64) *Balance {
	newID := uuid.New()

	balance := Balance{
		ID:      newID.String(),
		UserID:  userID,
		Symbol:  symbol,
		Name:    name,
		Amount:  Int64(amount),
		Locked:  Int64(locked),
		Address: address,
	}
	return &balance
}

type Balance struct {
	ID        string
	UserID    string
	Symbol    string
	Name      string
	Amount    Int64
	Locked    Int64
	Precision float64
	Address   string
}

type Currency struct {
	Symbol    string
	Name      string
	Precision float64
}

type Int64 int64

func (i *Int64) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("must be strings")
	}

	n, err := strconv.ParseInt(str, 10, 64)
	if err == nil {
		fmt.Printf("%d of type %T", n, n)
	}

	*i = Int64(n)
	return nil
}

// MarshalGQL implements the graphql.Marshaler interface
func (i Int64) MarshalGQL(w io.Writer) {
	s := strconv.FormatInt(int64(i), 10)
	w.Write([]byte(s))
}
