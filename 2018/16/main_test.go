package main

import (
	"testing"
)

// tests for the AdventOfCode 2018 day 16 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]`, "1"},
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
	//	{"", ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
