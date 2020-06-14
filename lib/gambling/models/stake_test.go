package models

import (
	"fmt"
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
		sells := []Stake{
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

		//size := bet.Amount / bet.Odds
		fmt.Println(sells[0])
		fmt.Println(sells[1])

		assert.Equal(t, true, true)
	})
}
