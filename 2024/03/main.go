package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 3, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	result := doSimpleMults(input)
	return fmt.Sprint(result)
}

var re = regexp.MustCompile(`mul\((\d+),(\d+)\)`)

func doSimpleMults(input string) int {
	result := 0
	matches := re.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		a, _ := strconv.Atoi(match[1])
		b, _ := strconv.Atoi(match[2])
		result += a * b
	}
	//	fmt.Println("working on section: ", input, "-> result: ", result)
	return result
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// find sections of input that are enabled,
	enabled := true
	result := 0
	for i := 0; i < len(input); {
		var idx int
		if enabled {
			// look for the next disable command
			idx = strings.Index(input[i:], "don't()")
			if idx == -1 {
				// to the end!
				result += doSimpleMults(input[i:])
				break
			} else {
				result += doSimpleMults(input[i : i+idx])
				i += idx
				enabled = false
			}
		} else {
			// look for the next enable command
			idx = strings.Index(input[i:], "do()")
			if idx == -1 {
				// to the end!
				break
			} else {
				i += idx
				enabled = true
			}
		}
	}
	return fmt.Sprint(result)
}
