package main

import (
	"testing"
)

// tests for the AdventOfCode 2025 day 12 solutions

type Case[T any] struct {
	In  string
	Out T
}

const example = `0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2`

var problem1cases = []Case[int]{
	// cases here
	{example, 2},
}

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%v', Actual: '%v'", c.Out, actual)
		}
	}
}

var problem2cases = []Case[int]{
	// cases here
	{example, 0},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%v', Actual: '%v'", c.Out, actual)
		}
	}
}
