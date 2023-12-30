package grid

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"unicode/utf8"
)

type Grid[T any] [][]T

func (g Grid[T]) NumCols() int {
	return len(g)
}

func (g Grid[T]) NumRows() int {
	return len(g[0])
}

func (g Grid[T]) Get(p GridPosition) T {
	return g[p.Col][p.Row]
}

func (g Grid[T]) Set(p GridPosition, t T) {
	g[p.Col][p.Row] = t
}

func (g Grid[T]) Positions() []GridPosition {
	r := make([]GridPosition, g.NumRows()*g.NumCols())
	k := 0
	for nc := 0; nc < g.NumCols(); nc++ {
		for nr := 0; nr < g.NumRows(); nr++ {
			r[k] = GridPosition{nc, nr}
			k++
		}
	}
	return r
}

func (g Grid[T]) HasPosition(gp GridPosition) bool {
	return gp.Col >= 0 && gp.Col < g.NumCols() && gp.Row >= 0 && gp.Row < g.NumRows()
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

type GridPosition struct {
	Col, Row int
}

func (gp GridPosition) North() GridPosition {
	return GridPosition{gp.Col, gp.Row - 1}
}
func (gp GridPosition) South() GridPosition {
	return GridPosition{gp.Col, gp.Row + 1}
}
func (gp GridPosition) West() GridPosition {
	return GridPosition{gp.Col - 1, gp.Row}
}
func (gp GridPosition) East() GridPosition {
	return GridPosition{gp.Col + 1, gp.Row}
}

func (gp GridPosition) FourNeighbours() []GridPosition {
	return []GridPosition{gp.North(), gp.South(), gp.West(), gp.East()}
}

func (gp GridPosition) EightNeighbours() []GridPosition {
	return []GridPosition{
		gp.North(),
		gp.North().West(),
		gp.North().East(),
		gp.South(),
		gp.South().West(),
		gp.South().East(),
		gp.West(),
		gp.East(),
	}
}
