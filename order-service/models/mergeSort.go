package models

import (
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

func MergeSort(slice []*protoOrder.Order) []*protoOrder.Order {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func merge(left, right []*protoOrder.Order) []*protoOrder.Order {

	size, i, j := len(left)+len(right), 0, 0
	slice := make([]*protoOrder.Order, size, size)

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

// should return index where Order.Price == price
func binarySearch(a []*protoOrder.Order, price float64) (index int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		index = -1 // not found
	case a[mid].Price > price:
		index = binarySearch(a[:mid], price)
	case a[mid].Price < price:
		index = binarySearch(a[mid+1:], price)
		index += mid + 1
	default: // a[mid] == search
		index = mid // found
	}
	return
}

// returns first index where Order.Price == price
func searchIndex(sorted []*protoOrder.Order, price float64) (index int) {
	idx := binarySearch(sorted, price)
	slice := sorted[:idx]
	index = idx
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i].Price == price {
			index = i
		}
		if slice[i].Price < price {
			break
		}
	}

	return
}

func MatchIndices(sorted []*protoOrder.Order, price, size float64) (indices []int) {
	first := searchIndex(sorted, price)
	sum := 0.0
	//orders := make([]*protoOrder.Order, 0)
	for i, order := range sorted[first:] {
		if order.Price > price {
			break
		}
		sum += order.Size
		indices = append(indices, first+i)
		if sum >= size {
			break
		}
	}
	return
}
