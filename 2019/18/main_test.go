package main

import (
	"testing"
)

// tests for the AdventOfCode 2019 day 18 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`#########
#b.A.@.a#
#########
`, "8"}, {`########################
#f.D.E.e.C.b.A.@.a.B.c.#
######################.#
#d.....................#
########################
`, "86"},
	{`########################
#...............b.C.D.f#
#.######################
#.....@.a.B.c.d.A.e.F.g#
########################`, "132"},
	{`#################
#i.G..c...e..H.p#
########.########
#j.A..b...f..D.o#
########@########
#k.E..a...g..B.n#
########.########
#l.F..d...h..C.m#
#################`, "136"},
	{`########################
#@..............ac.GI.b#
###d#e#f################
###A#B#C################
###g#h#i################
########################
`, "81"},
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
	//{"", ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
