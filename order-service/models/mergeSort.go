package models

import (
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

func MergeSort(slice []*protoOrder.Order) []*protoOrder.Order {

	if len(slice) < 2 {
		return slice
	}
	mid := len(slice) / 2
	return merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func merge(left, right []*protoOrder.Order) []*protoOrder.Order {

	size, leftIndex, rightIndex := len(left)+len(right), 0, 0
	sorted := make([]*protoOrder.Order, size, size)

	for k := 0; k < size; k++ {
		switch {
		case leftIndex > len(left)-1 && rightIndex <= len(right)-1:
			// no more items in left
			// therefore take next index in right
			sorted[k] = right[rightIndex]
			rightIndex++
		case rightIndex > len(right)-1 && leftIndex <= len(left)-1:
			// no more items in right
			// therefore take next index in left
			// increment index of left
			sorted[k] = left[leftIndex]
			leftIndex++
		case left[leftIndex].Price < right[rightIndex].Price:
			// lesser price goes first
			sorted[k] = left[leftIndex]
			leftIndex++
		case left[leftIndex].Price == right[rightIndex].Price && left[leftIndex].CreatedOn < right[rightIndex].CreatedOn:
			// if prices are equal do comparison with created on dates
			sorted[k] = left[leftIndex]
			leftIndex++
		default:
			sorted[k] = right[rightIndex]
			rightIndex++
		}
	}
	return sorted
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

// returns first index where Order.Price <= price
func searchLessThan(sorted []*protoOrder.Order, price float64) (index int) {
	idx := binarySearch(sorted, price)
	slice := sorted[idx:]
	index = idx
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i].Price < price {
			index += i
			break
		}
	}

	return
}

// returns first index where Order.Price > price
func searchGreaterThan(sorted []*protoOrder.Order, price float64) (index int) {
	idx := binarySearch(sorted, price)
	slice := sorted[idx:]
	index = idx
	for i := 0; i < len(slice); i++ {
		if slice[i].Price > price {
			index += i
			break
		}
	}

	return
}

func MatchIndices(sorted []*protoOrder.Order, price, size float64) (indices []int) {
	first := searchLessThan(sorted, price)
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
