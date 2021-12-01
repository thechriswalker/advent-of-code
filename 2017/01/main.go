package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 1, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	b := []byte(strings.TrimSpace(input))
	var sum int
	for i, c := range b {
		if c == b[(i+1)%len(b)] {
			sum += int(c - '0')
		}
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	b := []byte(strings.TrimSpace(input))
	var sum int
	for i, c := range b {
		if c == b[(i+len(b)/2)%len(b)] {
			sum += int(c - '0')
		}
	}
	return fmt.Sprintf("%d", sum)
}
