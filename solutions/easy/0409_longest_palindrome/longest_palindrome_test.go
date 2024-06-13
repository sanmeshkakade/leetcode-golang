package longestpalindrome

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestPalindrome(t *testing.T) {
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
			args: args{s: "abccccdd"},
			want: 7,
		},
		{
			name: "tow",
			args: args{s: "abc"},
			want: 1,
		},
		{
			name: "one",
			args: args{s: "ababab"},
			want: 5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := longestPalindrome(test.args.s)

			assert.Equal(t, test.want, got)
		})
	}
}
