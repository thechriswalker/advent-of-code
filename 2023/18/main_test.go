package main

import (
	"testing"
)

// tests for the AdventOfCode 2023 day 18 solutions

type Case struct {
	In  string
	Out string
}

const example = `R 6 (#70c710)
D 5 (#0dc571)
L 2 (#5713f0)
D 2 (#d2c081)
R 2 (#59c680)
D 2 (#411b91)
L 5 (#8ceee2)
U 2 (#caa173)
L 1 (#1b58a2)
U 2 (#caa171)
R 2 (#7807d2)
U 3 (#a77fa3)
L 2 (#015232)
U 2 (#7a21e3)`

var problem1cases = []Case{
	// cases here
	{ex2, "91"}, //100-8-1
	{example, "62"},
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
	{example, "952408144115"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

// I need to find where my real input differs from the test
// i.e. there is a case I haven't considered.
// the test grid has left and right indents, but not vertical ones. Let us make a box like this:

const _ = `
>>>>  ^###
#  V>>^  #
#    V<^ #
##   V ^<#
 #   >>>V
 #      #
##      ##
#        #
#  ####  #
####  ####
`
const ex2 = `R 3 X
D 1 X
R 3 X
U 1 X
R 3 X
D 3 X
L 2 X
U 1 X
L 2 X
D 2 X
R 3 X
D 2 X
R 1 X
D 3 X
L 3 X
U 1 X
L 3 X
D 1 X
L 3 X
U 3 X
R 1 X
U 3 X
L 1 X
U 3 X`
