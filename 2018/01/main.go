package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2018, 1, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	r := strings.NewReader(input)
	sum := 0
	var f int
	for {
		if n, _ := fmt.Fscanln(r, &f); n != 1 {
			break
		}
		sum += f
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	r := strings.NewReader(input)
	sum := 0
	visited := map[int]bool{0: true}
	var f int
	for {
		if n, _ := fmt.Fscanln(r, &f); n != 1 {
			// need to reset
			r.Reset(input)
		} else {
			sum += f
			if _, ok := visited[sum]; ok {
				return fmt.Sprintf("%d", sum)
			} else {
				visited[sum] = true
			}
		}
	}
}
