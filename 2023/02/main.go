package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 2, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {

	limits := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	re := regexp.MustCompile(`(\d+) (red|green|blue)`)

	sum := 0

	aoc.MapLines(input, func(line string) error {
		var impossible bool
		for _, m := range re.FindAllStringSubmatch(line, -1) {
			// match should be ["13 green", "13", "green"]
			n, _ := strconv.Atoi(m[1])
			// mark game impossible
			if limits[m[2]] < n {
				impossible = true
				break
			}
		}
		if !impossible {
			var id int
			fmt.Sscanf(line, "Game %d:", &id)
			sum += id
		}
		return nil
	})

	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	var sum int
	re := regexp.MustCompile(`(\d+) (red|green|blue)`)
	aoc.MapLines(input, func(line string) error {
		// need to split into games.
		s1 := strings.Split(line, ": ")
		s2 := strings.Split(s1[1], "; ")
		// now each s2 is a draw.

		// find biggest RGB values in this draw
		var r, g, b int
		for _, d := range s2 {
			for _, m := range re.FindAllStringSubmatch(d, -1) {
				// match should be ["13 green", "13", "green"]
				n, _ := strconv.Atoi(m[1])
				switch m[2] {
				case "red":
					if n > r {
						r = n
					}
				case "green":
					if n > g {
						g = n
					}
				case "blue":
					if n > b {
						b = n
					}
				}
			}
		}
		// power set it the multiplied values
		p := r * g * b
		sum += p
		return nil
	})

	return fmt.Sprint(sum)
}
