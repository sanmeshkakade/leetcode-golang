package validanagram

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				s: "rat",
				t: "cat",
			},
			want: false,
		},
		{
			name: "three",
			args: args{
				s: "a",
				t: "ab",
			},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isAnagram(test.args.s, test.args.t)

			assert.Equal(t, test.want, got)
		})
	}
}
