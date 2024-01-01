package day14

import (
	"bytes"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"strings"
)

func SolvePart1(s string) int {
	lines := input.ReadLines(s)
	template := lines[0]
	rules := make(Rules)

	for _, line := range lines[2:] {
		runes := bytes.Runes([]byte(line))
		rules.addRule(runes[0], runes[1], runes[6])
	}

	for i := 0; i < 10; i++ {
		template = expand(template, rules)
	}

	counts := make([]int, 100)
	for _, r := range bytes.Runes([]byte(template)) {
		counts[r-'A']++
	}

	maxCount := utils.Max(utils.Filter(counts, func(i int) bool {
		return i != 0
	}))
	minCount := utils.Min(utils.Filter(counts, func(i int) bool {
		return i != 0
	}))

	return maxCount - minCount
}

func expand(template string, rules Rules) string {
	sb := strings.Builder{}
	runes := bytes.Runes([]byte(template))
	for i := range runes {
		current := runes[i]
		sb.WriteRune(current)
		if i == len(runes)-1 {
			break
		}
		next := runes[i+1]
		toInsert := rules.getRule(current, next)
		if toInsert != 0 {
			sb.WriteRune(toInsert)
		}
	}
	return sb.String()
}

func SolvePart2(s string) int {
	return 0
}

type Rules map[rune]map[rune]rune

func (r Rules) addRule(a, b, c rune) {
	nestedMap, ok := r[a]
	if !ok {
		nestedMap = make(map[rune]rune)
		r[a] = nestedMap
	}
	nestedMap[b] = c
}

func (r Rules) getRule(a, b rune) rune {
	nestedMap, ok := r[a]
	if !ok {
		return 0
	}
	result, ok := nestedMap[b]
	if !ok {
		return 0
	}
	return result
}
