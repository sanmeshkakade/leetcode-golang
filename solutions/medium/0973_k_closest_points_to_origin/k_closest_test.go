package kclosest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKClosest(t *testing.T) {
	type args struct {
		points [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one",
			args: args{
				points: [][]int{{1, 3}, {-2, 2}},
				k:      1,
			},
			want: [][]int{{-2, 2}},
		},
		{
			name: "two",
			args: args{
				points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
				k:      2,
			},
			want: [][]int{{3, 3}, {-2, 4}},
		},
		{
			name: "three",
			args: args{
				points: [][]int{{9, 0}, {7, 10}, {-4, -2}, {3, -9}, {9, 1}, {-5, -1}},
				k:      5,
			},
			want: [][]int{{-4, -2}, {-5, -1}, {9, 0}, {9, 1}, {3, -9}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := kClosest(test.args.points, test.args.k)

			assert.ElementsMatch(t, test.want, got)
		})
	}
}
