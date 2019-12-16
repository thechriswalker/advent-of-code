package main

import (
	"testing"
)

// tests for the AdventOfCode 2019 day 16 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"80871224585914546619083218645595", "24176176"},
	{"19617804207202209144916044189917", "73745418"},
	{"69317163492948606335995924319873", "52432133"},
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
	{"03036732577212944063491565474664", "84462026"},
	{"02935109699940807407585447034323", "78725270"},
	{"03081770884921959731165446850517", "53553731"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
