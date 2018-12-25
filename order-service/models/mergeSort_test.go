package models

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	example := []float64{0.04, 0.02, 0.03, 0.01, 0.007}
	sorted := MergeSort(example)

	fmt.Println(sorted)
	//assert.Equal(t, 1, len(book.BuyQ), "should be 1 order in buys")
	//assert.Equal(t, 0, len(book.SellQ), "should be 0 order in sells")
}
