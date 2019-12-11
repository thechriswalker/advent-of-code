package main

import (
	"fmt"

	"../../aoc"
)

func main() {
	aoc.Run(2015, 1, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	f := 0
	for _, c := range input {
		switch c {
		case '(':
			f++
		case ')':
			f--
		default:
			// ignore
		}
	}
	return fmt.Sprint(f)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	f := 0
	for i, c := range input {
		switch c {
		case '(':
			f++
		case ')':
			f--
		default:
			// ignore
		}
		if f == -1 {
			return fmt.Sprint(i + 1)
		}
	}
	return "<unsolved>"
}
