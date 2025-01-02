package main

import (
	"testing"
)

// tests for the AdventOfCode 2024 day 18 solutions

type Case struct {
	In  string
	Out string
}

const example = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

var problem1cases = []Case{
	// cases here
	{example, "22"},
}

func TestProblem1(t *testing.T) {
	GridSize = 7 // 0-6
	Part1Bytes = 12
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	{example, "6,1"},
}

func TestProblem2(t *testing.T) {
	GridSize = 7 // 0-6
	Part1Bytes = 12

	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
