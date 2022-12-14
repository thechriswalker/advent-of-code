package main

import (
	"fmt"
	"math"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 12, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, 'Z')

	var start, finish [2]int

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		switch b {
		case 'S':
			// get start pos
			start = [2]int{x, y}
			g.Set(x, y, 'a')
		case 'E':
			// get end and mark position as elevation 'z'
			finish = [2]int{x, y}
			g.Set(x, y, 'z')
		}
	})

	return fmt.Sprint(solveAt(g, start, finish))
}

type Cache map[[2]int]struct{}

func (c Cache) Has(x, y int) bool {
	_, ok := c[[2]int{x, y}]
	return ok
}
func (c Cache) Set(x, y int) {
	c[[2]int{x, y}] = struct{}{}
}

type Step struct {
	Previous  *Step
	Position  [2]int
	Elevation byte
	Count     int
}

func getNextOptions(g aoc.ByteGrid, p *Step, cache Cache) []*Step {
	var options []*Step
	max := p.Elevation + 1

	check := func(x, y int) {
		if b, oob := g.At(x, y); !oob && b <= max && !cache.Has(x, y) {
			// OK step this way
			cache.Set(x, y)
			options = append(options, &Step{
				Previous:  p,
				Position:  [2]int{x, y},
				Elevation: b,
				Count:     p.Count + 1,
			})
		}
	}

	check(p.Position[0]+1, p.Position[1])
	check(p.Position[0]-1, p.Position[1])
	check(p.Position[0], p.Position[1]+1)
	check(p.Position[0], p.Position[1]-1)

	return options
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, 'Z')

	var finish [2]int
	var lowpoints [][2]int

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		switch b {
		case 'a':
			lowpoints = append(lowpoints, [2]int{x, y})
		case 'S':
			// get start pos
			lowpoints = append(lowpoints, [2]int{x, y})
			g.Set(x, y, 'a')
		case 'E':
			// get end and mark position as elevation 'z'
			finish = [2]int{x, y}
			g.Set(x, y, 'z')
		}
	})

	// now solve all the low points.
	min := math.MaxInt

	for _, start := range lowpoints {
		n := solveAt(g, start, finish)
		if n < min {
			min = n
		}
	}

	return fmt.Sprint(min)
}

func solveAt(g aoc.ByteGrid, start, finish [2]int) int {
	// now from start, we need to walk every way possible until we hit a dead end or the finish!
	// the first one to hit will be the shortest path.

	// classic breath first search
	// where we have already been (don't double back)
	cache := map[[2]int]struct{}{
		start: {},
	}
	// the current list of steps we are checking
	current := []*Step{{Position: start, Elevation: 'a'}}
	// the list of possible steps next.
	var next []*Step
	for {
		for _, p := range current {
			// for each current step, see what options we have
			for _, s := range getNextOptions(g, p, cache) {
				// if one of these is the finish we are done!
				if s.Position == finish {
					return s.Count
				}
				next = append(next, s)
			}
		}
		if len(next) == 0 {
			// crap.
			// just bail.
			return math.MaxInt
		}
		// switch them up
		current = next
		next = nil
	}
}
