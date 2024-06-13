package balancedbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				root: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val:   20,
						Left:  &TreeNode{Val: 15},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: true,
		},
		{
			name: "two",
			args: args{
				root: &TreeNode{
					Val:  3,
					Left: &TreeNode{Val: 9},
					Right: &TreeNode{
						Val: 20,
						Left: &TreeNode{
							Val:  15,
							Left: &TreeNode{Val: 21},
						},
						Right: &TreeNode{Val: 7},
					},
				},
			},
			want: false,
		},
		{
			name: "three",
			args: args{
				root: &TreeNode{
					Val: 3,
					Right: &TreeNode{
						Val: 20,
						Right: &TreeNode{
							Val:   7,
							Right: &TreeNode{Val: 13},
						},
					},
				},
			},
			want: false,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := isBalanced(test.args.root)

			assert.Equal(t, test.want, got)
		})
	}
}
