package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 1, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var sum int

	aoc.MapLines(input, func(line string) error {
		var a, b int
		f := false
		// find the first and last int in the string
		for _, c := range []byte(line) {
			if c >= '0' && c <= '9' {
				n := int(c) - '0'
				if !f {
					a = n
					f = true
				}
				b = n
			}
		}

		sum += a*10 + b
		return nil
	})

	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	sum := 0

	// n = index + 1
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

	isNum := func(l string) (n int) {
		b := byte(l[0])
		if b >= '0' && b <= '9' {
			return int(b) - int('0')
		}
		for idx, t := range numbers {
			if len(l) < len(t) {
				continue
			}
			if l[0:len(t)] == t {
				return idx + 1
			}
		}
		return 0
	}

	aoc.MapLines(input, func(line string) error {
		var first, second int

		for i := 0; i < len(line); i++ {
			// find next number
			n := isNum(line[i:])
			if n > 0 {
				// we found one
				if first == 0 {
					first = n
				}
				second = n
			}
		}
		//fmt.Println("found", first, " and ", second, " in ", line)
		sum += (first * 10) + second
		return nil
	})

	// 53988 too low
	// 54258 too low
	return fmt.Sprint(sum)
}
