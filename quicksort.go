package l33t

import "golang.org/x/exp/constraints"

// Quicksort is a quicksort implementation that sorts Ordered values.
func Quicksort[T constraints.Ordered](A []T) []T {
	return quicksort(A, 0, len(A)-1)
}

func quicksort[T constraints.Ordered](A []T, lo, hi int) []T {
	if lo >= hi || lo < 0 {
		// this is actually invalid inputs but whatever
		return A
	}

	A, p := partition(A, lo, hi)
	A = quicksort(A, lo, p-1) // left of pivot
	A = quicksort(A, p+1, hi) // right of pivot

	return A
}

func partition[T constraints.Ordered](A []T, lo, hi int) ([]T, int) {
	pivot := A[hi]
	i := lo - 1
	for j := lo; j < hi-1; j++ {
		if A[j] <= pivot {
			i++
			// swap values at i and j
			A[j], A[i] = A[i], A[j]
		}
	}

	i++
	// swap values at i and hi
	A[hi], A[i] = A[i], A[hi]
	return A, i
}
