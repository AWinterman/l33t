package l33t_test

import (
	"fmt"
	"log"
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

	envelopes = []Envelope{
		{1, 10},
		{2, 11},
		{3, 12},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
	}

	is.Equal(maxEnvelopes(envelopes), 4)
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

type Envelope [2]int

func (e Envelope) Fits(smaller Envelope) bool {
	return e[0] > smaller[0] && e[1] > smaller[1]
}

func (e Envelope) Less(a any) bool {
	other, ok := a.(Envelope)

	if !ok {
		return false
	}
	return other.Fits(e)
}

func (e Envelope) String() string {
	return fmt.Sprintf("%dx%d", e[0], e[1])
}

func maxEnvelopes(envelopes []Envelope) int {
	envelopes = l33t.QuicksortLess(envelopes)
	sequences := [][]Envelope{}
	// for each envelope
	for n := len(envelopes) - 1; n >= 0; n-- {
		log.Printf("n:%v env:%s seqs:%v", n, envelopes, sequences)
		// find all the existing sequences the envelope fits in
		fits := []int{}
		for i, sequence := range sequences {
			if sequence[0].Fits(envelopes[n]) {
				fits = append(fits, i)
			}
		}

		log.Printf("fits:%v", fits)

		// if it doesnt fit in any it is a new envelope
		if len(fits) == 0 {
			sequences = append(sequences, []Envelope{envelopes[n]})
		} else {
			// otherwise, add it to the sequences it fits in
			for _, i := range fits {
				log.Printf("adding %v to %v", envelopes[n], sequences[i])
				sequences[i] = append([]Envelope{envelopes[n]}, sequences[i]...)
			}
		}
	}

	log.Println(sequences)

	// find the longest sequence
	var max int

	for _, sequence := range sequences {
		if len(sequence) > max {
			max = len(sequence)
		}
	}
	return max
}

func TestLess(t *testing.T) {
	a := Envelope{5, 4}
	b := Envelope{3, 4}
	c := Envelope{2, 3}
	d := Envelope{5, 3}

	is := is.New(t)

	is.True(!a.Less(b)) // "a" > "b"
	is.True(!b.Less(a)) // "a" > "b"
	is.True(c.Less(a))  // "c" < "a"
	is.True(c.Less(b))  // "c" < "b"

	is.True(!d.Less(a))
	is.True(!b.Less(d))
}
