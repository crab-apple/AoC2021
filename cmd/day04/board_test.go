package day04_test

import (
	"github.com/crab-apple/AoC2021/cmd/day04"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoard_RemoveRow(t *testing.T) {
	board := day04.NewBoard([5][5]int{
		{100, 101, 102, 103, 104},
		{110, 111, 112, 113, 114},
		{120, 121, 122, 123, 124},
		{130, 131, 132, 133, 134},
		{140, 141, 142, 143, 144},
	})

	assert.False(t, board.Remove(100))
	assert.False(t, board.Remove(101))
	assert.False(t, board.Remove(102))
	assert.False(t, board.Remove(103))

	assert.False(t, board.Remove(121))
	assert.False(t, board.Remove(122))

	assert.True(t, board.Remove(104))
}

func TestBoard_RemoveCol(t *testing.T) {
	board := day04.NewBoard([5][5]int{
		{100, 101, 102, 103, 104},
		{110, 111, 112, 113, 114},
		{120, 121, 122, 123, 124},
		{130, 131, 132, 133, 134},
		{140, 141, 142, 143, 144},
	})

	assert.False(t, board.Remove(100))
	assert.False(t, board.Remove(110))
	assert.False(t, board.Remove(120))
	assert.False(t, board.Remove(130))

	assert.False(t, board.Remove(121))
	assert.False(t, board.Remove(122))

	assert.True(t, board.Remove(140))
}

func TestBoard_SumOfNumbersLeft(t *testing.T) {

	board := day04.NewBoard([5][5]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	})

	assert.Equal(t, 325, board.SumOfNumbersLeft())

	board.Remove(5)

	assert.Equal(t, 320, board.SumOfNumbersLeft())

	board.Remove(2)

	assert.Equal(t, 318, board.SumOfNumbersLeft())

	board.Remove(2)

	assert.Equal(t, 318, board.SumOfNumbersLeft())
}
