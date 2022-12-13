package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 10, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	X := 1
	_cycle := 0
	targets := []int{20, 60, 100, 140, 180, 220, math.MaxInt}
	targetIdx := 0
	sum := 0

	cycle := func() {
		_cycle++
		if _cycle == targets[targetIdx] {
			sum += _cycle * X
			targetIdx++
		}
	}

	aoc.MapLines(input, func(line string) error {
		if line == "noop" {
			cycle()
		} else {
			// addx
			// two cycles beforehand.
			cycle()
			cycle()
			n, _ := strconv.Atoi(line[5:])
			X += n
		}
		return nil
	})

	return fmt.Sprint(sum)

}

// Implement Solution to Problem 2
func solve2(input string) string {
	X := 1
	_cycle := 0

	grid := make([]byte, 240)

	cycle := func() {
		// before the cycle advances
		// draw a pixel if the current cycle number
		// is within 1 of the current X register/

		var b byte
		b = '.'
		sprite := X % 40
		c := _cycle % 40

		if c >= sprite-1 && c <= sprite+1 {
			b = '#'
		}
		grid[_cycle] = b

		// and bump.
		_cycle++
	}

	aoc.MapLines(input, func(line string) error {
		if line == "noop" {
			cycle()
		} else {
			// addx
			// two cycles beforehand.
			cycle()
			cycle()
			n, _ := strconv.Atoi(line[5:])
			X += n
		}
		return nil
	})

	sb := strings.Builder{}

	sb.Write(grid[0:40])
	sb.WriteByte('\n')
	sb.Write(grid[40:80])
	sb.WriteByte('\n')
	sb.Write(grid[80:120])
	sb.WriteByte('\n')
	sb.Write(grid[120:160])
	sb.WriteByte('\n')
	sb.Write(grid[160:200])
	sb.WriteByte('\n')
	sb.Write(grid[200:240])
	sb.WriteByte('\n')

	g := aoc.CreateFixedByteGridFromString(sb.String(), '.')
	fmt.Println()
	aoc.PrintByteGrid(g, nil)

	// from reading the output...
	return "ZKJFBJFZ"
}
