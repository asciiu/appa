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

	assert.Equal(t, 5, len(sorted), "should be 5 sorted orders")
	assert.Equal(t, float64(0.2), sorted[1].Amount, "order 2 size did not match")
	assert.Equal(t, 5, sorted[3].ID, "order 3 order ID did not match")
	assert.Equal(t, float64(2.7), sorted[0].Amount, "order 1 size did not match")
}

// func TestSearchLessThan(t *testing.T) {
// 	now := time.Now().UTC()
// 	order1 := Stake{
// 		OrderID:   "#1",
// 		Price:     0.01,
// 		Size:      1.2,
// 		Side:      "buy",
// 		CreatedOn: now.String(),
// 	}
// 	order2 := Stake{
// 		OrderID:   "#2",
// 		Price:     0.0081,
// 		Size:      0.2,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 1).String(),
// 	}
// 	order3 := Stake{
// 		OrderID:   "#4",
// 		Price:     0.0073,
// 		Size:      2.7,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 20).String(),
// 	}
// 	order4 := Stake{
// 		OrderID:   "#3",
// 		Price:     0.0072,
// 		Size:      0.9,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 2).String(),
// 	}
// 	order5 := Stake{
// 		OrderID:   "#0",
// 		Price:     0.00034,
// 		Size:      0.9,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 100).String(),
// 	}

// 	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
// 	sorted := MergeSort(orders)

// 	buyPrice := 0.00735
// 	// find index where Order.Price <= searchPrice
// 	index := searchLessThan(sorted, buyPrice)
// 	for _, o := range sorted {
// 		fmt.Printf("%+v\n", o)
// 	}
// 	fmt.Println(index)
// 	assert.Equal(t, 2, index, "less than should be index 2")
// }

// func TestSearchIndexGreaterThan(t *testing.T) {
// 	now := time.Now().UTC()
// 	order1 := Stake{
// 		OrderID:   "#1",
// 		Price:     0.01,
// 		Size:      1.2,
// 		Side:      "buy",
// 		CreatedOn: now.String(),
// 	}
// 	order2 := Stake{
// 		OrderID:   "#2",
// 		Price:     0.0081,
// 		Size:      0.2,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 1).String(),
// 	}
// 	order3 := Stake{
// 		OrderID:   "#4",
// 		Price:     0.0073,
// 		Size:      2.7,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 20).String(),
// 	}
// 	order4 := Stake{
// 		OrderID:   "#3",
// 		Price:     0.0072,
// 		Size:      0.9,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 2).String(),
// 	}
// 	order5 := Stake{
// 		OrderID:   "#0",
// 		Price:     0.00034,
// 		Size:      0.9,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 100).String(),
// 	}

// 	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
// 	sorted := MergeSort(orders)

// 	sellPrice := 0.00735
// 	// find index where Order.Price > searchPrice
// 	index := searchGreaterThan(sorted, sellPrice)
// 	for _, o := range sorted {
// 		fmt.Printf("%+v\n", o)
// 	}
// 	fmt.Println(index)
// 	assert.Equal(t, 3, index, "index 3 should be greater than price")
// }

// func TestFindOrder(t *testing.T) {
// 	now := time.Now().UTC()
// 	order1 := Stake{
// 		OrderID:   "#1",
// 		Price:     0.01,
// 		Size:      1.2,
// 		Side:      "buy",
// 		CreatedOn: now.String(),
// 	}
// 	order2 := Stake{
// 		OrderID:   "#2",
// 		Price:     0.007,
// 		Size:      0.2,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 1).String(),
// 	}
// 	order3 := Stake{
// 		OrderID:   "#4",
// 		Price:     0.007,
// 		Size:      2.7,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 20).String(),
// 	}
// 	order4 := Stake{
// 		OrderID:   "#3",
// 		Price:     0.007,
// 		Size:      0.9,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 2).String(),
// 	}
// 	order5 := Stake{
// 		OrderID:   "#0",
// 		Price:     0.00034,
// 		Size:      0.9,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 100).String(),
// 	}

// 	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
// 	sorted := MergeSort(orders)

// 	index := FindOrder(sorted, order4)
// 	assert.Equal(t, 2, index, "index should be 2")
// }

// func TestFindOrderNotFound(t *testing.T) {
// 	now := time.Now().UTC()
// 	order1 := Stake{
// 		OrderID:   "#1",
// 		Price:     0.01,
// 		Size:      1.2,
// 		Side:      "buy",
// 		CreatedOn: now.String(),
// 	}
// 	order2 := Stake{
// 		OrderID:   "#2",
// 		Price:     0.007,
// 		Size:      0.2,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 1).String(),
// 	}
// 	order3 := Stake{
// 		OrderID:   "#4",
// 		Price:     0.007,
// 		Size:      2.7,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 20).String(),
// 	}
// 	order4 := Stake{
// 		OrderID:   "#3",
// 		Price:     0.007,
// 		Size:      0.9,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 2).String(),
// 	}
// 	order5 := Stake{
// 		OrderID:   "#0",
// 		Price:     0.00034,
// 		Size:      0.9,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 100).String(),
// 	}

// 	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
// 	sorted := MergeSort(orders)

// 	order6 := Stake{
// 		OrderID:   "#6",
// 		Price:     0.00034,
// 		Size:      0.9,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 100).String(),
// 	}
// 	index := FindOrder(sorted, order6)
// 	assert.Equal(t, -1, index, "index should be -1")
// }

// func TestBinarySearch(t *testing.T) {
// 	now := time.Now().UTC()
// 	order1 := Stake{
// 		OrderID:   "#1",
// 		Price:     0.01,
// 		Size:      1.2,
// 		Side:      "buy",
// 		CreatedOn: now.String(),
// 	}
// 	order2 := Stake{
// 		OrderID:   "#2",
// 		Price:     0.007,
// 		Size:      0.2,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 1).String(),
// 	}
// 	order3 := Stake{
// 		OrderID:   "#4",
// 		Price:     0.007,
// 		Size:      2.7,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 20).String(),
// 	}
// 	order4 := Stake{
// 		OrderID:   "#3",
// 		Price:     0.007,
// 		Size:      0.9,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 2).String(),
// 	}
// 	order5 := Stake{
// 		OrderID:   "#0",
// 		Price:     0.00034,
// 		Size:      0.9,
// 		Side:      "buy",
// 		CreatedOn: now.Add(time.Second * 100).String(),
// 	}

// 	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
// 	sorted := MergeSort(orders)

// 	searchPrice := 0.011
// 	// index of order where Order.Price <= searchPrice
// 	index := binarySearch(sorted, searchPrice)
// 	//for _, o := range sorted {
// 	//	fmt.Printf("%+v\n", o)
// 	//}
// 	//fmt.Println(index)
// 	assert.Equal(t, 4, index, "should return index of last item")
// }
