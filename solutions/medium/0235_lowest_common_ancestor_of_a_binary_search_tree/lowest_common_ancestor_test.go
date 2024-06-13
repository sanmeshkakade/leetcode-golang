package lca

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}

	tree1 := createBinaryTree([]int{2, 1, 3})
	tree2 := createBinaryTree([]int{2, 1})
	tree3 := createBinaryTree([]int{6, 2, 8, 0, 4, 7, 9, -1, 1, 3, 5})

	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "one",
			args: args{
				root: tree1,
				p:    tree1.Left,
				q:    tree1.Right,
			},
			want: tree1,
		},
		{
			name: "two",
			args: args{
				root: tree2,
				p:    tree2,
				q:    tree2.Left,
			},
			want: tree2,
		},
		{
			name: "three",
			args: args{
				root: tree3,
				p:    tree3.Left.Right.Left,
				q:    tree3.Left.Right.Right,
			},
			want: tree3.Left.Right,
		},
		{
			name: "four",
			args: args{
				root: tree3,
				p:    tree3.Left.Right.Left,
				q:    tree3.Right.Right,
			},
			want: tree3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := lowestCommonAncestor(test.args.root, test.args.p, test.args.q)

			assert.Equal(t, test.want, got)
		})
	}
}

func createBinaryTree(nums []int) *TreeNode {

	return performCreate(nums, 0)
}

func performCreate(nums []int, index int) *TreeNode {
	var root *TreeNode
	if index < len(nums) {
		root = &TreeNode{Val: nums[index]}
		left := 2*index + 1
		right := 2*index + 2
		root.Left = performCreate(nums, left)
		root.Right = performCreate(nums, right)
	}
	return root
}
