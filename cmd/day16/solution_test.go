package day16_test

import (
	"fmt"
	"github.com/crab-apple/AoC2021/cmd/day16"
	"github.com/crab-apple/AoC2021/internal/input"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPart1(t *testing.T) {

	examples := []Part1TestCase{
		{
			"8A004A801A8002F478",
			16,
		},
		{
			"620080001611562C8802118E34",
			12,
		},
		{
			"C0015000016115A2E0802F182340",
			23,
		},
		{
			"A0016C880162017C3686B18A3D4780",
			31,
		},
	}

	for i, testCase := range examples {
		t.Run(fmt.Sprintf("Example %v", i), func(t *testing.T) {
			assert.Equal(t, testCase.wantOutput, day16.SolvePart1(testCase.input))
		})
	}
}

func TestPart1RealInput(t *testing.T) {
	assert.Equal(t, 904, day16.SolvePart1(input.ReadInputFile()))
}

func TestPart2(t *testing.T) {
	assert.Equal(t, 0, day16.SolvePart2(""))
}

func TestPart2RealInput(t *testing.T) {
	assert.Equal(t, 0, day16.SolvePart2(input.ReadInputFile()))
}

type Part1TestCase struct {
	input      string
	wantOutput int
}

func TestPacketVersion(t *testing.T) {
	packet := day16.NewPacket("D2FE28")
	assert.Equal(t, 6, packet.Version())
}

func TestTypeId(t *testing.T) {
	packet := day16.NewPacket("D2FE28")
	assert.Equal(t, 4, packet.TypeId())
}

func TestLiteralValue(t *testing.T) {
	packet := day16.NewPacket("D2FE28")
	assert.Equal(t, 2021, packet.LiteralValue())
}

func TestLength(t *testing.T) {
	t.Run("Literal value", func(t *testing.T) {
		packet := day16.NewPacket("D2FE28")
		assert.Equal(t, 21, packet.Length())
	})

	t.Run("Operator length type 0", func(t *testing.T) {
		packet := day16.NewPacket("38006F45291200")
		assert.Equal(t, 49, packet.Length())
	})

	t.Run("Operator length type 1", func(t *testing.T) {
		packet := day16.NewPacket("EE00D40C823060")
		assert.Equal(t, 51, packet.Length())
	})

}
