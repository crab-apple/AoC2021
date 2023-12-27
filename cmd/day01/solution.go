package day01

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
)

func SolvePart1(s string) int {
	depths := input.Read(s, input.ToInt)
	return utils.Count(utils.Pairs(depths), func(p utils.Pair[int]) bool {
		return p.Left < p.Right
	})
}

func SolvePart2(s string) int {
	depths := input.Read(s, input.ToInt)
	c := 0
	for i := 0; i < len(depths)-3; i++ {
		if depths[i] < depths[i+3] {
			c++
		}
	}
	return c
}
