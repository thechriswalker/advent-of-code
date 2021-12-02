package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/2019/intcode"
	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 2, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	p := intcode.New(input)
	// so simulate the actual problem, not the tests
	if len(input) > 50 {
		p.Set(1, 12)
		p.Set(2, 2)
	}
	p.Run()
	return fmt.Sprintf("%d", p.Get(0))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	target := int64(19690720)
	clean := intcode.New(input)
	// naive exhaustive search
	for n := int64(0); n < 100; n++ {
		for v := int64(0); v < 100; v++ {
			p := clean.Copy()
			p.Set(1, n)
			p.Set(2, v)
			if p.Run() {
				if p.Get(0) == target {
					// found answer
					return fmt.Sprintf("%d", 100*n+v)
				}
			}
		}
	}
	return "<unsolved>"
}
