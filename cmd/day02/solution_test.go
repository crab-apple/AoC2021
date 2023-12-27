package day02_test

import (
	"github.com/crab-apple/AoC2021/cmd/day02"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
forward 5
down 5
forward 8
up 3
down 8
forward 2
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 150, day02.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 2150351, day02.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 900, day02.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 1842742223, day02.SolvePart2(input.ReadInputFile()))
}
