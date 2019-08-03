package models

import (
	"fmt"
	"testing"
	"time"

	protoOrder "github.com/asciiu/appa/order-service/proto/order"
	"github.com/stretchr/testify/assert"
)

func TestMergeSort(t *testing.T) {
	now := time.Now().UTC()
	//example := []float64{0.04, 0.02, 0.03, 0.01, 0.007}
	order1 := &protoOrder.Order{
		OrderID:   "#1",
		Price:     0.007,
		Amount:    1.2,
		Side:      "buy",
		CreatedOn: now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:   "#2",
		Price:     0.007,
		Amount:    0.2,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 1).String(),
	}
	order3 := &protoOrder.Order{
		OrderID:   "#4",
		Price:     0.007,
		Amount:    2.7,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 20).String(),
	}
	order4 := &protoOrder.Order{
		OrderID:   "#3",
		Price:     0.007,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 2).String(),
	}
	order5 := &protoOrder.Order{
		OrderID:   "#0",
		Price:     0.00034,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 100).String(),
	}

	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
	sorted := MergeSort(orders)

	// for _, order := range sorted {
	// 	fmt.Printf("%+v\n", order)
	// }
	assert.Equal(t, 5, len(sorted), "should be 5 sorted orders")
	assert.Equal(t, 1.2, sorted[1].Amount, "order 2 Amount did not match")
	assert.Equal(t, "#3", sorted[3].OrderID, "order 3 order ID did not match")
	assert.Equal(t, 0.9, sorted[0].Amount, "order 1 Amount did not match")
}

func TestSearchLessThan(t *testing.T) {
	now := time.Now().UTC()
	order1 := &protoOrder.Order{
		OrderID:   "#1",
		Price:     0.01,
		Amount:    1.2,
		Side:      "buy",
		CreatedOn: now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:   "#2",
		Price:     0.0081,
		Amount:    0.2,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 1).String(),
	}
	order3 := &protoOrder.Order{
		OrderID:   "#4",
		Price:     0.0073,
		Amount:    2.7,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 20).String(),
	}
	order4 := &protoOrder.Order{
		OrderID:   "#3",
		Price:     0.0072,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 2).String(),
	}
	order5 := &protoOrder.Order{
		OrderID:   "#0",
		Price:     0.00034,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 100).String(),
	}

	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
	sorted := MergeSort(orders)

	buyPrice := 0.00735
	// find index where Order.Price <= searchPrice
	index := searchLessThan(sorted, buyPrice)
	for _, o := range sorted {
		fmt.Printf("%+v\n", o)
	}
	fmt.Println(index)
	assert.Equal(t, 2, index, "less than should be index 2")
}

func TestSearchIndexGreaterThan(t *testing.T) {
	now := time.Now().UTC()
	order1 := &protoOrder.Order{
		OrderID:   "#1",
		Price:     0.01,
		Amount:    1.2,
		Side:      "buy",
		CreatedOn: now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:   "#2",
		Price:     0.0081,
		Amount:    0.2,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 1).String(),
	}
	order3 := &protoOrder.Order{
		OrderID:   "#4",
		Price:     0.0073,
		Amount:    2.7,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 20).String(),
	}
	order4 := &protoOrder.Order{
		OrderID:   "#3",
		Price:     0.0072,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 2).String(),
	}
	order5 := &protoOrder.Order{
		OrderID:   "#0",
		Price:     0.00034,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 100).String(),
	}

	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
	sorted := MergeSort(orders)

	sellPrice := 0.00735
	// find index where Order.Price > searchPrice
	index := searchGreaterThan(sorted, sellPrice)
	for _, o := range sorted {
		fmt.Printf("%+v\n", o)
	}
	fmt.Println(index)
	assert.Equal(t, 3, index, "index 3 should be greater than price")
}

func TestFindOrder(t *testing.T) {
	now := time.Now().UTC()
	order1 := &protoOrder.Order{
		OrderID:   "#1",
		Price:     0.01,
		Amount:    1.2,
		Side:      "buy",
		CreatedOn: now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:   "#2",
		Price:     0.007,
		Amount:    0.2,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 1).String(),
	}
	order3 := &protoOrder.Order{
		OrderID:   "#4",
		Price:     0.007,
		Amount:    2.7,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 20).String(),
	}
	order4 := &protoOrder.Order{
		OrderID:   "#3",
		Price:     0.007,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 2).String(),
	}
	order5 := &protoOrder.Order{
		OrderID:   "#0",
		Price:     0.00034,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 100).String(),
	}

	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
	sorted := MergeSort(orders)

	index := FindOrder(sorted, order4)
	assert.Equal(t, 2, index, "index should be 2")
}

func TestFindOrderNotFound(t *testing.T) {
	now := time.Now().UTC()
	order1 := &protoOrder.Order{
		OrderID:   "#1",
		Price:     0.01,
		Amount:    1.2,
		Side:      "buy",
		CreatedOn: now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:   "#2",
		Price:     0.007,
		Amount:    0.2,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 1).String(),
	}
	order3 := &protoOrder.Order{
		OrderID:   "#4",
		Price:     0.007,
		Amount:    2.7,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 20).String(),
	}
	order4 := &protoOrder.Order{
		OrderID:   "#3",
		Price:     0.007,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 2).String(),
	}
	order5 := &protoOrder.Order{
		OrderID:   "#0",
		Price:     0.00034,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 100).String(),
	}

	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
	sorted := MergeSort(orders)

	order6 := &protoOrder.Order{
		OrderID:   "#6",
		Price:     0.00034,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 100).String(),
	}
	index := FindOrder(sorted, order6)
	assert.Equal(t, -1, index, "index should be -1")
}

func TestBinarySearch(t *testing.T) {
	now := time.Now().UTC()
	order1 := &protoOrder.Order{
		OrderID:   "#1",
		Price:     0.01,
		Amount:    1.2,
		Side:      "buy",
		CreatedOn: now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:   "#2",
		Price:     0.007,
		Amount:    0.2,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 1).String(),
	}
	order3 := &protoOrder.Order{
		OrderID:   "#4",
		Price:     0.007,
		Amount:    2.7,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 20).String(),
	}
	order4 := &protoOrder.Order{
		OrderID:   "#3",
		Price:     0.007,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 2).String(),
	}
	order5 := &protoOrder.Order{
		OrderID:   "#0",
		Price:     0.00034,
		Amount:    0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 100).String(),
	}

	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
	sorted := MergeSort(orders)

	searchPrice := 0.011
	// index of order where Order.Price <= searchPrice
	index := binarySearch(sorted, searchPrice)
	//for _, o := range sorted {
	//	fmt.Printf("%+v\n", o)
	//}
	//fmt.Println(index)
	assert.Equal(t, 4, index, "should return index of last item")
}
