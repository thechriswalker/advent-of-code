package main

import (
	"fmt"

	"../../aoc"
)

func main() {
	aoc.Run(2015, 3, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var x, y int
	visits := map[[2]int]struct{}{}
	visit := func() {
		visits[[2]int{x, y}] = struct{}{}
	}
	for _, c := range input {
		visit()
		switch c {
		case '^':
			y++
		case '<':
			x--
		case '>':
			x++
		case 'v':
			y--
		}
	}
	visit()

	return fmt.Sprint(len(visits))

}

// Implement Solution to Problem 2
func solve2(input string) string {
	var santa, robot [2]int
	visits := map[[2]int]struct{}{}
	move := func(p [2]int, c rune) [2]int {
		switch c {
		case '^':
			p[1]++
		case '<':
			p[0]--
		case '>':
			p[0]++
		case 'v':
			p[1]--
		}
		return p
	}
	visit := func(p [2]int) {
		visits[p] = struct{}{}
	}

	for i, c := range input {
		if i%2 == 0 {
			visit(santa)
			santa = move(santa, c)
		} else {
			visit(robot)
			robot = move(robot, c)
		}
	}
	visit(santa)
	visit(robot)

	return fmt.Sprint(len(visits))
}
