package l33t_test

import (
	"fmt"
	"testing"

	"github.com/matryer/is"
	"github.com/wintersparkle/l33t"
)

func TestDolls(t *testing.T) {
	is := is.New(t)
	envelopes := []Envelope{
		{5, 4},
		{6, 4},
		{6, 7},
		{2, 3},
	}

	is.Equal(maxEnvelopes(envelopes), 3)
}

func HeapSort(arr []Envelope) []Envelope {
	n := len(arr)
	// heapify the entire array
	for i := n/2 - 1; i >= 0; i-- {
		arr = l33t.Heapify(arr, n, i)
	}

	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		arr = l33t.Heapify(arr, i, 0)
	}
	return arr
}

func TestHeapify(t *testing.T) {

	is := is.New(t)

	is.Equal(
		l33t.Heapify(
			[]Envelope{
				{5, 4},
				{6, 4},
				{1, 2},
				{3, 4},
				{4, 5},
			},
			5,
			0,
		),
		[]Envelope{
			{6, 4},
			{5, 4},
			{1, 2},
			{3, 4},
			{4, 5},
		},
	)
}

type Envelope [2]int

func (e Envelope) Less(a any) bool {
	other, ok := a.(Envelope)

	if !ok {
		return false
	}

	isLess := (e[0] < other[0] && e[1] < other[1]) || e[0]*e[1] < other[0]*other[1]
	return isLess
}

func (e Envelope) Fits(smaller Envelope) bool {
	return e[0] > smaller[0] && e[1] > smaller[1]
}

func (e Envelope) String() string {
	return fmt.Sprintf("%dx%d", e[0], e[1])
}

// Envelopes is a slice of Envelope
type Envelopes []Envelope

// Left returns the left subtree of the root node if the left child fits inside the root node.k
func (e Envelopes) Left(idx int) (tree Envelopes, fits bool, childIndex int) {
	childIndex = 2*idx + 1
	if len(e) < childIndex {
		return nil, false, 0
	}
	return e[childIndex:], e[0].Fits(e[childIndex]), childIndex
}

func (e Envelopes) Right(idx int) (Envelopes, bool, int) {
	childIndex := 2*idx + 2
	if len(e) < childIndex {
		return nil, false, 0
	}
	return e[childIndex:], e[0].Fits(e[childIndex]), childIndex
}

func maxEnvelopes(envelopes []Envelope) int {

	tree := Envelopes(envelopes)
	n := len(tree)
	// heapify the entire array
	for i := n/2 - 1; i >= 0; i-- {
		// tree = Heapify(tree, n, i)
	}
	return 0

}

func TestLess(t *testing.T) {
	a := Envelope{5, 4}
	b := Envelope{3, 4}
	c := Envelope{2, 3}
	d := Envelope{5, 3}

	is := is.New(t)

	is.True(!a.Less(b)) // "a" > "b"
	is.True(b.Less(a))  // "a" > "b"
	is.True(c.Less(a))  // "c" < "a"
	is.True(c.Less(b))  // "c" < "b"

	is.True(d.Less(a))
	is.True(b.Less(d))
}
