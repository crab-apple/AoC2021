package day03

import (
	"github.com/crab-apple/AoC2021/internal/input"
)

func SolvePart1(s string) int {

	lines := input.ReadLines(s)
	numBits := len(lines[0])
	var cnts [64]int

	for _, line := range lines {
		for i := range line {
			if line[i] == '1' {
				cnts[i]++
			}
		}
	}

	var gamma, epsilon int
	for i := 0; i < numBits; i++ {
		if cnts[i] > len(lines)/2 {
			gamma |= (1 << (numBits - 1)) >> i
		} else {
			epsilon |= (1 << (numBits - 1)) >> i
		}
	}

	return epsilon * gamma
}

func SolvePart2(s string) int {
	return 0
}
