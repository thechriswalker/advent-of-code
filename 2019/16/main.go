package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 16, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	curr := parseInput(input, 1)
	next := make([]uint8, len(curr))
	for i := 0; i < 100; i++ {
		phase(curr, next)
		next, curr = curr, next
	}
	// first 8 digits of curr.
	return fmt.Sprintf("%d%d%d%d%d%d%d%d", curr[0], curr[1], curr[2], curr[3], curr[4], curr[5], curr[6], curr[7])
}

// Implement Solution to Problem 2
func solve2(input string) string {
	offset, _ := strconv.Atoi(input[0:7])
	curr := parseInput(input, 10000)
	if offset < len(curr)/2 {
		return "<bad offset cannot solve>"
	}
	curr = curr[offset:]
	// I have a feeling the naive method is going to fail...
	// yep!

	for i := 0; i < 100; i++ {
		lessNaivePhase(curr)
	}
	return fmt.Sprintf("%d%d%d%d%d%d%d%d", curr[0], curr[1], curr[2], curr[3], curr[4], curr[5], curr[6], curr[7])
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

var basePattern = [4]int{0, 1, 0, -1}

func patternMultiplier(a, b int) int {
	// a represents the repetition constant.
	// that is the index in the array -1 of the digit being calculated
	// b is the offset which is the index of the digit -2 of the digit being multiplied
	// base pattern is 0,1,0,-1 (so 4 digits)
	// so the index based on a and b in that 4 pattern is: floor(b+1/a) %4

	// so for digit 3 in the input, multipled by index 2 in the input
	// we get (3+1) = 4 reps of each base digit

	// we get (2+1) = 4 the digit in the pattern list
	// is floor(3/4) = 0 in the base pattern = `0`
	//                             1 1  1  1  1  1
	// index:  0 1 2 3 4 5 6 7 8 9 0 1  2  3  4  5
	// base:   0 0 0 0 1 1 1 1 0 0 0 0 -1 -1 -1 -1
	// index:    0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6
	// f(i+1/4): 0 0 0 1 1 1 1 2 2 2 2 3 3 3 3 0
	baseIndex := ((b + 1) / (a + 1)) % 4
	return basePattern[baseIndex]
}

func phase(curr, next []uint8) {
	// do a phase and save the output in next.
	// hmmm, O(n^2)....
	for i := 0; i < len(curr); i++ {
		sum := 0
		for j := i; j < len(curr); j++ {
			sum += int(curr[j]) * patternMultiplier(i, j)
		}
		next[i] = uint8(abs(sum) % 10)
	}
}

func parseInput(input string, multipler int) []uint8 {
	s := strings.TrimSpace(input)
	l := len(s)
	ints := make([]uint8, l*multipler)
	for i, c := range s {
		n, _ := strconv.Atoi(string(c))
		un := uint8(n)
		for j := 0; j < multipler; j++ {
			ints[(l*j)+i] = un
		}
	}
	return ints
}

// the clue was in the name, this is an FFT problem, or fast-fourier-transform
// which transforms the n^2 complexity problem to n*log(n)
// but I don't know have to do that...

// OK, there is a pattern. leading zeros in the transformation matrix mean that
// each value is not dependant on any previous value, which implies:
// A: we can do the transformation IN PLACE we don't need 2 arrays.
// B: we don't need to care about anything before our offset.

// this only really means that we are still N^2 but our n is smaller...

// so lets look again. for each digit we start with that digit and sum (*1) the next
// N digits where N is the index of the digit we are working out. then we skip N
// and subtract the sum of the next N digits
// then skip N
// then sum N
// then skip N
// then substract sum of next N

// and for the big reveal, the final sequence (from the second half on, so we need
// offset > N/2) is actually like this (in reverse order)
//
// N:   0 0 0 0 0 0 0 0 0 0 0 0 0 x = x const
// N-1: 0 0 0 0 0 0 0 0 0 0 0 0 y x = (y+x) %10
// N-2: 0 0 0 0 0 0 0 0 0 0 0 z y x = (z + (y+x) %10) %10

// i.e the value of x(n) is only dependent on x(n+1)
// so we can work backwards with a linear algorithm. (if our offset is big enough)

func lessNaivePhase(inplace []uint8) {
	// assume our offset is > len(inplace) /2
	// start from the end (well -2, the final digit is constant)
	//  we assume we have just cut our input
	for j := len(inplace) - 2; j >= 0; j-- {
		inplace[j] = (inplace[j] + inplace[j+1]) % 10
	}
}
