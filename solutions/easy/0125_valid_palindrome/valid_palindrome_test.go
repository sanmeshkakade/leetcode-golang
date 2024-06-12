package validpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanatise(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected string
	}{
		{
			name:     "one",
			s:        "A man, a plan, a canal: Panama",
			expected: "amanaplanacanalpanama",
		},
		{
			name:     "two",
			s:        "race a car",
			expected: "raceacar",
		},
		{
			name:     "three",
			s:        " ",
			expected: "",
		},
		{
			name:     "four",
			s:        "0P",
			expected: "0p",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res := sanatise(test.s)

			assert.Equal(t, test.expected, res)
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		name     string
		s        string
		expected bool
	}{
		{
			name:     "one",
			s:        "A man, a plan, a canal: Panama",
			expected: true,
		},
		{
			name:     "two",
			s:        "race a car",
			expected: false,
		},
		{
			name:     "three",
			s:        " ",
			expected: true,
		},
		{
			name:     "four",
			s:        "0P",
			expected: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := isPalindrome(test.s)

			assert.Equal(t, test.expected, result)
		})
	}
}
