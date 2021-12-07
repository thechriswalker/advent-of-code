package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 7, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	crabs := aoc.ToIntSlice(input, ',')

	// lets sort them so the lowest is first
	sort.Ints(crabs)

	// we have to find the Minimum movement
	// basically the sum of the differences
	// the current position to desired.
	min := math.MaxInt64
	for x := crabs[0]; x < crabs[len(crabs)-1]; x++ {
		sum := 0
		for _, n := range crabs {
			d := x - n
			if d < 0 {
				sum -= d
			} else {
				sum += d
			}
		}
		if sum < min {
			min = sum
		}
	}

	return fmt.Sprintf("%d", min)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	crabs := aoc.ToIntSlice(input, ',')

	// lets sort them so the lowest is first
	sort.Ints(crabs)

	// we have to find the Minimum movement
	// basically the sum of the differences
	// the current position to desired.
	min := math.MaxInt64
	for x := crabs[0]; x < crabs[len(crabs)-1]; x++ {
		sum := 0
		for _, n := range crabs {
			d := x - n
			if d < 0 {
				d *= -1
			}
			// sum of integers is n*(n+1) / 2
			sum += d * (d + 1) / 2
		}
		if sum < min {
			min = sum
		}
	}

	return fmt.Sprintf("%d", min)
}
