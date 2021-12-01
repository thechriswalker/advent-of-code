package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 18, solve1, solve2)
}

// how many rows to produce
var Rows = 40

// Implement Solution to Problem 1
func solve1(input string) string {
	prev := parseInput(input)
	next := make(Tiles, len(prev))
	safe := prev.Safe()
	for rows := 1; rows < Rows; rows++ {
		prev.Next(next)
		safe += next.Safe()
		prev, next = next, prev // swap them
	}

	return fmt.Sprintf("%d", safe)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	Rows = 400000
	return solve1(input)
}

func parseInput(input string) Tiles {
	tiles := make(Tiles, 0, len(input))
	for _, c := range input {
		switch c {
		case '.':
			// safe
			tiles = append(tiles, 0)
		case '^':
			//trap
			tiles = append(tiles, 1)
		}
	}
	return tiles
}

type Tiles []uint8

// to avoid allocation, we write into the given tiles
func (prev Tiles) Next(next Tiles) {
	var l, r uint8
	for x := 0; x < len(prev); x++ {
		switch x {
		case 0:
			// assume left is safe
			l, r = 0, prev[x+1]
		case len(prev) - 1:
			// assume right is safe
			l, r = prev[x-1], 0
		default:
			l, r = prev[x-1], prev[x+1]
		}
		next[x] = l ^ r
	}
}

func (t Tiles) Traps() int {
	sum := 0
	for _, s := range t {
		sum += int(s)
	}
	return sum
}

func (t Tiles) Safe() int {
	return len(t) - t.Traps()
}
