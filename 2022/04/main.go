package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 4, solve1, solve2)
}

type Assignment struct {
	Start, Finish int
}

func (a Assignment) FullyContains(b Assignment) bool {
	// b contained in a?
	if a.Start <= b.Start && a.Finish >= b.Finish {
		return true
	}
	return false
}

func (a Assignment) Overlaps(b Assignment) bool {
	if a.Finish >= b.Start && a.Start <= b.Finish {
		return true
	}
	return false
}

// Implement Solution to Problem 1
func solve1(input string) string {
	n := 0
	aoc.MapLines(input, func(line string) error {
		a, b := Assignment{}, Assignment{}
		_, _ = fmt.Sscanf(line, "%d-%d,%d-%d", &(a.Start), &(a.Finish), &(b.Start), &(b.Finish))
		if a.FullyContains(b) || b.FullyContains(a) {
			n++
		}
		return nil
	})
	return fmt.Sprint(n)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	n := 0
	aoc.MapLines(input, func(line string) error {
		a, b := Assignment{}, Assignment{}
		_, _ = fmt.Sscanf(line, "%d-%d,%d-%d", &(a.Start), &(a.Finish), &(b.Start), &(b.Finish))
		if a.Overlaps(b) || b.Overlaps(a) {
			n++
		}
		return nil
	})
	return fmt.Sprint(n)
}
