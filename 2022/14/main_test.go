package main

import (
	"testing"
)

// tests for the AdventOfCode 2022 day 14 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`, "24"},
}

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	{`498,4 -> 498,6 -> 496,6
503,4 -> 502,4 -> 502,9 -> 494,9`, "93"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
