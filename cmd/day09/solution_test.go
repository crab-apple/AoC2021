package day09_test

import (
	"github.com/crab-apple/AoC2021/cmd/day09"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
2199943210
3987894921
9856789892
8767896789
9899965678
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 15, day09.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 558, day09.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, day09.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 0, day09.SolvePart2(input.ReadInputFile()))
}
