package main

import (
	"../intcode"

	"testing"
)

// tests for the AdventOfCode 2019 day 9 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	//{"109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99", "99"},
	{"1102,34915192,34915192,7,4,7,99,0", "1219070632396864"},
	{"104,1125899906842624,99", "1125899906842624"},
}

func TestProblem1(t *testing.T) {
	for _, c := range problem1cases {
		actual := solve1(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
	// test opcode 203 which is failing.
	p := intcode.New("109,1,203,0,4,0,4,1,99")
	p.EnqueueInput(999)
	done := p.RunAsync()
	if x := <-p.Output; x != 109 {
		t.Fatalf("Expected: '%d', Actual: '%d'", 109, x)
	}
	if x := <-p.Output; x != 999 {
		t.Fatalf("Expected: '%d', Actual: '%d'", 999, x)
	}
	<-done
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
