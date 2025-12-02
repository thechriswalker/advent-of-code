package main

import (
	"testing"
)

// tests for the AdventOfCode 2025 day 2 solutions

type Case struct {
	In  string
	Out string
}

const example = `11-22,95-115,998-1012,1188511880-1188511890,222220-222224,` +
	`1698522-1698528,446443-446449,38593856-38593862,565653-565659,` +
	`824824821-824824827,2121212118-2121212124`

var problem1cases = []Case{
	// cases here
	{example, "1227775554"},
	{"3433355031-3433496616", "3433434334"},
}

// ----- Range 3433355031 - 3433496616
// ----- Range 3433355031 - 3433496616
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
	{example, "4174379265"},
	{"95-115", "210"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
