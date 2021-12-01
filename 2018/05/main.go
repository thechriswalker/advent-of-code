package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2018, 5, solve1, solve2)
}

func solve1(s string) string {
	return solve1v2(s)
}

func solve2(s string) string {
	return solve2v2(s)
}

// Implement Solution to Problem 1
func solve1v1(input string) string {
	p := createPolymer(input)
	p.ReactAll()
	return fmt.Sprintf("%d", p.Length())
}

// Implement Solution to Problem 2
func solve2v1(input string) string {
	p := createPolymer(input)
	//s := p.String()
	p.ReactAll()
	min := p.Length() // this is the baseline
	//fmt.Printf("Baseline: %s => %s (%d)\n", s, p, min)
	for i := 'a'; i <= 'z'; i++ {
		lower := i
		upper := rune(strings.ToUpper(string(i))[0])
		p.Reset()
		p.IgnoreNeg = int8(-1 - int(lower) + lowerA)
		p.IgnorePos = int8(1 + int(upper) - upperA)
		//	s = p.String()
		p.ReactAll()
		//	fmt.Printf("'%c'-reduced: %s => %s (%d)\n", i, s, p, p.Length())
		if p.Length() < min {
			min = p.Length()
		}
	}
	return fmt.Sprintf("%d", min)
}

type Polymer struct {
	// this contains 0 if nothing at that index, -n or +n for type n pairs.
	Units     []int8
	Reactions []int8
	IgnoreNeg int8
	IgnorePos int8
	length    int
}

func (p *Polymer) Length() int {
	if p.length > -1 {
		return p.length
	}
	sum := 0
	var val int8
	for i := 0; i < len(p.Reactions); i++ {
		if p.Reactions[i] == 0 {
			val = p.Units[i]
			if val != p.IgnoreNeg && val != p.IgnorePos {
				sum++
			}
		}
	}
	p.length = sum
	return sum
}

func (p *Polymer) Reset() {
	for i := 0; i < len(p.Reactions); i++ {
		p.Reactions[i] = 0
	}
	p.length = -1
}

func (p *Polymer) String() string {
	// convert back to string.
	s := strings.Builder{}
	for _, i := range p.Units {
		if i == p.IgnoreNeg || i == p.IgnorePos {
			continue
		}
		switch {
		case i > 0:
			// Upper
			s.WriteByte(byte(upperA + int(i) - 1))
		case i < 0:
			// Lower
			s.WriteByte(byte(lowerA - 1 + int(-1*i)))
		}
	}
	return s.String()
}

func (p *Polymer) ReactAll() {
	var index = 0
	for index >= 0 {
		index = p.React(index)
		if index > 0 {
			index-- // go back one because the previous char may now react
		}
	}
}

// find the next reaction.
func (p *Polymer) React(startIndex int) int {
	// indexes
	var idx1, idx2 int
	var val1, val2 int8
	for idx1 < len(p.Units) {
		idx1, val1 = p.First(idx1)
		if idx1 == -1 {
			return -1
		}
		idx2, val2 = p.Next(idx1)
		if idx2 == -1 {
			return -1
		}
		if val1+val2 == 0 {
			// boom, set them reacted
			p.Reactions[idx1], p.Reactions[idx2] = 1, 1
			p.length -= 2
			return idx2
		}
		idx1 = idx2
	}
	return -1
}

// find the first value on or after the index.
func (p *Polymer) First(start int) (idx int, val int8) {
	for idx = start; idx < len(p.Units); idx++ {
		if p.Reactions[idx] == 0 {
			val = p.Units[idx]
			if val == p.IgnoreNeg || val == p.IgnorePos {
				// skip
			} else {
				return
			}
		}
	}
	return -1, 0
}

// find the first value  after the index.
func (p *Polymer) Next(start int) (idx int, val int8) {
	return p.First(start + 1)
}

var lowerA = int('a')
var upperA = int('A')

func createPolymer(input string) *Polymer {
	p := &Polymer{
		Units:     make([]int8, 0, len(input)),
		Reactions: make([]int8, 0, len(input)),
		length:    -1,
	}
	fillPolymer(p, input)
	return p
}

// returns true if ignore was ever true
// meaning this polymer *was* reduced, so don't bother processing it
func fillPolymer(p *Polymer, input string) {
	for _, c := range input {
		switch {
		case c >= 'a' && c <= 'z':
			// 'a' should be -1
			// - 1 - c + 'a'
			p.Units = append(p.Units, int8(-1-int(c)+lowerA))
			p.Reactions = append(p.Reactions, 0)
		case c >= 'A' && c <= 'Z':
			// 'A' should be +1
			// c - 'A' + 1
			p.Units = append(p.Units, int8(1+int(c)-upperA))
			p.Reactions = append(p.Reactions, 0)
		default:
			// skip
		}
	}
}

// version 2 build a stack rather than shortening one

func solve1v2(input string) string {
	stack := make([]byte, 0, len(input))
	rd := strings.NewReader(input)
	var curr, next byte
	var err error
	var index int
	curr, err = readNextGoodByte(rd)
	for {
		index = len(stack)
		// first work backwards to shorten the stack if needs be
		if index > 0 && isReactive(stack[index-1], curr) {
			// annihilate both
			stack = stack[0 : index-1]
			curr, err = readNextGoodByte(rd)
			if err != nil {
				break
			}
		} else {
			next, err = readNextGoodByte(rd)
			if err != nil {
				// the end, but we need to add the current to the stack
				stack = append(stack, curr)
				break
			}
			// if current and next collide, annihilate
			// which means popping the stack
			if isReactive(curr, next) {
				if index > 0 {
					curr = stack[index-1]
					stack = stack[0 : index-1]
				} else {
					// nothing on the stack so load a new curr
					curr, err = readNextGoodByte(rd)
					if err != nil {
						break
					}
				}
			} else {
				// current and next not reactive, so current can be put on the stack (for now)
				stack = append(stack, curr)
				// and the new curr is next
				curr = next
			}
		}
	}
	// do we need to append
	return fmt.Sprintf("%d", len(stack))
}

func solve2v2(input string) string {
	stack := make([]byte, 0, len(input))
	min := 1000000
	for c := byte('a'); c <= byte('z'); c++ {
		stack = stack[0:0] // reset stack
		l := solveWithSkippedPair(stack, input, c, c-32)
		if l < min {
			min = l
		}
	}
	return fmt.Sprintf("%d", min)
}

func solveWithSkippedPair(stack []byte, input string, a, b byte) int {
	rd := strings.NewReader(input)
	var curr, next byte
	var err error
	var index int

	curr, err = readNextGoodByteExcept(rd, a, b)
	for {
		index = len(stack)
		// first work backwards to shorten the stack if needs be
		if index > 0 && isReactive(stack[index-1], curr) {
			// annihilate both
			stack = stack[0 : index-1]
			curr, err = readNextGoodByteExcept(rd, a, b)
			if err != nil {
				break
			}
		} else {
			next, err = readNextGoodByteExcept(rd, a, b)
			if err != nil {
				// the end, but we need to add the current to the stack
				stack = append(stack, curr)
				break
			}
			// if current and next collide, annihilate
			// which means popping the stack
			if isReactive(curr, next) {
				if index > 0 {
					curr = stack[index-1]
					stack = stack[0 : index-1]
				} else {
					// nothing on the stack so load a new curr
					curr, err = readNextGoodByteExcept(rd, a, b)
					if err != nil {
						break
					}
				}
			} else {
				// current and next not reactive, so current can be put on the stack (for now)
				stack = append(stack, curr)
				// and the new curr is next
				curr = next
			}
		}
	}
	// do we need to append
	return len(stack)
}
func readNextGoodByteExcept(rd io.ByteReader, a, b byte) (byte, error) {
	for {
		n, err := readNextGoodByte(rd)
		if err == nil && (n == a || n == b) {
			continue
		}
		return n, err
	}
}

func readNextGoodByte(r io.ByteReader) (b byte, err error) {
	for {
		b, err = r.ReadByte()
		if err != nil {
			return 0, err
		}
		if (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') {
			return b, nil
		}
	}
}

func isReactive(a, b byte) bool {
	switch {
	case a >= 'a' && a <= 'z':
		return b == a-32
	default:
		return b == a+32
	}
}
