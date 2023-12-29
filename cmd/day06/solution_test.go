package day06_test

import (
	"github.com/crab-apple/AoC2021/cmd/day06"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
3,4,3,1,2
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 5934, day06.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 360610, day06.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, day06.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 0, day06.SolvePart2(input.ReadInputFile()))
}
