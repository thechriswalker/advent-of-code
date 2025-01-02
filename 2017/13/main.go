package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 13, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	scanners := map[int]int{}

	aoc.MapLines(input, func(line string) error {
		d := aoc.ToIntSlice(line, ':')
		scanners[d[0]] = d[1]
		return nil
	})

	s, _ := severityFor(scanners, 0)
	return fmt.Sprint(s)
}

func severityFor(scanners map[int]int, offset int) (int, bool) {
	sev := 0
	caught := false
	for i, d := range scanners {
		// you enter this layer at time i+offset
		t := i + offset
		// if there is a scanner at this layer,
		if d > 0 {
			period := d*2 - 2
			// then it will be at the top if the time is a multiple of the period -1 (because it is zero based)
			if t%period == 0 { // we are caught
				sev += i * d
				caught = true
			}
		}
	}
	return sev, caught
}

func failEarly(scanners map[int]int, offset int) (caught bool) {
	for i, d := range scanners {
		// you enter this layer at time i+offset
		t := i + offset
		// if there is a scanner at this layer,
		if t%(d*2-2) == 0 {
			return true
		}
	}
	return false
}

// Implement Solution to Problem 2
func solve2(input string) string {
	scanners := map[int]int{}
	aoc.MapLines(input, func(line string) error {
		d := aoc.ToIntSlice(line, ':')
		scanners[d[0]] = d[1]
		return nil
	})

	// to be fair there is probably a smarter way to do this using GCD or something.
	// this takes 2.7 seconds to get to 3861798.

	// simply switching to a "fail early" approach makes it run in ~500ms, better, but not great
	delay := 0
	for {
		if !failEarly(scanners, delay) {
			return fmt.Sprint(delay)
		}
		delay++
	}
}
