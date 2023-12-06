package main

import (
	"testing"
)

// tests for the AdventOfCode 2023 day 6 solutions

type Case struct {
	In  string
	Out string
}

const example = `Time:      7  15   30
Distance:  9  40  200`

var problem1cases = []Case{
	// cases here
	{example, "288"},
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
	{example, "71503"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
