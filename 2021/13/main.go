package main

import (
	"fmt"
	"strconv"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 13, solve1, solve2)
}

type Dot struct{ x, y int }

type Paper struct {
	xMax, xMin, yMax, yMin int
	Dots                   map[Dot]struct{}
}

func (p *Paper) Len() int {
	return len(p.Dots)
}

func (p *Paper) Print() {

	for y := p.yMin; y <= p.yMax; y++ {
		for x := p.xMin; x <= p.xMax; x++ {
			if _, ok := p.Dots[Dot{x: x, y: y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func (p *Paper) Set(dot Dot) {
	p.Dots[dot] = struct{}{}
	if p.xMax < dot.x {
		p.xMax = dot.x
	}
	if p.yMax < dot.y {
		p.yMax = dot.y
	}
	if p.xMin > dot.x {
		p.xMin = dot.x
	}
	if p.yMin > dot.y {
		p.yMin = dot.y
	}
}

// the fold is "along X/Y = value"
type Fold struct {
	X, Y int
}

func (p *Paper) Fold(f Fold) *Paper {
	// new paper with the fold
	// the transform is towards 0.
	// so if dot.x < fold.x
	next := &Paper{Dots: map[Dot]struct{}{}}

	var xform func(d Dot) Dot
	if f.X > 0 {
		//fmt.Printf("Fold is along x=%d\n", f.X)
		// this is a vertical fold, x values change
		xform = func(d Dot) Dot {
			if d.x < f.X {
				// this is on the left half, don't move.
				return d
			}
			// need to move the same distance Above the line as it was below

			dx := 2*f.X - d.x
			return Dot{x: dx, y: d.y}
		}
	} else {
		// horizontal fold
		//fmt.Printf("Fold is along y=%d\n", f.Y)
		xform = func(d Dot) Dot {
			if d.y < f.Y {
				// this is on the top half, don't move.
				//		fmt.Printf("Dot at (%d,%d) moves to (%d,%d)\n", d.x, d.y, d.x, d.y)
				return d
			}
			// need to move the same distance aboive as it was below
			dy := 2*f.Y - d.y
			//	fmt.Printf("Dot at (%d,%d) moves to (%d,%d)\n", d.x, d.y, d.x, dy)
			return Dot{x: d.x, y: dy}
		}
	}

	for dot := range p.Dots {
		next.Set(xform(dot))
	}
	return next
}

func parseInput(input string) (*Paper, []Fold) {
	p := &Paper{Dots: map[Dot]struct{}{}}
	f := []Fold{}
	aoc.MapLines(input, func(line string) error {
		switch {
		case line == "":
			return nil
		case line[0] == 'f':
			// fold along %c=%d
			fold := Fold{}
			if line[11] == 'x' {
				fold.X, _ = strconv.Atoi(line[13:])
			} else {
				fold.Y, _ = strconv.Atoi(line[13:])
			}
			f = append(f, fold)
		default:
			// assume a x,y pair.
			dot := Dot{}
			fmt.Sscanf(line, "%d,%d", &(dot.x), &(dot.y))
			p.Set(dot)

		}
		return nil
	})
	return p, f
}

// Implement Solution to Problem 1
func solve1(input string) string {
	p, f := parseInput(input)

	n := p.Fold(f[0])

	c := n.Len()

	return fmt.Sprint(c)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	p, f := parseInput(input)

	for i := range f {
		p = p.Fold(f[i])
	}
	fmt.Println()
	p.Print()

	return "^"
}
