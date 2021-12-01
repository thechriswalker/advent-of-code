package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2018, 11, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// in a 300x300 grid find the
	// top left of a 3x3 square with
	// the highest power level.
	// lets just do it the easy way and see how fast that is...
	// just the top left's
	var g Grid
	fmt.Sscanf(input, "%d", &g)
	maxPower := 0
	maxPowerXY := [2]int{}
	for x := 1; x <= 298; x++ {
		for y := 1; y <= 298; y++ {
			// sum the nine squares.
			power := g.PowerLevelAt(x, y)
			power += g.PowerLevelAt(x+1, y)
			power += g.PowerLevelAt(x+2, y)
			power += g.PowerLevelAt(x, y+1)
			power += g.PowerLevelAt(x+1, y+1)
			power += g.PowerLevelAt(x+2, y+1)
			power += g.PowerLevelAt(x, y+2)
			power += g.PowerLevelAt(x+1, y+2)
			power += g.PowerLevelAt(x+2, y+2)

			if power > maxPower {
				maxPower = power
				maxPowerXY = [2]int{x, y}
			}
		}
	}

	return fmt.Sprintf("%d,%d", maxPowerXY[0], maxPowerXY[1])
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// arbitrary sized square
	var g Grid
	fmt.Sscanf(input, "%d", &g)
	maxPower := 0
	maxPowerXYS := [3]int{}
	for x := 1; x <= 300; x++ {
		for y := 1; y <= 300; y++ {
			// what are the square sizes? the min of 1+300-x and 1+300-y
			dx := 1 + 300 - x
			dy := 1 + 300 - y
			size := dx
			if dy < dx {
				size = dy
			}
			runningTotal := 0
			for s := 0; s < size; s++ {
				// sum the squares. at least we can use the previous values
				// and add one row/col at a time/ corner square first
				runningTotal += g.PowerLevelAt(x+s, y+s)
				// now add the row/col
				for i := 0; i < s; i++ {
					runningTotal += g.PowerLevelAt(x+s, y+i)
					runningTotal += g.PowerLevelAt(x+i, y+s)
				}
				if runningTotal > maxPower {
					maxPower = runningTotal
					maxPowerXYS = [3]int{x, y, s + 1}
				}
			}
		}
	}
	return fmt.Sprintf("%d,%d,%d", maxPowerXYS[0], maxPowerXYS[1], maxPowerXYS[2])
}

type Grid int

func (g Grid) PowerLevelAt(x, y int) int {
	// x,y are 1-based
	//rack = x+10, (rack*y + g) * rack
	power := ((x+10)*y + int(g)) * (x + 10)
	// keep only the hundred digit
	power = (power % 1000) / 100
	// subtract 5
	return power - 5
}
