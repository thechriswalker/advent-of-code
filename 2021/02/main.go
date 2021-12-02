package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 2, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	h := 0
	d := 0

	var cmd string
	var x int
	aoc.MapLines(input, func(line string) error {
		_, err := fmt.Sscanf(line, "%s %d", &cmd, &x)
		if err == nil {
			switch cmd {
			case "forward":
				h += x
			case "down":
				d += x
			case "up":
				d -= x
			default:
				fmt.Printf("Unknown Command: %q\n", line)
			}
		}
		return nil
	})

	return fmt.Sprintf("%d", h*d)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	h := 0
	d := 0
	aim := 0

	var cmd string
	var x int
	aoc.MapLines(input, func(line string) error {
		_, err := fmt.Sscanf(line, "%s %d", &cmd, &x)
		if err == nil {
			switch cmd {
			case "forward":
				h += x
				d += aim * x
			case "down":
				aim += x
			case "up":
				aim -= x
			default:
				fmt.Printf("Unknown Command: %q\n", line)
			}
		}
		return nil
	})

	return fmt.Sprintf("%d", h*d)
}
