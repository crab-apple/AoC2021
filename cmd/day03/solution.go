package day03

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"strconv"
)

func SolvePart1(s string) int {
	numBits := len(input.ReadLines(s)[0])
	nums := parseInput(s)
	cnts := findFrequencies(nums, numBits)

	var gamma, epsilon int
	for i := range cnts {
		if cnts[i][1] > len(nums)/2 {
			gamma |= 1 << i
		} else {
			epsilon |= 1 << i
		}
	}
	return epsilon * gamma
}

func SolvePart2(s string) int {
	numBits := len(input.ReadLines(s)[0])
	nums := parseInput(s)

	oxigenRating := funnelInts(nums, numBits, func(freqs [2]int) int {
		if freqs[1] >= freqs[0] {
			return 1
		} else {
			return 0
		}
	})

	co2Rating := funnelInts(nums, numBits, func(freqs [2]int) int {
		if freqs[1] < freqs[0] {
			return 1
		} else {
			return 0
		}
	})

	return oxigenRating * co2Rating
}

func funnelInts(ints []int, numBits int, bitCriteria func(freqs [2]int) int) int {
	for len(ints) != 1 {
		cnts := findFrequencies(ints, numBits)
		var filter int
		filter = bitCriteria(cnts[len(cnts)-1])
		ints = utils.Filter(ints, func(i int) bool {
			return digitAt(i, numBits-1) == filter
		})
		numBits--
	}
	return ints[0]
}

func digitAt(n int, digitPosition int) int {
	return 1 & (n >> digitPosition)
}

func parseInput(s string) []int {
	return input.Read(s, func(line string) int {
		i, _ := strconv.ParseInt(line, 2, 64)
		return int(i)
	})
}

func findFrequencies(nums []int, numSignificantBits int) [][2]int {

	cnts := make([][2]int, numSignificantBits)

	for _, line := range nums {
		for i := range cnts {
			if line&(1<<i) != 0 {
				cnts[i][1]++
			} else {
				cnts[i][0]++
			}
		}
	}
	return cnts
}
