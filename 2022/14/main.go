package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 14, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := parseGrid(input)
	// let's "set" the sand entry so the bounds shift.
	g.Set(500, 0, '.')

	i := 0
	for {
		//aoc.PrintByteGrid(g, nil)
		if dropSand(g, pos{x: 500, y: 0}, 0) {
			break
		}
		i++
	}

	return fmt.Sprint(i)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := parseGrid(input)
	// let's "set" the sand entry so the bounds shift.
	seed := pos{x: 500, y: 0}

	g.Set(seed.x, seed.y, '+')

	// now we find the y bound and put a piece there to extend the "in bounds" grid to the floor.
	_, _, _, y := g.Bounds()

	floor := y + 2
	g.Set(seed.x, floor, '=')
	// for x := seed.x - 100; x <= seed.x+100; x++ {
	// 	g.Set(x, floor, '=')
	// }

	i := 0
	for {
		// if i%10 == 0 {
		// 	fmt.Println("Iteration:", i, "Floor level:", floor)
		// 	aoc.PrintByteGrid(g, nil)
		// }
		// we don't care about OOB, here.
		dropSand(g, pos{x: 500, y: 0}, floor)
		i++
		if b, _ := g.At(seed.x, seed.y); b == 'o' {
			break
		}
	}

	return fmt.Sprint(i)
}

type pos struct{ x, y int }

func parseGrid(input string) aoc.ByteGrid {
	g := aoc.NewSparseByteGrid('.')

	// parse the lines of blocks.
	aoc.MapLines(input, func(line string) error {
		pieces := strings.Split(line, " -> ")
		coords := make([]pos, len(pieces))
		for i, p := range pieces {
			c := coords[i]
			fmt.Sscanf(p, "%d,%d", &(c.x), &(c.y))
			coords[i] = c
		}
		current := coords[0]
		for i := 1; i < len(coords); i++ {
			// layout from current to next.
			next := coords[i]
			//...
			// one of the coords should match, so we are going either horizontal or vertical.
			if current.x == next.x {
				// vertical
				a, b := current.y, next.y
				if a > b {
					a, b = b, a
				}
				for y := a; y <= b; y++ {
					g.Set(current.x, y, '#')
				}
			} else if current.y == next.y {
				// horizontal
				a, b := current.x, next.x
				if a > b {
					a, b = b, a
				}
				for x := a; x <= b; x++ {
					g.Set(x, current.y, '#')
				}
			}
			current = next
		}
		return nil
	})
	return g
}

const (
	OK    = 0
	BLOCK = 1
	OOB   = 2
)

func dropSand(g aoc.ByteGrid, from pos, floor int) bool {
	// return whether we stopped or whether we fell into the abyss
	// sand spawns at "from"
	sand := from

	tryMove := func(x, y int) int {
		p, oob := g.At(x, y)
		if oob {
			return OOB
		}
		if p == '.' {
			return OK
		}
		return BLOCK // either sand or rock
	}

	// if floor > 0
	// then we need to check for "on the floor" and ignore x bounds

fall:
	for {
		if floor > 0 && floor == sand.y+1 {
			// we reached the floor.
			g.Set(sand.x, sand.y, 'o')
			return false
		}
		// next move is one down if possible.
		switch tryMove(sand.x, sand.y+1) {
		case OOB:
			// we don't care about OOB here. treat as "empty"
			if floor > 0 {
				// move that way
				sand.y++
				continue fall
			}
			return true
		case BLOCK:
			// down blocked, we still have some to try
		case OK:
			// move that way
			sand.y++
			continue fall
		}
		// at this point we know DOWN wasn't oob.
		// so if floor > 0 then we don't break on OOB.

		// try left.
		switch tryMove(sand.x-1, sand.y+1) {
		case OOB:
			if floor > 0 {
				// move that way
				// and extend our grid out that way.
				sand.y++
				sand.x--
				continue fall
			}
			return true
		case BLOCK:
			// down blocked, we still have some to try
		case OK:
			// move that way
			sand.y++
			sand.x--
			continue fall
		}
		// try right
		switch tryMove(sand.x+1, sand.y+1) {
		case OOB:
			if floor > 0 {
				// move that way
				sand.y++
				sand.x++
				continue fall
			}
			return true
		case BLOCK:
			// down blocked, nowhere to GO
			g.Set(sand.x, sand.y, 'o')
			return false
		case OK:
			// move that way
			sand.y++
			sand.x++
			continue fall
		}
		panic("unreachable")
	}
}
