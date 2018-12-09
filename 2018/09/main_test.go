package main

import (
	"testing"
)

// tests for the AdventOfCode 2018 day 9 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"9 players; last marble is worth 25 points", "32"},
	{"10 players; last marble is worth 1618 points", "8317"},
	{"13 players; last marble is worth 7999 points", "146373"},
	{"17 players; last marble is worth 1104 points", "2764"},
	{"21 players; last marble is worth 6111 points", "54718"},
	{"30 players; last marble is worth 5807 points", "37305"},
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
	//	{"", ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
