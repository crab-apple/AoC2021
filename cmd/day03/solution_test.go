package day03_test

import (
	"github.com/crab-apple/AoC2021/cmd/day03"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 198, day03.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 693486, day03.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, day03.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 0, day03.SolvePart2(input.ReadInputFile()))
}
