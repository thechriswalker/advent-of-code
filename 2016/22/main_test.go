package main

import (
	"testing"
)

// tests for the AdventOfCode 2016 day 22 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`root@ebhq-gridcenter# df -h
Filesystem              Size  Used  Avail  Use%
/dev/grid/node-x0-y0     91T   66T    25T   72%
/dev/grid/node-x0-y1     87T   25T    19T   78%`, "1"},
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
	// 	{`root@ebhq-gridcenter# df -h
	// Filesystem            Size  Used  Avail  Use%
	// /dev/grid/node-x0-y0   10T    8T     2T   80%
	// /dev/grid/node-x0-y1   11T    6T     5T   54%
	// /dev/grid/node-x0-y2   32T   28T     4T   87%
	// /dev/grid/node-x1-y0    9T    7T     2T   77%
	// /dev/grid/node-x1-y1    8T    0T     8T    0%
	// /dev/grid/node-x1-y2   11T    7T     4T   63%
	// /dev/grid/node-x2-y0   10T    6T     4T   60%
	// /dev/grid/node-x2-y1    9T    8T     1T   88%
	// /dev/grid/node-x2-y2    9T    6T     3T   66%`, "7"},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
