package main

import (
	"testing"
)

// tests for the AdventOfCode 2017 day 10 solutions

type Case struct {
	In  string
	Out string
}

const example = "3, 4, 1, 5"

var problem1cases = []Case{
	// cases here
	{example, "12"},
}

func TestProblem1(t *testing.T) {
	ringSize = 5
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	//{example, ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

func TestKnotHash(t *testing.T) {
	tests := [][2]string{
		{"", "a2582a3a0e66e6e86e3812dcb672a272"},
		{"AoC 2017", "33efeb34ea91902bb2f59c9920caa6cd"},
		{"1,2,3", "3efbe78a8d82f29979031a4aa0b16a9d"},
		{"1,2,4", "63960835bcdc130f0b66d7ff4f6a5a8e"},
	}
	for _, tt := range tests {
		input, expected := tt[0], tt[1]
		actual := KnotHash(input)
		if actual != expected {
			t.Errorf("Input=%q Expected=%s Actual=%s", input, expected, actual)
		}
	}
}
