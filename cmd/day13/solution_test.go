package day13_test

import (
	"github.com/crab-apple/AoC2021/cmd/day13"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils/grid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 17, day13.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 655, day13.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, day13.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 0, day13.SolvePart2(input.ReadInputFile()))
}

func TestApplyFold(t *testing.T) {
	p := grid.GridPosition{3, 4}
	assert.Equal(t, grid.GridPosition{3, 4}, day13.ApplyFold(p, day13.Fold{"y", 5}))
	assert.Equal(t, grid.GridPosition{3, 2}, day13.ApplyFold(p, day13.Fold{"y", 3}))
	assert.Equal(t, grid.GridPosition{3, 0}, day13.ApplyFold(p, day13.Fold{"y", 2}))
}
