package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

// tests for the AdventOfCode 2016 day 21 solutions

type Case struct {
	In   string
	Pass string
	Out  string
}

var instructions = `swap position 4 with position 0
swap letter d with letter b
reverse positions 0 through 4
rotate left 1 step
move position 1 to position 4
move position 3 to position 0
rotate based on position of letter b
rotate based on position of letter d
`

var problem1cases = []Case{
	// cases here
	{instructions, "abcde", "decab"},
}

func TestProblem1(t *testing.T) {

	for _, c := range problem1cases {
		Password1 = c.Pass
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	{instructions, "decab", "abcde"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		Password2 = c.Pass
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

func TestLineByLine(t *testing.T) {
	mainInput, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(mainInput), "\n")
	prev := "abcdefgh"
	var next string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) < 8 {
			continue
		}
		// one way
		Password1 = prev
		next = solve1(line)
		Password2 = next
		prev = solve2(line)
		if prev != Password1 {
			t.Fatalf("Line: %s\nConversion: %s => %s => %s\n", line, Password1, next, prev)
		}
		prev = next
	}
}
