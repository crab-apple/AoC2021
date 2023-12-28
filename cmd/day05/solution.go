package day05

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"log"
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
	return 0
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
	if !v.IsHorizontalOrVertical() {
		log.Panic("Not implemented")
	}
	return contained(v.Start.X, v.End.X, c.X) && contained(v.Start.Y, v.End.Y, c.Y)
}

func contained(rangeStart int, rangeEnd int, i int) bool {
	if rangeStart > rangeEnd {
		temp := rangeEnd
		rangeEnd = rangeStart
		rangeStart = temp
	}
	return rangeStart <= i && rangeEnd >= i
}

func (v Vent) IsHorizontalOrVertical() bool {
	return v.Start.X == v.End.X || v.Start.Y == v.End.Y
}

type Coords struct{ X, Y int }
