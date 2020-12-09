package main

import (
	"testing"
)

// tests for the AdventOfCode 2020 day 9 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`, "127"},
}

func TestProblem1(t *testing.T) {
	preambleLength = 5
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	{`35
20
15
25
47
40
62
55
65
95
102
117
150
182
127
219
299
277
309
576
`, "62"},
}

func TestProblem2(t *testing.T) {
	preambleLength = 5

	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
