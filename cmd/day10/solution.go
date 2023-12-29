package day10

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"sort"
)

var corruptedScoreSheet = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var incompleteScoreSheet = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func SolvePart1(s string) int {
	return utils.Sum(utils.Map(input.ReadLines(s), scoreCorrupted))
}

func SolvePart2(s string) int {
	scores := utils.Filter(utils.Map(input.ReadLines(s), scoreIncomplete), func(i int) bool {
		return i != 0
	})
	sort.Ints(scores)
	return scores[len(scores)/2]
}

func scoreCorrupted(line string) int {
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
				return corruptedScoreSheet[r]
			}
		}
	}
	return 0
}
func scoreIncomplete(line string) int {
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
				return 0
			}
		}
	}
	score := 0
	for len(expected) != 0 {
		score *= 5
		score += incompleteScoreSheet[expected.Pop()]
	}
	return score
}
