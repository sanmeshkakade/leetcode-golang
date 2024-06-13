/*
232. Implement Queue using Stacks
https://leetcode.com/problems/implement-queue-using-stacks/

Implement a first in first out (FIFO) queue using only two stacks.
The implemented queue should support all the functions of a normal queue (push, peek, pop, and empty).

Implement the MyQueue class:

void push(int x) Pushes element x to the back of the queue.
int pop() Removes the element from the front of the queue and returns it.
int peek() Returns the element at the front of the queue.
boolean empty() Returns true if the queue is empty, false otherwise.

Notes:
You must use only standard operations of a stack, which means only push to top, peek/pop from top, size,
and is empty operations are valid.
Depending on your language, the stack may not be supported natively.
You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.
*/
package implementqueueusingstack

type MyQueue struct {
	pushStk, popStk *Stack
}

func Constructor() MyQueue {
	return MyQueue{
		pushStk: &Stack{},
		popStk:  &Stack{},
	}
}

// Time Complexity: O(1)
// Space Complexity: O(1)

func (q *MyQueue) Push(x int) {
	q.pushStk.Push(x)
}

// Time Complexity: O(1) (amortized)
// Space Complexity: O(1)

func (q *MyQueue) Pop() int {

	if q.popStk.Empty() {
		for !q.pushStk.Empty() {
			q.popStk.Push(q.pushStk.Pop())
		}

	}

	return q.popStk.Pop()
}

// Time Complexity: O(1) (amortized)
// Space Complexity: O(1)

func (q *MyQueue) Peek() int {
	if q.popStk.Empty() {
		for !q.pushStk.Empty() {
			q.popStk.Push(q.pushStk.Pop())
		}
	}
	return q.popStk.Peek()
}

func (q *MyQueue) Empty() bool {
	return q.pushStk.Empty() && q.popStk.Empty()
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
