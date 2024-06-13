package implementqueueusingstack

// stack implementation

type Stack []int

func (s *Stack) Push(e int) {
	*s = append(*s, e)
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() int {
	e := s.Peek()
	*s = (*s)[:len(*s)-1]
	return e
}

func (s *Stack) Empty() bool {
	return len(*s) == 0
}
