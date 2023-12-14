package main

import (
	"fmt"
	"slices"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 10, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	_, s := parseGrid(input)

	// find the length of the loop
	l := 1
	prev, curr := s.Follow(nil)
	for {
		prev, curr = curr.Follow(prev)
		l++
		if curr == s {
			break
		}
	}

	d := l/2 + l%2

	return fmt.Sprint(d)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g, s := parseGrid(input)

	// to find the interior we will iterate the loop
	// and get all the pipe that is part of it.
	// then we iterate the grid. Note that we iterate
	// all interior pieces will have an ODD number of
	// pipe around them, so we just keep track of that
	//
	// Took me ages to realise that we can ignore the `-`
	// and that `L7` is 1 boundary
	// but `LJ` is 2 (or 0, an even number)
	// same with `F7` (2 boundaries) vs. `FJ` => 1
	loop := map[[2]int]*Node{
		s.pos: s,
	}

	prev, curr := s.Follow(nil)
	loop[curr.pos] = curr
	for {
		prev, curr = curr.Follow(prev)
		if curr != s {
			loop[curr.pos] = curr
		} else {
			break
		}
	}
	inside := 0
	row := 0
	boundaryCount := 0
	boundaryOpener := (*Node)(nil)

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		if row != y {
			row = y
			boundaryCount = 0
			boundaryOpener = nil
		}

		node, isPartOfLoop := loop[[2]int{x, y}]
		// is this a boundary cross? we are travelling west->east
		// a boundary is a pipe, or if the opener was a L then
		if isPartOfLoop {
			// was this an opening or closing of the boundary.
			switch node.Tile() {
			case '|':
				boundaryCount += 1 // definitely a boundary cross.
			case 'L', 'F': // start of a boundary
				boundaryOpener = node
			case '7': // end of boundary, the count
				// depends on the orientation.
				if boundaryOpener.Tile() == 'L' {
					boundaryCount++
				}
				boundaryOpener = nil
			case 'J':
				if boundaryOpener.Tile() == 'F' {
					boundaryCount++
				}
				boundaryOpener = nil
			}
		} else {
			if boundaryCount%2 == 1 {
				inside++
				g.Set(x, y, 'I')
			}
		}
	})

	// aoc.PrintByteGrid(g, nil)

	return fmt.Sprint(inside)
}

type Node struct {
	tile       byte
	pos        [2]int
	n, s, e, w *Node // nodes in other directions easier that just using 2
}

func (n *Node) Tile() byte {
	if n.tile == 'S' {
		// work out which one!
		north := n.n != nil
		south := n.s != nil
		east := n.e != nil
		west := n.w != nil

		switch {
		case north && south:
			return '|'
		case east && west:
			return '-'
		case north && east:
			return 'L'
		case north && west:
			return 'J'
		case south && east:
			return 'F'
		case south && west:
			return '7'
		}
	}
	return n.tile
}

func (curr *Node) Follow(prev *Node) (*Node, *Node) {
	// new curr/prev
	if curr.n != nil && curr.n != prev {
		return curr, curr.n
	}
	if curr.s != nil && curr.s != prev {
		return curr, curr.s
	}
	if curr.e != nil && curr.e != prev {
		return curr, curr.e
	}
	if curr.w != nil && curr.w != prev {
		return curr, curr.w
	}
	panic("no connections!")
}

// directions certain things can connect
var (
	north = []byte("|F7")
	south = []byte("|LJ")
	east  = []byte("-J7")
	west  = []byte("-FL")
)

func parseGrid(input string) (g aoc.ByteGrid, start *Node) {
	g = aoc.CreateFixedByteGridFromString(input, '.')

	// now we have to iterate the grid and build all the nodes.
	// the different type of nodes can connect in different ways.
	// we might as well join all possible joins, then our start node
	// should contain the loop we want.

	// holds nodes we have already parsed.
	cache := map[[2]int]*Node{}

	createOrRetrieve := func(x, y int, b byte) *Node {
		k := [2]int{x, y}
		n, ok := cache[k]
		if !ok {
			n = &Node{tile: b, pos: k}
			cache[k] = n
		}
		return n
	}

	connectIfPossible := func(x, y int, connectors []byte) *Node {
		b, _ := g.At(x, y)
		if b == 'S' || slices.Contains(connectors, b) {
			// yes.
			return createOrRetrieve(x, y, b)
		}
		return nil
	}

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		// create the node, or retrieve from cache
		curr := createOrRetrieve(x, y, b)

		// find its connections.
		switch b {
		case '|': // n,s
			curr.n = connectIfPossible(x, y-1, north)
			curr.s = connectIfPossible(x, y+1, south)
		case '-': // e,w
			curr.e = connectIfPossible(x+1, y, east)
			curr.w = connectIfPossible(x-1, y, west)
		case 'L': // n, e
			curr.n = connectIfPossible(x, y-1, north)
			curr.e = connectIfPossible(x+1, y, east)
		case 'J': // n, w
			curr.w = connectIfPossible(x-1, y, west)
			curr.n = connectIfPossible(x, y-1, north)
		case '7': // w, s
			curr.w = connectIfPossible(x-1, y, west)
			curr.s = connectIfPossible(x, y+1, south)
		case 'F': // e, s
			curr.s = connectIfPossible(x, y+1, south)
			curr.e = connectIfPossible(x+1, y, east)
		case '.':
			// nothing
		case 'S':
			// could be anywhere, but we have assurance that there will only be 2!
			curr.n = connectIfPossible(x, y-1, north)
			curr.e = connectIfPossible(x+1, y, east)
			curr.s = connectIfPossible(x, y+1, south)
			curr.w = connectIfPossible(x-1, y, west)
			start = curr
		}
	})

	return
}
