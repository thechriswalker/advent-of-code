package main

import (
	"testing"
)

// tests for the AdventOfCode 2016 day 7 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"abba[mnop]qrst", "1"},
	{"abcd[bddb]xyyx", "0"},
	{"aaaa[qwer]tyui", "0"},
	{"ioxxoj[asdfgh]zxcvbn", "1"},
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
	{"aba[bab]xyz", "1"},
	{"xyx[xyx]xyx", "0"},
	{"aaa[kek]eke", "1"},
	{"zazbz[bzb]cdb", "1"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
