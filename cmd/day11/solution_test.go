package day11_test

import (
	"github.com/crab-apple/AoC2021/cmd/day11"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 1656, day11.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 1625, day11.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 195, day11.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 244, day11.SolvePart2(input.ReadInputFile()))
}
