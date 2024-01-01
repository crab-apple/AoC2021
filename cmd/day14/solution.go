package day14

import (
	"bytes"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
)

func SolvePart1(s string) int {
	return solve(s, 10)
}

func SolvePart2(s string) int {
	return solve(s, 40)
}

func solve(s string, steps int) int {
	lines := input.ReadLines(s)
	template := lines[0]
	rules := make(Rules)

	for _, line := range lines[2:] {
		runes := bytes.Runes([]byte(line))
		rules.addRule(runes[0], runes[1], runes[6])
	}

	counts := RuneCount{}
	insertCounter := InsertCounter{rules: rules}
	runes := bytes.Runes([]byte(template))
	for i := range runes {
		counts[runes[i]-'A']++
		if i < len(runes)-1 {
			counts.add(insertCounter.countInserted(InsertCounterParam{runes[i], runes[i+1], steps}))
		}
	}

	return counts.diff()
}

type InsertCounter struct {
	rules Rules
	cache map[InsertCounterParam]RuneCount
}

type InsertCounterParam struct {
	a, b     rune
	numSteps int
}

func (ic InsertCounter) countInserted(param InsertCounterParam) RuneCount {

	if ic.cache == nil {
		ic.cache = make(map[InsertCounterParam]RuneCount)
	}

	result, ok := ic.cache[param]
	if ok {
		return result
	}

	result = RuneCount{}
	toInsert := ic.rules.getRule(param.a, param.b)
	if toInsert == 0 {
		return result
	}
	result[toInsert-'A']++
	if param.numSteps > 1 {
		result.add(ic.countInserted(InsertCounterParam{param.a, toInsert, param.numSteps - 1}))
		result.add(ic.countInserted(InsertCounterParam{toInsert, param.b, param.numSteps - 1}))
	}

	ic.cache[param] = result

	return result
}

type RuneCount [32]int

func (r *RuneCount) add(other RuneCount) {
	for i := range r {
		r[i] += other[i]
	}
}

func (r *RuneCount) plus(other *RuneCount) RuneCount {
	result := RuneCount{}
	for i := range result {
		result[i] = r[i] + other[i]
	}
	return result
}

func (r *RuneCount) diff() int {

	maxCount := utils.Max(utils.Filter(r[:], func(i int) bool {
		return i != 0
	}))
	minCount := utils.Min(utils.Filter(r[:], func(i int) bool {
		return i != 0
	}))

	return maxCount - minCount
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
