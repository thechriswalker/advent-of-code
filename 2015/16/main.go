package main

import (
	"errors"
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 16, solve1, solve2)
}

var result = map[string]int{
	"children:":    3,
	"cats:":        7,
	"samoyeds:":    2,
	"pomeranians:": 3,
	"akitas:":      0,
	"vizslas:":     0,
	"goldfish:":    5,
	"trees:":       3,
	"cars:":        2,
	"perfumes:":    1,
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var s1, s2, s3 string
	var x, v1, v2, v3 int
	aoc.MapLines(input, func(line string) error {
		fmt.Sscanf(line, "Sue %d: %s %d, %s %d, %s %d", &x, &s1, &v1, &s2, &v2, &s3, &v3)
		if result[s1] == v1 && result[s2] == v2 && result[s3] == v3 {
			return errors.New("")
		}
		return nil
	})

	return fmt.Sprintf("%d", x)
}

// Implement Solution to Problem 2
func solve2(input string) string {

	var match = func(s string, v int) bool {
		switch s {
		case "cats:", "trees:":
			return result[s] < v
		case "pomeranians:", "goldfish:":
			return result[s] > v
		default:
			return result[s] == v
		}
	}

	var s1, s2, s3 string
	var x, v1, v2, v3 int
	aoc.MapLines(input, func(line string) error {
		fmt.Sscanf(line, "Sue %d: %s %d, %s %d, %s %d", &x, &s1, &v1, &s2, &v2, &s3, &v3)
		if match(s1, v1) && match(s2, v2) && match(s3, v3) {
			return errors.New("")
		}
		return nil
	})

	return fmt.Sprintf("%d", x)
}
