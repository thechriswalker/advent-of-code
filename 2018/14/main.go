package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2018, 14, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	recipes := []uint8{3, 7}

	var count int
	// elves (current index)
	e1, e2 := 0, 1
	fmt.Sscanf(input, "%d", &count)

	for {
		if len(recipes) >= count+10 {
			// we have enough
			break
		}
		// create new recipes.
		// add current and sum the digits. min 0, max 18
		sum := recipes[e1] + recipes[e2]
		if sum > 9 {
			// add 1
			recipes = append(recipes, 1, sum%10)
		} else {
			recipes = append(recipes, sum)
		}
		// move elves.
		e1 = (e1 + int(recipes[e1]) + 1) % len(recipes)
		e2 = (e2 + int(recipes[e2]) + 1) % len(recipes)
	}

	// slice the 10 numbers
	s := &strings.Builder{}
	for i := 0; i < 10; i++ {
		fmt.Fprintf(s, "%d", recipes[count+i])
	}
	return s.String()
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// similar but we are looking for a pattern in the last X recipes
	input = strings.TrimSpace(input)
	pattern := make([]uint8, len(input))
	for i, c := range input {
		pattern[i] = uint8(c) - '0'
	}
	patternIndex := 0
	recipes := []uint8{3, 7}
	// elves (current index)
	e1, e2 := 0, 1
	if recipes[0] == pattern[0] {
		patternIndex = 1
		if recipes[1] == pattern[1] {
			patternIndex = 2
		} else {
			patternIndex = 0
		}
	}
	for {
		// create new recipes.cd 2018
		// add current and sum the digits. min 0, max 18
		sum := recipes[e1] + recipes[e2]
		if sum > 9 {
			// add 1
			recipes = append(recipes, 1)
			if pattern[patternIndex] == 1 {
				patternIndex++
				if patternIndex == len(pattern) {
					break
				}
			} else {
				patternIndex = 0
				if pattern[0] == 1 {
					patternIndex++
				}
			}
		}
		sum = sum % 10
		recipes = append(recipes, sum)
		if pattern[patternIndex] == sum {
			patternIndex++
			if patternIndex == len(pattern) {
				break
			}
		} else {
			patternIndex = 0
			if pattern[0] == sum {
				patternIndex++
			}

		}

		// move elves.
		e1 = (e1 + int(recipes[e1]) + 1) % len(recipes)
		e2 = (e2 + int(recipes[e2]) + 1) % len(recipes)
	}
	// now count is len(recipes) - len(pattern)

	return fmt.Sprintf("%d", len(recipes)-len(pattern))
}
