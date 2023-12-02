package main

import (
	"testing"
)

// tests for the AdventOfCode 2023 day 1 solutions

type Case struct {
	In  string
	Out string
}

const example = `1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet
`

var problem1cases = []Case{
	// cases here
	{example, "142"},
}

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var ex2 = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

var problem2cases = []Case{
	// cases here
	{ex2, "281"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
