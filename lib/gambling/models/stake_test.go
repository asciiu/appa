package models

import (
	"log"
	"testing"

	"gotest.tools/assert"
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func TestStake(t *testing.T) {
	t.Run("Wager basic", func(t *testing.T) {
		stakes := []Stake{
			Stake{
				ID:     1,
				Odds:   2.35,
				Amount: 40,
				Side:   "sell",
			},
			Stake{
				ID:     2,
				Odds:   3.33,
				Amount: 20,
				Side:   "sell",
			},
		}

		assert.Equal(t, float64(17.02), stakes[0].Payout(), "payout wrong")
		assert.Equal(t, float64(6), stakes[1].Payout(), "payout wrong")
	})
}
