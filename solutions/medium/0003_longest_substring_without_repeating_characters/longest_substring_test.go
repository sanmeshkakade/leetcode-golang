package longestsubstring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLongestSubstring(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{s: "abcabcbb"},
			want: 3,
		},
		{
			name: "two",
			args: args{s: "bbbbb"},
			want: 1,
		},
		{
			name: "three",
			args: args{s: "pwwkew"},
			want: 3,
		},
		{
			name: "four",
			args: args{s: " "},
			want: 1,
		},
		{
			name: "five",
			args: args{s: "abba"},
			want: 2,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := lengthOfLongestSubstring(test.args.s)

			assert.Equal(t, test.want, got)
		})
	}
}
