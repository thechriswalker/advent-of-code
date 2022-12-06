package main

import (
	"testing"
)

// tests for the AdventOfCode 2022 day 6 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", "5"},
	{"nppdvjthqldpwncqszvftbrmjlhg", "6"},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "10"},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "11"},
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
	{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", "19"},
	{"bvwbjplbgvbhsrlpgdmjqwftvncz", "23"},
	{"nppdvjthqldpwncqszvftbrmjlhg", "23"},
	{"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", "29"},
	{"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", "26"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
