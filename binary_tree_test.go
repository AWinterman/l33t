package l33t

import (
	"reflect"
	"testing"

	"github.com/matryer/is"
)

func (i Int) Less(other any) bool {
	integer, ok := other.(Int)
	if !ok {
		return false
	}
	return i < integer
}

type Int int

func TestTreeify(t *testing.T) {
	type args struct {
		arr []Int
	}
	tests := []struct {
		name string
		args args
		want *Tree[Int]
	}{
		{
			"empty array",
			args{[]Int{}},
			&Tree[Int]{},
		},
		{
			"single element array",
			args{[]Int{1}},
			&Tree[Int]{value: Int(1)},
		},
		{
			"sorted array one level deep",
			args{[]Int{1, 2, 3}},
			&Tree[Int]{
				value: Int(2),
				Left:  &Tree[Int]{value: Int(1)},
				Right: &Tree[Int]{value: Int(3)},
			},
		},
		{
			"sorted array two levels deep",
			args{[]Int{1, 2, 3, 4, 5}},
			&Tree[Int]{
				value: Int(3), // we start with the middle value, in this case index 2.
				Left: &Tree[Int]{
					value: Int(2), // then we make two subtrees of the left side of 3 and the right side of 3.
					// we pick two because we arbitrarily select the higher number when
					// there's an even-sized array, which means the other value must be on the left to maintain the sorting constraint.
					Left: &Tree[Int]{
						value: Int(1),
					},
				},
				Right: &Tree[Int]{
					value: Int(5),
					Left: &Tree[Int]{
						value: Int(4),
					},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Treeify(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Treeify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTreeSort(b *testing.T) {
	is := is.New(b)

	b.Run("pair", func(b *testing.T) {
		is := is.New(b)
		pair := []Int{2, 1}
		sorted := TreeSort(pair)
		is.Equal(sorted, []Int{1, 2}) // should sort a pair
	})

	b.Run("palimdrone", func(b *testing.T) {

		is := is.New(b)
		palimdrome := []Int{1, 2, 3, 3, 2, 1}
		sorted := TreeSort(palimdrome)
		is.Equal(sorted, []Int{1, 1, 2, 2, 3, 3}) // should sort palimdrome
	})

}

func TestTree_Insert(t *testing.T) {
	type fields struct {
		value Int
		left  *Tree[Int]
		right *Tree[Int]
	}
	type args struct {
		value Int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tree[Int]
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree[Int]{
				value: tt.fields.value,
				Left:  tt.fields.left,
				Right: tt.fields.right,
			}
			if got := tr.Insert(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Len(t *testing.T) {
	type fields struct {
		value Int
		left  *Tree[Int]
		right *Tree[Int]
	}
	tests := []struct {
		name       string
		fields     fields
		wantLength int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree[Int]{
				value: tt.fields.value,
				Left:  tt.fields.left,
				Right: tt.fields.right,
			}
			if gotLength := tr.Len(); gotLength != tt.wantLength {
				t.Errorf("Tree.Len() = %v, want %v", gotLength, tt.wantLength)
			}
		})
	}
}

func TestTree_Height(t *testing.T) {
	type fields struct {
		value Int
		left  *Tree[Int]
		right *Tree[Int]
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree[Int]{
				value: tt.fields.value,
				Left:  tt.fields.left,
				Right: tt.fields.right,
			}
			if got := tr.Height(); got != tt.want {
				t.Errorf("Tree.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Bal(t *testing.T) {
	type fields struct {
		value Int
		left  *Tree[Int]
		right *Tree[Int]
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Tree[Int]{
				value: tt.fields.value,
				Left:  tt.fields.left,
				Right: tt.fields.right,
			}
			if got := n.Bal(); got != tt.want {
				t.Errorf("Tree.Bal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Traverse(t *testing.T) {
	tests := []struct {
		name   string
		fields *Tree[Int]
		want   []Int
	}{
		{
			name: "traverses the tree",
			fields: &Tree[Int]{
				value: Int(2),
				Left:  &Tree[Int]{value: Int(1)},
				Right: &Tree[Int]{value: Int(3)},
			},
			want: []Int{1, 2, 3},
		},
		{
			name: "traverses nested",
			fields: &Tree[Int]{
				value: Int(4),
				Left: &Tree[Int]{
					value: Int(2),
					Left: &Tree[Int]{
						value: Int(1),
					},
					Right: &Tree[Int]{
						value: Int(3),
					},
				},
				Right: &Tree[Int]{
					value: Int(6),
					Left:  &Tree[Int]{value: 5},
					Right: &Tree[Int]{value: 7},
				},
			},
			want: []Int{1, 2, 3, 4, 5, 6, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Tree[Int]{
				value: tt.fields.value,
				Left:  tt.fields.Left,
				Right: tt.fields.Right,
			}
			if got := n.Traverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
