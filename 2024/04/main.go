package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 4, solve1, solve2)
}

func check(g aoc.ByteGrid, x, y, dx, dy int) bool {
	x += dx
	y += dy
	b, _ := g.At(x, y)
	if b != 'M' {
		return false
	}
	x += dx
	y += dy
	b, _ = g.At(x, y)
	if b != 'A' {
		return false
	}
	x += dx
	y += dy
	b, _ = g.At(x, y)
	return b == 'S'
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '.')

	found := 0

	// find the word XMAS the as many times as possible.
	// so if we hit an X, we can check around it.
	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		if b == 'X' {
			// north
			if check(g, x, y, 0, -1) {
				found++
			}
			// NE
			if check(g, x, y, 1, -1) {
				found++
			}
			// East
			if check(g, x, y, 1, 0) {
				found++
			}
			// SE
			if check(g, x, y, 1, 1) {
				found++
			}
			// South
			if check(g, x, y, 0, 1) {
				found++
			}
			// SW
			if check(g, x, y, -1, 1) {
				found++
			}
			// West
			if check(g, x, y, -1, 0) {
				found++
			}
			// NW
			if check(g, x, y, -1, -1) {
				found++
			}
		}
	})

	return fmt.Sprint(found)

}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '.')
	found := 0

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		// this time we check for 'A' which will be the center,
		// the around it we need 'M' and 'S' at opposite corners.
		if b != 'A' {
			return
		}
		// try NE and SW
		c1, _ := g.At(x+1, y-1)
		c2, _ := g.At(x-1, y+1)
		if c1 == 'M' && c2 == 'S' || c1 == 'S' && c2 == 'M' {
			// OK, now the other corners
			c1, _ = g.At(x-1, y-1)
			c2, _ = g.At(x+1, y+1)
			if c1 == 'M' && c2 == 'S' || c1 == 'S' && c2 == 'M' {
				found++
			}
		}

	})

	return fmt.Sprint(found)
}
