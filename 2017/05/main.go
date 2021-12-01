package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 5, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	list := parseInput(input)
	i := 0
	n := 0
	for i >= 0 && i < len(list) {
		j := list[i]
		list[i]++
		i += j
		n++
	}
	return fmt.Sprint(n)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	list := parseInput(input)
	i := 0
	n := 0
	for i >= 0 && i < len(list) {
		j := list[i]
		if j > 2 {
			list[i]--
		} else {
			list[i]++
		}
		i += j
		n++
	}
	return fmt.Sprint(n)
}

func parseInput(input string) []int {
	rd := strings.NewReader(input)
	list := []int{}
	var i int
	for {
		if _, err := fmt.Fscanln(rd, &i); err == nil {
			list = append(list, i)
		} else {
			break
		}
	}
	return list
}
