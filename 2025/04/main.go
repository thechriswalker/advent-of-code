package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 4, solve1, solve2)
}

var eightDirs = []aoc.V2{
	{-1, -1},
	{0, -1},
	{1, -1},
	{-1, 0},
	{1, 0},
	{-1, 1},
	{0, 1},
	{1, 1},
}

// Implement Solution to Problem 1
func solve1(input string) string {

	g := aoc.CreateFixedByteGridFromString(input, '.')

	count := 0
	aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
		if b == '@' && canBeRemoved(g, v) {
			count++
		}
	})

	return fmt.Sprint(count)
}

func canBeRemoved(g aoc.ByteGrid, pos aoc.V2) bool {
	adj := 0
	for _, d := range eightDirs {
		b, _ := g.Atv(pos.Add(d))
		if b == '@' {
			adj++
			if adj > 3 {
				return false
			}
		}
	}
	return true
}

// Implement Solution to Problem 2
func solve2(input string) string {

	g := aoc.CreateFixedByteGridFromString(input, '.')
	sum := 0
	for {
		mark := []aoc.V2{}
		aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
			if b == '@' && canBeRemoved(g, v) {
				mark = append(mark, v)
			}
		})
		sum += len(mark)
		if len(mark) == 0 {
			break
		}
		for _, m := range mark {
			g.Setv(m, '.')
		}

	}

	return fmt.Sprint(sum)
}
