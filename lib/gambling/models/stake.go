package models

import "fmt"

type Stake struct {
	ID     int
	Odds   float32
	Amount float32
	Side   string
}

func (stake Stake) String() string {
	return fmt.Sprintf("play: %.2f to win: %.2f odds@%.2f", stake.payout(), stake.Amount, stake.Odds)
}

func (stake Stake) payout() float32 {
	// payout amount using odds and wagered amount
	return stake.Amount / stake.Odds
}
