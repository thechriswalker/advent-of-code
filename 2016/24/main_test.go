package main

import (
	"testing"
)

// tests for the AdventOfCode 2016 day 24 solutions

type Case struct {
	In  string
	Out string
}

const mini = `###########
#0.1.....2#
#.#######.#
#4.......3#
###########`

var problem1cases = []Case{
	// cases here
	{mini, "14"},
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
	//{"", ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
