package day09

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"github.com/crab-apple/AoC2021/internal/utils/grid"
	"sort"
)

func SolvePart1(s string) int {

	g := parseGrid(s)

	count := 0

	for _, p := range findLowPoints(g) {
		h := g.Get(p)
		count += h + 1
	}

	return count
}

func SolvePart2(s string) int {
	g := parseGrid(s)

	lowPoints := findLowPoints(g)
	basinSizes := utils.Map(lowPoints, func(lowPoint grid.GridPosition) int {
		return findBasinSize(g, lowPoint)
	})

	sort.Ints(basinSizes)
	basinSizes = basinSizes[len(basinSizes)-3:]

	return utils.Reduce(1, basinSizes, func(a int, b int) int {
		return a * b
	})
}

func parseGrid(s string) grid.Grid[int] {
	return grid.NewGrid(s, func(r rune) int {
		return input.ToInt(string(r))
	})
}
func findLowPoints(g grid.Grid[int]) []grid.GridPosition {
	return utils.Filter(g.Positions(), func(p grid.GridPosition) bool {
		h := g.Get(p)
		lowest := true
		for _, n := range p.FourNeighbours() {
			if g.HasPosition(n) && g.Get(n) <= h {
				lowest = false
			}
		}
		return lowest
	})
}

func findBasinSize(g grid.Grid[int], lowPoint grid.GridPosition) int {
	seen := make(utils.Set[grid.GridPosition])
	pending := make([]grid.GridPosition, 0)
	pending = append(pending, lowPoint)
	for len(pending) != 0 {
		p := pending[0]
		pending = pending[1:]
		if !seen.Contains(p) && g.HasPosition(p) && g.Get(p) != 9 {
			seen.Add(p)
			for _, n := range p.FourNeighbours() {
				pending = append(pending, n)
			}
		}
	}
	return len(seen)
}
