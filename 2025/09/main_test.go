package main

import (
	"testing"
)

// tests for the AdventOfCode 2025 day 9 solutions

type Case struct {
	In  string
	Out string
}

const example = `7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3`

var problem1cases = []Case{
	// cases here
	{example, "50"},
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
	{example, "24"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
