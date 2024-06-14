/*
876. Middle of the Linked List
https://leetcode.com/problems/middle-of-the-linked-list/

Given the head of a singly linked list, return the middle node of the linked list.

If there are two middle nodes, return the second middle node.
*/

package middleoflinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(N)
// Space Complexity: O(1)

func middleNode(head *ListNode) *ListNode {

	var slow, fast = head, head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	return slow
}
