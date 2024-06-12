package twosum

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {

	funcs := []func([]int, int) []int{twoSum, twoSum1}

	testCases := []struct {
		name     string
		nums     []int
		target   int
		expected []int
	}{
		{
			name:     "one",
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		{
			name:     "two",
			nums:     []int{2, 7, 11, 15},
			target:   6,
			expected: []int{},
		},
	}

	for _, f := range funcs {
		for _, test := range testCases {
			t.Run(test.name, func(t *testing.T) {

				result := f(test.nums, test.target)

				assert.ElementsMatch(t, result, test.expected)
			})
		}
	}

}
