/*
206. Reverse Linked List
https://leetcode.com/problems/reverse-linked-list/

Given the head of a singly linked list, reverse the list, and return the reversed list.
*/

package reverselinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(N)
// Space Complexity: O(1)

func reverseList(head *ListNode) *ListNode {
	var (
		cur   = head
		dummy *ListNode
	)

	for cur != nil {
		temp := cur.Next
		cur.Next = dummy
		dummy = cur
		cur = temp
	}

	return dummy
}
