package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 18, solve1, solve2)
}

const (
	On  = '#'
	Off = '.'
)

type GoL struct {
	Grid   []rune
	Stride int
	tmp    []rune
}

func (g *GoL) String() string {
	b := &strings.Builder{}
	for r := 0; r < g.Stride; r++ {
		b.WriteString(string(g.Grid[r*g.Stride : (r+1)*g.Stride]))
		b.WriteRune('\n')
	}
	fmt.Fprintf(b, "Stride:%d, On:%d\n", g.Stride, g.CountOn())
	return b.String()
}

func (g *GoL) Index(x, y int) int {
	if x < 0 || x >= g.Stride || y < 0 || y >= g.Stride {
		return -1
	}
	return x*g.Stride + y
}

func (g *GoL) GetNeighboursOn(idx int, fixed bool) int {
	x := idx / g.Stride
	y := idx % g.Stride

	if fixed && (x == 0 || x == g.Stride-1) && (y == 0 || y == g.Stride-1) {
		// it's a corner
		return 3 // 3 will always keep this on
	}

	sum := 0
	check := func(xx, yy int) {
		i := g.Index(xx, yy)
		if i >= 0 && g.Grid[i] == On {
			sum++
		}
	}
	check(x-1, y-1)
	check(x-1, y)
	check(x-1, y+1)
	check(x, y-1)
	// check(x, y) // don't need this one
	check(x, y+1)
	check(x+1, y-1)
	check(x+1, y)
	check(x+1, y+1)

	return sum
}

func (g *GoL) CountOn() int {
	sum := 0
	for _, r := range g.Grid {
		if r == On {
			sum++
		}
	}
	return sum
}

// update this one by one tick
func (g *GoL) Tick(fixed bool) {
	for i, r := range g.Grid {
		n := g.GetNeighboursOn(i, fixed)
		if r == On {
			if n == 2 || n == 3 {
				g.tmp[i] = On
			} else {
				g.tmp[i] = Off
			}
		} else {
			if n == 3 {
				g.tmp[i] = On
			} else {
				g.tmp[i] = Off
			}
		}
	}
	g.Grid, g.tmp = g.tmp, g.Grid
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return solve1a(input, 100)
}

func solve1a(input string, steps int) string {
	gol := &GoL{
		Grid:   make([]rune, 0, len(input)),
		Stride: 0,
	}
	aoc.MapLines(input, func(line string) error {
		if gol.Stride == 0 {
			gol.Stride = len(line)
		}
		gol.Grid = append(gol.Grid, []rune(line)...)
		return nil
	})
	gol.tmp = make([]rune, len(gol.Grid))

	// now tick
	//fmt.Printf("Initial State\n%v\n", gol)

	for x := 0; x < steps; x++ {
		gol.Tick(false)
		//	fmt.Printf("Tick %d\n%v\n", x+1, gol)
	}

	return fmt.Sprintf("%d", gol.CountOn())
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return solve2a(input, 100)
}

func solve2a(input string, steps int) string {
	gol := &GoL{
		Grid:   make([]rune, 0, len(input)),
		Stride: 0,
	}
	aoc.MapLines(input, func(line string) error {
		if gol.Stride == 0 {
			gol.Stride = len(line)
		}
		gol.Grid = append(gol.Grid, []rune(line)...)
		return nil
	})
	gol.tmp = make([]rune, len(gol.Grid))

	// now fix on the corners.
	gol.Grid[gol.Index(0, 0)] = On
	gol.Grid[gol.Index(0, gol.Stride-1)] = On
	gol.Grid[gol.Index(gol.Stride-1, 0)] = On
	gol.Grid[gol.Index(gol.Stride-1, gol.Stride-1)] = On

	// now tick
	//fmt.Printf("Initial State\n%v\n", gol)

	for x := 0; x < steps; x++ {
		gol.Tick(true)
		//	fmt.Printf("Tick %d\n%v\n", x+1, gol)
	}

	return fmt.Sprintf("%d", gol.CountOn())
}
