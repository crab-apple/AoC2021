package day15

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"github.com/crab-apple/AoC2021/internal/utils/grid"
)

func SolvePart1(s string) int {

	g := grid.NewGrid(s, func(r rune) int {
		return input.ToInt(string(r))
	})

	costs := make(map[grid.GridPosition]int)
	pending := utils.Queue[grid.GridPosition]{}

	startingPosition := grid.GridPosition{0, 0}
	costs[startingPosition] = 0
	pending.Add(startingPosition)

	for pending.Len() != 0 {
		current := pending.RemoveFirst()
		for _, n := range current.FourNeighbours() {
			if g.HasPosition(n) {
				cost := costs[current] + g.Get(n)
				oldCost, ok := costs[n]
				if !ok || oldCost > cost {
					costs[n] = cost
					pending.Add(n)
				}
			}
		}
	}

	return costs[grid.GridPosition{g.NumCols() - 1, g.NumRows() - 1}]

}

func SolvePart2(s string) int {
	return 0
}
