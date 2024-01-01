package day12_test

import (
	"github.com/crab-apple/AoC2021/cmd/day12"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput1 = `
start-A
start-b
A-c
A-b
b-d
A-end
b-end
`

var testInput2 = `
dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc
`
var testInput3 = `
fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW
`

func TestPart1(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, 10, day12.SolvePart1(testInput1))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, 19, day12.SolvePart1(testInput2))
	})
	t.Run("Example 3", func(t *testing.T) {
		assert.Equal(t, 226, day12.SolvePart1(testInput3))
	})
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 3887, day12.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		assert.Equal(t, 36, day12.SolvePart2(testInput1))
	})
	t.Run("Example 2", func(t *testing.T) {
		assert.Equal(t, 103, day12.SolvePart2(testInput2))
	})
	t.Run("Example 3", func(t *testing.T) {
		assert.Equal(t, 3509, day12.SolvePart2(testInput3))
	})
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 104834, day12.SolvePart2(input.ReadInputFile()))
}
