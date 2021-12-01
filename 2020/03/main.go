package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 3, solve1, solve2)
}

const (
	EMPTY = 0
	TREE  = 1
)

type grid struct {
	// first slice is 0 at the top, second is zero at the left
	points [][]uint8 // EMPTY or TREE
	height int
}

// get the thing at the point in the grid
func (g *grid) At(x, y int) uint8 {
	if y < 0 || y >= g.height {
		// assume empty!
		return EMPTY
	}
	row := g.points[y]
	dx := x % len(row)
	return row[dx]
}

func makeGrid(in string) *grid {
	entries := strings.Split(in, "\n")
	stack := make([][]uint8, 0, len(entries))
	for _, entry := range entries {
		if entry == "" {
			continue
		}
		row := make([]uint8, len(entry))
		for n, c := range entry {
			switch c {
			case '.':
				row[n] = EMPTY
			case '#':
				row[n] = TREE
			default:
				panic(fmt.Sprintf("Invalid Character in grid: %c", c))
			}
		}
		stack = append(stack, row)
	}

	return &grid{
		points: stack,
		height: len(stack),
	}

}

func (g *grid) CountTreesOnPath(dx, dy int) int {
	x, y, c := 0, 0, 0
	for y < g.height {
		if g.At(x, y) == TREE {
			c++
		}
		x += dx
		y += dy
	}
	return c
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := makeGrid(input)
	// starting at 0,0
	// move right 3 (x+=3)
	// move down 1 (y+=1)
	return fmt.Sprintf("%d", g.CountTreesOnPath(3, 1))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := makeGrid(input)

	c1 := g.CountTreesOnPath(1, 1)
	c2 := g.CountTreesOnPath(3, 1)
	c3 := g.CountTreesOnPath(5, 1)
	c4 := g.CountTreesOnPath(7, 1)
	c5 := g.CountTreesOnPath(1, 2)

	return fmt.Sprintf("%d", c1*c2*c3*c4*c5)
}
