package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 9, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	sequences := [][]int{}

	aoc.MapLines(input, func(line string) error {
		sequences = append(sequences, aoc.ToIntSlice(line, ' '))
		return nil
	})

	sum := 0
	for _, s := range sequences {
		sum += extrapolateValue(s, true)
	}

	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	sequences := [][]int{}

	aoc.MapLines(input, func(line string) error {
		sequences = append(sequences, aoc.ToIntSlice(line, ' '))
		return nil
	})

	sum := 0
	for _, s := range sequences {
		sum += extrapolateValue(s, false)
	}

	return fmt.Sprint(sum)
}

func extrapolateValue(s []int, prev bool) int {
	values := [][]int{s}
	// start with one less than the length
	// of the initial sequence, and work
	// back until we have a "length 1"
	// sequence or a new sequence of "all 0s"
	for n := len(s) - 1; n > 2; n-- {
		allZeros := true
		next := make([]int, n)
		prev := tail(values)
		for i := range next {
			// difference between i+1 and i of prev seq
			next[i] = prev[i+1] - prev[i]
			if next[i] != 0 {
				allZeros = false
			}
		}
		values = append(values, next)
		if allZeros {
			break
		}
	}

	//fmt.Println(values)

	// now work back up using the end or beginning of every sequence
	x := 0
	if prev {
		//fmt.Println("last value", x)
		for i := len(values) - 1; i >= 0; i-- {
			x = tail(values[i]) + x
			//	fmt.Println("next last value", x)
		}
	} else {
		for i := len(values) - 1; i >= 0; i-- {
			x = values[i][0] - x
			//	fmt.Println("next first value", x)
		}
		//fmt.Println(x)
	}

	return x
}

func tail[T any](s []T) T {
	return s[len(s)-1]
}
