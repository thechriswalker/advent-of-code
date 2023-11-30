package main

import (
	"testing"
)

// tests for the AdventOfCode 2015 day 25 solutions

type Case struct {
	In  string
	Out string
}

const example = ""

var problem1cases = []Case{
	// cases here
	//{example, ""},
}

type rcCase struct{ row, col, num int }

func TestRowCol(t *testing.T) {
	cases := []rcCase{
		{1, 1, 1},
		{1, 2, 3},
		{6, 1, 16},
		{1, 6, 21},
		{4, 3, 18},
	}
	for _, c := range cases {
		actual := rowXColY(c.row, c.col)
		if c.num != actual {
			t.Errorf("failed case row=%d, col=%d, expected=%d, got=%d", c.row, c.col, c.num, actual)
		}
	}
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
	{example, ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
