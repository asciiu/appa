package models

import (
	"fmt"
	"testing"
	"time"

	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

func TestMergeSort(t *testing.T) {
	now := time.Now().UTC()
	//example := []float64{0.04, 0.02, 0.03, 0.01, 0.007}
	order1 := protoOrder.Order{
		OrderID:   "#1",
		Price:     0.007,
		Size:      1.2,
		Side:      "buy",
		CreatedOn: now.String(),
	}
	order2 := protoOrder.Order{
		OrderID:   "#2",
		Price:     0.007,
		Size:      0.2,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 1).String(),
	}
	order3 := protoOrder.Order{
		OrderID:   "#4",
		Price:     0.007,
		Size:      2.7,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 20).String(),
	}
	order4 := protoOrder.Order{
		OrderID:   "#3",
		Price:     0.007,
		Size:      0.9,
		Side:      "buy",
		CreatedOn: now.Add(time.Second * 2).String(),
	}

	orders := []protoOrder.Order{order1, order2, order3, order4}
	sorted := MergeSort(orders)

	for _, order := range sorted {
		fmt.Printf("%+v\n", order)
	}
	//assert.Equal(t, 1, len(book.BuyQ), "should be 1 order in buys")
	//assert.Equal(t, 0, len(book.SellQ), "should be 0 order in sells")
}
