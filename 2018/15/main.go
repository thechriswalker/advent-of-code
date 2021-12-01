package main

import (
	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2018, 15, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return "<unsolved>"
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return "<unsolved>"
}

const (
	ELF    = 'E'
	GOBLIN = 'G'

	WALL  = '#'
	EMPTY = '.'
)

type Tile struct {
	Kind   byte
	AP, HP int
	Index int
}

func (t *Tile) IsEmpty() bool {
	return t.Kind == EMPTY
}

func (t *Tile) CanAttack(other *Tile) bool {
	if t.Kind == EMPTY || t.Kind == WALL || other.Kind == EMPTY || other.Kind == WALL {
		return false
	}
	// the only kinds left are elves and goblins
	return t.Kind != other.Kind
}

type Cavern struct {
	Width  int
	Height int
	Layout []*Tile
}

var (
	WallTile = &Tile{Kind: WALL}
	OpenTile = &Tile{Kind: EMPTY}
)

func (c *Cavern) TileAt(x, y int) *Tile {
	if x < 0 || x > c.Width-1 || y < 0 || y > c.Height-1 {
		return WallTile // pretend the walls extend infinitely
	}
	idx := x + y*c.Width
	return c.Layout[idx]
}

func (c *Cavern) Move(src, dst int) {
	srcTile := c.Layout[src]
	dstTile := c.Layout[dst]
	if dstTile.Kind != EMPTY {
		panic("Cannot move into non-empty space")
	}
	switch srcTile.Kind {
	case ELF, GOBLIN:
		// ok
	default:
		panic("Only elves and goblins can move!")
	}
	c.Layout[src], c.Layout[dst] = c.Layout[dst], c.Layout[src]
	c.Layout[src].Index = src
	c.Layout[dst].Index = dst
}

// storing the tiles in this order means they are in reading order.
// so we can compare indexes

// process one tick of the game, return true if we bailed early (i.e. game over)
func (c *Cavern) Tick() bool {
	// tiles are in reading order.
	// but we need to get just the units in starting position reading order.
	order := []*Tile{}
	var elves, goblins int
	for _, t := range c.Layout {
		switch t.Kind {
		case ELF:
			elves++
		case GOBLIN:
			goblins++
		default:
			continue
		}
		order = append(order, t)
	}

	for _, t := range order {
		// can we attack
		if t.
	}
}
