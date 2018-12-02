package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 3, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// this is the horizontal so we scan line by line
	count := 0
	r := strings.NewReader(input)
	var a, b, c int
	for {
		n, _ := fmt.Fscanf(r, "%d %d %d\n", &a, &b, &c)
		if n != 3 {
			break
		}
		if IsTrianglePossible(a, b, c) {
			count++
		}
	}
	return fmt.Sprintf("%d", count)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// vertical, so run in threes.
	var a, b, c, d, e, f, g, h, i int
	count := 0
	r := strings.NewReader(input)
	for {
		n, _ := fmt.Fscanf(r, "%d %d %d\n%d %d %d\n%d %d %d\n",
			&a, &d, &g,
			&b, &e, &h,
			&c, &f, &i,
		)
		if n != 9 {
			break
		}
		if IsTrianglePossible(a, b, c) {
			count++
		}
		if IsTrianglePossible(d, e, f) {
			count++
		}
		if IsTrianglePossible(g, h, i) {
			count++
		}
	}
	return fmt.Sprintf("%d", count)
}

func IsTrianglePossible(a, b, c int) bool {
	if a > b {
		if a > c {
			return a < b+c
		}
	} else if b > c {
		return b < a+c
	}
	return c < b+a
}
