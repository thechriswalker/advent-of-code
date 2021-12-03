package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 10, solve1, solve2)
}

func step(prev string) string {
	b := strings.Builder{}
	runOf := 'x'
	runStart := 0
	for i, c := range prev {
		//fmt.Println("i:", i, "c:", string(c), "runOf:", string(runOf), "runStart:", runStart)
		if c == runOf {
			continue
		}
		// run came to an end.
		// write the current run into the builder
		l := i - runStart
		if l > 0 {
			fmt.Fprintf(&b, "%d%c", l, runOf)
		}
		runOf = c
		runStart = i
	}
	// and write the final log.
	l := len(prev) - runStart
	if l > 0 {
		fmt.Fprintf(&b, "%d%c", l, runOf)
	}
	return b.String()
}

func repeat(input string, steps int) string {
	s := strings.TrimSpace(input)
	for i := 0; i < steps; i++ {
		//fmt.Printf("%03d: %s\n", i, s)
		s = step(s)
	}
	return fmt.Sprintf("%d", len(s))
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return repeat(input, 40)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// should this have been harder?
	return repeat(input, 50)
}
