package mergetwosortedlists

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func createSingleLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}

	return head.Next
}

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		l1       *ListNode
		l2       *ListNode
		expected *ListNode
	}{
		{
			l1:       createSingleLinkedList([]int{1, 3, 5}),
			l2:       createSingleLinkedList([]int{2, 4, 8, 9}),
			expected: createSingleLinkedList([]int{1, 2, 3, 4, 5, 8, 9}),
		},
		{
			l1:       createSingleLinkedList([]int{1, 3, 5}),
			l2:       nil,
			expected: createSingleLinkedList([]int{1, 3, 5}),
		},
		{
			l1:       createSingleLinkedList([]int{1, 2, 3}),
			l2:       createSingleLinkedList([]int{2, 4}),
			expected: createSingleLinkedList([]int{1, 2, 2, 3, 4}),
		},
	}

	for _, test := range testCases {

		t.Run(t.Name(), func(t *testing.T) {

			result := mergeTwoLists(test.l1, test.l2)

			assert.Equal(t, test.expected, result)
		})
	}
}
