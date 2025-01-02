package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 16, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '#')
	v := countEnergy(g, Beam{x: 0, y: 0, dir: EAST})
	return fmt.Sprint(v)
}

func countEnergy(g aoc.ByteGrid, init Beam) int {
	cache := map[Beam]struct{}{}

	curr := []Beam{init}
	cache[init] = struct{}{}

	var next []Beam

	for {
		next = []Beam{}
		for _, c := range curr {
			next = append(next, nextOptions(g, c, cache)...)
		}
		if len(next) == 0 {
			break
		}
		curr = next
	}

	// now count unique grid points.
	count := map[[2]int]struct{}{}

	for b := range cache {
		count[[2]int{b.x, b.y}] = struct{}{}
	}

	return len(count)
}

func nextOptions(g aoc.ByteGrid, c Beam, cache map[Beam]struct{}) (opts []Beam) {
	b, _ := g.At(c.x, c.y)
	// b is where we are.
	// where we can go next depends on our direction and the tile we are on.
	push := func(x, y int, dir byte) {
		// add this option if it is valid and not in the cache
		if !aoc.OOB(g, x, y) {
			beam := Beam{x: x, y: y, dir: dir}
			if _, seen := cache[beam]; !seen {
				cache[beam] = struct{}{}
				opts = append(opts, beam)
			}
		}
	}

	switch {
	case b == '-' && (c.dir == NORTH || c.dir == SOUTH):
		// east and west
		push(c.x-1, c.y, WEST)
		push(c.x+1, c.y, EAST)
	case b == '|' && (c.dir == EAST || c.dir == WEST):
		push(c.x, c.y-1, NORTH)
		push(c.x, c.y+1, SOUTH)
	case b == '/':
		switch c.dir {
		case NORTH: // now east
			push(c.x+1, c.y, EAST)
		case EAST:
			push(c.x, c.y-1, NORTH)
		case SOUTH:
			push(c.x-1, c.y, WEST)
		case WEST:
			push(c.x, c.y+1, SOUTH)
		}
	case b == '\\':
		switch c.dir {
		case NORTH: // now east
			push(c.x-1, c.y, WEST)
		case EAST:
			push(c.x, c.y+1, SOUTH)
		case SOUTH:
			push(c.x+1, c.y, EAST)
		case WEST:
			push(c.x, c.y-1, NORTH)
		}
	default:
		// continue in the same direction
		switch c.dir {
		case NORTH: // now east
			push(c.x, c.y-1, NORTH)
		case EAST:
			push(c.x+1, c.y, EAST)
		case SOUTH:
			push(c.x, c.y+1, SOUTH)
		case WEST:
			push(c.x-1, c.y, WEST)
		}
	}
	return
}

const (
	NORTH = 'N'
	SOUTH = 'S'
	EAST  = 'E'
	WEST  = 'W'
)

type Beam struct {
	x, y int
	dir  byte
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// here we need to find the MAX enegry
	g := aoc.CreateFixedByteGridFromString(input, '#')
	max := 0
	xmin, ymin, xmax, ymax := g.Bounds()

	check := func(init Beam) {
		if v := countEnergy(g, init); v > max {
			max = v
		}
	}

	// x up and down
	for x := xmin; x <= xmax; x++ {
		check(Beam{x: x, y: ymin, dir: SOUTH})
		check(Beam{x: x, y: ymax, dir: NORTH})
	}
	for y := ymin; y <= ymax; y++ {
		check(Beam{x: xmin, y: y, dir: EAST})
		check(Beam{x: xmax, y: y, dir: WEST})
	}

	return fmt.Sprint(max)
}
