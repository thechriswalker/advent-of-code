package main

import (
	"testing"
)

// tests for the AdventOfCode 2019 day 1 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`12`, "2"},
	{`14`, "2"},
	{`1969`, "654"},
	{`100756`, "33583"},
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
	{`12`, "2"},
	{`14`, "2"},
	{`1969`, "966"},
	{`100756`, "50346"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
