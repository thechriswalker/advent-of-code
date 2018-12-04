package main

import (
	"testing"
)

// tests for the AdventOfCode 2016 day 11 solutions

type Case struct {
	In  string
	Out string
}

var problem1cases = []Case{
	// cases here
	{`The first floor contains a hydrogen-compatible microchip and a lithium-compatible microchip.
The second floor contains a hydrogen generator.
The third floor contains a lithium generator.
The fourth floor contains nothing relevant.`, "11"},
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
	//	{"", ""},
}

func TestProblem2(t *testing.T) {
	for _, c := range problem2cases {
		actual := solve2(c.In)
		if c.Out != actual {
			t.Fatalf("Expected: '%s', Actual: '%s'", c.Out, actual)
		}
	}
}

func generateAllStates() []*State {
	// generate all possible state combinations for 2 pairs of G/M.
	// double loop creating test cases for all state combinations.
	// check the hash function doesn't collide
	states := []*State{}
	// 5 items and 4 floors (yes some are illegal but we check them anyway)
	// so for each
	for lift := 0; lift < FLOORS; lift++ {
		for gen1 := 0; gen1 < FLOORS; gen1++ {
			for chp1 := 0; chp1 < FLOORS; chp1++ {
				for gen2 := 0; gen2 < FLOORS; gen2++ {
					for chp2 := 0; chp2 < FLOORS; chp2++ {
						states = append(states, &State{
							Lift:  uint8(lift),
							Items: []uint8{uint8(gen1), uint8(chp1), uint8(gen2), uint8(chp2)},
						})
					}
				}
			}
		}
	}
	return states
}

func TestHash(t *testing.T) {

	states := generateAllStates()
	//pristine := make([]*State, len(states))
	for _, s := range states {
		//pristine[i] = s.Clone()
		s.sorted = true // pretend to see how the hash function collisions work, there should be none
	}
	cache := map[string]*State{}

	for _, c := range states {
		h := hash(c)
		if s, ok := cache[h]; ok {
			t.Logf("Collision (%s): %s vs %s", h, s, c)
		} else {
			cache[h] = c
		}
	}
}
