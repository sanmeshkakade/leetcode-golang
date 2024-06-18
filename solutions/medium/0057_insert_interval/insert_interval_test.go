package insertinterval

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsert(t *testing.T) {
	type args struct {
		intervals   [][]int
		newInterval []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "one",
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "two",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name: "three",
			args: args{
				intervals:   [][]int{},
				newInterval: []int{5, 7},
			},
			want: [][]int{{5, 7}},
		},
	}

	funcs := []func([][]int, []int) [][]int{insert, insert1}

	for key, f := range funcs {
		for _, test := range tests {
			t.Run(fmt.Sprintf("%s_%d", test.name, key), func(t *testing.T) {
				got := f(test.args.intervals, test.args.newInterval)

				assert.Equal(t, test.want, got)
			})
		}
	}
}
