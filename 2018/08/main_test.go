package main

import (
	"testing"
)

// tests for the AdventOfCode 2018 day 8 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", "138"},
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
	{"2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2", "66"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
