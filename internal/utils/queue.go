package utils

type Queue[T any] struct {
	len   int
	first *node[T]
	last  *node[T]
}

func (s *Queue[T]) Len() int {
	return s.len
}

func (s *Queue[T]) Add(t T) {
	s.len++

	n := node[T]{t: t}
	if s.first == nil {
		s.first = &n
		s.last = &n
	} else {
		s.last.next = &n
		s.last = &n
	}
}

func (s *Queue[T]) RemoveFirst() T {
	s.len--

	if s.first == nil {
		panic("Empty queue")
	}

	if s.first == s.last {
		n := *s.first
		s.first = nil
		s.last = nil
		return n.t
	}

	n := *s.first
	s.first = s.first.next

	return n.t
}

type node[T any] struct {
	t    T
	next *node[T]
}
