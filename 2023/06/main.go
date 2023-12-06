package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 6, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var times, distances []int

	aoc.MapLines(input, func(line string) error {
		if strings.HasPrefix(line, "Time: ") {
			times = aoc.ToIntSlice(line[6:], ' ')
		}
		if strings.HasPrefix(line, "Distance: ") {
			distances = aoc.ToIntSlice(line[10:], ' ')
		}
		return nil
	})
	n := 1

	for i := range times {
		// how to win this game?
		n *= numWinningOptions(times[i], distances[i])
	}
	return fmt.Sprint(n)
}

func numWinningOptions(t, d int) int {
	n := 0

	for i := 1; i < t; i++ {
		di := i * (t - i)
		if di > d {
			n++
		}
	}
	//fmt.Println("t=", t, "d=", d, "n=", n)
	return n
}

// Implement Solution to Problem 2
func solve2(input string) string {
	var t, d int
	aoc.MapLines(input, func(line string) error {
		if strings.HasPrefix(line, "Time: ") {
			line := strings.ReplaceAll(line[6:], " ", "")
			t, _ = strconv.Atoi(line)
		}
		if strings.HasPrefix(line, "Distance: ") {
			line := strings.ReplaceAll(line[10:], " ", "")
			d, _ = strconv.Atoi(line)
		}
		return nil
	})
	n := numWinningOptions(t, d)
	return fmt.Sprint(n)
}
