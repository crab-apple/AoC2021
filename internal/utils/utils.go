package utils

import (
	"cmp"
	"log"
)

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

func Single[T any](ts []T, p func(T) bool) T {
	ts = Filter(ts, p)
	if len(ts) != 1 {
		log.Panicf("Espected one element, found %v", len(ts))
	}
	return ts[0]
}

func Reduce[T, U any](init U, ts []T, f func(U, T) U) U {
	u := init
	for _, t := range ts {
		u = f(u, t)
	}
	return u
}

func Values[T comparable, U any](m map[T]U) []U {
	r := make([]U, len(m))
	for _, v := range m {
		r = append(r, v)
	}
	return r
}

func Keys[T comparable, U any](m map[T]U) []T {
	r := make([]T, len(m))
	for k, _ := range m {
		r = append(r, k)
	}
	return r
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

func Max[T cmp.Ordered](ts []T) T {
	return MaxBy(ts, Identity[T])
}

func Min[T cmp.Ordered](ts []T) T {
	return MinBy(ts, Identity[T])
}

func MaxBy[T any, U cmp.Ordered](ts []T, mapper func(T) U) T {
	return MaxByComp(ts, func(a T, b T) int {
		return cmp.Compare(mapper(a), mapper(b))
	})
}

func MinBy[T any, U cmp.Ordered](ts []T, mapper func(T) U) T {
	return MinByComp(ts, func(a T, b T) int {
		return cmp.Compare(mapper(a), mapper(b))
	})
}

func MaxByComp[T any](ts []T, compare func(T, T) int) T {
	var maxResult T
	var haveResult bool
	for i := range ts {
		r := ts[i]
		if !haveResult || compare(r, maxResult) > 0 {
			maxResult = r
			haveResult = true
		}
	}
	return maxResult
}

func MinByComp[T any](ts []T, compare func(T, T) int) T {
	return MaxByComp(ts, func(a T, b T) int {
		return -compare(a, b)
	})
}

func Identity[T any](t T) T {
	return t
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
