package day17

import (
	"fmt"
	"github.com/crab-apple/AoC2021/internal/input"
	"strings"
)

func SolvePart1(s string) int {
	line := input.ReadLines(s)[0]
	line = strings.TrimPrefix(line, "target area: ")
	minX, maxX := parseRange(strings.Split(line, ", ")[0])
	minY, _ := parseRange(strings.Split(line, ", ")[1])

	// Can we hit the target at x velocity 0?
	candidateX := 0
	for stop(candidateX) < minX {
		candidateX++
	}
	if stop(candidateX) > maxX {
		panic("CANNOT hit the target at x speed 0")
	}

	// We can hit the target at X speed 0, so we can ignore the X now. We just need to find the max Y we can achieve
	if minY > 0 {
		panic("Not implemented")
	}

	maxYSPeed := -minY - 1

	return stop(maxYSPeed)
}

func SolvePart2(s string) int {
	line := input.ReadLines(s)[0]
	line = strings.TrimPrefix(line, "target area: ")
	minX, maxX := parseRange(strings.Split(line, ", ")[0])
	minY, maxY := parseRange(strings.Split(line, ", ")[1])

	minXSpeed := 0
	for stop(minXSpeed) < minX {
		minXSpeed++
	}
	maxXSpeed := maxX
	maxYSpeed := -minY - 1
	minYSpeed := minY

	fmt.Printf("minX: %v, maxX:%v, minY:%v, maxY:%v\n", minXSpeed, maxXSpeed, minYSpeed, maxYSpeed)

	evaluate := func(xSpeed, ySpeed int) bool {
		var x, y int
		for {
			if x >= minX && x <= maxX && y >= minY && y <= maxY {
				return true
			}
			if x > maxX || y < minY {
				return false
			}
			x += xSpeed
			y += ySpeed
			if xSpeed > 0 {
				xSpeed--
			}
			ySpeed--
		}
	}

	count := 0
	for xSpeed := minXSpeed; xSpeed <= maxXSpeed; xSpeed++ {
		for ySpeed := minYSpeed; ySpeed <= maxYSpeed; ySpeed++ {
			if evaluate(xSpeed, ySpeed) {
				count++
			}
		}
	}

	return count

}

func parseRange(s string) (int, int) {
	min := input.ToInt(strings.TrimLeft(strings.Split(s, "..")[0], "xy="))
	max := input.ToInt(strings.TrimLeft(strings.Split(s, "..")[1], "xy="))
	return min, max
}

func stop(speed int) int {
	r := 0
	for speed > 0 {
		r += speed
		speed--
	}
	return r
}
