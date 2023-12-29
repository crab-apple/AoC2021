package search_test

import (
	"github.com/crab-apple/AoC2021/internal/mathutils"
	"github.com/crab-apple/AoC2021/internal/utils/search"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMin(t *testing.T) {

	foundVal, _ := search.FindMin(-100000000, 100000000, func(i int) int {
		return mathutils.Abs(i - 12345)
	})

	assert.Equal(
		t,
		12345,
		foundVal,
	)
}

func TestFindMin_Flat(t *testing.T) {

	mapper := func(i int) int {
		if i == 12345 {
			return 0
		}
		return (mathutils.Abs(i-12345) + 100) / 100
	}
	foundVal, _ := search.FindMin(-100000000, 100000000, mapper)

	assert.Equal(
		t,
		12345,
		foundVal,
	)
}
