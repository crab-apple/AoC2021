package day15

import (
	"bytes"
	"fmt"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"github.com/crab-apple/AoC2021/internal/utils/grid"
	"strconv"
	"strings"
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

	lines := input.ReadLines(s)
	lines = utils.Map(lines, func(line string) string {
		return line + inc(line, 1) + inc(line, 2) + inc(line, 3) + inc(line, 4)
	})

	orig := lines
	lines = append(lines, incLines(orig, 1)...)
	lines = append(lines, incLines(orig, 2)...)
	lines = append(lines, incLines(orig, 3)...)
	lines = append(lines, incLines(orig, 4)...)

	sb := strings.Builder{}

	for _, line := range lines {
		sb.WriteString(line + "\n")
	}

	fmt.Println(sb.String())

	return SolvePart1(sb.String())

}

func inc(line string, i int) string {
	sb := strings.Builder{}
	for _, r := range bytes.Runes([]byte(line)) {
		val := input.ToInt(string(r))
		sb.WriteString(strconv.Itoa((val-1+i)%9 + 1))
	}

	return sb.String()
}

func incLines(lines []string, i int) []string {
	return utils.Map(lines, func(line string) string {
		return inc(line, i)
	})
}
