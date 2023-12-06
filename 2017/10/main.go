package main

import (
	"encoding/hex"
	"fmt"
	"slices"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 10, solve1, solve2)
}

type Ring struct {
	size int
	list []uint8
	head int
	skip int
}

func NewRing(size int) *Ring {
	r := &Ring{
		size: size,
		list: make([]uint8, size),
		head: 0,
		skip: 0,
	}
	// fill in the initial positions.
	for i := range r.list {
		r.list[i] = uint8(i)
	}
	return r
}

func (r *Ring) PinchAndTwist(l int) {
	// take a subslice of the ring, length l
	// and reverse it.
	// if the whole length doesn't wrap, then we can do this
	// in-place.
	if l < r.size-r.head {
		ss := r.list[r.head : r.head+l]
		slices.Reverse(ss)
	} else {
		// if not, we have to pluck it out and put it back.
		ss := make([]uint8, l)
		copy(ss, r.list[r.head:])
		copy(ss[r.size-r.head:], r.list)
		// now revers that
		slices.Reverse(ss)
		// and set it back into the ring
		copy(r.list[r.head:], ss)
		copy(r.list, ss[r.size-r.head:])
	}
	// now move the current position l+skip
	r.head = (r.head + l + r.skip) % r.size
	r.skip++
}

var ringSize = 256

// Implement Solution to Problem 1
func solve1(input string) string {
	lengths := aoc.ToIntSlice(input, ',')

	r := NewRing(ringSize)

	for _, l := range lengths {
		r.PinchAndTwist(l)
	}

	return fmt.Sprint(r.list[0] * r.list[1])
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// now input is bytes
	// we should trim the string though

	return KnotHash(strings.TrimSpace(input))
}

func KnotHash(input string) string {
	lengths := []byte(strings.TrimSpace(input))
	lengths = append(lengths, 17, 31, 73, 47, 23)
	// noq 64 rounds of twists
	r := NewRing(256)
	for x := 0; x < 64; x++ {
		for _, l := range lengths {
			r.PinchAndTwist(int(l))
		}
	}
	// now compact into the "dense" version
	dense := make([]uint8, 16)
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			dense[i] ^= r.list[(i*16)+j]
		}
	}
	return hex.EncodeToString(dense)
}
