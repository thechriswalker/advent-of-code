package main

import (
	"testing"
)

// tests for the AdventOfCode 2018 day 19 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	// 	{`#ip 0
	// seti 5 0 1
	// seti 6 0 2
	// addi 0 1 0
	// addr 1 2 3
	// setr 1 0 0
	// seti 8 0 4
	// seti 9 0 5
	// `, "6"},
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
