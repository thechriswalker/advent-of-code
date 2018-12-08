package main

import (
	"testing"
)

// tests for the AdventOfCode 2016 day 19 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"5", "3"},
}

func TestProblem1(t *testing.T) {
	///print a few cases...
	// for n := 1; n < 17; n++ {
	// 	x := NaiveSolution(n, false)
	// 	fmt.Println(n, formula(n), x)
	// }

	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	{"5", "2"},
}

func TestProblem2(t *testing.T) {
	//print a few cases...
	// for n := 2; n < 256; n++ {
	// 	x := NaiveSolution2(n)
	// 	fmt.Println(n, formula2(n), x)
	// }
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
