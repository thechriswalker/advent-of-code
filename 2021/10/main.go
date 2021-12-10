package main

import (
	"fmt"
	"sort"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 10, solve1, solve2)
}

// return and space if no error: ' '
func firstUnexpectedRune(line string) (unexpected rune) {
	// we push a closing rune on here for every opening rune
	// when we find a closing rune, it should be the top of the stack
	// or we found a syntax error.
	stack := []rune{}
	for _, r := range line {
		switch r {
		// openers
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '<':
			stack = append(stack, '>')
		// closers
		case ')', '}', ']', '>':
			if stack[len(stack)-1] == r {
				// good, expected.
				stack = stack[:len(stack)-1]
			} else {
				return r
			}
		}
	}
	return ' '
}

// Implement Solution to Problem 1
func solve1(input string) string {
	points := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	score := 0
	aoc.MapLines(input, func(line string) error {
		r := firstUnexpectedRune(line)
		//fmt.Printf("line %q unexpected: %c (score: %d)\n", line, r, points[r])
		score += points[r]
		return nil
	})

	return fmt.Sprintf("%d", score)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	scores := []int{}
	aoc.MapLines(input, func(line string) error {
		s := scoreIncompleteCompletion(line)
		if s > 0 {
			scores = append(scores, s)
			//fmt.Printf("line %q (score: %d)\n", line, s)
		}
		return nil
	})

	// sort and take the center value.
	sort.Ints(scores)
	mid := scores[len(scores)/2]

	return fmt.Sprintf("%d", mid)
}

var autopoints = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

// return and space if no error: ' '
func scoreIncompleteCompletion(line string) (score int) {
	// same as the other, but we score 0 for an invalid line.
	// and use the stack for the scoring (it will hold the correct closing tags)
	stack := []rune{}
	for _, r := range line {
		switch r {
		// openers
		case '(':
			stack = append(stack, ')')
		case '{':
			stack = append(stack, '}')
		case '[':
			stack = append(stack, ']')
		case '<':
			stack = append(stack, '>')
		// closers
		case ')', '}', ']', '>':
			if stack[len(stack)-1] == r {
				// good, expected.
				stack = stack[:len(stack)-1]
			} else {
				return
			}
		}
	}
	last := len(stack) - 1
	for i := 0; i <= last; i++ {
		// reading backwards
		score *= 5
		score += autopoints[stack[last-i]]
	}
	return
}
