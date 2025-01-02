package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 10, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {

	// make a grid. find the 0's
	// for each zero score it my finding 9s.
	g := aoc.CreateFixedByteGridFromString(input, '0')
	score := 0
	aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
		if b == '0' {
			score += scoreTrail(g, v, true)
		}
	})

	return fmt.Sprint(score)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// make a grid. find the 0's
	// for each zero score it my finding 9s.
	g := aoc.CreateFixedByteGridFromString(input, '0')
	score := 0
	aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
		if b == '0' {
			score += scoreTrail(g, v, false)
		}
	})

	return fmt.Sprint(score)
}

var directions = []aoc.V2{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func scoreTrail(g aoc.ByteGrid, zero aoc.V2, dedup bool) int {
	target := byte('1')
	// current trail positions
	curr := []aoc.V2{zero}
	var next []aoc.V2
	for {
		// can we find the target around the current position?
		for _, c := range curr {
			for _, d := range directions {
				if b, _ := g.Atv(c.Add(d)); b == target {
					// found it
					// add this point to the "next" list
					next = append(next, c.Add(d))
				}
			}
		}
		if target == '9' || len(next) == 0 {
			break
		}
		target++
		curr = next
		next = nil
	}
	// score is the number of unique "next" points.
	if !dedup || len(next) < 2 {
		return len(next) // 0 or one results must be unique.
	}
	// after that we should check.
	m := make(map[aoc.V2]struct{}, len(next))
	for _, n := range next {
		m[n] = struct{}{}
	}
	return len(m)
}
