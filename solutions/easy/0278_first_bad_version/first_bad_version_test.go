package firstbadversion

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_firstBadVersion(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name         string
		args         args
		versions     []int
		isBadVersion func(version int) bool
		want         int
	}{
		{
			name:     "one",
			args:     args{n: 5},
			versions: []int{0, 0, 0, 1, 1},
			want:     4,
		},
		{
			name:     "two",
			args:     args{n: 1},
			versions: []int{1},
			want:     1,
		},
		{
			name:     "three",
			args:     args{n: 7},
			versions: []int{0, 1, 1, 1, 1, 1, 1},
			want:     2,
		},
	}
	for _, test := range tests {
		isBadVersion = func(version int) bool {
			return test.versions[version-1] == 1
		}
		t.Run(test.name, func(t *testing.T) {
			got := firstBadVersion(test.args.n)

			assert.Equal(t, test.want, got)
		})
	}
}
