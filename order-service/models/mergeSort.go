package models

import (
	protoOrder "github.com/asciiu/appa/order-service/proto/order"
)

func MergeSort(slice []*protoOrder.Order) []*protoOrder.Order {

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return Merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func Merge(left, right []*protoOrder.Order) []*protoOrder.Order {

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

func BinarySearch(a []*protoOrder.Order, search float64) (index int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		index = -1 // not found
	case a[mid].Price > search:
		index = BinarySearch(a[:mid], search)
	case a[mid].Price < search:
		index = BinarySearch(a[mid+1:], search)
		index += mid + 1
	default: // a[mid] == search
		index = mid // found
	}
	return
}

func Search(sorted []*protoOrder.Order, search float64) (index int) {
	idx := BinarySearch(sorted, search)
	slice := sorted[:idx]
	index = idx
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i].Price == search {
			index = i
		}
		if slice[i].Price < search {
			break
		}
	}

	return
}
