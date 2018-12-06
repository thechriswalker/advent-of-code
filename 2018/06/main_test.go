package main

import (
	"testing"
)

// tests for the AdventOfCode 2018 day 6 solutions

type Case struct {
	In  string
	Out string
}

const testInput = `1, 1
1, 6
8, 3
3, 4
5, 5
8, 9
`

var problem1cases = []Case{
	// cases here
	{testInput, "17"},
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
	{testInput, "16"},
}

func TestProblem2(t *testing.T) {
	LIMIT = 32
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
