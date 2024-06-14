package maxdepthofbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one",
			args: args{root: &TreeNode{
				Val:  3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{
					Val:   20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7},
				},
			}},
			want: 3,
		},
		{
			name: "two",
			args: args{root: &TreeNode{
				Val:   1,
				Right: &TreeNode{Val: 2},
			}},
			want: 2,
		},
		{
			name: "three",
			args: args{root: nil},
			want: 0,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := maxDepth(test.args.root)

			assert.Equal(t, test.want, got)
		})
	}
}
