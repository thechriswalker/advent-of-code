package main

import (
	"fmt"
	"math"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 17, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '.')
	h := minHeatLossPath(g)
	return fmt.Sprint(h)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '.')
	h := minHeatLossPathUltra(g)
	// 993 too high?
	return fmt.Sprint(h)
}

const (
	NORTH = 'N'
	SOUTH = 'S'
	EAST  = 'E'
	WEST  = 'W'
)

type Pos struct {
	x, y       int
	dir        byte
	stepsInDir int
}

type Path struct {
	pos      Pos
	heatloss int
}

func minHeatLossPath(g aoc.ByteGrid) int {

	// walk the path, how do we cull options?
	// never touch the same square again during a run?
	// never touch the same square in the same direction?
	// neve touch the same square in the same direction with
	// the same number of squares moved in that direction?

	// unless your heatloss is lower in which case we
	// I think each path has to cache it self and not re-tread
	// a step. but a pos / step count cache might be useful

	// finish is xmax,ymax
	_, _, xmax, ymax := g.Bounds()

	cache := map[Pos]int{}

	//
	curr := []Path{
		// we pretend to be moving east with 0 steps.
		// this means we can turn south now or continue
		// east and have only one step
		{pos: Pos{x: 0, y: 0, dir: EAST, stepsInDir: 0}, heatloss: 0},
	}

	cache[curr[0].pos] = curr[0].heatloss

	var next []Path

	minHeatloss := math.MaxInt

	for {
		next = []Path{}

		for _, c := range curr {
			for _, o := range options(g, c) {
				//fmt.Printf("option: %#v\n", o)
				if o.pos.x == xmax && o.pos.y == ymax {
					// we are done stop this path
					if minHeatloss > o.heatloss {
						minHeatloss = o.heatloss
					}
				} else if o.heatloss < minHeatloss {
					// I feel like we need a per-path backtrack protection...
					// or use our cache
					if s, seen := cache[o.pos]; !seen || s > o.heatloss {
						cache[o.pos] = o.heatloss
						next = append(next, o)
					}
				}
			}
		}
		if len(next) == 0 {
			break
		}
		curr = next
	}
	return minHeatloss
}

func options(g aoc.ByteGrid, p Path) []Path {
	opts := []Path{}

	try := func(dir byte) {
		x, y := p.pos.x, p.pos.y
		switch dir {
		case NORTH:
			y--
		case SOUTH:
			y++
		case EAST:
			x++
		case WEST:
			x--
		}
		b, oob := g.At(x, y)
		if !oob {
			steps := p.pos.stepsInDir + 1
			if dir != p.pos.dir {
				steps = 1
			}
			opts = append(opts, Path{
				pos:      Pos{x: x, y: y, dir: dir, stepsInDir: steps},
				heatloss: p.heatloss + int(b-'0'),
			})
		}
	}

	// which direction are we going and how far have we gone.
	if p.pos.stepsInDir < 3 {
		// we can go straight
		try(p.pos.dir)
	}

	// otherwise we can turn left or right.
	switch p.pos.dir {
	case NORTH, SOUTH:
		try(EAST)
		try(WEST)
	case EAST, WEST:
		try(NORTH)
		try(SOUTH)
	}
	return opts
}

func minHeatLossPathUltra(g aoc.ByteGrid) int {

	// walk the path, how do we cull options?
	// never touch the same square again during a run?
	// never touch the same square in the same direction?
	// neve touch the same square in the same direction with
	// the same number of squares moved in that direction?

	// unless your heatloss is lower in which case we
	// I think each path has to cache it self and not re-tread
	// a step. but a pos / step count cache might be useful

	// finish is xmax,ymax
	_, _, xmax, ymax := g.Bounds()

	cache := map[Pos]int{}

	//
	curr := []Path{
		// we pretend to be moving east with 0 steps.
		// this means we can turn south now or continue
		// east and have only one step
		{pos: Pos{x: 0, y: 0, dir: EAST, stepsInDir: 0}, heatloss: 0},
	}

	cache[curr[0].pos] = curr[0].heatloss

	var next []Path

	minHeatloss := math.MaxInt

	for {
		next = []Path{}

		for _, c := range curr {
			for _, o := range optionsUltra(g, c) {
				//fmt.Printf("option: %#v\n", o)
				if o.pos.x == xmax && o.pos.y == ymax && o.pos.stepsInDir >= 4 {
					// we are done stop this path
					if minHeatloss > o.heatloss {
						minHeatloss = o.heatloss
					}
				} else if o.heatloss < minHeatloss {
					// I feel like we need a per-path backtrack protection...
					// or use our cache
					if s, seen := cache[o.pos]; !seen || s > o.heatloss {
						cache[o.pos] = o.heatloss
						next = append(next, o)
					}
				}
			}
		}
		if len(next) == 0 {
			break
		}
		curr = next
	}
	return minHeatloss
}

func optionsUltra(g aoc.ByteGrid, p Path) []Path {
	opts := []Path{}

	try := func(dir byte) {
		x, y := p.pos.x, p.pos.y
		switch dir {
		case NORTH:
			y--
		case SOUTH:
			y++
		case EAST:
			x++
		case WEST:
			x--
		}
		b, oob := g.At(x, y)
		if !oob {
			steps := p.pos.stepsInDir + 1
			if dir != p.pos.dir {
				steps = 1
			}
			opts = append(opts, Path{
				pos:      Pos{x: x, y: y, dir: dir, stepsInDir: steps},
				heatloss: p.heatloss + int(b-'0'),
			})
		}
	}

	// we must go 4 in a row (with an exception for 0 so our starting point works...)
	if p.pos.stepsInDir > 0 && p.pos.stepsInDir < 4 {
		// we have to keep going!
		try(p.pos.dir)
		return opts
	}

	// other wise our max is 10

	// which direction are we going and how far have we gone.
	if p.pos.stepsInDir < 10 {
		// we can go straight
		try(p.pos.dir)
	}

	// otherwise we can turn left or right.
	switch p.pos.dir {
	case NORTH, SOUTH:
		try(EAST)
		try(WEST)
	case EAST, WEST:
		try(NORTH)
		try(SOUTH)
	}
	return opts
}
