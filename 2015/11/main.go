package main

import (
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 11, solve1, solve2)
}

// in place increment
// 8 chars a-z
func incrementPassword(p []byte) {
	for i := 7; i >= 0; i-- {
		// try to increment this byte
		if p[i] == 'z' {
			p[i] = 'a'
			continue
		} else {
			p[i]++
			return
		}
	}
}

func isValid(pass []byte) bool {
	// must contain a run of three in sequence
	runOfThree := false
	// and 2 pairs of non-overlapping matches
	firstPair := 0
	secondPair := false
	for i := 0; i < 8; i++ {
		switch pass[i] {
		case 'o', 'i', 'l': // non of these allowed
			return false
		}
		if !runOfThree && i <= 5 {
			if pass[i]+1 == pass[i+1] && pass[i+1]+1 == pass[i+2] {
				// run of 3
				runOfThree = true
			}
		}
		if firstPair == 0 && i <= 4 {
			if pass[i] == pass[i+1] {
				firstPair = i + 1
			}
		}
		if firstPair != 0 && !secondPair && i <= 6 && i > firstPair {
			if pass[i] == pass[i+1] {
				secondPair = true
			}
		}
	}
	return runOfThree && secondPair
}

// Implement Solution to Problem 1
func solve1(input string) string {
	pass := []byte(strings.TrimSpace(input))
	for {
		incrementPassword(pass)
		if isValid(pass) {
			return string(pass)
		}
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// here we start with the answer to the previous
	pass := []byte(solve1(input))
	for {
		incrementPassword(pass)
		if isValid(pass) {
			return string(pass)
		}
	}
}
