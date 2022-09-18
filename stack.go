package aho

type Stack struct {
	data []*node
}

func NewStack() *Stack {
	return &Stack{}

}

func (s *Stack) Pop() *node {
	if s.Empty() {
		panic("stack is empty")
	}

	lastIdx := len(s.data) - 1
	last := s.data[lastIdx]
	s.data = s.data[:lastIdx]
	return last
}

func (s *Stack) Push(n *node) {
	s.data = append(s.data, n)
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Empty() bool {
	return s.Size() == 0
}
