package main

import (
	"testing"
)

// tests for the AdventOfCode 2021 day 9 solutions

type Case struct {
	In  string
	Out string
}

var ex = `2199943210
3987894921
9856789892
8767896789
9899965678
`

var problem1cases = []Case{
	// cases here
	{ex, "15"},
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
	{ex, "1134"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
