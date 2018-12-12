package main

import (
	"fmt"
	"sort"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2017, 4, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	sum := run(input, func(fields []string) bool {
		cache := map[string]struct{}{}
		for _, s := range fields {
			if _, ok := cache[s]; ok {
				return false
			}
			cache[s] = struct{}{}
		}
		return true
	})
	return fmt.Sprintf("%d", sum)
}

type sortbytes []byte

func (s sortbytes) Len() int           { return len(s) }
func (s sortbytes) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s sortbytes) Less(i, j int) bool { return s[i] < s[j] }

// Implement Solution to Problem 2
func solve2(input string) string {
	sum := run(input, func(fields []string) bool {
		cache := map[string]struct{}{}
		for _, s := range fields {
			b := sortbytes(s)
			sort.Sort(b)
			s = string(b)
			if _, ok := cache[s]; ok {
				return false
			}
			cache[s] = struct{}{}
		}
		return true
	})
	return fmt.Sprintf("%d", sum)
}

func run(input string, f func(s []string) bool) int {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	var sum int
	for _, passphrase := range lines {
		fields := strings.Fields(passphrase)
		if f(fields) {
			sum++
		}
	}
	return sum
}
