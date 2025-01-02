package main

import (
	"fmt"
	"sort"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 1, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var a, b []int
	aoc.MapLines(input, func(line string) error {
		var x, y int
		_, err := fmt.Sscanf(line, "%d %d", &x, &y)
		if err != nil {
			return err
		}
		a = append(a, x)
		b = append(b, y)
		return nil
	})
	sort.IntSlice(a).Sort()
	sort.IntSlice(b).Sort()
	diff := 0
	for i := range a {
		aa, bb := a[i], b[i]
		if aa > bb {
			aa, bb = bb, aa
		}
		diff += bb - aa
	}

	return fmt.Sprint(diff)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	xx := []int{}
	yy := map[int]int{}

	aoc.MapLines(input, func(line string) error {
		var x, y int
		_, err := fmt.Sscanf(line, "%d %d", &x, &y)
		if err != nil {
			return err
		}
		xx = append(xx, x)
		yy[y] = yy[y] + 1
		return nil
	})

	score := 0
	for _, x := range xx {
		score += x * yy[x]
	}
	return fmt.Sprint(score)

}
