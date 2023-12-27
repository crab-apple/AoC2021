package input

import (
	"github.com/crab-apple/AoC2021/internal/utils"
	"log"
	"os"
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
	return utils.Filter(
		utils.Map(
			strings.Split(input, "\n"),
			strings.TrimSpace,
		),
		utils.Not(utils.IsEmpty),
	)
}

func ToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Panic(err)
	}
	return i
}
