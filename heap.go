package l33t

type Lesser interface {
	Less(other any) bool
}

// Heapify makes an array into a heap
func Heapify[T Lesser](arr []T, size, root int) []T {
	largest := root     // largest candiadte value (guess it is sorted)
	left := 2*root + 1  //left child index
	right := 2*root + 2 //right child index

	// check if either parent is less than its children

	if left < size && arr[largest].Less(arr[left]) {
		// if the index of the left child is inbounds, and
		// it is greater than the largest, then the left child is actually the largest
		largest = left
	}

	if right < size && arr[largest].Less(arr[right]) {
		// if the index of the right child is inbounds, and
		// it is greater than or equal to the largest, then the right child is actually the largest
		largest = right
	}

	if largest != root {
		arr[root], arr[largest] = arr[largest], arr[root]

		arr = Heapify(arr, size, largest)
	}
	return arr
}

func Heapsort[T Lesser](arr []T) []T {
	n := len(arr)
	// heapify the entire array
	// we could just heapify from 0, but this means we only take log(n) time instead of n timne.
	for i := n/2 - 1; i >= 0; i-- {
		arr = Heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		// now that it is heapified, we know that
		// 0 is the largest value, and it belongs at the end.
		arr[0], arr[i] = arr[i], arr[0]

		// now we need to ensure the subtree is still a heap given the new root--
		// this will take i time, for each if i, i-1, i-2 ... 0
		arr = Heapify(arr, i, 0)
	}
	return arr

}
