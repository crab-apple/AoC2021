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

	var cnts [][2]int
	var funnel []int
	var wNumBits int

	funnel = nums
	wNumBits = numBits
	for len(funnel) != 1 {
		cnts = findFrequencies(funnel, wNumBits)
		var filter int
		if cnts[len(cnts)-1][1] >= cnts[len(cnts)-1][0] {
			filter = 1
		} else {
			filter = 0
		}
		funnel = utils.Filter(funnel, func(i int) bool {
			return digitAt(i, wNumBits-1) == filter
		})
		wNumBits--
	}

	oxigenRating := funnel[0]

	funnel = nums
	wNumBits = numBits
	for len(funnel) != 1 {
		cnts = findFrequencies(funnel, wNumBits)
		var filter int
		if cnts[len(cnts)-1][1] < cnts[len(cnts)-1][0] {
			filter = 1
		} else {
			filter = 0
		}
		funnel = utils.Filter(funnel, func(i int) bool {
			return digitAt(i, wNumBits-1) == filter
		})
		wNumBits--
	}

	co2Rating := funnel[0]

	return oxigenRating * co2Rating
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
