package input

import (
	"github.com/crab-apple/AoC2021/internal/utils"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Read[T any](input string, f func(string) T) []T {
	return utils.Map(ReadLines(input), f)
}

func ReadInputFile() string {
	b, err := os.ReadFile("input.txt")
	if err != nil {
		log.Panic(err)
	}
	return string(b)
}

func ReadLines(input string) []string {
	lines := utils.Map(
		strings.Split(input, "\n"),
		strings.TrimSpace,
	)
	if utils.IsEmpty(lines[0]) {
		lines = lines[1:]
	}
	if utils.IsEmpty(lines[len(lines)-1]) {
		lines = lines[:len(lines)-1]
	}
	return lines
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}

func RegexSplit(s string, regex string) []string {
	return regexp.MustCompile(regex).Split(s, -1)
}
