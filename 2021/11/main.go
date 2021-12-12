package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 11, solve1, solve2)
}

var directions = [][2]int{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}
var flashpoint = byte('9') + 1

// Implement Solution to Problem 1
func solve1(input string) string {
	tick := makeTicker(input)
	flashes := 0
	for i := 0; i < 100; i++ {
		flashes += tick()
		//fmt.Println("After", i+1, "ticks we have", flashes, "flashes")
	}

	return fmt.Sprint(flashes)
}
func makeTicker(input string) func() int {
	var height, stride int
	grid := make([]byte, 0, 100)
	aoc.MapLines(input, func(line string) error {
		height++
		stride = len(line)
		grid = append(grid, []byte(line)...)
		return nil
	})

	adjacent := func(idx int) []int {
		a := make([]int, 0, 8)
		x, y := aoc.GridCoords(idx, stride)
		for _, d := range directions {
			ii := aoc.GridIndex(x+d[0], y+d[1], stride, height)
			if ii != -1 {
				a = append(a, ii)
			}
		}
		return a
	}

	tick := func() int {
		// first inc all entries
		// if exactly 9 then flash
		flashed := map[int]struct{}{}
		for i := 0; i < height*stride; i++ {
			grid[i] += 1
			if grid[i] == flashpoint {
				flashed[i] = struct{}{}
			}
		}
		// now we iterate those that flashed and
		// do the chain reaction.
		var chain func(int)
		chain = func(idx int) {
			for _, x := range adjacent(idx) {
				grid[x] += 1
				if grid[x] == flashpoint {
					chain(x)
				}
			}
		}
		for x := range flashed {
			chain(x)
		}

		// now set any 9 or over to 0
		count := 0
		for i := 0; i < height*stride; i++ {
			if grid[i] >= flashpoint {
				count++
				grid[i] = '0'
			}
		}
		return count
	}

	return tick
}

// Implement Solution to Problem 2
func solve2(input string) string {
	tick := makeTicker(input)
	for i := 0; ; i++ {
		if tick() == 100 {
			return fmt.Sprint(i + 1)
		}
	}
}
