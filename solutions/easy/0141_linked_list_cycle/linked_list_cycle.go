/*
141. Linked List Cycle
https://leetcode.com/problems/linked-list-cycle/

Given head, the head of a linked list, determine if the linked list has a cycle in it.

There is a cycle in a linked list if there is some node in the list that can be reached again by continuously following the next pointer. Internally, pos is used to denote the index of the node that tail's next pointer is connected to. Note that pos is not passed as a parameter.

Return true if there is a cycle in the linked list. Otherwise, return false.
*/

package linkedlistcycle

type ListNode struct {
	Val  int
	Next *ListNode
}

// Time Complexity: O(N)
// Space Complexity: O(1)

func hasCycle(head *ListNode) bool {
	var (
		slow = head
		fast = head
	)

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}

// Time Complexity: O(N)
// Space Complexity: O(N)

func hasCycle1(head *ListNode) bool {
	var (
		cur = head
		set = map[*ListNode]struct{}{}
	)

	for cur != nil {
		set[cur] = struct{}{}
		cur = cur.Next
		if _, exist := set[cur]; exist {
			return true
		}
	}
	return false
}
