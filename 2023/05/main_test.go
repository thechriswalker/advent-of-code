package main

import (
	"testing"
)

// tests for the AdventOfCode 2023 day 5 solutions

type Case struct {
	In  string
	Out string
}

const example = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

var problem1cases = []Case{
	// cases here
	{example, "35"},
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
	{example, "46"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

func TestGetRange(t *testing.T) {
	// m := Map{}
	// m.SetRange(MapRange{src: 100, dst: 200, len: 10})
	// m.SetRange(MapRange{src: 120, dst: 0, len: 10})
	// m.Sort()
	// fmt.Println(m.GetRanges(ValueRange{min: 50, max: 150}))
}
