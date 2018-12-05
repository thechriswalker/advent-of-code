package main

import (
	"testing"
)

// tests for the AdventOfCode 2016 day 12 solutions

type Case struct {
	In  string
	Out string
}

const prog = `cpy 41 a
inc a
inc a
dec a
jnz a 2
dec a
`

var problem1cases = []Case{
	// cases here
	{prog, "42"},
	{"cpy 111 a", "111"},
	{"cpy 111 b\ncpy b a", "111"},
	{"jnz b 2\ncpy 123 a", "123"},
	{"cpy 111 b\njnz b 2\ncpy 123 a", "0"},
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
	//	{"", ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
