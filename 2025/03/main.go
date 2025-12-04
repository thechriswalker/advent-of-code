package main

import (
	"fmt"
	"math"
	"slices"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 3, solve1, solve2)
}

type Cell struct {
	jolts int64
	idx   int
}

func sortCells(c []Cell) {
	slices.SortFunc(c, func(a, b Cell) int {
		if a.jolts == b.jolts {
			return a.idx - b.idx
		}
		if a.jolts < b.jolts {
			return 1
		} else if b.jolts < a.jolts {
			return -1
		}
		return 0
	})
}

func findLargest(c []Cell, idx int) Cell {
	for _, c := range c {
		if c.idx > idx {
			return c
		}
	}
	panic("fail")
}

func findLargestInRange(c []Cell, minIdx int, maxIdx int) (Cell, bool) {
	for _, c := range c {
		if c.idx >= minIdx && c.idx <= maxIdx {
			return c, true
		}
	}
	return Cell{}, false
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// previously had a single use solution for just 2 digits
	return solver(input, 2)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// refactor to parameterise the solution.
	return solver(input, 12)
}

func solver(input string, numDigits int) string {
	sum := int64(0)
	aoc.MapLines(input, func(line string) error {
		cells := []Cell{}
		for i, j := range line {
			cells = append(cells, Cell{
				jolts: int64(j - '0'),
				idx:   i,
			})
		}
		sortCells(cells)
		digits := []int64{}
		minIdx := 0
		l := len(cells)
		aoc.Debug("line", line, "cells", cells)
		for i := 0; i < numDigits; i++ {
			maxIdx := l - (numDigits - i)
			cell, ok := findLargestInRange(cells, minIdx, maxIdx)
			if !ok {
				// fail?
				aoc.Debug("digits", digits)
				panic("failed")
			}
			digits = append(digits, cell.jolts)
			minIdx = cell.idx + 1
		}

		n := int64(0)
		for i, x := range digits {
			n += int64(math.Pow10(numDigits-i-1)) * x
		}
		aoc.Debug("digits", digits, "n", n)
		sum += n
		return nil
	})
	return fmt.Sprint(sum)
}
