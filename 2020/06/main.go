package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2020, 6, solve1, solve2)
}

type QA struct {
	answers map[rune]int
	size    int
}

func newQA() *QA {
	return &QA{
		answers: map[rune]int{},
		size:    0,
	}
}

func (qa *QA) addAnswers(line string) {
	qa.size++
	for _, c := range line {
		if c >= 'a' && c <= 'z' {
			qa.answers[c] = qa.answers[c] + 1
		} else {
			panic("BAD CHARACTER IN LINE: " + line)
		}
	}
}

func (qa *QA) countAnyone() int {
	return len(qa.answers)
}
func (qa *QA) countEveryone() int {
	count := 0
	for _, n := range qa.answers {
		if n == qa.size {
			count++
		}
	}
	return count
}

func parseQuestions(in string) []*QA {
	lines := strings.Split(in, "\n")
	stack := make([]*QA, 0, len(lines)/2)

	curr := newQA()
	for _, line := range lines {
		if line == "" {
			stack = append(stack, curr)
			curr = newQA()
			continue
		}
		curr.addAnswers(line)
	}
	if curr.countAnyone() > 1 {
		stack = append(stack, curr)
	}
	return stack
}

// Implement Solution to Problem 1
func solve1(input string) string {
	answers := parseQuestions(input)

	sum := 0
	for _, a := range answers {
		sum += a.countAnyone()
	}

	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	answers := parseQuestions(input)

	sum := 0
	for _, a := range answers {
		sum += a.countEveryone()
	}

	return fmt.Sprintf("%d", sum)
}
