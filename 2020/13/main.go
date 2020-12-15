package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2020, 13, solve1, solve2)
}

func parseInput(in string) (int, map[int]int) {
	//earliest time on first line.
	//buses on second.
	lines := strings.Split(in, "\n")
	// need at least 2
	estimate, err := strconv.Atoi(lines[0])
	if err != nil {
		panic(err)
	}
	buses := strings.Split(lines[1], ",")
	m := make(map[int]int, len(buses))
	// gonna put in the "out of service" as 1
	// as the actual part 2 says "anytime"
	i := 0
	for _, b := range buses {
		if b != "x" {
			n, _ := strconv.Atoi(b)
			m[i] = n
		}
		i++
	}
	return estimate, m
}

// Implement Solution to Problem 1
func solve1(input string) string {
	estimate, buses := parseInput(input)
	x := estimate
	for {
		// find a bus that starts leaves at the given time
		// i.e. x % number == 0
		for _, b := range buses {
			if x%b == 0 {
				// this one leaves now!
				return fmt.Sprintf("%d", b*(x-estimate))
			}
		}
		x++
		if x == estimate+10000 {
			break
		}
	}
	return "<unsolved>"
}

// Implement Solution to Problem 2
func solve2(input string) string {
	_, buses := parseInput(input)
	// need to find the first time that the buses
	// depart in order.
	// i.e. `t` such that `(t + x)%bus[x] == 0` for all x.
	t := 1
	inc := 1

	for i, b := range buses {
		for (t+i)%b != 0 {
			t += inc
		}
		inc *= b
	}
	return fmt.Sprintf("%d", t)
}
