package day24_test

import (
	"github.com/crab-apple/AoC2021/cmd/day24"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
Test input here
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 0, day24.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 0, day24.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, day24.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 0, day24.SolvePart2(input.ReadInputFile()))
}
