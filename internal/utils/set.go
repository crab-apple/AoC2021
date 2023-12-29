package utils

type Set[T comparable] map[T]interface{}

func (s Set[T]) Add(t T) {
	s[t] = nil
}

func (s Set[T]) Contains(t T) bool {
	_, ok := s[t]
	return ok
}
