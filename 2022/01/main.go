package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 1, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	curr, max := 0, 0

	aoc.MapLines(input, func(line string) error {
		if len(line) == 0 {
			if curr > max {
				max = curr
			}
			curr = 0
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				return err
			}
			curr += i
		}
		return nil
	})
	// add the final one!
	if curr > max {
		max = curr
	}
	return fmt.Sprint(max)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	ints := sort.IntSlice{}
	curr := 0
	aoc.MapLines(input, func(line string) error {
		if len(line) == 0 {
			ints = append(ints, curr)
			curr = 0
		} else {
			i, err := strconv.Atoi(line)
			if err != nil {
				return err
			}
			curr += i
		}
		return nil
	})
	// add the final one
	ints = append(ints, curr)
	sort.Sort(sort.Reverse(ints))
	//fmt.Print(ints)
	// assume there will be at least three.
	sum := ints[0] + ints[1] + ints[2]
	return fmt.Sprint(sum)
}
