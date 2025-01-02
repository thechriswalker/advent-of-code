package main

import (
	"testing"
)

// tests for the AdventOfCode 2024 day 12 solutions

type Case struct {
	In  string
	Out string
}

const example = `AAAA
BBCD
BBCC
EEEC`

const ex1 = `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`

const ex2 = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

var problem1cases = []Case{
	// cases here
	{example, "140"},
	{ex1, "772"},
	{ex2, "1930"},
}

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

const (
	ex3 = `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`
	ex4 = `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`
)

var problem2cases = []Case{
	// cases here
	{example, "80"},
	{ex1, "436"},
	{ex3, "236"},
	{ex4, "368"},
	{ex2, "1206"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
