package main

import (
	"testing"
)

// tests for the AdventOfCode 2016 day 9 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"A(1x5)BC", "7"},
	{"ADVENT", "6"},
	{"(3x3)XYZ", "9"},
	{"A(2x2)BCD(2x2)EFG", "11"},
	{"(6x1)(1x3)A", "6"},
	{"X(8x2)(3x3)ABCY", "18"},
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
	{"(3x3)XYZ", "9"},
	{"X(8x2)(3x3)ABCY", "20"},
	{"(27x12)(20x12)(13x14)(7x10)(1x12)A", "241920"},
	{"(25x3)(3x3)ABC(2x3)XY(5x2)PQRSTX(18x9)(3x2)TWO(5x7)SEVEN", "445"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
