package main

import (
	"fmt"
	"sort"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 9, solve1, solve2)
}

type Grid struct {
	Points []byte
	Stride int
	Height int
}

func (g *Grid) Index(x, y int) int {
	return aoc.GridIndex(x, y, g.Stride, g.Height)
}

func (g *Grid) At(x, y int) int {
	idx := g.Index(x, y)
	if idx == -1 {
		return 9 // highest possible
	}
	return int(g.Points[idx] - '0')
}

func (g *Grid) IsLowPoint(x, y int) (int, bool) {
	v := g.At(x, y)
	n, s, e, w := g.At(x-1, y), g.At(x+1, y), g.At(x, y+1), g.At(x, y-1)
	return v, v < n && v < s && v < e && v < w
}

func (g *Grid) BasinSize(x, y int) int {
	n, ok := g.IsLowPoint(x, y)
	if !ok {
		return 0
	}
	// the points we have covered.
	basin := map[[2]int]int{{x, y}: n}
	curr := [][2]int{{x, y}}
	var next [][2]int
	adjacent := [][2]int{{-1, 0}, {1, 0}, {0, 1}, {0, -1}}
	for {
		// start fresh
		next = [][2]int{}
		// go in all four directions. for each current point
		for _, xy := range curr {
			// this should be in the basin already.
			v := basin[xy]
			for _, d := range adjacent {
				p := [2]int{xy[0] + d[0], xy[1] + d[1]}
				if _, ok := basin[p]; ok {
					// already done this one
					continue
				}
				// check it
				n := g.At(p[0], p[1])
				// we went up, but not the top
				if n > v && n != 9 {
					// this is a new candidate.
					basin[p] = n
					next = append(next, p)
				}
			}
		}
		// any more
		if len(next) == 0 {
			// nope.
			return len(basin)
		}
		// otherwise we need to recurse
		curr = next
	}
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := Grid{
		Points: []byte{},
	}
	aoc.MapLines(input, func(line string) error {
		g.Stride = len(line)
		g.Height++
		g.Points = append(g.Points, []byte(line)...)
		return nil
	})

	// count the low points. technically, this is excessive, we don't
	// need to count ALL, the points as those adjacent to a low point
	// AREN'T by definition.
	sum := 0
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Stride; x++ {
			if n, ok := g.IsLowPoint(x, y); ok {
				sum += 1 + n
			}
		}
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := Grid{
		Points: []byte{},
	}
	aoc.MapLines(input, func(line string) error {
		g.Stride = len(line)
		g.Height++
		g.Points = append(g.Points, []byte(line)...)
		return nil
	})

	// count the low points. technically, this is excessive, we don't
	// need to count ALL, the points as those adjacent to a low point
	// AREN'T by definition.
	basins := []int{}
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Stride; x++ {
			b := g.BasinSize(x, y)
			if b > 0 {
				basins = append(basins, b)
			}
		}
	}
	// sort them with largest first
	sort.Sort(sort.Reverse(sort.IntSlice(basins)))

	// multiply the top three
	p := basins[0] * basins[1] * basins[2]

	return fmt.Sprintf("%d", p)
}
