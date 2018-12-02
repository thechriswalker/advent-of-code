package main

import (
	"testing"
)

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	{"R2, L3", "5"},
	{"R2, R2, R2", "2"},
	{"R5, L5, R5, R3", "12"},
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
	{"R8, R4, R4, R8", "4"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
