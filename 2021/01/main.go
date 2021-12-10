package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 1, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	readings := aoc.ToIntSlice(input, '\n')
	count := 0
	for i := 1; i < len(readings); i++ {
		if readings[i] > readings[i-1] {
			count++
		}
	}
	return fmt.Sprintf("%d", count)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// this will be easier to use memory...
	readings := aoc.ToIntSlice(input, '\n')

	// now create the windows.
	// we don't actually have to do this.
	// the windows always overlap by 2 readings
	// so the only chance of a difference is the
	// beginning and the end.
	count := 0
	// start at the "fourth" reading (the last of the second window)
	for i := 3; i < len(readings); i++ {
		if readings[i] > readings[i-3] {
			count++
		}
	}
	return fmt.Sprintf("%d", count)
}
