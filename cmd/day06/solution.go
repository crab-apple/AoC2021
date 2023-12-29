package day06

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"strings"
)

func SolvePart1(s string) int {

	line := input.ReadLines(s)[0]

	individualFish := utils.Map(strings.Split(line, ","), input.ToInt)
	var school School

	for _, fish := range individualFish {
		school[fish]++
	}

	for i := 0; i < 80; i++ {
		school.advance()
	}

	return school.total()
}

func SolvePart2(s string) int {
	return 0
}

type School [9]int

func (s *School) advance() {
	zeroes := s[0]
	for i := 1; i < len(s); i++ {
		s[i-1] = s[i]
	}
	s[len(s)-1] = zeroes
	s[len(s)-3] += zeroes
}

func (s *School) total() int {
	return utils.Sum(s[:])
}
