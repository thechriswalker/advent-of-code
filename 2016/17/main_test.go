package main

import (
	"testing"
)

// tests for the AdventOfCode 2016 day 17 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"hijkl", "<fail>"},
	{"ihgpwlah", "DDRRRD"},
	{"kglvqrro", "DDUDRLRRUDRD"},
	{"ulqzkmiv", "DRURDRUDDLLDLUURRDULRLDUUDDDRR"},
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
	{"ihgpwlah", "370"},
	{"kglvqrro", "492"},
	{"ulqzkmiv", "830"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
