package binarytree

type stack struct {
	top  int
	size int
	data []any
}

func (s *stack) enlargeCap() {
	s.size *= 2
	var res = make([]any, s.size)

	copy(res, s.data)

	s.data = res
}

func (s *stack) pop() any {
	if s.isEmpty() {
		return nil
	}

	v := s.data[s.top]
	s.top--
	return v
}

func (s *stack) push(v any) {
	if s.isFull() {
		s.enlargeCap()
	}

	s.top++
	s.data[s.top] = v
}

func (s *stack) isEmpty() bool {
	return s.top == -1
}

func (s *stack) isFull() bool {
	return s.top == s.size-1
}
