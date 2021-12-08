package main

import (
	"testing"
)

// tests for the AdventOfCode 2015 day 15 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
FAKE:  capacity 0, durability 0, flavor 0, texture 0, calories 0
FAKE:  capacity 0, durability 0, flavor 0, texture 0, calories 0
`, "62842880"},
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
	{`Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3
FAKE:  capacity 0, durability 0, flavor 0, texture 0, calories 0
FAKE:  capacity 0, durability 0, flavor 0, texture 0, calories 0
`, "57600000"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
