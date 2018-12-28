package models

import (
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

func MergeSort(slice []protoOrder.Order) []protoOrder.Order {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func Merge(left, right []protoOrder.Order) []protoOrder.Order {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]protoOrder.Order, size, size)

	for k := 0; k < size; k++ {
		if i > len(left)-1 && j <= len(right)-1 {
			slice[k] = right[j]
			j++
		} else if j > len(right)-1 && i <= len(left)-1 {
			slice[k] = left[i]
			i++
		} else if left[i].Price < right[j].Price {
			slice[k] = left[i]
			i++
		} else if left[i].Price == right[j].Price && left[i].CreatedOn < right[j].CreatedOn {
			slice[k] = left[i]
			i++
		} else {
			slice[k] = right[j]
			j++
		}
	}
	return slice
}
