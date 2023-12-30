package day11

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"github.com/crab-apple/AoC2021/internal/utils/grid"
)

func SolvePart1(s string) int {

	octopuses := grid.NewGrid(s, func(r rune) int {
		return input.ToInt(string(r))
	})

	count := 0
	for i := 0; i < 100; i++ {
		count += doStep(octopuses)
	}
	return count
}

func SolvePart2(s string) int {
	return 0
}

func doStep(octopuses grid.Grid[int]) int {

	toIncrease := utils.Queue[grid.GridPosition]{}
	flashed := utils.Set[grid.GridPosition]{}

	for _, p := range octopuses.Positions() {
		toIncrease.Add(p)
	}

	for toIncrease.Len() != 0 {
		current := toIncrease.RemoveFirst()

		val := octopuses.Get(current)
		if flashed.Contains(current) {
			continue
		}
		if val == 9 {
			flashed.Add(current)
			for _, n := range current.EightNeighbours() {
				if octopuses.HasPosition(n) {
					toIncrease.Add(n)
				}
			}
			octopuses.Set(current, 0)
			continue
		}
		octopuses.Set(current, val+1)
	}

	return len(flashed)
}
