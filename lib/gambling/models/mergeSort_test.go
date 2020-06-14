package models

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	now := time.Now().UTC()

	o1 := Stake{
		ID:        1,
		Odds:      2.34,
		Amount:    1.2,
		CreatedAt: now.Add(1 * time.Hour),
	}
	o2 := Stake{
		ID:        2,
		Odds:      1.1,
		Amount:    0.2,
		CreatedAt: now.Add(1 * time.Minute),
	}
	o3 := Stake{
		ID:        3,
		Odds:      0.7,
		Amount:    2.7,
		CreatedAt: now.Add(2 * time.Minute),
	}
	o4 := Stake{
		ID:        4,
		Odds:      2.2,
		Amount:    0.9,
		CreatedAt: now,
	}
	o5 := Stake{
		ID:        5,
		Odds:      2.2,
		Amount:    100.0,
		CreatedAt: now.Add(1 * time.Millisecond),
	}

	orders := []Stake{o1, o2, o3, o4, o5}
	sorted := MergeSort(orders)

	for _, order := range sorted {
		fmt.Printf("%+v\n", order)
	}

	t.Run("sort order", func(t *testing.T) {
		assert.Equal(t, 5, len(sorted), "should be 5 sorted orders")
		assert.Equal(t, float64(0.2), sorted[1].Amount, "order 2 size did not match")
		assert.Equal(t, 5, sorted[3].ID, "order 3 order ID did not match")
		assert.Equal(t, float64(2.7), sorted[0].Amount, "order 1 size did not match")
	})

	t.Run("search less than", func(t *testing.T) {
		odds := 2.01
		index := searchLessThan(sorted, odds)
		assert.Equal(t, 2, index, "less than should be index 1")
	})

	t.Run("search greater than", func(t *testing.T) {
		odds := 2.01
		index := searchGreaterThan(sorted, odds)
		assert.Equal(t, 2, index, "greater than should be index 2")
	})

	t.Run("find order", func(t *testing.T) {
		index := FindOrder(sorted, o4)
		assert.Equal(t, 2, index, "found index should be 2")
	})

	t.Run("not find order", func(t *testing.T) {
		o6 := Stake{ID: 6}
		index := FindOrder(sorted, o6)
		assert.Equal(t, -1, index, "index should be -1")
	})

	t.Run("binary search", func(t *testing.T) {
		searchOdds := 3.3
		index := binarySearch(sorted, searchOdds)
		assert.Equal(t, 4, index, "should return index of last item")
	})
}
