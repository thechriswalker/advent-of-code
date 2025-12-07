package main

import (
	"testing"
)

// tests for the AdventOfCode 2025 day 6 solutions

type Case struct {
	In  string
	Out string
}

const example = `123 328  51 64 
 45 64  387 23 
  6 98  215 314
*   +   *   +  `

var problem1cases = []Case{
	// cases here
	{example, "4277556"},
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
	{example, "3263827"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
