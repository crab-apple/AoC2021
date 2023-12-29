package day07

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/mathutils"
	"github.com/crab-apple/AoC2021/internal/utils"
	"github.com/crab-apple/AoC2021/internal/utils/search"
	"strings"
)

func SolvePart1(s string) int {
	return solve(s, utils.Identity[int])
}

func SolvePart2(s string) int {
	return solve(s, func(i int) int {
		return i * (i + 1) / 2
	})
}

func solve(s string, fModel FuelConsumptionModel) int {
	line := input.ReadLines(s)[0]
	crabs := CrabArmy(utils.Map(strings.Split(line, ","), input.ToInt))
	_, fuel := search.FindMin(0, utils.Max(crabs), func(i int) int {
		return crabs.collectiveDistanceTo(i, fModel)
	})
	return fuel
}

type FuelConsumptionModel func(int) int

type CrabArmy []int

func (ca CrabArmy) collectiveDistanceTo(i int, model FuelConsumptionModel) int {
	r := 0
	for _, c := range ca {
		r += model(mathutils.Abs(c - i))
	}
	return r
}
