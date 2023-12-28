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

	var gamma2, epsilon2 int
	for i := range cnts {
		if cnts[i] == 0 {
			break
		}
		if cnts[i] > len(nums)/2 {
			gamma2 |= 1 << i
		} else {
			epsilon2 |= 1 << i
		}
	}
	gamma, epsilon := gamma2, epsilon2
	return epsilon * gamma
}

func SolvePart2(s string) int {
	numBits := len(input.ReadLines(s)[0])
	nums := parseInput(s)

	var cnts []int
	var funnel []int
	var wNumBits int

	funnel = nums
	wNumBits = numBits
	for len(funnel) != 1 {
		cnts = findFrequencies(funnel, wNumBits)
		var threshold int
		if len(funnel)%2 == 0 {
			threshold = len(funnel) / 2
		} else {
			threshold = (len(funnel) + 1) / 2
		}

		if cnts[len(cnts)-1] >= threshold {
			funnel = utils.Filter(funnel, func(i int) bool {
				return digitAt(i, wNumBits-1) == 1
			})
		} else {
			funnel = utils.Filter(funnel, func(i int) bool {
				return digitAt(i, wNumBits-1) == 0
			})
		}
		wNumBits--
	}

	oxigenRating := funnel[0]

	funnel = nums
	wNumBits = numBits
	for len(funnel) != 1 {
		cnts = findFrequencies(funnel, wNumBits)
		var threshold int
		if len(funnel)%2 == 0 {
			threshold = len(funnel) / 2
		} else {
			threshold = (len(funnel) + 1) / 2
		}

		if cnts[len(cnts)-1] < threshold {
			funnel = utils.Filter(funnel, func(i int) bool {
				return digitAt(i, wNumBits-1) == 1
			})
		} else {
			funnel = utils.Filter(funnel, func(i int) bool {
				return digitAt(i, wNumBits-1) == 0
			})
		}
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

func findFrequencies(nums []int, numSignificantBits int) []int {

	cnts := make([]int, numSignificantBits)

	for _, line := range nums {
		for i := range cnts {
			if line&(1<<i) != 0 {
				cnts[i]++
			}
		}
	}
	return cnts
}
