package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 19, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var patterns []string
	possible := 0

	aoc.MapLines(input, func(line string) error {
		line = strings.TrimSpace(line)
		if len(patterns) == 0 {
			patterns = strings.Split(line, ", ")
			return nil
		}
		if line == "" {
			return nil
		}
		// a design.
		if isPossible(line, patterns) {
			possible++
		}
		return nil
	})

	return fmt.Sprint(possible)
}

func isPossible(design string, patterns []string) bool {
	// can we find any combination of patterns that will match the design?
	// I think I will do this recursively.
	// we will try and match the patterns to the beginning of the design.
	// for all that "fit", we will try and match the rest of the design by slicing off the beginning and trying to match the rest.
	// I think this might fail for the larger designs... but let's see.
	// correct, so now we keep track of the patterns we have matched in a cache, and try "longest first"
	// or to start with, we will have a shared cache for the "offsets" where we know it is possible to match the design.
	// the prevents excess work.
	cache := make(map[int]bool)
	// try largest first
	slices.SortFunc(patterns, func(a, z string) int { return len(z) - len(a) })
	return rIsPossible(design, 0, cache, patterns)
}

func rIsPossible(design string, offset int, cache map[int]bool, patterns []string) bool {
	if result, ok := cache[offset]; ok {
		return result
	}
	for _, pattern := range patterns {
		if strings.HasPrefix(design[offset:], pattern) {
			// we have a match, now we need to try and match the rest of the design.
			if len(pattern) == len(design[offset:]) {
				// we have a match
				cache[offset] = true
				return true
			}
			// we can also add the whole thing to our "known" patterns. later...
			poss := rIsPossible(design, offset+len(pattern), cache, patterns)
			cache[offset] = poss
			if poss {
				return true
			}
		}
	}
	cache[offset] = false
	return false
}

// Implement Solution to Problem 2
func solve2(input string) string {
	var patterns []string
	possibilities := 0

	aoc.MapLines(input, func(line string) error {
		line = strings.TrimSpace(line)
		if len(patterns) == 0 {
			patterns = strings.Split(line, ", ")
			return nil
		}
		if line == "" {
			return nil
		}
		// a design.
		possibilities += countPossibilities(line, patterns)
		return nil
	})

	return fmt.Sprint(possibilities)
}

func countPossibilities(design string, patterns []string) int {
	// similar to the isPossible, but we need to cache the "number" of cache hits at a given offset,
	// then when we find the "next" number of possibilities, we can multiply
	cache := make(map[int]int)
	// try largest first
	slices.SortFunc(patterns, func(a, z string) int { return len(z) - len(a) })
	return rCountPossibilities(design, 0, cache, patterns)
}

func rCountPossibilities(design string, offset int, cache map[int]int, patterns []string) int {
	if result, ok := cache[offset]; ok {
		return result
	}
	possibilities := 0
	for _, pattern := range patterns {
		if strings.HasPrefix(design[offset:], pattern) {
			// we have a match, now we need to try and match the rest of the design.
			if len(pattern) == len(design[offset:]) {
				// we have a match
				possibilities++
			} else {
				possibilities += rCountPossibilities(design, offset+len(pattern), cache, patterns)
			}
		}
	}
	cache[offset] = possibilities
	return possibilities

}
