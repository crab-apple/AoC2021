package day12_test

import (
	"github.com/crab-apple/AoC2021/cmd/day12"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
start-A
start-b
A-c
A-b
b-d
A-end
b-end
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 10, day12.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 3887, day12.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, day12.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 0, day12.SolvePart2(input.ReadInputFile()))
}
