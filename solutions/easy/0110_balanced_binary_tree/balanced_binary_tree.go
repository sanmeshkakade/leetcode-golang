/*
110. Balanced Binary Tree
https://leetcode.com/problems/balanced-binary-tree/

Given a binary tree, determine if it is height-balanced
A height-balanced binary tree is a binary tree in which the depth of the two subtrees of every node never differs by more than one.
*/

package balancedbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	return height(root) != -1
}

// Recursion
// Time Complexity: O(N) where N is no. of nodes in the tree
// Space Complexity: O(h) where h is height of the three

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := height(root.Left)
	right := height(root.Right)

	if left == -1 || right == -1 {
		return -1
	}

	if abs(left-right) > 1 {
		return -1
	}

	return 1 + max(left, right)
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
