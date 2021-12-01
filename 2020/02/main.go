package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 2, solve1, solve2)
}

type password struct {
	char     string
	min, max int
	plain    string
}

func (p *password) valid() bool {
	count := strings.Count(p.plain, p.char)
	return count >= p.min && count <= p.max
}

func (p *password) valid2() bool {
	c := p.char[0]
	// 1-based indexing!
	pa := p.min - 1
	pb := p.max - 1
	a := pa < len(p.plain) && p.plain[pa] == c
	b := pb < len(p.plain) && p.plain[pb] == c
	return (a && !b) || (b && !a)
}

func passwords(in string) []password {
	entries := strings.Split(in, "\n")
	stack := make([]password, 0, len(entries))
	var min, max int
	var char, plain string
	for _, entry := range entries {
		n, err := fmt.Sscanf(entry, "%d-%d %1s: %s\n", &min, &max, &char, &plain)
		//		log.Printf("min:%d, max:%d, char:'%s', plain:'%s'\n", min, max, char, plain)
		if n != 4 || err != nil {
			continue
		}
		p := password{
			min:   min,
			max:   max,
			char:  char,
			plain: plain,
		}
		//log.Printf("%#v\n", p)
		stack = append(stack, p)

	}
	return stack
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// now count how many are valid.
	sum := 0
	for _, p := range passwords(input) {
		if p.valid() {
			sum++
		}
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// now count how many are valid.
	sum := 0
	for _, p := range passwords(input) {
		//	fmt.Printf("min:%d, max:%d, char:'%s', plain:'%s', valid:%v\n", p.min, p.max, p.char, p.plain, p.valid2())
		if p.valid2() {
			sum++
		}
	}
	return fmt.Sprintf("%d", sum)
}
