package main

import (
	"testing"
)

// tests for the AdventOfCode 2015 day 14 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
`, "1120"},
}

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := maxDistance(c.In, 1000)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	{`Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.
`, "689"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := maxPoints(c.In, 1000)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
