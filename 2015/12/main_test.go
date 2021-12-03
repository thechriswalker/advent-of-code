package main

import (
	"testing"
)

// tests for the AdventOfCode 2015 day 12 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"[1,2,3]", "6"},
	{`{"a":2,"b":4}`, "6"},
	{`[[[3]]]`, "3"},
	{`{"a":{"b":4},"c":-1}`, "3"},
	{`{"a":[-1,1]}`, "0"},
	{`[-1,{"a":1}]`, "0"},
	{`[]`, "0"},
	{`{}`, "0"},
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
	{`{"a":[{"e":147},-1,142]}`, "288"},
	{`[[{"x":"red","b":1},2],{"red":2},2]`, "6"},
	{"[1,2,[3]]", "6"},
	{`[1,{"c":"red","b":4},3]`, "4"},
	{`{"d":"red","e":[1,2,3,4],"f":5}`, "0"},
	{`[1,"red",5]`, "6"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
