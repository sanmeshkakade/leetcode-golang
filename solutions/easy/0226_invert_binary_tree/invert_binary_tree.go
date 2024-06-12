/*
226. Invert Binary Tree
https://leetcode.com/problems/invert-binary-tree/

Given the root of a binary tree, invert the tree, and return its root.
*/

package invertbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursion
// Time Complexity: O(N) where N is number of nodes in the tree
// Space Complexity: O(h) in call stack where h is the height of the tree

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	invertTree(root.Left)
	invertTree(root.Right)
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	return root
}
