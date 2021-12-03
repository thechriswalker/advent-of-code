package main

import (
	"fmt"
	"strconv"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 8, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	code := 0
	mem := 0
	aoc.MapLines(input, func(s string) error {
		// we could just json decode, but that is
		code += len(s)
		s, _ = strconv.Unquote(s)
		mem += len(s)
		return nil
	})

	return fmt.Sprintf("%d", code-mem)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	orig := 0
	enc := 0
	aoc.MapLines(input, func(s string) error {
		// we could just json decode, but that is
		orig += len(s)
		s = strconv.Quote(s)
		enc += len(s)
		return nil
	})

	return fmt.Sprintf("%d", enc-orig)
}
