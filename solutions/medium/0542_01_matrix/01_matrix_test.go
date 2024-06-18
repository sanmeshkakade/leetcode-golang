package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateMatrix(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one",
			args: args{mat: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}},
		},
		{
			name: "two",
			args: args{mat: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
			want: [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}},
		},
		{
			name: "three",
			args: args{mat: [][]int{{0, 1, 1}, {1, 1, 1}}},
			want: [][]int{{0, 1, 2}, {1, 2, 3}},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := updateMatrix(test.args.mat)

			assert.Equal(t, test.want, got)
		})
	}
}
