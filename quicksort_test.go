package l33t_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/wintersparkle/l33t"
	"golang.org/x/exp/constraints"
)

func quicksort[T constraints.Ordered](items []T, t *testing.T) []T {
	return l33t.Quicksort(items)
}

func BenchmarkQuickSort_equalValues(t *testing.B) {
	want := []string{"b", "b"}
	got := l33t.Quicksort([]string{"b", "b"})
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Quicksort() = %v, want %v", got, want)
	}
}

func TestQuickSort_withStringsAlreadyInOrder(t *testing.T) {
	want := []string{"b", "c"}
	got := quicksort([]string{"b", "c"}, t)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Quicksort() = %v, want %v", got, want)
	}
}

func TestQuickSort_withStrings(t *testing.T) {
	want := []string{"a", "b", "b", "c", "d", "e"}
	got := quicksort([]string{"c", "b", "a", "b", "e", "d"}, t)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Quicksort() = %v, want %v", got, want)
	}
}

func TestQuickSort_AnotherPalimdrome(t *testing.T) {
	got := quicksort([]rune("cbab"), t)
	want := "abbc"
	// pick pivot idx 2 = c
	// a < c (cool)
	// b < c (cool)
	// c == c (not cool, need to find a swap)
	// descend from top to find a swap
	// a < c (not cool swap found)
	// swap idx 2 and idx 5

	if !reflect.DeepEqual(string(got), want) {
		t.Errorf("Quicksort() = %v, want %v", string(got), want)
	}
}

func TestQuickSort_withStringsWhereLessThanPivotAreAllInOrder(t *testing.T) {
	want := "aabbc"
	got := quicksort([]rune("abcba"), t)
	// pick pivot idx 2 = c
	// a < c (cool)
	// b < c (cool)
	// c == c (not cool, need to find a swap)
	// descend from top to find a swap
	// a < c (not cool swap found)
	// swap idx 2 and idx 5

	if !reflect.DeepEqual(string(got), want) {
		t.Errorf("Quicksort() = %v, want %v", string(got), want)
	}
}

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "reverse order case",
			args: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "reverse order case",
			args: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "random order",
			args: []int{10, 1, 3, 11, 10, 10, 10, 15, 7, 9},
			want: []int{1, 3, 7, 9, 10, 10, 10, 10, 11, 15},
		},
		{
			name: "empty case",
			args: []int{},
			want: []int{},
		},

		{
			name: "two elements",
			args: []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "two elements in order",
			args: []int{1, 2},
			want: []int{1, 2},
		},
		{
			name: "three elements; reverse order",
			args: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := quicksort(tt.args, t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Quicksort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkQuicksort(b *testing.B) {
	var elements = 1000
	A := []int{}
	for i := 0; i < elements; i += 1 {
		A = append(A, i, rand.Int())
	}
	for i := 0; i < b.N; i++ {
		l33t.Quicksort(A)
	}
}
