package diameterofbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDiameterOfBinaryTree(t *testing.T) {
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
				Val: 1,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 4},
					Right: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 3},
			}},
			want: 3,
		},
		{
			name: "two",
			args: args{root: &TreeNode{
				Val:  1,
				Left: &TreeNode{Val: 2},
			}},
			want: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := diameterOfBinaryTree(test.args.root)

			assert.Equal(t, test.want, got)
		})
	}
}
