package main

import (
	"testing"
)

// tests for the AdventOfCode 2022 day 9 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`, "13"},
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
	{`R 4
U 4
L 3
D 1
R 4
D 1
L 5
R 2`, "1"},
	{`R 5
U 8
L 8
D 3
R 17
D 10
L 25
U 20`, "36"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
