package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 2, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {

	var f = func(row []int) int {
		// find the diff between biggest and smallest
		var min, max = row[0], row[0]
		for i := 1; i < len(row); i++ {
			if row[i] < min {
				min = row[i]
			}
			if row[i] > max {
				max = row[i]
			}
		}
		return max - min
	}

	return fmt.Sprintf("%d", run(input, f))
}

// Implement Solution to Problem 2
func solve2(input string) string {

	var f = func(row []int) int {
		var a, b int
		for i := 0; i < len(row)-1; i++ {
			for j := i + 1; j < len(row); j++ {
				a, b = row[i], row[j]
				if a < b {
					a, b = b, a
				}
				if b*(a/b) == a {
					return a / b
				}
			}
		}
		panic("no divisble entries in row")
	}

	return fmt.Sprintf("%d", run(input, f))
}

func run(input string, f func(row []int) int) int {
	lines := strings.Split(input, "\n")
	var sum, v int
	row := make([]int, 64) // max 64 entries per line (i think we get 16)
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		row = row[0:0]
		rd := strings.NewReader(line)

		for {
			if _, err := fmt.Fscanf(rd, "%d", &v); err == nil {
				row = append(row, v)
			} else {
				break
			}
		}
		if len(row) > 0 {
			// handle line n long
			sum += f(row)
		}
	}

	return sum
}
