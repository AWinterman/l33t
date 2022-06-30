package l33t

import "testing"

func TestIsXMatrix(t *testing.T) {
	type args struct {
		matrix Matrix
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "a test case",
			args: args{
				matrix: [][]int{
					{2, 0, 0, 1},
					{0, 3, 1, 0},
					{0, 5, 2, 0},
					{4, 0, 0, 2},
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsXMatrix(tt.args.matrix); got != tt.want {
				t.Errorf("IsXMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
