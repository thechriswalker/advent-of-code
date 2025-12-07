package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 5, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	ranges := [][2]int{}
	ingredients := []int{}
	aoc.MapLines(input, func(line string) error {
		lo, hi, found := strings.Cut(line, "-")
		if found {
			//.a range
			l, _ := strconv.Atoi(lo)
			h, _ := strconv.Atoi(hi)
			ranges = append(ranges, [2]int{l, h})
		} else {
			//ingredient
			i, _ := strconv.Atoi(line)
			ingredients = append(ingredients, i)
		}
		return nil
	})

	count := 0

	for _, i := range ingredients {
		for _, r := range ranges {
			if i >= r[0] && i <= r[1] {
				count++
				break
			}
		}
	}

	return fmt.Sprint(count)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	ranges := []Range{}
	aoc.MapLines(input, func(line string) error {
		lo, hi, found := strings.Cut(line, "-")
		if found {
			//.a range
			l, _ := strconv.ParseUint(lo, 10, 64)
			h, _ := strconv.ParseUint(hi, 10, 64)
			ranges = append(ranges, Range{lo: l, hi: h})
		}
		return nil
	})

	// we need to flatten the ranges.
	// best if we sort by size.
	slices.SortFunc(ranges, func(a, b Range) int {
		if a.Size() == b.Size() {
			return 0
		}
		if a.Size() < b.Size() {
			return -1
		}
		return 1
	})
	ranges = flatten(ranges)

	sum := uint64(0)
	for _, r := range ranges {
		sum += r.Size()
	}
	// 350405915068134 too high
	// 350405915068134
	// 344813017450467
	return fmt.Sprint(sum)
}

type Range struct {
	lo, hi uint64
}

func (r Range) Size() uint64 {
	// it's inclusive!
	return 1 + r.hi - r.lo
}

// adds this range to the current or returns false.
func (r Range) Combine(other Range) (Range, bool) {
	if r.hi < other.lo || r.lo > other.hi {
		// no overlap.
		return Range{}, false
	}
	// other take the outer limits.
	if r.lo > other.lo {
		r.lo = other.lo
	}
	if r.hi < other.hi {
		r.hi = other.hi
	}
	return r, true
}

func flatten(ranges []Range) []Range {
	r := ranges
	var h Range
	out := []Range{}
	for {
		h, r = flathead(r)
		aoc.Debug("\nout", out, "\nh", h, "\nrng", r)
		out = append(out, h)
		if len(r) == 0 {
			break
		}
	}
	return out
}

func flathead(r []Range) (Range, []Range) {
	head := r[0]
	tail := r[1:]
	next := []Range{}
	for {
		for i := range tail {
			c, ok := head.Combine(tail[i])
			if ok {
				head = c
			} else {
				next = append(next, tail[i])
			}
		}
		if len(next) == len(tail) {
			break
		}
		tail = next
		next = []Range{}
	}

	return head, next
}
