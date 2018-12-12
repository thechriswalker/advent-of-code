package main

import (
	"testing"
)

// tests for the AdventOfCode 2017 day 1 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"1122", "3"},
	{"1111", "4"},
	{"1234", "0"},
	{"91212129", "9"},
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
	{"1212", "6"},
	{"1221", "0"},
	{"123425", "4"},
	{"123123", "12"},
	{"12131415", "4"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
