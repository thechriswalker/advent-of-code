package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 6, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return solveN(input, 4)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return solveN(input, 14)
}

// could have abstracted that...
func solveN(input string, n int) string {
	// iterate over the "trimmed" string.
	input = strings.TrimSpace(input)
	buf := make([]rune, n)
	//fmt.Println("input:", input)
iteration:
	for i, r := range input {
		buf[i%n] = r
		if i < n-1 {
			continue
		}
		// if the buf contains ALL different values we are good.
		//fmt.Println("index:", i, "buf:", string(buf[]))
		for x := 0; x < n-1; x++ {
			for y := x + 1; y < n; y++ {
				//fmt.Println("testing, x:", x, "y:", y, "buf[x]:", buf[x], "buf[y]:", buf[y])
				if buf[x] == buf[y] {
					continue iteration
				}
			}
		}
		// if we got here, we are done.
		return fmt.Sprint(i + 1)
	}
	panic("no solution")
}
