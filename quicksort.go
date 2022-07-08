package l33t

import (
	"log"

	"golang.org/x/exp/constraints"
)

// Quicksort is a quicksort implementation that sorts Ordered values.
func Quicksort[T constraints.Ordered](A []T) []T {

	// log.Println("sorting", A)
	sorted := quicksort(A)
	// log.Println("done", sorted)
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

	var lo int = 0
	var hi int = length - 1

	var index = length / 2
	var pivot = items[index]

	log.Printf("checking %v at pivot items[%d] = %v", items, index, pivot)
	for {

		for items[lo] < pivot {
			lo++
		}

		for items[hi] > pivot {
			hi--
		}

		if lo >= hi {
			// they crossed, time to exit
			// log.Printf("	lo %d and hi %d crossed", lo, hi)
			break
		}

		log.Printf("	swap items[%d] = %v and items[%d] = %v", lo, items[lo], hi, items[hi])
		items[lo], items[hi] = items[hi], items[lo]
		log.Printf("	swap result %v", items)

		lo++
		hi--
	}

	bottom := items[:hi+1]
	log.Printf("	checking bottom half of %v: %v", items, bottom)
	bottomhalf := quicksort(bottom)

	top := items[hi+1:]
	log.Printf("	checking top half of %v: %v", items, top)
	tophalf := quicksort(top)

	items = append(bottomhalf, tophalf...)

	// log.Println("	Done ", items)
	return items
}
