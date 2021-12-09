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
	options := [][2]string{}
	var molecule, a, b string
	var optionsDone bool
	aoc.MapLines(input, func(line string) error {
		if optionsDone {
			molecule = line
		} else if line == "" {
			optionsDone = true
		} else {
			fmt.Sscanf(line, "%s => %s", &a, &b)
			options = append(options, [2]string{a, b})

		}
		return nil
	})

	// we also need to sort them to the longest first.
	sort.Sort(sortOptions(options))
	// we need a breadth first search now.
	// make each replacement one by one until we hit the
	// target molecule.
	//
	// Hmmm, this works for the small sample, but fails (too slow)
	// on the actual problem. There must be a trick to reduce the number
	// of iterations.
	// OK, none of the "replacements" are smaller than that they replace
	// so we can stop iterating if the new string is too long!
	//
	// also we should track all points we have been to before, so we don't
	// try them in later interations.

	// hmm, still too slow. what can we do next....
	// ditch this and go the other way.
	// starting with the molecule, can we get to e?
	depth := 0
	test := molecule
	found := false
	for test != "e" {
		// try a round of replacements
		found = false
		for _, option := range options {
			if s, ok := replaceNext(test, option[1], option[0]); ok {
				depth++
				test = s
				found = true
				//fmt.Println("next iteration:", s, "depth", depth)
				// if depth > 10 {
				// 	return ""
				// }
				break
			}
		}
		if !found {
			// ran out of options
			return fmt.Sprintf("no more options at depth %d, final: %s", depth, test)
		}
	}
	return fmt.Sprintf("%d", depth)
}

func replaceNext(m, find, replace string) (string, bool) {
	re := regexp.MustCompile(find)
	for _, match := range re.FindAllStringIndex(m, -1) {
		return m[:match[0]] + replace + m[match[1]:], true
	}
	return "", false
}
