package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 6, solve1, solve2)
}

type Point struct {
	x, y int
}

type Direction Point

var (
	North = Direction{0, -1}
	South = Direction{0, 1}
	East  = Direction{1, 0}
	West  = Direction{-1, 0}
)

// Implement Solution to Problem 1
func solve1(input string) string {

	g := aoc.CreateFixedByteGridFromString(input, '.')

	var pos Point
	dir := North

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		if b == '^' {
			pos = Point{x, y}
		}
	})

	// record the points we have visited
	visited := make(map[Point]struct{})
	visited[pos] = struct{}{}

	for {
		// move in "direction" until we hit a wall, then turn right and continue.
		// if we hit OOB, we are done.

		b, oob := g.At(pos.x+dir.x, pos.y+dir.y)
		if oob {
			break
		}
		if b == '#' {
			// turn right.
			dir = dir.TurnRight()
		} else {
			// empty space.
			// we can move there and add it to our visited list.
			pos.x += dir.x
			pos.y += dir.y
			visited[pos] = struct{}{}
		}
	}

	return fmt.Sprint(len(visited))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '.')

	// this is where we could place a block to create a loop.
	var blockOptions = map[Point]struct{}{}
	var pos Point
	dir := North
	// we can't put an obstruction where we have already visited (or the guard would have walked into it the first time.)
	visited := make(map[Point]struct{})
	// find the start position.
	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		if b == '^' {
			pos = Point{x, y}
		}
	})
	visited[pos] = struct{}{}

	for {
		// move in "direction" until we hit a wall, then turn right and continue.
		// if we hit OOB, we are done.
		b, oob := g.At(pos.x+dir.x, pos.y+dir.y)
		if oob {
			break
		}
		switch b {
		case '#':
			// turn right.
			dir = dir.TurnRight()
		case '.':
			// empty space.
			// see if adding a block here (and turning instead) would lead to a loop.
			blockPos := Point{pos.x + dir.x, pos.y + dir.y}
			if _, ok := visited[blockPos]; !ok {
				// first time stepping here, we can test if it would create a loop.
				if checkLoop(g, blockPos, pos, dir) {
					blockOptions[blockPos] = struct{}{}
				}
			}
			// we can move there and add it to our visited list.
			pos.x += dir.x
			pos.y += dir.y
			visited[pos] = struct{}{}
		default:
			// this will be the '^', treat as a space, but not a block option
			pos.x += dir.x
			pos.y += dir.y
			visited[pos] = struct{}{}
		}
	}

	// 1798 first guess - too high! then I realised that we couldn't place an obstacle in a visited square.
	// 1720 second guess - still too high! ?? lol. I tried 1719 in case it was an off-by-one error. it was.
	// but what is the problem!??? It was that I was attempting to place a block in the starting position, when travelling in
	// a different direction.

	return fmt.Sprint(len(blockOptions))
}

func (d Direction) TurnRight() Direction {
	switch d {
	case North:
		return East
	case East:
		return South
	case South:
		return West
	case West:
		return North
	}
	panic("unreachable")
}

var debug = false

func init() {
	if os.Getenv("AOC_DEBUG") != "" {
		debug = true
	}
}

// basically set a block at the next position and then use a cache to see if we find a loop.

func checkLoop(g aoc.ByteGrid, blockPos, currentPos Point, currentDir Direction) (loopFound bool) {
	pos := currentPos
	dir := currentDir
	visited := map[[2]Point]struct{}{}
	visited[[2]Point{pos, Point(dir)}] = struct{}{}
	if _, oob := g.At(blockPos.x, blockPos.y); oob {
		return false
	}

	if debug {
		defer func() {
			if loopFound {
				g2 := g.Clone() // so we don't muck up the original
				g2.Set(blockPos.x, blockPos.y, 'O')
				for p := range visited {
					switch Direction(p[1]) {
					case North, South:
						g2.Set(p[0].x, p[0].y, '|')
					case East, West:
						g2.Set(p[0].x, p[0].y, '-')
					}
				}
				aoc.PrintByteGridFunc(g2, func(x, y int, b byte) aoc.Color {
					switch b {
					case '-', '|', '^':
						return aoc.BoldCyan
					case '#':
						return aoc.BoldWhite
					case 'O':
						return aoc.BoldRed
					default:
						return aoc.NoColor
					}
				})
				fmt.Println()
				bufio.NewReader(os.Stdin).ReadBytes('\n')
			}
		}()
	}
	var nextPos Point
	for {
		nextPos.x = pos.x + dir.x
		nextPos.y = pos.y + dir.y
		b, oob := g.At(nextPos.x, nextPos.y)
		if oob {
			return
		}
		if b == '#' || nextPos == blockPos {
			// turn right.
			dir = dir.TurnRight()
		} else {
			// empty space.
			// we can move there
			pos.x = nextPos.x
			pos.y = nextPos.y
		}
		c := [2]Point{pos, Point(dir)}
		//_, alreadyVisited := visited[c]
		//fmt.Println(c, alreadyVisited)
		if _, ok := visited[c]; ok {
			loopFound = true
			return
		}
		visited[c] = struct{}{}
	}
}
