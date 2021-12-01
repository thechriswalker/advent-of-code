package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 1, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	masses := strings.Split(input, "\n")
	sum := 0
	for _, s := range masses {
		n, err := strconv.Atoi(s)
		if err == nil {
			sum += fuelReqs(n)
		}
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	masses := strings.Split(input, "\n")
	sum := 0
	for _, s := range masses {
		n, err := strconv.Atoi(s)
		if err == nil {
			sum += recursiveFuelReqs(n)
		}
	}
	return fmt.Sprintf("%d", sum)
}

func fuelReqs(n int) int {
	d := float64(n) / 3.0
	rd := math.Floor(d)
	return int(rd) - 2
}

func recursiveFuelReqs(n int) int {
	r := fuelReqs(n)
	sum := 0
	for r > 0 {
		sum += r
		r = fuelReqs(r)
	}
	return sum
}
