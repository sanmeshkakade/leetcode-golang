/*
543. Diameter of Binary Tree
https://leetcode.com/problems/diameter-of-binary-tree/

Given the root of a binary tree, return the length of the diameter of the tree.

The diameter of a binary tree is the length of the longest path between any two nodes in a tree. This path may or may not pass through the root.

The length of a path between two nodes is represented by the number of edges between them.
*/
package diameterofbinarytree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Time Complexity: O(N) N -> no. of nodes in the tree
// Space Complexity: O(h) in call stack where h -> height of tree

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var dia = 0
	height(root, &dia)
	return dia
}

func height(root *TreeNode, dia *int) int {
	if root == nil {
		return 0
	}
	left := height(root.Left, dia)
	right := height(root.Right, dia)
	*dia = max(*dia, left+right)
	return 1 + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
