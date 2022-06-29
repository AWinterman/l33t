package l33t_test

import (
	"math/rand"
	"reflect"
	"testing"

	"github.com/awinterman/l33t"
)

func TestQuicksort(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{
			name: "empty case",
			args: []int{},
			want: []int{},
		},
		{
			name: "reverse order case",
			args: []int{5, 4, 3, 2, 1},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "random order",
			args: []int{10, 1, 3, 11, 10, 10, 10, 15, 7, 9},
			want: []int{1, 2, 3, 7, 9, 10, 10, 10, 11, 15},
		},
		{
			name: "two elements",
			args: []int{2, 1},
			want: []int{1, 2},
		},
		{
			name: "three elements",
			args: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := l33t.Quicksort(tt.args); !reflect.DeepEqual(got, tt.want) {
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
