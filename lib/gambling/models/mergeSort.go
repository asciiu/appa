package models

func MergeSort(slice []Stake) []Stake {

	if len(slice) < 2 {
		return slice
	}
	mid := len(slice) / 2
	return merge(MergeSort(slice[:mid]), MergeSort(slice[mid:]))
}

// Merges left and right slice into newly created slice
func merge(left, right []Stake) []Stake {

	size, leftIndex, rightIndex := len(left)+len(right), 0, 0
	sorted := make([]Stake, size, size)

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
		case left[leftIndex].Odds < right[rightIndex].Odds:
			// lesser price goes first
			sorted[k] = left[leftIndex]
			leftIndex++
		case left[leftIndex].Odds == right[rightIndex].Odds && left[leftIndex].CreatedAt.Before(right[rightIndex].CreatedAt):
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

// should return index where Stake.odds == odds
func binarySearch(a []Stake, odds float64) (index int) {
	mid := len(a) / 2
	switch {
	case len(a) == 0:
		index = -1 // not found
	case a[mid].Odds > odds:
		index = binarySearch(a[:mid], odds)
	case a[mid].Odds < odds:
		index = binarySearch(a[mid+1:], odds)
		index += mid + 1
	default: // a[mid] == search
		index = mid // found
	}
	return
}

// returns first index where Order.Price <= price
func searchLessThan(sorted []Stake, odds float64) (index int) {
	idx := binarySearch(sorted, odds)
	index = -1
	if idx < 0 {
		return
	}

	slice := sorted[idx:]
	index = idx
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i].Odds < odds {
			index += i
			break
		}
	}

	return
}

// returns first index where Order.Price > price
func searchGreaterThan(sorted []Stake, odds float64) (index int) {
	idx := binarySearch(sorted, odds)
	slice := sorted[idx:]
	index = idx
	for i := 0; i < len(slice); i++ {
		if slice[i].Odds > odds {
			index += i
			break
		}
	}

	return
}

func FindOrder(sorted []Stake, stake Stake) (index int) {
	start := searchLessThan(sorted, stake.Odds)
	index = -1
	if start < 0 {
		return
	}

	for i, o := range sorted[start:] {
		if o.ID == stake.ID {
			index = start + i
		}
	}
	return
}
