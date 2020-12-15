package main

import (
	"testing"
)

// tests for the AdventOfCode 2020 day 13 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`939
7,13,x,x,59,x,31,19
`, "295"},
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
	{`1
17,x,13,19`, `3417`},
	{`1
67,7,59,61`, `754018`},
	{`1
67,x,7,59,61`, `779210`},
	{`1
67,7,x,59,61`, `1261476`},
	{`1
1789,37,47,1889`, `1202161486`},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
