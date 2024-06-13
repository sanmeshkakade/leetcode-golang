/*
235. Lowest Common Ancestor of a Binary Search Tree
https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/

Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes p and q as
the lowest node in T that has both p and q as descendants (where we allow a node to be a descendant of itself).”
*/

package lca

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursion
// Time Complexity: O(h) where h is height of binary tree
// Space Complexity: O(h) at max h calls will be present in call stack

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q {
		return root
	}
	if p.Val > q.Val {
		return lowestCommonAncestor(root, q, p)
	}
	if root.Val > p.Val && root.Val < q.Val {
		return root
	} else if root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	} else {
		return lowestCommonAncestor(root.Right, p, q)
	}
}
