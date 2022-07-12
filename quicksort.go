package l33t

import (
	"golang.org/x/exp/constraints"
)

// Quicksort is a quicksort implementation that sorts Ordered values.
func Quicksort[T constraints.Ordered](A []T) []T {
	sorted := quicksort(A)
	return sorted
}

func quicksort[T constraints.Ordered](items []T) []T {
	length := len(items)

	if length < 2 {
		return items
	}

	if length == 2 {
		if items[0] < items[1] {
			return items
		}

		return []T{items[1], items[0]}
	}

	var lo = 0
	var hi = length - 1

	var index = length / 2
	var pivot = items[index]

	for {

		for items[lo] < pivot {
			lo++
		}

		for items[hi] > pivot {
			hi--
		}

		if lo >= hi {
			// they crossed, time to exit
			break
		}

		items[lo], items[hi] = items[hi], items[lo]

		lo++
		hi--
	}

	bottom := items[:hi+1]
	bottomhalf := quicksort(bottom)

	top := items[hi+1:]
	tophalf := quicksort(top)

	items = append(bottomhalf, tophalf...)

	return items
}


func QuicksortLess[T Lesser](items []T) []T {
	length := len(items)

	if length < 2 {
		return items
	}

	if length == 2 {
		if items[0].Less(items[1]) {
			return items
		}

		return []T{items[1], items[0]}
	}

	var lo = 0
	var hi = length - 1

	var index = length / 2
	var pivot = items[index]

	for {

		for items[lo].Less(pivot) {
			lo++
		}

		for pivot.Less(items[hi]) {
			hi--
		}

		if lo >= hi {
			// they crossed, time to exit
			break
		}

		items[lo], items[hi] = items[hi], items[lo]

		lo++
		hi--
	}

	bottom := items[:hi+1]
	bottomhalf := QuicksortLess(bottom)

	top := items[hi+1:]
	tophalf := QuicksortLess(top)

	items = append(bottomhalf, tophalf...)

	return items
}
