package l33t

import (
	"log"

	"golang.org/x/exp/constraints"
)

// Quicksort is a quicksort implementation that sorts Ordered values.
func Quicksort[T constraints.Ordered](A []T) []T {
	return quicksort(A, 0, len(A)-1)
}

func quicksort[T constraints.Ordered](A []T, lo, hi int) []T {
	// log.Printf("sort %v[%d:%d]", A, lo, hi)
	if lo >= 0 && hi >= 0 && lo < hi {
		A, p := partition(A, lo, hi)

		A = quicksort(A, lo, p)   // left of pivot
		A = quicksort(A, p+1, hi) // right of pivot
		return A
	}
	// log.Printf("	sorted %v[%d:%d]", A, lo, hi)
	return A

}

func partition[T constraints.Ordered](A []T, lo, hi int) ([]T, int) {

	middle := ((hi + lo) / 2)
	pivot := A[middle] // The value in the middle of the array
	left := lo
	right := hi

	for {
		// log.Printf("	find first swappables %v[%v:%v]", A, left, right)
		for A[left] < pivot {
			left += 1
		}

		for A[right] > pivot {
			right -= 1
		}

		// log.Printf("	found: A[%v] = %v < pivot A[%v] = %v < A[%v] = %v; ", left, A[left], middle, pivot, right, A[right])

		if left >= right {
			log.Printf("		because cursors met; returning %v and pivot %v", A, right)
			return A, right
		}

		// log.Printf("		swap A[%v] = %v and A[%v] = %v", left, A[left], right, A[right])
		A[left], A[right] = A[right], A[left]

		// move both cursors-- we don't want to swap these elements again
		left += 1
		right -= 1
	}
}
