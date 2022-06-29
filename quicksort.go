package l33t

import "golang.org/x/exp/constraints"

/*
// Sorts a (portion of an) array, divides it into partitions, then sorts those
algorithm quicksort(A, lo, hi) is
  // Ensure indices are in correct order
  if lo >= hi || lo < 0 then
    return

  // Partition array and get the pivot index
  p := partition(A, lo, hi)

  // Sort the two partitions
  quicksort(A, lo, p - 1) // Left side of pivot
  quicksort(A, p + 1, hi) // Right side of pivot

// Divides array into two partitions
algorithm partition(A, lo, hi) is
  pivot := A[hi] // Choose the last element as the pivot

  // Temporary pivot index
  i := lo - 1

  for j := lo to hi - 1 do
    // If the current element is less than or equal to the pivot
    if A[j] <= pivot then
      // Move the temporary pivot index forward
      i := i + 1

      // Swap the current element with the element at the temporary pivot index
      swap A[i] with A[j]
  // Move the pivot element to the correct pivot position (between the smaller and larger elements)
  i := i + 1
  swap A[i] with A[hi]
  return i // the pivot index
*/
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
