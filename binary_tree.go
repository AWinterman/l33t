package l33t

import "fmt"

// Tree is a struct that creates a binary search tree obeying the constraints:
// 1. The left subtree of a node is always less than the value of the node
// 2. The right subtree of a node is always greater than the value of the node
type Tree[T Lesser] struct {
	value T
	Left  *Tree[T]
	Right *Tree[T]
}

func (t *Tree[T]) String() string {
	if t.Left != nil && t.Right != nil {
		return fmt.Sprintf("(%s) <- %v -> (%s)", t.Left, t.value, t.Right)
	}
	if t.Left != nil {
		return fmt.Sprintf("(%s) <- %v", t.Left, t.value)
	}

	if t.Right != nil {
		return fmt.Sprintf("%v -> (%s)", t.value, t.Right)
	}

	return fmt.Sprintf("%v", t.value)
}

func TreeSort[T Lesser](arr []T) []T {
	var tree *Tree[T]
	for _, l := range arr {
		tree = tree.Insert(l)
	}

	return tree.Traverse()
}

func (t *Tree[T]) Insert(value T) *Tree[T] {
	if t == nil {
		t = &Tree[T]{value: value}
		return t
	}

	if value.Less(t.value) {
		t.Left = t.Left.Insert(value)
		return t
	}

	t.Right = t.Right.Insert(value)
	return t
}

func (t *Tree[T]) Len() (length int) {
	length += 1
	if t.Left != nil {
		length += t.Left.Len()
	}

	if t.Right != nil {
		length += t.Right.Len()
	}

	return
}

func (t *Tree[T]) Height() int {
	heights := make([]int, 2)
	if t.Left != nil {
		heights[0] = t.Left.Height() + 1
	}

	if t.Right != nil {
		heights[1] = t.Right.Height() + 1
	}

	max := 1
	for _, h := range heights {
		if h > max {
			max = h
		}
	}

	return max
}

func (n *Tree[T]) Bal() int {
	return n.Right.Height() - n.Left.Height()
}

// rotate true means left, false means right
func (n *Tree[T]) Traverse() []T {
	arr := []T{}
	if n.Left != nil {
		arr = append(arr, n.Left.Traverse()...)
	}

	arr = append(arr, n.value)

	if n.Right != nil {
		arr = append(arr, n.Right.Traverse()...)
	}

	return arr
}

// Treeify inserts a sorted array into the tree.
// If your array is not sorted, you will get a tree that is not sorted, which will produce incorrect results.
// if you are relying on tis sorting characteristics.
//
// this works as follows:
// 1. find the middle value
// 2. make a new tree with the middle value as root
// 3. then repeat with the arrays left of middle and right of middle
func Treeify[T Lesser](arr []T) *Tree[T] {
	middle := len(arr) / 2
	if len(arr) <= middle {
		return &Tree[T]{}
	}

	tree := &Tree[T]{value: arr[middle]}

	if middle > 0 {
		tree.Left = Treeify(arr[:middle])
	}
	if middle+1 < len(arr) {
		tree.Right = Treeify(arr[middle+1:])
	}

	return tree
}
