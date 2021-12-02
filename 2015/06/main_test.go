package main

import (
	"testing"
)

// tests for the AdventOfCode 2015 day 6 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`turn on 0,0 through 999,999
toggle 0,0 through 999,0
turn off 499,499 through 500,500
`, "998996"},
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
	{"turn on 0,0 through 0,0", "1"},
	{"toggle 0,0 through 999,999", "2000000"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
