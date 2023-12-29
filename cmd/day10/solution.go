package day10

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
)

var scoreSheet = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func SolvePart1(s string) int {
	return utils.Sum(utils.Map(input.ReadLines(s), score))
}

func SolvePart2(s string) int {
	return 0
}

func score(line string) int {
	expected := utils.Stack[rune]{}
	for _, r := range line {
		switch r {
		case '(':
			expected.Push(')')
		case '[':
			expected.Push(']')
		case '{':
			expected.Push('}')
		case '<':
			expected.Push('>')
		default:
			if expected.Pop() != r {
				return scoreSheet[r]
			}
		}
	}
	return 0
}
