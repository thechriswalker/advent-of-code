package main

import (
	"testing"
)

// tests for the AdventOfCode 2023 day 7 solutions

type Case struct {
	In  string
	Out string
}

const example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

var problem1cases = []Case{
	// cases here
	{example, "6440"},
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
	{example, "5905"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

// func TestJoker(t *testing.T) {
// 	games := []*Game{
// 		parseGame("2282K 1", true),
// 		parseGame("228JK 1", true),
// 		parseGame("J2A23 1", true),
// 	}
// 	slices.SortFunc(games, gameSort)

// 	fmt.Println(games)
// 	t.Fail()

// }
