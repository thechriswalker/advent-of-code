package main

import (
	"testing"
)

// tests for the AdventOfCode 2018 day 14 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"5", "0124515891"},
	{"9", "5158916779"},
	{"18", "9251071085"},
	{"2018", "5941429882"},
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
	{"51589", "9"},
	{"01245", "5"},
	{"92510", "18"},
	{"59414", "2018"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
