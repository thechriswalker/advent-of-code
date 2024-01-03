package main

import (
	"testing"
)

// tests for the AdventOfCode 2023 day 17 solutions

type Case struct {
	In  string
	Out string
}

const example = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

var problem1cases = []Case{
	// cases here
	{example, "102"},
}

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

const ex2 = `111111111111
999999999991
999999999991
999999999991
999999999991`

var problem2cases = []Case{
	// cases here
	{example, "94"},
	{ex2, "71"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
