package sortcolors

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	testCases := []struct {
		name     string
		nums     []int
		expected []int
	}{
		{
			name:     "one",
			nums:     []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		{
			name:     "two",
			nums:     []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		{
			name:     "three",
			nums:     []int{2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 0},
			expected: []int{0, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2},
		},
	}

	for _, test := range testCases{
		t.Run(test.name, func (t *testing.T) {

			sortColors(test.nums)

			if !reflect.DeepEqual(test.nums, test.expected){
				t.Errorf("expected %v, got %v", test.nums, test.expected)
			}
		})
	}
}
