package validparentheses

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {

	testCases := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "one",
			s:        "{}()[]",
			expected: true,
		},
		{
			name:     "two",
			s:        "{([])}",
			expected: true,
		},
		{
			name:     "three",
			s:        "{]",
			expected: false,
		},
		{
			name:     "four",
			s:        "{",
			expected: false,
		},
	}

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {
			result := isValid(test.s)

			assert.Equal(t, result, test.expected)
		})
	}
}
