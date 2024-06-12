package invertbinarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	testCases := []struct {
		name     string
		root     *TreeNode
		inverted *TreeNode
	}{
		{
			name: "one",
			root: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 1},
					Right: &TreeNode{Val: 3},
				},
				Right: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 6},
					Right: &TreeNode{Val: 9},
				},
			},
			inverted: &TreeNode{
				Val: 4,
				Left: &TreeNode{
					Val:   7,
					Left:  &TreeNode{Val: 9},
					Right: &TreeNode{Val: 6},
				},
				Right: &TreeNode{
					Val:   2,
					Left:  &TreeNode{Val: 3},
					Right: &TreeNode{Val: 1},
				},
			},
		},
		{
			name: "two",
			root: &TreeNode{
				Val:   4,
				Left:  &TreeNode{Val: 2},
				Right: nil,
			},
			inverted: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: &TreeNode{Val: 2},
			},
		},
		{
			name:     "three",
			root:     nil,
			inverted: nil,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			res := invertTree(test.root)

			assert.Equal(t, test.inverted, res)
		})
	}
}
