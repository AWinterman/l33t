package l33t

import (
	"reflect"
	"testing"
)

func TestTreeify(t *testing.T) {
	type args struct {
		arr []Lesser
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Treeify(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Treeify() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Insert(t *testing.T) {
	type fields struct {
		value Lesser
		left  *Tree
		right *Tree
	}
	type args struct {
		value Lesser
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *Tree
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tree{
				value: tt.fields.value,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := tr.Insert(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Len(t *testing.T) {
	type fields struct {
		value Lesser
		left  *Tree
		right *Tree
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
			tr := &Tree{
				value: tt.fields.value,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if gotLength := tr.Len(); gotLength != tt.wantLength {
				t.Errorf("Tree.Len() = %v, want %v", gotLength, tt.wantLength)
			}
		})
	}
}

func TestTree_Height(t *testing.T) {
	type fields struct {
		value Lesser
		left  *Tree
		right *Tree
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
			tr := &Tree{
				value: tt.fields.value,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := tr.Height(); got != tt.want {
				t.Errorf("Tree.Height() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Bal(t *testing.T) {
	type fields struct {
		value Lesser
		left  *Tree
		right *Tree
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
			n := &Tree{
				value: tt.fields.value,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := n.Bal(); got != tt.want {
				t.Errorf("Tree.Bal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTree_Traverse(t *testing.T) {
	type fields struct {
		value Lesser
		left  *Tree
		right *Tree
	}
	tests := []struct {
		name   string
		fields fields
		want   []Lesser
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := &Tree{
				value: tt.fields.value,
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := n.Traverse(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.Traverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
