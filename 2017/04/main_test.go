package main

import (
	"testing"
)

// tests for the AdventOfCode 2017 day 4 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"aa bb cc dd ee", "1"},
	{"aa bb cc dd aa", "0"},
	{"aa bb cc dd aaa", "1"},
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
	{"abcde fghij", "1"},
	{"abcde xyz ecdab", "0"},
	{"a ab abc abd abf abj", "1"},
	{"iiii oiii ooii oooi oooo", "1"},
	{"oiii ioii iioi iiio", "0"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
