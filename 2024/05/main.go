package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 5, solve1, solve2)
}

func parse(input string) (rules map[int][]int, lists [][]int) {
	rules = map[int][]int{} // key = page, val = before

	aoc.MapLines(input, func(line string) error {
		if line == "" {
			return nil
		}
		if strings.Contains(line, "|") {
			s := aoc.ToIntSlice(line, '|')
			rules[s[0]] = append(rules[s[0]], s[1])
			return nil
		}
		if strings.Contains(line, ",") {
			lists = append(lists, aoc.ToIntSlice(line, ','))
			return nil
		}
		return nil
	})

	return
}

func isOrdered(list []int, rules map[int][]int) bool {
	for i := 0; i < len(list); i++ {
		p := list[i]
		// we need to find
		// is there a rule for this number?
		if r, ok := rules[p]; ok {
			// there is a rule, and we need to make sure that this number is "before" all the other numbers.
			// i.e. we need to make sure none of those numbers appear _before_ this number.
			for j := 0; j < i; j++ {
				if slices.Contains(r, list[j]) {
					// bad
					return false
				}
			}
		}
	}
	return true
}

// Implement Solution to Problem 1
func solve1(input string) string {
	sum := 0
	rules, lists := parse(input)
	for _, list := range lists {
		// and check each number to see if it's in right place compared to the other numbers;
		if isOrdered(list, rules) {
			// find the middle number an add it.
			sum += list[len(list)/2]
		}
		//fmt.Println(list, isOrdered(list), list[len(list)/2])
	}
	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	sum := 0
	rules, lists := parse(input)
	for _, list := range lists {
		if fixOrdering(list, rules) {
			sum += list[len(list)/2]
		}
	}
	return fmt.Sprint(sum)
}

func fixOrdering(list []int, rules map[int][]int) (fixed bool) {
	// same as isOrdered, but we need to find the first number that is out of order
	// and switch it
	for i := 0; i < len(list); i++ {
		p := list[i]
		// we need to find
		// is there a rule for this number?
		if r, ok := rules[p]; ok {
			// there is a rule, and we need to make sure that this number is "before" all the other numbers.
			// i.e. we need to make sure none of those numbers appear _before_ this number.
			for j := 0; j < i; j++ {
				if slices.Contains(r, list[j]) {
					// bad
					// need to swap position
					list[i], list[j] = list[j], list[i]
					fixed = true
				}
			}
		}
	}
	return
}
