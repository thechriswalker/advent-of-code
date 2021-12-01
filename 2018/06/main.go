package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2018, 6, solve1, solve2)
}

// bigger than any coordinate
var BIG = 10000

// how much wider than the actual bounds to return
var PAD = 1

// the limit for problem 2 (to change in the test)
var LIMIT = 10000

// Implement Solution to Problem 1
func solve1(input string) string {
	grid := parseInput(input)
	// area of each points closest
	counts := map[Point]int{}
	// the set of points that has their closest on the boundary
	// they will have infinite areas, so we discount them
	inifinites := map[Point]struct{}{}

	// get bounds
	tl, br := grid.Bounds()

	// the image is awesome!
	fmt.Println(grid.Points)
	fmt.Println(grid)

	var p Point
	var isInfinite bool
	var index int
	for x := tl[0]; x <= br[0]; x++ {
		for y := tl[1]; y <= br[1]; y++ {
			p, index = grid.ClosestPointTo(Point{x, y})
			if index == -1 {
				continue
			}
			isInfinite = x == tl[0] || x == br[0] || y == tl[1] || y == br[1]
			if isInfinite {
				inifinites[p] = struct{}{}
			} else {
				counts[p] = counts[p] + 1
			}
		}
	}

	// ok now we find the max non-inifinite count.
	max := 0
	for p, c := range counts {
		if _, isInfinite = inifinites[p]; isInfinite {
			continue
		}
		if max < c {
			max = c
		}
	}

	return fmt.Sprintf("%d", max)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	grid := parseInput(input)
	tl, br := grid.Bounds()

	// the image is awesome!
	// fmt.Println(grid.Points)
	// fmt.Println(grid.PrintLimitArea())

	var inArea bool
	sum := 0
	for x := tl[0]; x <= br[0]; x++ {
		for y := tl[1]; y <= br[1]; y++ {
			if inArea, _ = grid.IsInLimitArea(Point{x, y}); inArea {
				sum++
			}
		}
	}
	return fmt.Sprintf("%d", sum)
}

func parseInput(input string) *Grid {
	rd := strings.NewReader(strings.TrimSpace(input))
	grid := &Grid{
		Points: []Point{},
	}
	var x, y int
	var err error
	for {
		_, err = fmt.Fscanf(rd, "%d, %d\n", &x, &y)
		if err != nil {
			break
		}
		grid.Points = append(grid.Points, Point{x, y})
	}
	return grid
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

type Point [2]int

func (p Point) ManhattanDistance(q Point) int {
	return abs(p[0]-q[0]) + abs(p[1]-q[1])
}

type Grid struct {
	Points []Point
}

// finds the bounds of the points in the grid, we don't need
// to look further than this.
func (g *Grid) Bounds() (topleft, bottonright Point) {
	min := Point{BIG, BIG}
	max := Point{-1 * BIG, -1 * BIG}
	for _, p := range g.Points {
		if p[0] < min[0] {
			min[0] = p[0]
		}
		if p[1] < min[1] {
			min[1] = p[1]
		}
		if p[0] > max[0] {
			max[0] = p[0]
		}
		if p[1] > max[1] {
			max[1] = p[1]
		}
	}
	// add some padding
	min[0] -= PAD
	min[1] -= PAD
	max[0] += PAD
	max[1] += PAD

	return min, max
}

// finds the closest point, and it's index.
// if a dupe then the index is -1
func (g *Grid) ClosestPointTo(q Point) (Point, int) {
	min := BIG
	closest := Point{}
	index := -1
	for i, p := range g.Points {
		if p == q {
			return p, i
		}
		d := p.ManhattanDistance(q)
		if d < min {
			closest = p
			min = d
			index = i
		} else if d == min {
			index = -1
		}
	}
	return closest, index
}

func (g *Grid) String() string {
	// find bounds, add 1 in each direction and
	// generate the picture

	tl, br := g.Bounds()

	width := br[1] - tl[1]
	height := br[0] - tl[0]

	s := &strings.Builder{}
	var index int
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			_, index = g.ClosestPointTo(Point{tl[0] + x, tl[1] + y})
			switch {
			case index == -1:
				s.WriteByte('.')
			default:
				s.WriteByte(chars[index])
			}
		}
		s.WriteByte('\n')
	}
	return s.String()
}

func (g *Grid) PrintLimitArea() string {
	// find bounds, add 1 in each direction and
	// generate the picture

	tl, br := g.Bounds()

	width := br[1] - tl[1]
	height := br[0] - tl[0]

	s := &strings.Builder{}
	var index int
	var inLimitArea bool
	for y := 0; y <= height; y++ {
		for x := 0; x <= width; x++ {
			inLimitArea, index = g.IsInLimitArea(Point{tl[0] + x, tl[1] + y})
			if index != -1 {
				s.WriteByte(chars[index])
			} else if inLimitArea {
				s.WriteByte('#')
			} else {
				s.WriteByte('.')
			}
		}
		s.WriteByte('\n')
	}
	return s.String()
}

// returns bool isInArea and index if this is a point, the index, else -1
func (g *Grid) IsInLimitArea(q Point) (bool, int) {
	sum := 0
	index := -1
	for i, p := range g.Points {
		sum += p.ManhattanDistance(q)
		if p == q {
			index = i
		}
	}
	return sum < LIMIT, index
}

var chars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqstruvwxyz0123456789")
