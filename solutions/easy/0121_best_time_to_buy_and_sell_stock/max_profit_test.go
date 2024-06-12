package maxprofit

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxprofit(t *testing.T) {
	testCases := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "one",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "two",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "three",
			prices:   []int{7, 7, 7, 7, 7},
			expected: 0,
		},
	}

	funcs := []func([]int) int{maxProfit, maxProfit1}

	for key, f := range funcs {
		for _, test := range testCases {
			t.Run(fmt.Sprintf("%s_%d", test.name, key), func(t *testing.T) {

				result := f(test.prices)

				assert.Equal(t, test.expected, result)

			})
		}
	}
}
