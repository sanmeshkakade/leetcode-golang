package maxprofit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfitAssignment(t *testing.T) {
	type args struct {
		difficulty []int
		profit     []int
		worker     []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				difficulty: []int{2, 4, 6, 8, 10},
				profit:     []int{10, 20, 30, 40, 50},
				worker:     []int{4, 5, 6, 7},
			},
			want: 100,
		},
		{
			name: "two",
			args: args{
				difficulty: []int{85, 47, 57},
				profit:     []int{24, 66, 99},
				worker:     []int{40, 25, 25},
			},
			want: 0,
		},
		{
			name: "three",
			args: args{
				difficulty: []int{66, 1, 28, 73, 53, 35, 45, 60, 100, 44, 59, 94, 27, 88, 7, 18, 83, 18, 72, 63},
				profit:     []int{66, 20, 84, 81, 56, 40, 37, 82, 53, 45, 43, 96, 67, 27, 12, 54, 98, 19, 47, 77},
				worker:     []int{61, 33, 68, 38, 63, 45, 1, 10, 53, 23, 66, 70, 14, 51, 94, 18, 28, 78, 100, 16},
			},
			want: 1392,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxProfitAssignment(test.args.difficulty, test.args.profit, test.args.worker)

			assert.Equal(t, test.want, got)
		})
	}
}
