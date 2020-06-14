package models

import (
	"fmt"
	"math"
	"time"
)

type Stake struct {
	ID        int
	Odds      float64
	Amount    float64
	Side      string
	CreatedAt time.Time
}

func (stake Stake) String() string {
	return fmt.Sprintf("id: %d play: %.2f to win: %.2f odds@%.2f", stake.ID, stake.Payout(), stake.Amount, stake.Odds)
}

func (stake Stake) Payout() float64 {
	// payout amount using odds and wagered amount
	payout := stake.Amount / stake.Odds

	return math.Floor(payout*100) / 100
}
