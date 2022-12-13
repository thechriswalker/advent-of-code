package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 9, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	head, tail := Pos{0, 0}, Pos{0, 0}

	tailCache := map[Pos]int{tail: 1}

	// parse instructions and move
	aoc.MapLines(input, func(line string) error {
		var d rune
		var n int
		fmt.Sscanf(line, "%c %d", &d, &n)
		//fmt.Printf("moving %d lines in direction %c\n", n, d)
		for i := 0; i < n; i++ {
			switch d {
			case 'R':
				head[0]++
			case 'L':
				head[0]--
			case 'U':
				head[1]++
			case 'D':
				head[1]--
			}
			tail = follow(head, tail)
			tailCache[tail]++
			//fmt.Printf(" - head: %v, tail: %v\n", head, tail)
		}
		return nil
	})

	return fmt.Sprint(len(tailCache))
}

func absDiff(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func follow(head, tail Pos) Pos {
	x := absDiff(head[0], tail[0])
	y := absDiff(head[1], tail[1])

	// we are in the same column/row and more than
	// 1 from head, move.
	switch {
	case x == 0:
		if y > 1 {
			if head[1] > tail[1] {
				tail[1]++
			} else {
				tail[1]--
			}
		}
	case y == 0:
		if x > 1 {
			if head[0] > tail[0] {
				tail[0]++
			} else {
				tail[0]--
			}
		}
	default:
		if y > 1 || x > 1 {
			if head[0] > tail[0] {
				tail[0]++
			} else {
				tail[0]--
			}
			if head[1] > tail[1] {
				tail[1]++
			} else {
				tail[1]--
			}
		}
	}
	return tail
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// now we have a slice of ten and they follow each other.
	snake := [10]Pos{}
	// we still keep track of the final one
	cache := map[Pos]int{snake[9]: 1}
	// parse instructions and move
	aoc.MapLines(input, func(line string) error {
		var d rune
		var n int
		fmt.Sscanf(line, "%c %d", &d, &n)
		//fmt.Printf("moving %d lines in direction %c\n", n, d)
		for i := 0; i < n; i++ {
			switch d {
			case 'R':
				snake[0][0]++
			case 'L':
				snake[0][0]--
			case 'U':
				snake[0][1]++
			case 'D':
				snake[0][1]--
			}

			for j := 1; j < 10; j++ {
				snake[j] = follow(snake[j-1], snake[j])
			}

			cache[snake[9]]++
			//fmt.Printf(" - head: %v, tail: %v\n", head, tail)
		}
		return nil
	})

	return fmt.Sprint(len(cache))
}

type Pos [2]int
