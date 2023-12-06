package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 4, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	pts := 0

	aoc.MapLines(input, func(line string) error {
		colon := strings.Index(line, ": ")
		pipe := strings.Index(line, " | ")

		winning := aoc.ToIntSlice(line[colon+2:pipe], ' ')
		yours := aoc.ToIntSlice(line[pipe+3:], ' ')

		hits := 0

		for _, x := range yours {
			if slices.Contains(winning, x) {
				hits++
			}
		}
		if hits > 0 {
			score := 1
			for x := 1; x < hits; x++ {
				score *= 2
			}
			pts += score
		}
		return nil
	})

	return fmt.Sprint(pts)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	cards := map[int]int{}

	card := 0
	aoc.MapLines(input, func(line string) error {
		card++
		colon := strings.Index(line, ": ")
		pipe := strings.Index(line, " | ")

		winning := aoc.ToIntSlice(line[colon+2:pipe], ' ')
		yours := aoc.ToIntSlice(line[pipe+3:], ' ')

		cards[card]++ // original card

		hits := 0

		for _, x := range yours {
			if slices.Contains(winning, x) {
				hits++
				// win a copy of the card+hits card (for each version we have)
				cards[card+hits] += cards[card]
			}
		}
		return nil
	})

	//fmt.Println(cards)

	// count all cards
	sum := 0
	for _, n := range cards {
		sum += n
	}

	return fmt.Sprint(sum)
}
