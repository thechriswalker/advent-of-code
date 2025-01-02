package main

import (
	"testing"
)

// tests for the AdventOfCode 2017 day 16 solutions

type Case struct {
	In  string
	Out string
}

const example = `s1,x3/4,pe/b`

var problem1cases = []Case{
	// cases here
	{example, "baedc"},
}

func TestProblem1(t *testing.T) {
	LineSize = 5
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	//{example, ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
