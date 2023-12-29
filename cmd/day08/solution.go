package day08

import (
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/crab-apple/AoC2021/internal/utils"
	"log"
	"math"
	"sort"
	"strings"
)

func SolvePart1(s string) int {
	count := 0
	for _, line := range input.ReadLines(s) {
		for _, i := range DecodeLine(line) {
			if i == 1 || i == 4 || i == 7 || i == 8 {
				count++
			}
		}
	}
	return count
}

func SolvePart2(s string) int {
	count := 0
	for _, line := range input.ReadLines(s) {
		for i, n := range DecodeLine(line) {
			decimalPos := 3 - i
			count += n * int(math.Pow10(decimalPos))
		}
	}
	return count
}

var repToDigit = map[string]int{
	"abcefg":  0,
	"cf":      1,
	"acdeg":   2,
	"acdfg":   3,
	"bcdf":    4,
	"abdfg":   5,
	"abdefg":  6,
	"acf":     7,
	"abcdefg": 8,
	"abcdfg":  9,
}

func DecodeLine(line string) []int {
	key := DecodeDigits(strings.Split(strings.Split(line, " | ")[0], " "))
	outputDigits := utils.Map(strings.Split(strings.Split(line, " | ")[1], " "), getChars)

	decodeDigit := func(digit utils.Set[rune]) int {
		decodedChars := make([]rune, 0)
		for ch := range digit {
			decodedChars = append(decodedChars, key[ch])
		}
		sort.Slice(decodedChars, func(i, j int) bool {
			return decodedChars[i] < decodedChars[j]
		})
		sb := strings.Builder{}
		for _, r := range decodedChars {
			sb.WriteRune(r)
		}
		i, ok := repToDigit[sb.String()]
		if !ok {
			log.Panicf("Missing: %s", sb.String())
		}
		return i
	}

	return utils.Map(outputDigits, decodeDigit)
}

func DecodeDigits(strings []string) map[rune]rune {

	encodedDigits := utils.Map(strings, getChars)

	encodedFrequencies := findFrequencies(encodedDigits)
	key := make(map[rune]rune)

	assign := func(encoded, unencoded rune) {
		key[encoded] = unencoded
		delete(encodedFrequencies, encoded)
	}

	assignByFreq := func(freq int, unencoded rune) {
		for char, f := range encodedFrequencies {
			if f == freq {
				assign(char, unencoded)
			}
		}
	}

	// Assign the digits with a unique frequency
	assignByFreq(4, 'e')
	assignByFreq(9, 'f')
	assignByFreq(6, 'b')

	encoded1 := utils.Single(encodedDigits, func(s utils.Set[rune]) bool {
		return len(s) == 2
	})
	encoded7 := utils.Single(encodedDigits, func(s utils.Set[rune]) bool {
		return len(s) == 3
	})
	encoded4 := utils.Single(encodedDigits, func(s utils.Set[rune]) bool {
		return len(s) == 4
	})

	for char := range encoded7 {
		if !encoded1.Contains(char) {
			assign(char, 'a')
		}
	}

	assignByFreq(8, 'c')

	// Only remaining are 'd' and 'g'. Number 4 contains 'd' but doesn't contain 'g'
	for char := range encoded4 {
		_, ok := encodedFrequencies[char]
		if ok {
			assign(char, 'd')
		}
	}

	for char := range encodedFrequencies {
		assign(char, 'g')
	}

	return key
}

func findFrequencies(sets []utils.Set[rune]) map[rune]int {
	freqs := make(map[rune]int)
	for char := 'a'; char <= 'g'; char++ {
		freqs[char] = 0
	}
	for i := range sets {
		for char := range sets[i] {
			freqs[char]++
		}
	}
	return freqs
}

func getChars(s string) utils.Set[rune] {
	set := make(utils.Set[rune])
	for _, char := range s {
		set.Add(char)
	}
	return set
}
