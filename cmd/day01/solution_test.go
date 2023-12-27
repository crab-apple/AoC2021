package day01_test

import (
	"github.com/crab-apple/AoC2021/cmd/day01"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
199
200
208
210
200
207
240
269
260
263
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 7, day01.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 1557, day01.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 5, day01.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 1608, day01.SolvePart2(input.ReadInputFile()))
}
