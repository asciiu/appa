package models

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrderBook(t *testing.T) {
	book := NewOrderBook("McGreggor vs Kabib - McGreggor ML")

	now := time.Now().UTC()

	o1 := Stake{
		ID:        1,
		Odds:      2.34,
		Amount:    1.2,
		Side: "buy",
		CreatedAt: now.Add(1 * time.Hour),
	}
	// o2 := Stake{
	// 	ID:        2,
	// 	Odds:      1.17,
	// 	Amount:    0.2,
	// 	CreatedAt: now.Add(1 * time.Minute),
	// }
	// o3 := Stake{
	// 	ID:        3,
	// 	Odds:      1.1,
	// 	Amount:    2.7,
	// 	CreatedAt: now.Add(2 * time.Minute),
	// }
	// o4 := Stake{
	// 	ID:        4,
	// 	Odds:      2.2,
	// 	Amount:    0.9,
	// 	CreatedAt: now,
	// }
	// o5 := Stake{
	// 	ID:        5,
	// 	Odds:      2.2,
	// 	Amount:    100.0,
	// 	CreatedAt: now.Add(1 * time.Millisecond),
	// }

	t.Run("add order", func(t *testing.T) {
		book.AddStake(o1)

		assert.Equal(t, 1, len(book.buys), "should be 1 order in buys")
		assert.Equal(t, 0, len(book.sells), "should be 0 order in sells")
	})
}
