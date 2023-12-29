package search

import (
	"cmp"
	"github.com/crab-apple/AoC2021/internal/utils"
)

func FindMin[T cmp.Ordered](start, end int, mapper func(int) T) (int, T) {

	midPoint1 := start + (end-start)/3
	midPoint2 := start + (end-start)/3*2

	midValue1 := mapper(midPoint1)
	midValue2 := mapper(midPoint2)

	if midValue1 == midValue2 {
		results := make([]result[T], 0)
		for i := start; i <= end; i++ {
			results = append(results, result[T]{i, mapper(i)})
		}
		minResult := utils.MinBy(results, func(result result[T]) T { return result.v })
		return minResult.i, minResult.v
	}

	if midValue1 > midValue2 {
		return FindMin(midPoint1, end, mapper)
	}
	return FindMin(start, midPoint2, mapper)
}

type result[T any] struct {
	i int
	v T
}
