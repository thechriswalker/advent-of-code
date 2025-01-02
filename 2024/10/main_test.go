package main

import (
	"testing"
)

// tests for the AdventOfCode 2024 day 10 solutions

type Case struct {
	In  string
	Out string
}

const example = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

var problem1cases = []Case{
	// cases here
	{example, "36"},
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
	{example, "81"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
