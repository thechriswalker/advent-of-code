package main

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 1, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	scn := bufio.NewScanner(strings.NewReader(input))
	prev := uint64(math.MaxUint64)
	count := 0
	for scn.Scan() {
		curr, err := strconv.ParseUint(scn.Text(), 10, 64)
		if err != nil {
			panic(err)
		}
		if curr > prev {
			count++
		}
		prev = curr
	}
	return fmt.Sprintf("%d", count)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// this will be easier to use memory...
	strs := strings.Split(input, "\n")
	readings := make([]uint64, 0, len(strs))
	for _, s := range strs {
		i, err := strconv.ParseUint(s, 10, 64)
		if err == nil {
			readings = append(readings, i)
		}
	}
	// now create the windows.
	windows := make([]uint64, len(readings))
	for i := range readings {
		for j := i; j >= 0 && j > i-3; j-- {
			windows[i] += readings[j]
		}
	}
	count := 0
	// we start at index 2 (first 3 readings)
	prev := windows[2]
	//fmt.Println(readings)
	//fmt.Println(windows)
	for _, w := range windows[3:] {
		if prev < w {
			count++
		}
		prev = w
	}
	return fmt.Sprintf("%d", count)
}
