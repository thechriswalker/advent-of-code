package main

import (
	"testing"
)

// tests for the AdventOfCode 2025 day 1 solutions

type Case struct {
	In  string
	Out string
}

const example = `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
`

var problem1cases = []Case{
	// cases here
	{example, "3"},
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
	{example, "6"},
	{"L300", "3"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
