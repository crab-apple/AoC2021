package day05_test

import (
	"github.com/crab-apple/AoC2021/cmd/day05"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 5, day05.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 5585, day05.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, day05.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 0, day05.SolvePart2(input.ReadInputFile()))
}

func TestVent_Contains(t *testing.T) {
	vent := day05.Vent{day05.Coords{2, 3}, day05.Coords{20, 3}}

	assert.False(t, vent.Contains(day05.Coords{1, 3}))
	assert.False(t, vent.Contains(day05.Coords{21, 3}))
	assert.True(t, vent.Contains(day05.Coords{2, 3}))
	assert.True(t, vent.Contains(day05.Coords{20, 3}))
	assert.True(t, vent.Contains(day05.Coords{15, 3}))
	assert.False(t, vent.Contains(day05.Coords{15, 2}))
}

func TestVent_ContainsReversed(t *testing.T) {
	vent := day05.Vent{day05.Coords{20, 3}, day05.Coords{2, 3}}

	assert.False(t, vent.Contains(day05.Coords{1, 3}))
	assert.False(t, vent.Contains(day05.Coords{21, 3}))
	assert.True(t, vent.Contains(day05.Coords{2, 3}))
	assert.True(t, vent.Contains(day05.Coords{20, 3}))
	assert.True(t, vent.Contains(day05.Coords{15, 3}))
	assert.False(t, vent.Contains(day05.Coords{15, 2}))
}

func TestVentSet_ThereIsOverlap(t *testing.T) {

	vs := day05.VentSet{
		day05.Vent{day05.Coords{2, 3}, day05.Coords{2, 8}},
		day05.Vent{day05.Coords{2, 6}, day05.Coords{2, 10}},
	}

	assert.False(t, vs.ThereIsOverlap(day05.Coords{2, 0}))
	assert.False(t, vs.ThereIsOverlap(day05.Coords{2, 5}))
	assert.True(t, vs.ThereIsOverlap(day05.Coords{2, 6}))
	assert.True(t, vs.ThereIsOverlap(day05.Coords{2, 7}))
	assert.True(t, vs.ThereIsOverlap(day05.Coords{2, 8}))
	assert.False(t, vs.ThereIsOverlap(day05.Coords{2, 9}))
}
