package main

import (
	"testing"
)

// tests for the AdventOfCode 2023 day 14 solutions

type Case struct {
	In  string
	Out string
}

const example = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....
`

var problem1cases = []Case{
	// cases here
	{example, "136"},
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
	{example, "64"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

// func TestLongestRepeatingPatter(t *testing.T) {
// 	s := aoc.ToIntSlice("65 64 65 63 68 69 69 65 64 65 63 68 69 69 65 64 65 63 68 69 69 65 64 65 63 68 69 69 65 64 65 63 68 69 69 65 64 65 63 68 69 69 65 64 65 63 68 69 69 65 64 65 63 68 69 69 65 64 65 63 68 69 69 65 64 65 63 68 69 69 65 64", ' ')

// 	o, l := findLongestRepeatingSubPattern(s)

// 	t.Errorf("o=%d, l=%d", o, l)

// }
