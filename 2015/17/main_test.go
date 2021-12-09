package main

import (
	"testing"
)

// tests for the AdventOfCode 2015 day 17 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"20\n15\n10\n5\n5\n", "4"},
}

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := solveSane(c.In, 25)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	{"20\n15\n10\n5\n5\n", "3"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solveSane2(c.In, 25)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
