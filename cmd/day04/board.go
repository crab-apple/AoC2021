package day04

import (
	"github.com/crab-apple/AoC2021/internal/utils"
)

type Board struct {
	nums      map[int]utils.Pair[int]
	rowCounts [5]int
	colCounts [5]int
}

func (b *Board) Remove(number int) bool {
	position, ok := b.nums[number]
	if !ok {
		return false
	}
	delete(b.nums, number)
	b.rowCounts[position.Left] -= 1
	b.colCounts[position.Right] -= 1
	if b.rowCounts[position.Left] == 0 {
		return true
	}
	if b.colCounts[position.Right] == 0 {
		return true
	}
	return false
}

func (b *Board) SumOfNumbersLeft() int {
	return utils.Reduce(0, utils.Keys(b.nums), func(a int, b int) int {
		return a + b
	})
}

func NewBoard(nums [5][5]int) Board {
	mapped := make(map[int]utils.Pair[int])
	for rowNum := range nums {
		for colNum := range nums[rowNum] {
			num := nums[rowNum][colNum]
			mapped[num] = utils.Pair[int]{rowNum, colNum}
		}
	}
	return Board{
		mapped,
		[5]int{5, 5, 5, 5, 5},
		[5]int{5, 5, 5, 5, 5},
	}
}
