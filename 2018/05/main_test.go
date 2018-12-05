package main

import (
	"testing"
)

// tests for the AdventOfCode 2018 day 5 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"aA", "0"},
	{"abBA", "0"},
	{"abAB", "4"},
	{"aabAAB", "6"},
	{"dabAcCaCBAcCcaDA", "10"},
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
	{"dabAcCaCBAcCcaDA", "4"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
