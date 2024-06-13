package reverselinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "one",
			args: args{head: createSingleLinkedList([]int{1, 2, 3, 4, 5, 6})},
			want: createSingleLinkedList([]int{6, 5, 4, 3, 2, 1}),
		},
		{
			name: "two",
			args: args{head: createSingleLinkedList([]int{1})},
			want: createSingleLinkedList([]int{1}),
		},
		{
			name: "three",
			args: args{head: createSingleLinkedList([]int{})},
			want: createSingleLinkedList([]int{}),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := reverseList(test.args.head)

			assert.Equal(t, test.want, got)
		})
	}
}

func createSingleLinkedList(nums []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, num := range nums {
		cur.Next = &ListNode{Val: num}
		cur = cur.Next
	}

	return head.Next
}
