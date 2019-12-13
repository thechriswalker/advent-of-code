package main

import (
	"testing"
)

// tests for the AdventOfCode 2019 day 12 solutions

type Case struct {
	In    string
	Out   string
	Steps int
}

var problem1cases = []Case{
	// cases here
	{`<x=-1, y=0, z=2>
<x=2, y=-10, z=-7>
<x=4, y=-8, z=8>
<x=3, y=5, z=-1>
`, "179", 10},
	{`<x=-8, y=-10, z=0>
<x=5, y=5, z=10>
<x=2, y=-7, z=3>
<x=9, y=-8, z=-3>
`, "1940", 100},
}

// func TestAAA_Trie(t *testing.T) {
// 	tr := Trie{}

// 	if tr.HasOrSet([3]int16{1, 2, 3}, [3]int16{4, 5, 6}) {
// 		t.Fatalf("setting unknown value should be false")
// 	}
// 	if tr.HasOrSet([3]int16{1, 2, 3}, [3]int16{4, 5, 7}) {
// 		t.Fatalf("setting unknown value should be false")
// 	}
// 	if tr.HasOrSet([3]int16{1, 2, 4}, [3]int16{4, 5, 6}) {
// 		t.Fatalf("setting unknown value should be false")
// 	}
// 	if !tr.HasOrSet([3]int16{1, 2, 3}, [3]int16{4, 5, 6}) {
// 		t.Fatalf("setting known value should be true")
// 	}
// 	if !tr.HasOrSet([3]int16{1, 2, 3}, [3]int16{4, 5, 6}) {
// 		t.Fatalf("setting known value should be true")
// 	}
// 	if !tr.HasOrSet([3]int16{1, 2, 3}, [3]int16{4, 5, 7}) {
// 		t.Fatalf("setting known value should be true")
// 	}
// }

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := solve1steps(c.In, c.Steps)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

var problem2cases = []Case{
	// cases here
	{"<x=-1, y=0, z=2>\n<x=2, y=-10, z=-7>\n<x=4, y=-8, z=8>\n<x=3, y=5, z=-1>", "2772", 0},
	{"<x=-8, y=-10, z=0>\n<x=5, y=5, z=10>\n<x=2, y=-7, z=3>\n<x=9, y=-8, z=-3>", "4686774924", 0},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}
