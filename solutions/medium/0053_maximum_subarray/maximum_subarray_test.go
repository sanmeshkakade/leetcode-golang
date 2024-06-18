package maximumsubarray

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSubArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}},
			want: 6,
		},
		{
			name: "two",
			args: args{nums: []int{1}},
			want: 1,
		},
		{
			name: "three",
			args: args{nums: []int{5, 4, -1, 7, 8}},
			want: 23,
		},
		{
			name: "four",
			args: args{nums: []int{-5, -4, -1, -7, -8}},
			want: -1,
		},
		{
			name: "five",
			args: args{nums: []int{5, 4, 1, 7, 8}},
			want: 25,
		},
	}

	funcs := []func([]int) int{maxSubArray, maxSubArray1, maxSubArray2}

	for key, f := range funcs {
		for _, test := range tests {
			t.Run(fmt.Sprintf("%s_%d", test.name, key), func(t *testing.T) {
				got := f(test.args.nums)

				assert.Equal(t, test.want, got)
			})
		}
	}
}
