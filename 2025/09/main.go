package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 9, solve1, solve2)
}

func abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func area(a, b aoc.V2) int {
	w := 1 + abs(a.X-b.X)
	h := 1 + abs(a.Y-b.Y)
	// aoc.Debug("rectangle", a, b)
	// aoc.Debug("area = ", w, "*", h)
	return w * h
}

// Implement Solution to Problem 1
func solve1(input string) string {
	list := []aoc.V2{}
	aoc.MapLines(input, func(line string) error {
		v := aoc.V2{}
		fmt.Sscanf(line, "%d,%d", &v.X, &v.Y)
		list = append(list, v)
		return nil
	})

	largest := 0
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			a := area(list[i], list[j])
			if a > largest {
				largest = a
			}
		}
	}

	return fmt.Sprint(largest)
}

// Implement Solution to Problem 2
// the red tiles form a boundary and we need to find a solution
// inside the boundary.
// the boundary lines as between the pairs of
// tiles in the input.
// how do we tell if a rectangle is "contained"
// inside the polygon. that is the key here.
func solve2(input string) string {
	poly := []aoc.V2{}

	aoc.MapLines(input, func(line string) error {
		v := aoc.V2{}
		fmt.Sscanf(line, "%d,%d", &v.X, &v.Y)
		poly = append(poly, v)

		return nil
	})

	edges := makeEdges(poly)
	rectangles := makeRectangles(poly)

	// we need to check all rectangles
	// for each rectangle, is it contained.
	// is so what's the area.
	largest := 0

	for _, r := range rectangles {
		a := r.Area()
		if a < largest {
			continue
		}
		// check containment.
		// if any point is inside, then do any of the edges stick into the rectanges?
		if r.Inside(edges) {
			largest = a
		}
	}

	return fmt.Sprint(largest)
}

func makeEdges(p []aoc.V2) []Rect {
	e := make([]Rect, 0, len(p))
	for i := 0; i < len(p); i++ {
		e = append(e, NewRect(p[i], p[(i+1)%len(p)]))
	}
	return e
}

func makeRectangles(p []aoc.V2) []Rect {
	e := []Rect{}
	for i := 0; i < len(p)-1; i++ {
		for j := i + 1; j < len(p); j++ {
			e = append(e, NewRect(p[i], p[j]))
		}
	}
	return e
}

type Rect struct {
	tl, br aoc.V2
}

// normalise the rectangle
func NewRect(a, b aoc.V2) Rect {
	if a.X > b.X {
		a.X, b.X = b.X, a.X
	}
	if a.Y > b.Y {
		a.Y, b.Y = b.Y, a.Y
	}
	return Rect{tl: a, br: b}
}

func (r Rect) Area() int {
	return (1 + r.br.X - r.tl.X) * (1 + r.br.Y - r.tl.Y)
}

// returns true is the rect is contained in the edges
func (r Rect) Inside(edges []Rect) bool {
	// other wise do any edges intersect the rectange?
	for _, e := range edges {
		if r.Intersects(e) {
			return false
		}
	}
	return true
}

func (r Rect) Intersects(e Rect) bool {
	// these lines are horizontal or vertical
	horizontal := e.tl.Y == e.br.Y
	if horizontal {
		// the line would stick in from the X axis.
		if e.tl.Y <= r.tl.Y || e.tl.Y >= r.br.Y {
			// can't intersect
			return false
		}
		// intersects if X in range
		if max(e.tl.X, r.tl.X) < min(e.br.X, r.br.X) {
			// intersects
			return true
		}
		return false
	}
	// vertical
	// the line would stick in from the Y axis.
	// so does this line point at the rectangle at all?
	if e.tl.X <= r.tl.X || e.tl.X >= r.br.X {
		// nope
		return false
	}
	if max(e.tl.Y, r.tl.Y) < min(e.br.Y, r.br.Y) {
		// intersects
		return true
	}
	return false
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b

}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
