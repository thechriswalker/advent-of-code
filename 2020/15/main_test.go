package main

import (
	"testing"
)

// tests for the AdventOfCode 2020 day 15 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`0,3,6`, `436`},
	{`1,3,2`, `1`},
	{`2,1,3`, `10`},
	{`1,2,3`, `27`},
	{`2,3,1`, `78`},
	{`3,2,1`, `438`},
	{`3,1,2`, `1836`},
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
	{`0,3,6`, `175594`},
	{`1,3,2`, `2578`},
	{`2,1,3`, `3544142`},
	{`1,2,3`, `261214`},
	{`2,3,1`, `6895259`},
	{`3,2,1`, `18`},
	{`3,1,2`, `362`},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
