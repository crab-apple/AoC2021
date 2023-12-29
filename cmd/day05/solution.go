package day05

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/mathutils"
	"github.com/crab-apple/AoC2021/internal/utils"
	"strings"
)

func SolvePart1(s string) int {
	vents := utils.Filter(parseInput(s), func(vent Vent) bool {
		return vent.IsHorizontalOrVertical()
	})
	return countOverlaps(vents)
}

func SolvePart2(s string) int {
	vents := parseInput(s)
	return countOverlaps(vents)
}

func countOverlaps(vents []Vent) int {

	counts := make(map[Coords]int)
	for i := range vents {
		for c, _ := range vents[i].points {
			counts[c]++
		}
	}

	overlapCount := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			coords := Coords{i, j}
			if counts[coords] > 1 {
				overlapCount++
			}
		}
	}
	return overlapCount
}

func parseInput(s string) []Vent {
	return input.Read(s, parseVent)
}

func parseVent(s string) Vent {
	return NewVent(
		parseCoords(strings.Split(s, " -> ")[0]),
		parseCoords(strings.Split(s, " -> ")[1]),
	)
}

func parseCoords(s string) Coords {
	return Coords{
		input.ToInt(strings.Split(s, ",")[0]),
		input.ToInt(strings.Split(s, ",")[1]),
	}
}

type Vent struct {
	Start, End Coords
	points     utils.Set[Coords]
}

func NewVent(start, end Coords) Vent {
	sign := end.Minus(start).Sign()

	points := make(utils.Set[Coords])

	current := start
	points.Add(current)
	for current != end {
		current = current.Plus(sign)
		points.Add(current)
	}

	return Vent{start, end, points}
}

func (v Vent) IsHorizontalOrVertical() bool {
	return v.Start.X == v.End.X || v.Start.Y == v.End.Y
}

type Coords struct{ X, Y int }

func (c Coords) Plus(other Coords) Coords {
	return Coords{c.X + other.X, c.Y + other.Y}
}

func (c Coords) Minus(other Coords) Coords {
	return Coords{c.X - other.X, c.Y - other.Y}
}

func (c Coords) Sign() Coords {
	return Coords{mathutils.Sign(c.X), mathutils.Sign(c.Y)}
}
