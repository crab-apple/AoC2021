package day07

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/mathutils"
	"github.com/crab-apple/AoC2021/internal/utils"
	"github.com/crab-apple/AoC2021/internal/utils/search"
	"strings"
)

func SolvePart1(s string) int {
	line := input.ReadLines(s)[0]
	crabs := crabArmy(utils.Map(strings.Split(line, ","), input.ToInt))
	_, fuel := search.FindMin(0, utils.Max(crabs), func(i int) int {
		return crabs.collectiveDistanceTo(i)
	})

	return fuel
}

func SolvePart2(s string) int {
	return 0
}

type crabArmy []int

func (ca crabArmy) collectiveDistanceTo(i int) int {
	r := 0
	for _, c := range ca {
		r += mathutils.Abs(i - c)
	}
	return r
}
