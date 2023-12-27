package day02

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"strings"
)

func SolvePart1(s string) int {
	var distance, depth int
	for _, command := range input.ReadLines(s) {
		cmdType := strings.Split(command, " ")[0]
		amt := input.ToInt(strings.Split(command, " ")[1])
		switch cmdType {
		case "forward":
			distance += amt
		case "up":
			depth -= amt
		case "down":
			depth += amt
		}
	}
	return depth * distance
}

func SolvePart2(s string) int {
	var aim, distance, depth int
	for _, command := range input.ReadLines(s) {
		cmdType := strings.Split(command, " ")[0]
		amt := input.ToInt(strings.Split(command, " ")[1])
		switch cmdType {
		case "forward":
			distance += amt
			depth += aim * amt
		case "up":
			aim -= amt
		case "down":
			aim += amt
		}
	}
	return depth * distance
}
