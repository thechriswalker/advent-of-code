package main

import (
	"testing"
)

// tests for the AdventOfCode 2016 day 8 solutions

type Case struct {
	In  string
	Out string
}

const testInput = `rect 3x2
rotate column x=1 by 1
rotate row y=0 by 4
rotate column x=1 by 1`

const testOutput = `
.#..#.#
#.#....
.#.....`

var problem1cases = []Case{
	// cases here
	{testInput, "6"},
}

func TestProblem1(t *testing.T) {
	Width = 7
	Height = 3
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	{testInput, testOutput},
}

func TestProblem2(t *testing.T) {
	Width = 7
	Height = 3
	Space = '.'
	for _, c := range problem2cases {
		actual := solve2raw(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
