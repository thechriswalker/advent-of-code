package main

import (
	"fmt"
	"io"
	"os"

	"../../aoc"
	"../intcode"
)

func main() {
	aoc.Run(2019, 11, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	pg := intcode.New(input)
	robot := &Robot{pg: pg, hull: map[[2]int]int64{}}
	robot.Run()
	return fmt.Sprintf("%d", robot.NumberOfPaintedPanels())
}

// Implement Solution to Problem 2
func solve2(input string) string {
	pg := intcode.New(input)
	// set the starting panel white
	robot := &Robot{pg: pg, hull: map[[2]int]int64{
		[2]int{0, 0}: 1,
	}}
	robot.Run()
	fmt.Println()
	robot.DrawHull(os.Stdout)
	fmt.Println()
	return "<see picture>"

}

const (
	black int64 = 0
	white       = 1
	right       = 0
	left        = 1

	north = 0 // start facing north
	east  = 1
	south = 2
	west  = 3
)

func turn(facing, turn int64) int64 {
	switch turn {
	case left:
		// add 4 (number of directions)
		// substract 1
		// mod 4
		return (facing + 4 - 1) % 4
	case right:
		// just add an mod
		return (facing + 1) % 4
	default:
		panic("invalid direction to turn")
	}
}

func move(from [2]int, facing int64) [2]int {
	switch facing {
	case north:
		// add one to X
		from[0]++
	case east:
		// add one to Y
		from[1]++
	case south:
		// sub one from X
		from[0]--
	case west:
		// sub one from Y
		from[1]--
	}
	return from
}

type Robot struct {
	// the brain
	pg *intcode.Program
	// the hull, represented as a sparse map
	// key is x,y co-ordinates
	// value is color, zero value will be black
	// and `len(r.hull)` will be number of squares
	// painted (black or white)
	hull map[[2]int]int64

	position [2]int // current x,y
	facing   int64

	xmin, xmax, ymin, ymax int
}

func (r *Robot) Run() {
	r.pg.RunAsync()
	input := func() int64 {
		return r.hull[r.position]
	}
	for {
		// the loop is send color at position
		// or halt if program halts
		select {
		case r.pg.Input <- input:
		case <-r.pg.Halted:
			return
		}
		// we gave an input, read the color to paint
		r.hull[r.position] = <-r.pg.Output
		// read the way to turn
		r.facing = turn(r.facing, <-r.pg.Output)
		// and move
		r.position = move(r.position, r.facing)
		r.UpdateBounds(r.position)
		// continue
	}
}

func (r *Robot) NumberOfPaintedPanels() int {
	return len(r.hull)
}

func (r *Robot) UpdateBounds(pos [2]int) {
	x, y := pos[0], pos[1]
	if x < r.xmin {
		r.xmin = x
	}
	if x > r.xmax {
		r.xmax = x
	}
	if y < r.ymin {
		r.ymin = y
	}
	if y > r.ymax {
		r.ymax = y
	}
}

func (r *Robot) DrawHull(w io.Writer) {
	// first we get the bounds.
	// then we iterate over the bounds an paint using the map.
	// it draws the letter upside down! so we iterate backwards.
	for x := r.xmax + 1; x >= r.xmin-1; x-- {
		for y := r.ymax + 1; y >= r.ymin-1; y-- {
			color, ok := r.hull[[2]int{x, y}]
			if !ok {
				w.Write([]byte{' '})
			} else {
				switch color {
				case black:
					w.Write([]byte("\x1b[1;90m#\x1b[0m"))
				case white:
					w.Write([]byte("\x1b[1;97m#\x1b[0m"))
				default:
					panic("unknown color")
				}
			}
		}
		// add a newline
		w.Write([]byte{'\n'})
	}
}
