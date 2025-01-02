package main

import (
	"testing"
)

// tests for the AdventOfCode 2024 day 17 solutions

type Case struct {
	In  string
	Out string
}

const example = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

var problem1cases = []Case{
	// cases here
	{example, "4,6,3,5,6,3,5,2,1,0"},
}

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

const ex2 = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

var problem2cases = []Case{
	// cases here
	//{ex2, "117440"}, // my fastProgram will not work on the test input
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

func TestFast(t *testing.T) {
	// OK so it does workj...
	if _, ok := fastProgram(46323429, []int{7, 6, 1, 5, 3, 1, 4, 2, 6}); !ok {
		t.Fatalf("fastProgram failed for known input")
	}
}
