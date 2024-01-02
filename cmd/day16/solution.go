package day16

import (
	"encoding/hex"
	"fmt"
)

func SolvePart1(s string) int {
	packet := NewPacket(s)
	//printInfo(packet, 0)
	return sumVersions(packet)
}

func printInfo(packet Packet, indent int) {

	printIndented := func(s string) {
		for i := 0; i < indent; i++ {
			fmt.Print("  ")
		}
		fmt.Println(s)
	}
	printIndented("------------")
	printIndented(fmt.Sprintf("Version: %v", packet.Version()))
	printIndented(fmt.Sprintf("Type: %v", packet.TypeId()))
	if packet.TypeId() == 4 {
		printIndented(fmt.Sprintf("Value: %v", packet.LiteralValue()))
	}
	for _, sp := range packet.SubPackets() {
		printInfo(sp, indent+1)
	}
}

func SolvePart2(s string) int {
	return 0
}

func sumVersions(p Packet) int {
	r := p.Version()
	for _, sp := range p.SubPackets() {
		r += sumVersions(sp)
	}
	return r
}

type Packet struct {
	bytes     []byte
	bitOffset int
}

func NewPacket(s string) Packet {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return Packet{bytes, 0}
}

func (p *Packet) Version() int {
	return int(p.GetBits(0, 3))
}

func (p *Packet) TypeId() int {
	return int(p.GetBits(3, 3))
}

func (p *Packet) GetBits(offset, len int) byte {
	if len > 8 {
		panic("unsupported")
	}

	effectiveOffset := p.bitOffset + offset
	leftLen := min(len, 8-effectiveOffset%8)
	rightLen := len - leftLen
	left := p.getByte(effectiveOffset/8, effectiveOffset%8, leftLen)
	right := byte(0)
	if rightLen > 0 {
		right = p.getByte(effectiveOffset/8+1, 0, rightLen)
	}
	return left<<rightLen + right
}

func (p *Packet) GetInt(offset, len int) int {
	if len > 64 {
		panic("unsupported")
	}

	if len <= 8 {
		return int(p.GetBits(offset, len))
	}
	return p.GetInt(offset, len-8)<<8 + p.GetInt(offset+len-8, 8)
}

func (p *Packet) getByte(numByte, offset, len int) byte {
	return p.bytes[numByte] << offset >> (8 - len)
}

func (p *Packet) LiteralValue() int {

	if p.TypeId() != 4 {
		panic("Unsupported (not literal)")
	}

	bytes := p.literalValueBits()

	result := 0
	for i, b := range bytes {
		result += +int(b) << ((len(bytes) - 1 - i) * 4)
	}
	return result
}

func (p *Packet) literalValueBits() []byte {

	if p.TypeId() != 4 {
		panic("Unsupported (not literal)")
	}

	bytes := make([]byte, 0)

	start := 6
	for {
		chunk := p.GetBits(start, 5)
		more := chunk >> 4
		bytes = append(bytes, chunk&0xF)
		if more == 0 {
			break
		}
		start += 5
	}

	return bytes
}

func (p *Packet) Length() int {

	length := 6

	if p.TypeId() == 4 {
		// Literal
		length += len(p.literalValueBits()) * 5
	} else {
		length += 1
		lengthTypeId := p.GetBits(6, 1)
		if lengthTypeId == 0 {
			length += 15
		} else {
			length += 11
		}
		for _, subpacket := range p.SubPackets() {
			length += subpacket.Length()
		}
	}
	return length
}

func (p *Packet) SubPackets() []Packet {

	subpackets := make([]Packet, 0)

	if p.TypeId() == 4 {
		// Literal
		return subpackets
	}

	lengthTypeId := p.GetBits(6, 1)
	if lengthTypeId == 0 {
		contentLength := p.GetInt(7, 15)
		subpacketOffset := 6 + 1 + 15
		for contentLength > 0 {
			subpacket := Packet{bytes: p.bytes, bitOffset: p.bitOffset + subpacketOffset}
			subpacketOffset += subpacket.Length()
			subpackets = append(subpackets, subpacket)
			contentLength -= subpacket.Length()
		}

	} else {
		numSubPackets := p.GetInt(7, 11)
		subPacketOffset := 6 + 1 + 11
		for {
			subpacket := Packet{bytes: p.bytes, bitOffset: p.bitOffset + subPacketOffset}
			subpackets = append(subpackets, subpacket)
			numSubPackets--
			if numSubPackets == 0 {
				break
			}
			subPacketOffset += subpacket.Length()
		}
	}
	return subpackets
}
