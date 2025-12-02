package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 2, solve1, solve2)
}

type IDRange struct {
	first, last string
	f, l        uint64
}

func parseInput(input string) []IDRange {
	ranges := []IDRange{}
	for _, r := range strings.Split(input, ",") {
		first, last, _ := strings.Cut(r, "-")
		f, _ := strconv.ParseUint(strings.TrimSpace(first), 10, 64)
		l, _ := strconv.ParseUint(strings.TrimSpace(last), 10, 64)
		ranges = append(ranges, IDRange{first: first, last: last, f: f, l: l})
	}
	return ranges
}

// Implement Solution to Problem 1
func solve1(input string) string {
	defer aoc.DisableDebug()()
	ranges := parseInput(input)

	sum := uint64(0)
	for _, r := range ranges {
		sum += findSumInvalidInRange(r)
	}
	// 18282222375 is too low
	// 21715656709 is too low
	return fmt.Sprint(sum)
}

func findSumInvalidInRange(r IDRange) uint64 {
	sum := uint64(0)
	// fuck it, I'm sure ill regret it in the second half, but let's take the naive view.
	// the range we are looking at is all n digit numbers

	// // we need to find all numbers in the range that are 2 repeated strings.
	lf := len(r.first)
	ll := len(r.last)
	// if either length is odd, we cant have any in that range so we bump up/down the length
	if lf%2 == 1 {
		lf++
	}
	if ll%2 == 1 {
		ll--
	}
	aoc.Debug("----- Range", r.first, "-", r.last, "(", r.f, "-", r.l, ")")
	for ll <= lf {
		// while the length of the far end is less or equal we can find numbers in the range.
		// find all numbers of the given length (it will be even)
		// basically how many numbers larger than lf
		// how many number are there in an n-digit number that have no leading zeros?
		// n = 1 -> 1-9 -> 9
		// n = 2 -> 10-99 -> 90
		// n = 3 -> 100-999 -> 900
		// so 10**n - 10 ** (n-1)
		// but we are looking at the "half range."
		n := lf / 2
		lo, factor := uint64(math.Pow10(n-1)), uint64(math.Pow10(n))
		hi := factor - 1
		aoc.Debug(">>> lf", lf, "ll", ll, "n", n, "lo", lo, "hi", hi, "factor", factor)
		for i := lo; i <= hi; i++ {
			x := i + i*factor
			if x >= r.f && x <= r.l {
				sum += x
				aoc.Debug(" + ", x)
			}
			if x > r.l {
				return sum
			}
		}
		lf += 2
	}
	return sum
}

// Implement Solution to Problem 2
func solve2(input string) string {
	ranges := parseInput(input)

	sum := uint64(0)
	for _, r := range ranges {
		sum += findSumInvalidInRange2(r)
	}
	// 32976912643 is too low
	return fmt.Sprint(sum)
}

func findSumInvalidInRange2(r IDRange) uint64 {
	sum := uint64(0)
	aoc.Debug("----- Range", r.first, "-", r.last, "(", r.f, "-", r.l, ")")

	// so now numbers can be any number of repetitions.
	// this means we have to start a 1 digit numbers, and move on to the max - len(l.last)/2

	// this means numbers can repeat...
	// i.e. 2 repeated 4 fours = 2222 = 22 repeated twice
	// so we use a cache.
	cache := map[uint64]struct{}{}
	minDigits := 1
	maxDigits := len(r.last) / 2
	for n := minDigits; n <= maxDigits; n++ {
		lo, factor := uint64(math.Pow10(n-1)), uint64(math.Pow10(n))
		hi := factor - 1
		aoc.Debug("n", n, "lo", lo, "hi", hi, "factor", factor)
		for i := lo; i <= hi; i++ {
			x := i + i*factor // at last one rep
			//aoc.Debug("i", i, "x", x)
			// while we are lower that the start keep repeating.
			for x < r.f {
				x = i + x*factor // this is a repition of the number
				//aoc.Debug("bump i", i, "x", x)
			}
			// while we are lower than the end sum and keep repeating
			for x <= r.l {
				aoc.Debug("found: ", x)
				if _, ok := cache[x]; !ok {
					sum += x
					cache[x] = struct{}{}
				}
				x = i + x*factor
				//aoc.Debug("bump i", i, "x", x)

			}
		}
	}
	return sum
}
