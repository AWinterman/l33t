package l33t

type Tree struct {
	value Lesser
	left  *Tree
	right *Tree
}

func (t *Tree) Insert(value Lesser) *Tree {
	if t == nil || t.value == nil {
		t = &Tree{value: value}
		return t
	}

	if t.value.Less(value) {
		if t.left == nil {
			t.left = &Tree{value, nil, nil}
			return t
		}
		t.left.Insert(value)
		return t
	}

	if t.right == nil {
		t.right = &Tree{value, nil, nil}
		return t
	}

	t.right.Insert(value)
	return t
}

func (t *Tree) Len() (length int) {
	length += 1
	if t.left != nil {
		length += t.left.Len()
	}

	if t.right != nil {
		length += t.right.Len()
	}

	return
}

func (t *Tree) Height() int {
	heights := make([]int, 2)
	if t.left != nil {
		heights[0] = t.left.Height()
	}

	if t.right != nil {
		heights[1] = t.right.Height()
	}

	max := 0
	for _, h := range heights {
		if h > max {
			max = h
		}
	}

	return max
}

func (n *Tree) Bal() int {
	return n.right.Height() - n.left.Height()
}

// rotate true means left, false means right
func (n *Tree) Traverse() []Lesser {
	arr := []Lesser{}
	if n.left != nil {
		arr = append(arr, n.left.Traverse()...)
	}

	arr = append(arr, n.value)

	if n.right != nil {
		arr = append(arr, n.right.Traverse()...)
	}

	return arr
}

// find the middle value
// make a new tree with the middle value as root
// then repeat with the arrays left of middle and right of middle
func Treeify(arr []Lesser) *Tree {

	middle := len(arr) / 2
	tree := &Tree{value: arr[middle]}

	if middle > 0 {
		tree.left = Treeify(arr[:middle])
	}
	if middle+1 < len(arr) {
		tree.right = Treeify(arr[middle+1:])
	}
	return tree
}
