package models

import (
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
		Size:      1.2,
		Side:      "buy",
		CreatedOn: now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:   "#2",
		Price:     0.007,
		Size:      0.2,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 1).String(),
	}
	order3 := &protoOrder.Order{
		OrderID:   "#4",
		Price:     0.007,
		Size:      2.7,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 20).String(),
	}
	order4 := &protoOrder.Order{
		OrderID:   "#3",
		Price:     0.007,
		Size:      0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 2).String(),
	}
	order5 := &protoOrder.Order{
		OrderID:   "#0",
		Price:     0.00034,
		Size:      0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 100).String(),
	}

	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
	sorted := MergeSort(orders)

	// for _, order := range sorted {
	// 	fmt.Printf("%+v\n", order)
	// }
	assert.Equal(t, 5, len(sorted), "should be 5 sorted orders")
	assert.Equal(t, 1.2, sorted[1].Size, "order 2 size did not match")
	assert.Equal(t, "#3", sorted[3].OrderID, "order 3 order ID did not match")
	assert.Equal(t, 0.9, sorted[0].Size, "order 1 size did not match")
}

func TestSearchIndex(t *testing.T) {
	now := time.Now().UTC()
	//example := []float64{0.04, 0.02, 0.03, 0.01, 0.007}
	order1 := &protoOrder.Order{
		OrderID:   "#1",
		Price:     0.01,
		Size:      1.2,
		Side:      "buy",
		CreatedOn: now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:   "#2",
		Price:     0.007,
		Size:      0.2,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 1).String(),
	}
	order3 := &protoOrder.Order{
		OrderID:   "#4",
		Price:     0.007,
		Size:      2.7,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 20).String(),
	}
	order4 := &protoOrder.Order{
		OrderID:   "#3",
		Price:     0.007,
		Size:      0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 2).String(),
	}
	order5 := &protoOrder.Order{
		OrderID:   "#0",
		Price:     0.00034,
		Size:      0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 100).String(),
	}

	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
	sorted := MergeSort(orders)

	searchPrice := 0.01
	index := SearchIndex(sorted, searchPrice)
	assert.Equal(t, 4, index, "match should be at index 4")
}

func TestMatchIndices(t *testing.T) {
	now := time.Now().UTC()
	//example := []float64{0.04, 0.02, 0.03, 0.01, 0.007}
	order1 := &protoOrder.Order{
		OrderID:   "#1",
		Price:     0.01,
		Size:      1.2,
		Side:      "buy",
		CreatedOn: now.String(),
	}
	order2 := &protoOrder.Order{
		OrderID:   "#2",
		Price:     0.007,
		Size:      0.2,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 1).String(),
	}
	order3 := &protoOrder.Order{
		OrderID:   "#4",
		Price:     0.007,
		Size:      2.7,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 20).String(),
	}
	order4 := &protoOrder.Order{
		OrderID:   "#3",
		Price:     0.007,
		Size:      0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 2).String(),
	}
	order5 := &protoOrder.Order{
		OrderID:   "#0",
		Price:     0.00034,
		Size:      0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 100).String(),
	}

	orders := []*protoOrder.Order{order1, order2, order3, order4, order5}
	sorted := MergeSort(orders)

	searchPrice := 0.007
	match := MatchIndices(sorted, searchPrice, 1.5)
	//for _, o := range match {
	//	fmt.Printf("%+v\n", o)
	//}
	assert.Equal(t, 3, len(match), "should be 3 matches")
}
