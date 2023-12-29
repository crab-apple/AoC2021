package day05

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/mathutils"
	"github.com/crab-apple/AoC2021/internal/utils"
	"strings"
)

func SolvePart1(s string) int {
	vents := VentSet(utils.Filter(parseInput(s), func(vent Vent) bool {
		return vent.IsHorizontalOrVertical()
	}))
	overlapCount := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			coords := Coords{i, j}
			if vents.ThereIsOverlap(coords) {
				overlapCount++
			}
		}
	}
	return overlapCount
}

func SolvePart2(s string) int {
	vents := VentSet(parseInput(s))
	overlapCount := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			coords := Coords{i, j}
			if vents.ThereIsOverlap(coords) {
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
	return Vent{
		parseCoords(strings.Split(s, " -> ")[0]),
		parseCoords(strings.Split(s, " -> ")[1]),
	}
}

func parseCoords(s string) Coords {
	return Coords{
		input.ToInt(strings.Split(s, ",")[0]),
		input.ToInt(strings.Split(s, ",")[1]),
	}
}

type VentSet []Vent

func (v VentSet) ThereIsOverlap(c Coords) bool {
	ventCount := utils.Count(v, func(vent Vent) bool {
		return vent.Contains(c)
	})
	return ventCount > 1
}

type Vent struct{ Start, End Coords }

func (v Vent) Contains(c Coords) bool {
	minX := min(v.Start.X, v.End.X)
	maxX := max(v.Start.X, v.End.X)
	minY := min(v.Start.Y, v.End.Y)
	maxY := max(v.Start.Y, v.End.Y)
	if minX == maxX || minY == maxY {
		if minX <= c.X && maxX >= c.X && minY <= c.Y && maxY >= c.Y {
			return true
		}
	}

	if c == v.Start {
		return true
	}

	vector := v.End.Minus(v.Start)
	relativeCoords := c.Minus(v.Start)

	return mathutils.Abs(relativeCoords.X) == mathutils.Abs(relativeCoords.Y) && relativeCoords.Sign() == vector.Sign() && mathutils.Abs(relativeCoords.X) <= mathutils.Abs(vector.X)
}

func (v Vent) IsHorizontalOrVertical() bool {
	return v.Start.X == v.End.X || v.Start.Y == v.End.Y
}

type Coords struct{ X, Y int }

func (c Coords) Minus(other Coords) Coords {
	return Coords{c.X - other.X, c.Y - other.Y}
}

func (c Coords) Sign() Coords {
	return Coords{mathutils.Sign(c.X), mathutils.Sign(c.Y)}
}
