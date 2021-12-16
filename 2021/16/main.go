package main

import (
	"encoding/hex"
	"fmt"
	"math"
	"sort"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 16, solve1, solve2)
}

type Packet struct {
	Version    int
	Type       int
	Literal    int
	SubPackets []*Packet
}

func (p *Packet) SumVersions() int {
	sum := p.Version
	if p.SubPackets != nil {
		for _, sp := range p.SubPackets {
			sum += sp.SumVersions()
		}
	}
	return sum
}

const (
	pSum     = 0
	pProd    = 1
	pMin     = 2
	pMax     = 3
	pLiteral = 4
	pGreater = 5
	pLess    = 6
	pEqual   = 7
)

func (p *Packet) Process() (v int) {
	switch p.Type {
	case pLiteral:
		v = p.Literal
	case pSum:
		for _, sp := range p.SubPackets {
			v += sp.Process()
		}
	case pProd:
		v = 1
		for _, sp := range p.SubPackets {
			v *= sp.Process()
		}
	case pMin:
		s := make([]int, len(p.SubPackets))
		for i := range s {
			s[i] = p.SubPackets[i].Process()
		}
		sort.Ints(s)
		return s[0]
	case pMax:
		s := make([]int, len(p.SubPackets))
		for i := range s {
			s[i] = p.SubPackets[i].Process()
		}
		sort.Sort(sort.Reverse(sort.IntSlice(s)))
		return s[0]
	case pGreater:
		if p.SubPackets[0].Process() > p.SubPackets[1].Process() {
			v = 1
		}
	case pLess:
		if p.SubPackets[0].Process() < p.SubPackets[1].Process() {
			v = 1
		}
	case pEqual:
		if p.SubPackets[0].Process() == p.SubPackets[1].Process() {
			v = 1
		}
	default:
		fmt.Printf("p.Type == %d\n", p.Type)
		panic("unknown type")
	}
	return
}

// offset is the offset IN BITS into the data
func readPacket(data []byte, offset int) (*Packet, int) {
	p := &Packet{}

	// packet version is the first three bits
	p.Version, offset = readUint3(data, offset)
	p.Type, offset = readUint3(data, offset)

	switch p.Type {
	case pLiteral:
		// a literal.
		p.Literal, offset = readVarint(data, offset)
		//fmt.Printf("readPacket: isLiteral, Value=%d\n", p.Literal)
	default:
		// ANYTHING ELSE is op
		var l bool
		l, offset = readBit(data, offset)
		//fmt.Printf("readPacket: isOperator, LengthType=%v\n", l)
		if l {
			// 11bits representing number of sub-packets to read
			var n int
			n, offset = readUint11(data, offset)
			//fmt.Printf("readPacket: should have %d subpackets\n", n)
			p.SubPackets = make([]*Packet, n)
			for i := 0; i < n; i++ {
				p.SubPackets[i], offset = readPacket(data, offset)
			}
		} else {
			// next 15 bits is length in BITS of the of the sub-packets
			var n int
			n, offset = readUint15(data, offset)
			target := n + offset
			//fmt.Printf("readPacket: should have subpacket length=%d (target offset=%d)\n", n, target)
			p.SubPackets = []*Packet{}
			var next *Packet
			for {
				if target == offset {
					// done
					break
				}
				if offset > target {
					panic("READ PAST TARGET OFFSET!")
				}
				next, offset = readPacket(data, offset)
				p.SubPackets = append(p.SubPackets, next)
			}
		}
	}
	return p, offset
}

func readBit(data []byte, offset int) (bool, int) {
	start := offset / 8
	index := offset % 8
	// shift the bit to the far right and mask it
	b := (data[start] >> (7 - index)) & 0x01
	return b == 1, offset + 1
}

const (
	maskUint3  = 0xFF >> 5           // b00000111
	maskUint4  = 0xFF >> 4           // b00001111
	maskUint11 = math.MaxUint16 >> 5 // b0000011111111111
	maskUint15 = math.MaxUint16 >> 1 // b0111111111111111
)

func readUint3(data []byte, offset int) (int, int) {
	// read 3 bits at the offset.
	// they could be in 1 or 2 bytes, depending on the offset
	start := offset / 8
	index := offset % 8
	offset += 3
	switch {
	case index == 5:
		// the stars have aligned!
		return int(data[start] & maskUint3), offset
	case index < 5:
		// all bits are in one byte
		// so we can read as a uint8
		// once we have shifted and masked (first 5 bits)
		v := (data[start] >> (5 - index))
		v &= maskUint3
		return int(v), offset
	default:
		// this is the case where the bits are in 2 bytes
		// the easiest way is to use a uint16
		v := (uint16(data[start]) << 8)
		v |= uint16(data[start+1])
		// now shift it to the right and mask
		v = (v >> (13 - index))
		v &= maskUint3
		return int(v), offset
	}
}

func readUint4(data []byte, offset int) (int, int) {
	// read 4 bits at the offset.
	// they could be in 1 or 2 bytes, depending on the offset
	start := offset / 8
	index := offset % 8
	offset += 4
	switch {
	case index == 4:
		// the stars have aligned!
		return int(data[start] & maskUint4), offset
	case index < 4:
		// all bits are in one byte
		// so we can read as a uint8
		// once we have shifted and masked
		v := (data[start] >> (4 - index))
		v &= maskUint4
		return int(v), offset
	default:
		// this is the case where the bits are in 2 bytes
		// the easiest way is to use a uint16
		v := (uint16(data[start]) << 8)
		v |= uint16(data[start+1])
		// now shift it to the right and mask
		v = (v >> (12 - index))
		v &= maskUint4
		return int(v), offset
	}
}

func readUint11(data []byte, offset int) (int, int) {
	// this could be over 3 bytes.
	// so the same as a uint3 we switch on the offset
	start := offset / 8
	index := offset % 8
	offset += 11
	switch {
	case index == 5:
		// aligned on 2 bytes
		v := uint16(data[start]) << 8
		v |= uint16(data[start+1])
		v &= maskUint11
		return int(v), offset
	case index < 5:
		// still only 2 bytes
		v := uint16(data[start]) << 8
		v |= uint16(data[start+1])
		// but we shift it to align
		v = (v >> (5 - index))
		v &= maskUint11
		return int(v), offset
	default:
		// this is over 3 bytes.
		// so we need to read into a uint32 first
		v := uint32(data[start]) << 16
		v |= uint32(data[start+1]) << 8
		v |= uint32(data[start+2])
		// now shift/mask it
		v = (v >> (13 - index))
		v &= maskUint11
		return int(v), offset
	}
}

func readUint15(data []byte, offset int) (int, int) {
	// this also could be 2 or 3 bytes
	start := offset / 8
	index := offset % 8
	offset += 15
	switch {
	case index == 1:
		// aligned!
		v := uint16(data[start]) << 8
		v |= uint16(data[start+1])
		v &= maskUint15
		return int(v), offset
	case index < 1:
		// index == 0
		v := uint16(data[start]) << 8
		v |= uint16(data[start+1])
		v = v >> 1
		v &= maskUint15
		return int(v), offset
	default:
		// we have 3 bytes like the uint11 case
		// so we need to read into a uint32 first
		v := uint32(data[start]) << 16
		v |= uint32(data[start+1]) << 8
		v |= uint32(data[start+2])
		// now shift/mask it
		v = (v >> (9 - index))
		v &= maskUint15
		return int(v), offset
	}
}

// now the tricky one
func readVarint(data []byte, offset int) (int, int) {
	sections := []int{}
	var more bool
	var v int
	for {
		more, offset = readBit(data, offset)

		// and the uint4
		v, offset = readUint4(data, offset)
		//fmt.Println("readVarint: uint4:", v, "hasMore?", more)
		sections = append(sections, v)
		if !more {
			break
		}
	}
	// reset to zero
	v = 0
	// now each section is shifted and or'd onto the number
	l := len(sections)
	for i, s := range sections {
		// l - (i+1) is the number of 4bit shifts we should do
		v |= s << (4 * (l - (i + 1)))
	}

	return v, offset
}

// Implement Solution to Problem 1
func solve1(input string) string {
	data, _ := hex.DecodeString(input)
	packet, _ := readPacket(data, 0)

	return fmt.Sprint(packet.SumVersions())
}

// Implement Solution to Problem 2
func solve2(input string) string {
	data, _ := hex.DecodeString(input)
	packet, _ := readPacket(data, 0)

	return fmt.Sprint(packet.Process())
}
