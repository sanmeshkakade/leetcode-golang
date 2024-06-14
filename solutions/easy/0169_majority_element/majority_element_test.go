package majorityelement

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElement(t *testing.T) {
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
			args: args{nums: []int{3, 2, 3}},
			want: 3,
		},
		{
			name: "two",
			args: args{nums: []int{2, 2, 1, 1, 1, 2, 2}},
			want: 2,
		},
		{
			name: "three",
			args: args{nums: []int{1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0}},
			want: 1,
		},
	}

	funcs := []func([]int) int{majorityElement, majorityElement1}

	for key, f := range funcs {
		for _, test := range tests {
			t.Run(fmt.Sprintf("%s_%d", test.name, key), func(t *testing.T) {
				got := f(test.args.nums)

				assert.Equal(t, test.want, got)
			})
		}
	}
}
