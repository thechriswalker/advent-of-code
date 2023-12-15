package main

import (
	"fmt"
	"slices"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 3, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	sum := 0
	aoc.MapLines(input, func(line string) error {
		sum += findCommon([]byte(line))
		return nil
	})
	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	sum := 0
	i := 0
	maps := make([]map[byte]struct{}, 3)
	aoc.MapLines(input, func(line string) error {
		maps[i] = lineToMap(line)
		if i == 2 {
			// get common value
			sum += getCommonValue(maps[0], maps[1], maps[2])
			i = 0
		} else {
			i++
		}
		return nil
	})
	return fmt.Sprint(sum)
}

func findCommon(rucksack []byte) int {
	cl := len(rucksack) / 2
	h1, h2 := rucksack[cl:], rucksack[:cl]
	for _, b := range h1 {
		if slices.Contains(h2, b) {
			return priority(b)
		}
	}
	panic("no match found")
}

func priority(b byte) int {
	switch {
	case b >= 'a' && b <= 'z':
		return int(b-'a') + 1
	case b >= 'A' && b <= 'Z':
		return int(b-'A') + 27
	default:
		panic("bad character!")
	}
}

func lineToMap(line string) map[byte]struct{} {
	m := map[byte]struct{}{}
	for _, b := range line {
		m[byte(b)] = struct{}{}
	}
	return m
}

func getCommonValue(a, b, c map[byte]struct{}) int {
	for aa := range a {
		for bb := range b {
			if aa == bb {
				for cc := range c {
					if bb == cc {
						return priority(aa)
					}
				}
			}
		}
	}
	panic("no matches!")
}
