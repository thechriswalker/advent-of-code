package main

import (
	"testing"
)

// tests for the AdventOfCode 2023 day 11 solutions

type Case struct {
	In  string
	Out string
}

const example = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....
`

var problem1cases = []Case{
	// cases here
	{example, "374"},
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
	//	{example, ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

func TestExpansion(t *testing.T) {
	cases := []struct {
		Expansion int
		Expected  string
	}{
		{2, "374"},
		{10, "1030"},
		{100, "8410"},
	}

	for _, c := range cases {
		actual := solveD(example, c.Expansion)
		if c.Expected != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Expected, actual)
		}
	}

}
