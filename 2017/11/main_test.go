package main

import (
	"testing"
)

// tests for the AdventOfCode 2017 day 11 solutions

type Case struct {
	In  string
	Out string
}

const example = "ne,ne,ne"

var problem1cases = []Case{
	// cases here
	{example, "3"},
	{"ne,ne,sw,sw", "0"},
	{"ne,ne,s,s", "2"},
	{"se,sw,se,sw,sw", "3"},
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
