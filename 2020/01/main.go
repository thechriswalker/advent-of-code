package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 1, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	entries := strings.Split(input, "\n")
	stack := make([]int, 0, len(entries))
	for _, s := range entries {
		n, err := strconv.Atoi(s)
		if err == nil {
			for _, ss := range stack {
				if ss+n == 2020 {
					return fmt.Sprintf("%d", ss*n)
				}
			}
			stack = append(stack, n)
		}
	}
	return "<unsolved>"
}

// Implement Solution to Problem 2
func solve2(input string) string {
	entries := strings.Split(input, "\n")
	stack := make([]int, 0, len(entries))
	for _, s := range entries {
		n, err := strconv.Atoi(s)
		if err == nil {
			stack = append(stack, n)
		}
	}
	for i := 0; i < len(stack)-2; i++ {
		for j := i + 1; j < len(stack)-1; j++ {
			for k := j + 1; k < len(stack); k++ {
				ii, jj, kk := stack[i], stack[j], stack[k]
				if ii+jj+kk == 2020 {
					return fmt.Sprintf("%d", ii*jj*kk)
				}
			}
		}
	}
	return "<unsolved>"
}
