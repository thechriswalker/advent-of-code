package main

import (
	"testing"
)

// tests for the AdventOfCode 2019 day 2 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"1,0,0,0,99", "2"},
	{"2,3,0,3,99", "2"},
	{"2,4,4,5,99,0", "2"},
	{"1,1,1,4,99,5,6,0,99", "30"},
	{"1,9,10,3,2,3,11,0,99,30,40,50", "3500"},
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
	//{"", ""}, no tests
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
