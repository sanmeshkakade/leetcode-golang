package implementqueueusingstack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackUsingQueueOne(t *testing.T) {

	q := Constructor()

	q.Push(1)
	assert.False(t, q.Empty())
	assert.Equal(t, 1, q.Peek())

	q.Push(2)
	assert.False(t, q.Empty())
	assert.Equal(t, 1, q.Peek())

	popped := q.Pop()
	assert.Equal(t, 1, popped)
	assert.False(t, q.Empty())
	assert.Equal(t, 2, q.Peek())

	popped2 := q.Pop()
	assert.Equal(t, 2, popped2)
	assert.True(t, q.Empty())

}
