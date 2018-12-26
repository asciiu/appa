package models

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	//example := []float64{0.04, 0.02, 0.03, 0.01, 0.007}
	order1 := Order{
		OrderID: "#1",
		Price:   0.04,
		Size:    1.2,
		Side:    "buy",
	}
	order2 := Order{
		OrderID: "#2",
		Price:   0.007,
		Size:    0.2,
		Side:    "buy",
	}
	order3 := Order{
		OrderID: "#3",
		Price:   0.03,
		Size:    2.7,
		Side:    "buy",
	}
	orders := []Order{order1, order2, order3}
	sorted := MergeSort(orders)

	for _, order := range sorted {
		fmt.Printf("%+v\n", order)
	}
	//assert.Equal(t, 1, len(book.BuyQ), "should be 1 order in buys")
	//assert.Equal(t, 0, len(book.SellQ), "should be 0 order in sells")
}
