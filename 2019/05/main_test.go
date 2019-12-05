package main

import (
	"testing"

	"../intcode"
)

// tests for the AdventOfCode 2019 day 5 solutions

// need to change the defaults here
type Case struct {
	Program  string
	Input    int
	Expected int
}

// var problem1cases = []Case{
// 	// cases here
// 	//	{"", ""},
// }

func TestProblem1(t *testing.T) {
	// for _, c := range problem1cases {
	// 	// actual := solve1(c.In)
	// 	// if c.Out != actual {
	// 	// 	t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
	// 	// }
	// }
}

var problem2cases = []Case{
	// cases here
	{"3,9,8,9,10,9,4,9,99,-1,8", 8, 1},
	{"3,9,8,9,10,9,4,9,99,-1,8", 6, 0},
	{"3,9,8,9,10,9,4,9,99,-1,8", 10, 0},
	{"3,9,7,9,10,9,4,9,99,-1,8", 8, 0},
	{"3,9,7,9,10,9,4,9,99,-1,8", 9, 0},
	{"3,9,7,9,10,9,4,9,99,-1,8", 7, 1},
	{"3,9,7,9,10,9,4,9,99,-1,8", -199, 1},
	{"3,3,1108,-1,8,3,4,3,99", 8, 1},
	{"3,3,1108,-1,8,3,4,3,99", 6, 0},
	{"3,3,1108,-1,8,3,4,3,99", 10, 0},
	{"3,3,1107,-1,8,3,4,3,99", 8, 0},
	{"3,3,1107,-1,8,3,4,3,99", 9, 0},
	{"3,3,1107,-1,8,3,4,3,99", 7, 1},
	{"3,3,1107,-1,8,3,4,3,99", -111, 1},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		p := intcode.New(c.Program)
		p.EnqueueInput(c.Input)
		p.RunAsync(false)
		actual := p.GetOutput()
		if c.Expected != actual {
			t.Fatalf("Expected: '%d', Actual: '%d'", c.Expected, actual)
		}
	}
}
