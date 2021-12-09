package main

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 19, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// we map the "string"  -> ["replace", "ments"]
	// then do 1 tick of replacements.
	// we ONLY perform 1 replacement per tick.
	options := map[string][]string{}
	var molecule, a, b string
	var optionsDone bool
	aoc.MapLines(input, func(line string) error {
		if optionsDone {
			molecule = line
		} else if line == "" {
			optionsDone = true
		} else {
			fmt.Sscanf(line, "%s => %s", &a, &b)
			if m, ok := options[a]; ok {
				options[a] = append(m, b)
			} else {
				options[a] = []string{b}
			}

		}
		return nil
	})

	// we discard duplicates, so lets make a set.
	set := map[string]struct{}{}
	for i := range molecule {
		for k, v := range options {
			if strings.HasPrefix(molecule[i:], k) {
				l := len(k)
				for _, ss := range v {
					set[molecule[:i]+ss+molecule[i+l:]] = struct{}{}
				}
			}
		}
	}

	return fmt.Sprintf("%d", len(set))
}

type sortOptions [][2]string

var _ sort.Interface = (sortOptions)(nil)

func (s sortOptions) Len() int           { return len(s) }
func (s sortOptions) Less(i, j int) bool { return s[i][1] > s[j][1] }
func (s sortOptions) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Implement Solution to Problem 2
func solve2(input string) string {
	var molecule string
	var optionsDone bool
	aoc.MapLines(input, func(line string) error {
		if optionsDone {
			molecule = line
		} else if line == "" {
			optionsDone = true
		}
		return nil
	})

	// so we need to count all the Elements
	total := 0
	// and the number of Rn,Ar,Y
	Rn, Ar, Y := 0, 0, 0

	re := regexp.MustCompile("[A-Z][a-z]?")
	for _, m := range re.FindAllString(molecule, -1) {
		total++
		switch m {
		case "Rn":
			Rn++
		case "Ar":
			Ar++
		case "Y":
			Y++
		}
	}

	// now the formula
	steps := total - Rn - Ar - 2*Y - 1

	return fmt.Sprintf("%d", steps)
}
