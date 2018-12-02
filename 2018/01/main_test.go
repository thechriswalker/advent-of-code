package main

import (
	"testing"
)

// tests for the AdventOfCode 2018 day 1 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"+1\n-2\n+3\n+1", "3"},
	{"+1\n+1\n+1", "3"},
	{"+1\n+1\n-2", "0"},
	{"-1\n-2\n-3", "-6"},
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
	{"+1\n-2\n+3\n+1", "2"},
	{"+1\n-1", "0"},
	{"+3\n+3\n+4\n-2\n-4", "10"},
	{"-6\n+3\n+8\n+5\n-6", "5"},
	{"+7\n+7\n-2\n-7\n-4", "14"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
