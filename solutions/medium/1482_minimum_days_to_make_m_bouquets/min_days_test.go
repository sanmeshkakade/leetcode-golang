package mindays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinDays(t *testing.T) {
	type args struct {
		bloomDay []int
		m        int
		k        int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				bloomDay: []int{1, 10, 3, 10, 2},
				m:        3,
				k:        1,
			},
			want: 3,
		},
		{
			name: "two",
			args: args{
				bloomDay: []int{1, 10, 3, 10, 2},
				m:        3,
				k:        2,
			},
			want: -1,
		},
		{
			name: "three",
			args: args{
				bloomDay: []int{7, 7, 7, 7, 12, 7, 7},
				m:        2,
				k:        3,
			},
			want: 12,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := minDays(test.args.bloomDay, test.args.m, test.args.k)

			assert.Equal(t, test.want, got)
		})
	}
}
