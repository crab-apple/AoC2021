package day09

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"unicode/utf8"
)

func SolvePart1(s string) int {

	grid := NewGrid(s, func(r rune) int {
		return input.ToInt(string(r))
	})

	count := 0

	for i := 0; i < grid.NumCols(); i++ {
		for j := 0; j < grid.NumRows(); j++ {
			h := grid[i][j]
			lowest := true
			if i > 0 && grid[i-1][j] <= h {
				lowest = false
			}
			if i < grid.NumCols()-1 && grid[i+1][j] <= h {
				lowest = false
			}
			if j > 0 && grid[i][j-1] <= h {
				lowest = false
			}
			if j < grid.NumRows()-1 && grid[i][j+1] <= h {
				lowest = false
			}
			if lowest {
				count += h + 1
			}
		}
	}

	return count
}

func SolvePart2(s string) int {
	return 0
}

type Grid[T any] [][]T

func (g Grid[T]) NumCols() int {
	return len(g)
}

func (g Grid[T]) NumRows() int {
	return len(g[0])
}

func NewGrid[T any](s string, mapper func(rune) T) Grid[T] {
	lines := input.ReadLines(s)
	numCols := utf8.RuneCountInString(lines[0])
	grid := make([][]T, numCols)
	for i := 0; i < numCols; i++ {
		grid[i] = make([]T, len(lines))
	}
	for rowNum, line := range lines {
		colNum := 0
		for _, r := range line {
			grid[colNum][rowNum] = mapper(r)
			colNum++
		}
	}
	return grid
}
