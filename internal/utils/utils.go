package utils

func Map[T, U any](ts []T, f func(T) U) []U {
	us := make([]U, len(ts))
	for i, t := range ts {
		us[i] = f(t)
	}
	return us
}

func Filter[T any](ts []T, p func(T) bool) []T {
	r := make([]T, 0)
	for _, t := range ts {
		if p(t) {
			r = append(r, t)
		}
	}
	return r
}

func Reduce[T, U any](init U, ts []T, f func(U, T) U) U {
	u := init
	for _, t := range ts {
		u = f(u, t)
	}
	return u
}

func Sum(is []int) int {
	return Reduce(0, is, func(a int, b int) int {
		return a + b
	})
}

func Count[T any](ts []T, p func(T) bool) int {
	c := 0
	for _, t := range ts {
		if p(t) {
			c++
		}
	}
	return c
}

func Not[T any](p func(T) bool) func(T) bool {
	return func(t T) bool {
		return !p(t)
	}
}

func IsEmpty(s string) bool {
	return s == ""
}
func Pairs[T any](ts []T) []Pair[T] {
	r := make([]Pair[T], len(ts))
	for i := range ts {
		if i < len(ts)-1 {
			r[i] = Pair[T]{ts[i], ts[i+1]}
		}
	}
	return r
}
