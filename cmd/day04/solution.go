package day04

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"log"
	"slices"
	"strings"
)

func SolvePart1(s string) int {
	drawnNumbers, boards := parseInput(s)

	for _, drawnNumber := range drawnNumbers {
		for i := range boards {
			if boards[i].Remove(drawnNumber) {
				return drawnNumber * boards[i].SumOfNumbersLeft()
			}
		}
	}

	log.Panic("Impossible")
	return 0
}

func SolvePart2(s string) int {
	drawnNumbers, boards := parseInput(s)
	score := 0
	var won []int

	for _, drawnNumber := range drawnNumbers {
		for i := range boards {
			if !slices.Contains(won, i) && boards[i].Remove(drawnNumber) {
				score = drawnNumber * boards[i].SumOfNumbersLeft()
				won = append(won, i)
			}
		}
	}

	return score
}

func parseInput(s string) ([]int, []Board) {
	lines := input.ReadLines(s)
	drawnNumbers := utils.Map(strings.Split(lines[0], ","), input.ToInt)

	numBoards := (len(lines) - 1) / 6
	boards := make([]Board, numBoards)

	for boardNum := 0; boardNum < numBoards; boardNum++ {
		var nums [5][5]int
		for i := 0; i < 5; i++ {
			line := lines[2+i+boardNum*6]
			rowNums := utils.Map(input.RegexSplit(line, "\\s+"), input.ToInt)
			copy(nums[i][:], rowNums)
		}
		boards[boardNum] = NewBoard(nums)
	}

	return drawnNumbers, boards
}
