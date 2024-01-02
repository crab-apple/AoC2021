package day18_test

import (
	"fmt"
	"github.com/crab-apple/AoC2021/cmd/day18"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

var testInput = `
[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]
`

func TestPart1(t *testing.T) {
	assert.Equal(t, 4140, day18.SolvePart1(testInput))
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 4072, day18.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, day18.SolvePart2(testInput))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 0, day18.SolvePart2(input.ReadInputFile()))
}

func TestParse(t *testing.T) {
	str := "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"
	snailNumber := day18.ParseSnailNumber(str)
	assert.Equal(t, str, snailNumber.String())
}

func TestExplode(t *testing.T) {
	cases := []TestCase{
		{"[[[[[9,8],1],2],3],4]", "[[[[0,9],2],3],4]"},
		{"[7,[6,[5,[4,[3,2]]]]]", "[7,[6,[5,[7,0]]]]"},
		{"[[6,[5,[4,[3,2]]]],1]", "[[6,[5,[7,0]]],3]"},
		{"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]"},
		{"[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]", "[[3,[2,[8,0]]],[9,[5,[7,0]]]]"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v should become %v after explosion", c.input, c.output), func(t *testing.T) {
			snailNumber := day18.ParseSnailNumber(c.input)
			snailNumber.Explode()
			assert.Equal(t, c.output, snailNumber.String())
		})
	}
}

func TestSplit(t *testing.T) {
	cases := []TestCase{
		{"[[[[0,7],4],[15,[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,13]]],[1,1]]"},
		{"[[[[0,7],4],[[7,8],[0,13]]],[1,1]]", "[[[[0,7],4],[[7,8],[0,[6,7]]]],[1,1]]"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v should become %v after split", c.input, c.output), func(t *testing.T) {
			snailNumber := day18.ParseSnailNumber(c.input)
			snailNumber.Split()
			assert.Equal(t, c.output, snailNumber.String())
		})
	}
}

func TestReduce(t *testing.T) {
	cases := []TestCase{
		{"[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]", "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]"},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v should become %v after reduction", c.input, c.output), func(t *testing.T) {
			snailNumber := day18.ParseSnailNumber(c.input)
			day18.Reduce(snailNumber)
			assert.Equal(t, c.output, snailNumber.String())
		})
	}
}

func TestMagnitude(t *testing.T) {
	cases := []struct {
		input  string
		output int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%v should become %v after reduction", c.input, c.output), func(t *testing.T) {
			snailNumber := day18.ParseSnailNumber(c.input)
			assert.Equal(t, c.output, snailNumber.Magnitude())
		})
	}
}

type TestCase struct {
	input, output string
}
