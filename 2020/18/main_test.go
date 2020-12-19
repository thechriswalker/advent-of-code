package main

import (
	"testing"
)

// tests for the AdventOfCode 2020 day 18 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{"1 + (2 * 3) + (4 * (5 + 6))", "51"},
	{`2 * 3 + (4 * 5)`, `26`},
	{`5 + (8 * 3 + 9 + 3 * 4 * 3)`, `437`},
	{`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`, `12240`},
	{`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`, `13632`},
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
	{"1 + (2 * 3) + (4 * (5 + 6))", "51"},
	{`2 * 3 + (4 * 5)`, `46`},
	{`5 + (8 * 3 + 9 + 3 * 4 * 3)`, `1445`},
	{`5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))`, `669060`},
	{`((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2`, `23340`},
	/*
		  ((  6   * 9) * (  15  *   14 ) + 6) +   6   * 2
		  ((     54  ) * (      210    ) + 6) + 6 * 2
		  (      54    *            216     ) + 6 * 2
		  (          11664                  ) + 6 * 2
							  11670 * 2
							  23340
	*/

	{`(2*1)+(1*1)+3*2`, `12`},
	{`(1 + 1 * 2) + (1 * 2 + 1) + 1 * 2`, `16`},
	//(  2   * 2) + (1 *   3  ) + 1 * 2
	//(      4  ) + (  3      ) + 1 * 2
	//            7             + 1 * 2
	//                          8   * 2
	//                              16
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
