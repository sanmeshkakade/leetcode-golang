package ransomnote

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_canConstruct(t *testing.T) {
	type args struct {
		ransomNote string
		magazine   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				ransomNote: "a",
				magazine:   "b",
			},
			want: false,
		},
		{
			name: "two",
			args: args{
				ransomNote: "a",
				magazine:   "",
			},
			want: false,
		},
		{
			name: "one",
			args: args{
				ransomNote: "aa",
				magazine:   "aab",
			},
			want: true,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := canConstruct(test.args.ransomNote, test.args.magazine)

			assert.Equal(t, test.want, got)
		})
	}
}
