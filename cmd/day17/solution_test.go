package day17_test

import (
	"github.com/crab-apple/AoC2021/cmd/day17"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
target area: x=20..30, y=-10..-5
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 45, day17.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 11175, day17.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 112, day17.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 3540, day17.SolvePart2(input.ReadInputFile()))
}
