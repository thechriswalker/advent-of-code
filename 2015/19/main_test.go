package main

import (
	"testing"
)

// tests for the AdventOfCode 2015 day 19 solutions

type Case struct {
	In  string
	Out string
}

const ex = `H => HO
H => OH
O => HH

HOHOHO
`

var problem1cases = []Case{
	// cases here
	{ex, "7"},
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
	// 	{`e => H
	// e => O
	// H => HO
	// H => OH
	// O => HH

	// HOHOHO
	// `, "6"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
