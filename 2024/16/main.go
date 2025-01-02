package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 16, solve1, solve2)
}

type Reindeer struct {
	Position  aoc.V2
	Direction aoc.V2 // wasteful, but easy to use
}

type State struct {
	Reindeer
	Cost    int
	Visited []aoc.V2
}

// Implement Solution to Problem 1
func solve1(input string) string {
	cost, _ := solve(input, false)
	return fmt.Sprint(cost)
}

func solve(input string, debug bool) (cost, tiles int) {

	// make a grid and do a depth-first search with postion/direction => score
	// as a cache. If we hit a match with a higher score, we can stop that branch.
	// after all paths are done, we will have the lowest score for the target.
	g := aoc.CreateFixedByteGridFromString(input, '#')
	var start, end aoc.V2
	aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
		if b == 'S' {
			start = v
		}
		if b == 'E' {
			end = v
		}
	})

	// cache is the minimum found cost at a given state.
	cache := map[Reindeer]int{}
	stepCost := 1
	turnCost := 1000
	curr := []State{{Reindeer: Reindeer{start, aoc.East}, Cost: 0, Visited: []aoc.V2{start}}}
	cache[curr[0].Reindeer] = 0

	completed := []State{}

	for {
		next := []State{}
		for _, c := range curr {
			if c.Position == end {
				continue
			}
			// try all the options.
			// we can turn anticlockwise, clockwise, or go straight.
			// try to go straight.
			straight := c.Position.Add(c.Direction)
			if b, _ := g.Atv(straight); b != '#' {
				// we could go straight.
				nextState := State{
					Reindeer: Reindeer{straight, c.Direction},
					Cost:     c.Cost + stepCost,
					Visited:  copyAndAppend(c.Visited, straight),
				}
				if min, ok := cache[nextState.Reindeer]; !ok || min >= nextState.Cost {
					cache[nextState.Reindeer] = nextState.Cost
					next = append(next, nextState)
				}
				if straight == end && cache[nextState.Reindeer] >= nextState.Cost {
					completed = append(completed, nextState)
				}
			}
			// could we turn left?
			left := aoc.CardinalAnticlockwise(c.Direction)
			if b, _ := g.Atv(c.Position.Add(left)); b != '#' {
				// we could go left, so let's try it.
				nextState := State{Reindeer: Reindeer{c.Position, left}, Cost: c.Cost + turnCost, Visited: c.Visited[:]}

				if min, ok := cache[nextState.Reindeer]; !ok || min >= nextState.Cost {
					cache[nextState.Reindeer] = nextState.Cost
					next = append(next, nextState)
				}
			}
			// could we turn right?
			right := aoc.CardinalClockwise(c.Direction)
			if b, _ := g.Atv(c.Position.Add(right)); b != '#' {
				// we could go right, so let's try it.
				nextState := State{Reindeer: Reindeer{c.Position, right}, Cost: c.Cost + turnCost, Visited: c.Visited[:]}
				if min, ok := cache[nextState.Reindeer]; !ok || min >= nextState.Cost {
					cache[nextState.Reindeer] = nextState.Cost
					next = append(next, nextState)
				}
			}
		}
		if len(next) == 0 {
			break
		}
		// swap.
		curr = next
	}

	// now the minimum of all the cache values
	min := 1 << 31
	for r, c := range cache {
		if r.Position == end && c < min {
			min = c
		}
	}
	visitCache := map[aoc.V2]struct{}{}
	// and the number of tiles visited in the "best" routes.
	for _, c := range completed {
		if c.Cost == min {
			for _, v := range c.Visited {
				visitCache[v] = struct{}{}
				g.Setv(v, 'O')
			}
		}
	}

	if debug {
		aoc.PrintByteGridC(g, map[byte]aoc.Color{
			'O': aoc.BoldMagenta,
			'#': aoc.BoldWhite,
		})
	}

	return min, len(visitCache)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// this time we need to keep track of our path.
	_, tiles := solve(input, true)
	return fmt.Sprint(tiles)
}

func copyAndAppend[T any](s []T, v T) []T {
	c := make([]T, len(s)+1)
	copy(c, s)
	c[len(s)] = v
	return c
}
