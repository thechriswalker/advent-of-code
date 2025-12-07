package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 7, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	defer aoc.DisableDebug()()
	g := aoc.CreateFixedByteGridFromString(input, '.')

	positions := aoc.FindInGrid(g, 'S')

	var next []aoc.V2
	sum := 0
	for {
		for _, p := range positions {
			p = p.Add(aoc.South)
			b, oob := g.Atv(p)
			if b == '^' {
				// splitter
				r, l := p.Add(aoc.East), p.Add(aoc.West)
				if x, _ := g.Atv(r); x != '|' {
					next = append(next, r)
					g.Setv(r, '|')
				}
				if x, _ := g.Atv(l); x != '|' {
					next = append(next, l)
					g.Setv(l, '|')
				}
				sum++
			} else if !oob && b != '|' {
				next = append(next, p)
				g.Setv(p, '|')
			}
		}
		if len(next) == 0 {
			break
		}
		positions = next
		next = []aoc.V2{}
		if aoc.IsDebug() {
			aoc.PrintByteGridC(g, map[byte]aoc.Color{'|': aoc.BoldCyan, '^': aoc.BoldGreen})
		}
	}

	// 1906 is too high
	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2#
// in this one we don't care about the duplicates.
// but the naive way is too slow...
// of course if we have points in the same place, we
// can just run the course from there and double up the
// answer.
func solve2(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '.')

	positions := aoc.FindInGrid(g, 'S')

	timelines := map[aoc.V2]int{
		positions[0]: 1,
	}
	var next []aoc.V2
	for {
		nextTimelines := map[aoc.V2]int{}
		for _, p := range positions {
			t := timelines[p]
			s := p.Add(aoc.South)
			b, oob := g.Atv(s)
			if b == '^' {
				// splitter
				r, l := s.Add(aoc.East), s.Add(aoc.West)
				nextTimelines[r] = nextTimelines[r] + t
				nextTimelines[l] = nextTimelines[l] + t
				if x, _ := g.Atv(r); x != '|' {
					next = append(next, r)
					g.Setv(r, '|')
				}
				if x, _ := g.Atv(l); x != '|' {
					next = append(next, l)
					g.Setv(l, '|')
				}
			} else if !oob {
				nextTimelines[s] = nextTimelines[s] + t
				if b != '|' {
					next = append(next, s)
					g.Setv(s, '|')
				}
			}
		}
		aoc.Debug("positions", positions)
		aoc.Debug("timelines", nextTimelines)
		if len(next) == 0 {
			sum := 0
			// nothing changed so our old timelines are the ones we need to count.
			for _, n := range timelines {
				sum += n
			}
			return fmt.Sprint(sum)
		}
		positions = next
		timelines = nextTimelines
		next = []aoc.V2{}
	}
}
