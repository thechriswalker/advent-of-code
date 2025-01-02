package main

import (
	"testing"
)

// tests for the AdventOfCode 2024 day 1 solutions

type Case struct {
	In  string
	Out string
}

const example = `3   4
4   3
2   5
1   3
3   9
3   3
`

var problem1cases = []Case{
	// cases here
	{example, "11"},
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
	{example, "31"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
