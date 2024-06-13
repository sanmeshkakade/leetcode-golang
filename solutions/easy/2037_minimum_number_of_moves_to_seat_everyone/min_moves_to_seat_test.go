package minmoves

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinMovesToSeat(t *testing.T) {

	funcs := []func([]int, []int) int{minMovesToSeat, minMovesToSeat1}

	type args struct {
		seats    []int
		students []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{
				seats:    []int{3, 1, 5},
				students: []int{2, 7, 4},
			},
			want: 4,
		},
		{
			name: "two",
			args: args{
				seats:    []int{4, 1, 5, 9},
				students: []int{1, 3, 2, 6},
			},
			want: 7,
		},
		{
			name: "three",
			args: args{
				seats:    []int{2, 2, 6, 6},
				students: []int{1, 3, 2, 6},
			},
			want: 4,
		},
	}

	for key, f := range funcs {
		for _, test := range tests {
			t.Run(fmt.Sprintf("%s_%d", test.name, key), func(t *testing.T) {
				got := f(test.args.seats, test.args.students)

				assert.Equal(t, got, test.want)
			})
		}
	}
}
