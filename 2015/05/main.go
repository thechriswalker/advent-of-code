package main

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 5, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	s := bufio.NewScanner(strings.NewReader(input))
	count := 0
	for s.Scan() {
		ok := isNice1(s.Text())
		//fmt.Printf("%s => %v\n", s.Text(), ok)
		if ok {
			count++
		}
	}

	return fmt.Sprintf("%d", count)
}

func isNice1(s string) bool {
	// for the vowel and double letter,
	// we will just mark a bool.
	// for the bad groups, we will return early
	l := len(s)
	var vowel_count int
	var has_double bool
	for i := 0; i < l; i++ {

		if vowel_count < 3 {
			switch s[i] {
			case 'a', 'e', 'i', 'o', 'u':
				vowel_count++
			}
		}
		if i != l-1 {
			switch s[i : i+2] {
			case "ab", "cd", "pq", "xy":
				return false
			}
			if !has_double {
				if s[i] == s[i+1] {
					has_double = true
				}
			}
		}
	}
	return vowel_count >= 3 && has_double
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return "<unsolved>"
}
