package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 20, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// create the enchancment
	enhance := parseEnhance(input[0:512])

	g := aoc.CreateFixedByteGridFromString(input[514:], '.')
	fmt.Println("original")
	aoc.PrintByteGrid(g, nil)

	g1 := runEnhance(enhance, g)
	fmt.Println("\n\nfirst pass")
	aoc.PrintByteGrid(g1, nil)
	g2 := runEnhance(enhance, g1)
	fmt.Println("\n\nsecond pass")
	aoc.PrintByteGrid(g2, nil)

	lit := 0
	aoc.IterateByteGrid(g2, func(_, _ int, b byte) {
		if b == '#' {
			lit++
		}
	})

	return fmt.Sprint(lit)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return "<unsolved>"
}

func parseEnhance(input string) (e int) {
	for i := 0; i < len(input); i++ {
		switch input[i] {
		case '#':
			e |= i
		}
	}
	return
}

func runEnhance(enhance int, g aoc.ByteGrid) aoc.ByteGrid {
	// we will make a Sparse one.
	next := aoc.NewSparseByteGrid('.')
	// we actually need to iterate from "outside" the bounds...
	// because we want all squares that have at least 1 pixel inside the image
	// which means 1 extra pixel each side.
	x1, y1, x2, y2 := g.Bounds()
	x1 -= 1
	y1 -= 1
	x2 += 1
	y2 += 1
	for x := x1; x < x2; x++ {
		for y := y1; y < y2; y++ {
			// we need the nine pixels around this one.
			code := find9(g, x, y)
			lit := ((enhance >> code) & 1) == 1
			if lit {
				next.Set(x, y, '#')
			} else {
				next.Set(x, y, '.')
			}
		}
	}
	return next
}

func find9(g aoc.ByteGrid, x, y int) int {
	n := 8
	v := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			b, _ := g.At(x+i, y+j)
			if b == '#' {
				v |= 1 << n
			}
			n--
		}
	}
	return v
}
