package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 8, solve1, solve2)
}

func left(x, y int) (int, int)  { return x - 1, y }
func right(x, y int) (int, int) { return x + 1, y }
func up(x, y int) (int, int)    { return x, y - 1 }
func down(x, y int) (int, int)  { return x, y + 1 }

// Implement Solution to Problem 1
func solve1(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '0')

	visible := 0

	w, h := g.Width()-1, g.Height()-1

	isExterior := func(x, y int) bool {
		return x == 0 || y == 0 || x == w || y == h
	}
	var isVisible func(x, y int, b byte, next func(i, j int) (int, int)) bool
	isVisible = func(x, y int, b byte, next func(i, j int) (int, int)) bool {
		x, y = next(x, y)
		v, oob := g.At(x, y)
		if oob {
			return true
		}
		if v >= b {
			return false
		}
		return isVisible(x, y, b, next)
	}

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		//fmt.Print("x:", x, " y:", y, " b:", string([]byte{b}))
		switch {
		// on the outside anyway
		case isExterior(x, y):
			//	fmt.Println(" - exterior")
			visible++
		// left
		case isVisible(x, y, b, left):
			//	fmt.Println(" - visible left")
			visible++
			// right
		case isVisible(x, y, b, right):
			//fmt.Println(" - visible right")
			visible++
			// up
		case isVisible(x, y, b, up):
			//fmt.Println(" - visible up")
			visible++
			// down
		case isVisible(x, y, b, down):
			//fmt.Println(" - visible down")
			visible++
		default:
			//fmt.Println(" - not visible")
		}
	})

	return fmt.Sprint(visible)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '0')

	var countVisible func(x, y int, b byte, next func(i, j int) (int, int)) int
	countVisible = func(x, y int, b byte, next func(i, j int) (int, int)) int {
		x, y = next(x, y)
		v, oob := g.At(x, y)
		if oob {
			return 0
		}
		if v >= b {
			// we can see this one, but it is the last one
			return 1
		}
		// we can see this one and can continue
		return 1 + countVisible(x, y, b, next)
	}

	best := 0

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		l := countVisible(x, y, b, left)
		r := countVisible(x, y, b, right)
		u := countVisible(x, y, b, up)
		d := countVisible(x, y, b, down)
		s := l * r * u * d
		if s > best {
			best = s
		}
	})

	return fmt.Sprint(best)
}
