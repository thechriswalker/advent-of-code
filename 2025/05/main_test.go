package main

import (
	"testing"
)

// tests for the AdventOfCode 2025 day 5 solutions

type Case struct {
	In  string
	Out string
}

const example = `3-5
10-14
16-20
12-18

1
5
8
11
17
32`

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
	{example, "14"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
