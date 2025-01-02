package main

import (
	"testing"
)

// tests for the AdventOfCode 2023 day 20 solutions

type Case struct {
	In  string
	Out string
}

const ex1 = `broadcaster -> a, b, c
%a -> b
%b -> c
%c -> inv
&inv -> a`

const ex2 = `broadcaster -> a
%a -> inv, con
&inv -> b
%b -> con
&con -> output
%output -> `

var problem1cases = []Case{
	// cases here
	{ex1, "32000000"},
	{ex2, "11687500"},
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
	// {ex1, ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
