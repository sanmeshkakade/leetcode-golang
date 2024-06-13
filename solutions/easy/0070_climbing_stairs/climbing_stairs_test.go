package climbingstairs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{n: 1},
			want: 1,
		},
		{
			name: "two",
			args: args{n: 2},
			want: 2,
		},
		{
			name: "three",
			args: args{n: 3},
			want: 3,
		},
		{
			name: "four",
			args: args{n: 8},
			want: 34,
		},
	}

	funcs := []func(int) int{climbStairs, climbStairs1, climbStairs2, climbStairs3}

	for key, f := range funcs {
		for _, test := range tests {
			t.Run(fmt.Sprintf("%s_%d", test.name, key), func(t *testing.T) {
				got := f(test.args.n)

				assert.Equal(t, test.want, got)
			})
		}
	}

}
