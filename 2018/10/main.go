package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2018, 10, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	prev := parseInput(input)
	next := make(Stars, len(prev))
	nextBounds := prev.Step(next) // step 1
	//fmt.Println(next.Print(nextBounds))
	prevBounds := next.Step(prev) // step 2
	//fmt.Println(prev.Print(prevBounds))
	prevBounds, nextBounds = nextBounds, prevBounds
	prev, next = next, prev

	steps := 1
	for nextBounds.Size() < prevBounds.Size() {
		// switch the buffers
		prevBounds, nextBounds = nextBounds, prevBounds
		prev, next = next, prev
		// step until the bounds start to grow
		nextBounds = prev.Step(next)
		steps++
	}
	fmt.Printf("\n%s\n", prev.Print(prevBounds))
	return fmt.Sprintf("%d", steps)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// turns out my code for problem 1 **was** the answer to problem 2...
	return solve1(input)
}

func parseInput(input string) Stars {
	rd := strings.NewReader(input)
	stars := Stars{}
	var x, y, dx, dy int
	var err error
	for {
		_, err = fmt.Fscanf(rd, "position=<%d,%d> velocity=<%d,%d>\n", &x, &y, &dx, &dy)
		if err != nil {
			break
		}
		stars = append(stars, [4]int{x, y, dx, dy})
	}
	return stars
}

// a single slice of [4]int arrays
// which is x,y,dx,dy
type Stars [][4]int

const (
	X  = 0
	Y  = 1
	DX = 2
	DY = 3
)

type Bounds [4]int

const (
	T = 0
	L = 1
	B = 2
	R = 3
)

func (b Bounds) Size() int {
	return (b[B] - b[T]) * (b[R] - b[L])
}

func (ss Stars) Step(next [][4]int) Bounds {
	bounds := [4]int{}
	// assume we won't see anything bigger/smaller than 1e6

	var x, y int
	for i, s := range ss {
		//new position
		x, y = s[X]+s[DX], s[Y]+s[DY]
		next[i] = [4]int{x, y, s[DX], s[DY]}
		if i == 0 {
			bounds[T], bounds[L], bounds[B], bounds[R] = y, x, y, x
			continue
		}
		if x < bounds[L] {
			bounds[L] = x
		}
		if x > bounds[R] {
			bounds[R] = x
		}
		if y < bounds[T] {
			bounds[T] = y
		}
		if y > bounds[B] {
			bounds[B] = y
		}
	}
	return bounds
}

func (ss Stars) Print(bounds Bounds) string {
	// need to put the stars into a map for printing
	starMap := map[[2]int]struct{}{}
	for _, s := range ss {
		starMap[[2]int{s[X], s[Y]}] = struct{}{}
	}
	s := strings.Builder{}
	for j := bounds[T] - 1; j <= bounds[B]+1; j++ {
		for i := bounds[L] - 1; i <= bounds[R]+1; i++ {
			// is it a star?
			if _, ok := starMap[[2]int{i, j}]; ok {
				s.WriteByte('#')
			} else {
				s.WriteByte('.')
			}
		}
		s.WriteByte('\n')
	}
	return s.String()
}
