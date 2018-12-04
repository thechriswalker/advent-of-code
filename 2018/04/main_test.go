package main

import (
	"testing"
)

// tests for the AdventOfCode 2018 day 4 solutions

type Case struct {
	In  string
	Out string
}

const example = `
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-03 00:29] wakes up
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-05 00:55] wakes up
[1518-11-04 00:46] wakes up
[1518-11-01 00:55] wakes up
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-01 00:30] falls asleep
[1518-11-05 00:45] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-03 00:24] falls asleep
[1518-11-02 00:40] falls asleep
[1518-11-01 00:00] Guard #10 begins shift
[1518-11-02 00:50] wakes up
[1518-11-01 00:05] falls asleep
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-04 00:36] falls asleep
`

var problem1cases = []Case{
	// cases here
	{example, "240"},
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
	{example, "4455"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
