package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 11, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	instructions := strings.Split(input, ",")
	pos := Tile{0, 0}
	for _, i := range instructions {
		switch strings.TrimSpace(i) {
		case "n":
			pos = pos.Add(N)
		case "ne":
			pos = pos.Add(NE)
		case "se":
			pos = pos.Add(SE)
		case "s":
			pos = pos.Add(S)
		case "sw":
			pos = pos.Add(SW)
		case "nw":
			pos = pos.Add(NW)
		default:
			fmt.Printf("bad direction: %q\n", i)
			panic("bad direction")
		}
	}

	// now we have the final position, we need to find the distance to the origin.
	return fmt.Sprint(pos.DistanceTo(Tile{0, 0}))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	instructions := strings.Split(input, ",")
	max := 0
	origin := Tile{0, 0}
	pos := origin
	for _, i := range instructions {
		switch strings.TrimSpace(i) {
		case "n":
			pos = pos.Add(N)
		case "ne":
			pos = pos.Add(NE)
		case "se":
			pos = pos.Add(SE)
		case "s":
			pos = pos.Add(S)
		case "sw":
			pos = pos.Add(SW)
		case "nw":
			pos = pos.Add(NW)
		default:
			fmt.Printf("bad direction: %q\n", i)
			panic("bad direction")
		}
		d := pos.DistanceTo(origin)
		if d > max {
			max = d
		}
	}
	return fmt.Sprint(max)
}

// we need to model a hexagonal grid.
// the how to traverse it
// then how to find the shortest path between two tiles.
// we are going to use an axial coordinate system, then we can represent is with an V2

type Tile aoc.V2

func (t *Tile) Z() int {
	return -t.X - t.Y
}

func (t *Tile) Add(o Tile) Tile {
	return Tile{t.X + o.X, t.Y + o.Y}
}

// we have postive X as north, Y as north-east, and Z as south-east
var (
	N  = Tile{X: 1, Y: 0}
	NE = Tile{X: 0, Y: 1}
	SE = Tile{X: -1, Y: 1}
	S  = Tile{X: -1, Y: 0}
	SW = Tile{X: 0, Y: -1}
	NW = Tile{X: 1, Y: -1}
)

func (t *Tile) DistanceTo(o Tile) int {
	return (abs(t.X-o.X) + abs(t.Y-o.Y) + abs(t.Z()-o.Z())) / 2
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
