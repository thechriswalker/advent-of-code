package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 5, solve1, solve2)
}

type Line struct {
	x1, x2, y1, y2 int
}

// x,y -> num lines
type Grid map[[2]int]int

func (g Grid) AddStraightLine(l Line, includeDiagonal bool) {
	// the lines are either horizontal or vertical.
	if l.x1 == l.x2 {
		y1, y2 := l.y1, l.y2
		if y2 < y1 {
			y1, y2 = y2, y1
		}
		for i := y1; i <= y2; i++ {
			g[[2]int{l.x1, i}]++
		}
	} else if l.y1 == l.y2 {
		x1, x2 := l.x1, l.x2
		if x2 < x1 {
			x1, x2 = x2, x1
		}
		for i := x1; i <= x2; i++ {
			g[[2]int{i, l.y1}]++
		}
	} else if includeDiagonal {
		// work out which way to increment x and y each time.
		// they are 45 degrees, so the number of steps is always equal
		// horizontal and vertical
		dx := 1
		if l.x1 > l.x2 {
			dx = -1
		}
		dy := 1
		if l.y1 > l.y2 {
			dy = -1
		}
		d := l.x1 - l.x2
		if d < 0 {
			d *= -1
		}
		for i := 0; i <= d; i++ {
			g[[2]int{l.x1 + i*dx, l.y1 + i*dy}]++
		}
	}
}

func (g Grid) CountIntersections() int {
	sum := 0
	for _, n := range g {
		if n > 1 {
			sum++
		}
	}
	return sum
}

// Implement Solution to Problem 1
func solve1(input string) string {
	//8,0 -> 0,8
	grid := make(Grid)
	aoc.MapLines(input, func(line string) error {
		// we want horizontal or vetical lines only.
		l := Line{}
		fmt.Sscanf(line, "%d,%d -> %d,%d", &(l.x1), &(l.y1), &(l.x2), &(l.y2))
		grid.AddStraightLine(l, false)

		return nil
	})

	return fmt.Sprintf("%d", grid.CountIntersections())
}

// Implement Solution to Problem 2
func solve2(input string) string {
	//8,0 -> 0,8
	grid := make(Grid)
	aoc.MapLines(input, func(line string) error {
		// we want horizontal or vetical lines only.
		l := Line{}
		fmt.Sscanf(line, "%d,%d -> %d,%d", &(l.x1), &(l.y1), &(l.x2), &(l.y2))
		grid.AddStraightLine(l, true)

		return nil
	})

	return fmt.Sprintf("%d", grid.CountIntersections())
}
