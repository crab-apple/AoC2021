package day13

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"github.com/crab-apple/AoC2021/internal/utils/grid"
	"strings"
)

func SolvePart1(s string) int {
	points, folds := parseInput(s)

	points = applyFold(points, folds[0])

	return len(points)
}

func SolvePart2(s string) int {
	return 0
}

func applyFold(points utils.Set[grid.GridPosition], fold Fold) utils.Set[grid.GridPosition] {
	r := utils.Set[grid.GridPosition]{}
	for _, point := range utils.Keys(points) {
		r.Add(ApplyFold(point, fold))
	}
	return r
}

func ApplyFold(point grid.GridPosition, fold Fold) grid.GridPosition {
	x := point.Col
	y := point.Row
	if fold.Axis == "x" {
		if x >= fold.Value {
			x = fold.Value*2 - x
		}
	}
	if fold.Axis == "y" {
		if y >= fold.Value {
			y = fold.Value*2 - y
		}
	}
	return grid.GridPosition{x, y}
}

func parseInput(s string) (utils.Set[grid.GridPosition], []Fold) {
	points := utils.Set[grid.GridPosition]{}
	folds := make([]Fold, 0)
	lines := input.ReadLines(s)
	for _, line := range lines {
		if utils.IsEmpty(line) {
			continue
		}
		if strings.HasPrefix(line, "fold") {
			fold := Fold{
				strings.Split(strings.Split(line, " ")[2], "=")[0],
				input.ToInt(strings.Split(strings.Split(line, " ")[2], "=")[1]),
			}
			folds = append(folds, fold)
		} else {
			point := grid.GridPosition{
				input.ToInt(strings.Split(line, ",")[0]),
				input.ToInt(strings.Split(line, ",")[1]),
			}
			points.Add(point)
		}
	}
	return points, folds
}

type Fold struct {
	Axis  string
	Value int
}
