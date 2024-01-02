package day18

import (
	"bytes"
	"fmt"
	"github.com/crab-apple/AoC2021/internal/input"
	"strconv"
	"strings"
)

func SolvePart1(s string) int {
	numbers := input.Read(s, func(s string) SnailNumber {
		return ParseSnailNumber(s)
	})
	current := numbers[0]
	for i := 1; i < len(numbers); i++ {
		current = sum(current, numbers[i])
	}
	return current.Magnitude()
}

func SolvePart2(s string) int {
	numberStrs := input.ReadLines(s)

	maxVal := 0
	for i, ni := range numberStrs {
		for j, nj := range numberStrs {
			if i == j {
				continue
			}
			maxVal = max(maxVal, sum(ParseSnailNumber(ni), ParseSnailNumber(nj)).Magnitude())
			maxVal = max(maxVal, sum(ParseSnailNumber(nj), ParseSnailNumber(ni)).Magnitude())
		}
	}
	return maxVal
}

func sum(a, b SnailNumber) SnailNumber {
	result := &PairNumber{a, b}
	Reduce(result)
	return result
}

func Reduce(number SnailNumber) {
	for {
		ok := number.Explode()
		if ok {
			continue
		}
		ok, _ = number.Split()
		if !ok {
			break
		}
	}
}

type SnailNumber interface {
	String() string
	Explode() bool
	ExplodeLevels(levels int) (int, int, bool)
	Split() (bool, SnailNumber)
	SingleValue() int
	AddToRightMost(int) bool
	AddToLeftMost(int) bool
	Magnitude() int
}

type PairNumber struct {
	left, right SnailNumber
}

func (n *PairNumber) String() string {
	return fmt.Sprintf("[%v,%v]", n.left, n.right)
}

func (n *PairNumber) Explode() bool {
	_, _, ok := n.ExplodeLevels(4)
	return ok
}

func (n *PairNumber) ExplodeLevels(levels int) (int, int, bool) {

	if levels == 0 {
		leftVal := n.left.SingleValue()
		rightVal := n.right.SingleValue()

		return leftVal, rightVal, true
	}

	left, right, ok := n.left.ExplodeLevels(levels - 1)
	if ok {
		if levels == 1 {
			n.left = &SingleNumber{0}
		}
		if right != 0 && n.right.AddToLeftMost(right) {
			right = 0
		}
		return left, right, true
	}
	left, right, ok = n.right.ExplodeLevels(levels - 1)
	if ok {
		if levels == 1 {
			n.right = &SingleNumber{0}
		}
		if left != 0 && n.left.AddToRightMost(left) {
			left = 0
		}
		return left, right, true
	}

	return 0, 0, false
}

func (n *PairNumber) Split() (bool, SnailNumber) {
	ok, replacement := n.left.Split()
	if ok {
		if replacement != nil {
			n.left = replacement
		}
		return true, nil
	}
	ok, replacement = n.right.Split()
	if ok {
		if replacement != nil {
			n.right = replacement
		}
		return true, nil
	}
	return false, nil
}

func (n *PairNumber) AddToRightMost(val int) bool {

	ok := n.right.AddToRightMost(val)
	if !ok {
		ok = n.left.AddToRightMost(val)
	}
	return ok
}

func (n *PairNumber) AddToLeftMost(val int) bool {

	ok := n.left.AddToLeftMost(val)
	if !ok {
		ok = n.right.AddToLeftMost(val)
	}
	return ok
}

func (n *PairNumber) SingleValue() int {
	panic("Not a single number")
}

func (n *PairNumber) Magnitude() int {
	return 3*n.left.Magnitude() + 2*n.right.Magnitude()
}

type SingleNumber struct {
	val int
}

func (n *SingleNumber) String() string {
	return strconv.Itoa(n.val)
}
func (n *SingleNumber) Explode() bool {
	return false
}

func (n *SingleNumber) ExplodeLevels(levels int) (int, int, bool) {
	return 0, 0, false
}

func (n *SingleNumber) Split() (bool, SnailNumber) {
	if n.val > 9 {
		leftVal := n.val / 2
		rightVal := n.val - leftVal
		return true, &PairNumber{&SingleNumber{leftVal}, &SingleNumber{rightVal}}
	}
	return false, nil
}

func (n *SingleNumber) SingleValue() int {
	return n.val
}

func (n *SingleNumber) AddToRightMost(val int) bool {
	n.val += val
	return true
}

func (n *SingleNumber) AddToLeftMost(val int) bool {
	n.val += val
	return true
}

func (n *SingleNumber) Magnitude() int {
	return n.val
}

func ParseSnailNumber(str string) SnailNumber {

	if strings.HasPrefix(str, "[") {
		str = strings.TrimPrefix(str, "[")
		str = strings.TrimSuffix(str, "]")
		// Now find where the split is
		runes := bytes.Runes([]byte(str))
		splitPosition := 0
		bracketCount := 0
		for i, r := range runes {
			if r == ',' && bracketCount == 0 {
				splitPosition = i
				break
			}
			if r == '[' {
				bracketCount++
			}
			if r == ']' {
				bracketCount--
			}
		}
		leftStr := string(runes[:splitPosition])
		rightStr := string(runes[splitPosition+1:])
		return &PairNumber{ParseSnailNumber(leftStr), ParseSnailNumber(rightStr)}
	}

	return &SingleNumber{input.ToInt(str)}
}
