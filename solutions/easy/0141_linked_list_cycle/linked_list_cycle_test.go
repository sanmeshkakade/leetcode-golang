package linkedlistcycle

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}

	l1 := generateLinkedList([]int{3, 2, 0, -4})
	l1.Next.Next.Next.Next = l1.Next
	l2 := generateLinkedList([]int{1, 2, 3, 4, 5, 6})

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "one",
			args: args{
				head: nil,
			},
			want: false,
		},
		{
			name: "two",
			args: args{
				head: l1,
			},
			want: true,
		},
		{
			name: "three",
			args: args{
				head: l2,
			},
			want: false,
		},
	}

	funcs := []func(*ListNode) bool{hasCycle, hasCycle1}

	for key, f := range funcs {
		for _, test := range tests {
			t.Run(fmt.Sprintf("%s_%d", test.name, key), func(t *testing.T) {
				got := f(test.args.head)

				assert.Equal(t, test.want, got)
			})
		}
	}
}

func generateLinkedList(nums []int) *ListNode {
	var dummy = &ListNode{Val: -1}
	var cur = dummy

	for _, val := range nums {
		cur.Next = &ListNode{Val: val}
		cur = cur.Next
	}

	return dummy.Next
}
