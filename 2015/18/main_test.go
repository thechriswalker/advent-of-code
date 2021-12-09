package main

import (
	"testing"
)

// tests for the AdventOfCode 2015 day 18 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`.#.#.#
...##.
#....#
..#...
#.#..#
####..
`, "4"},
}

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := solve1a(c.In, 4)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	{`.#.#.#
...##.
#....#
..#...
#.#..#
####..
`, "17"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2a(c.In, 5)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
