package containsduplicate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsDuplicate(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{nums: []int{1, 2, 3, 1}},
			want: true,
		},
		{
			name: "two",
			args: args{nums: []int{1, 2, 3, 4}},
			want: false,
		},
		{
			name: "three",
			args: args{nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := containsDuplicate(test.args.nums)

			assert.Equal(t, test.want, got)
		})
	}
}
