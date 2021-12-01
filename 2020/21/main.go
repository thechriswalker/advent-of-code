package main

import (
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 21, solve1, solve2)
}

type food struct {
	ingredients []string
	allergens   []string
}

func parseInput(in string) []*food {
	stack := []*food{}
	for _, line := range strings.Split(in, "\n") {
		if line == "" {
			continue
		}

		f := &food{
			ingredients: []string{},
			allergens:   []string{},
		}
		fields := strings.Fields(line)
		before := true
		for _, field := range fields {
			if field == "(contains" {
				before = false
				continue
			}
			if before {
				f.ingredients = append(f.ingredients, field)
			} else {
				f.allergens = append(f.allergens, strings.TrimRight(field, ",)"))
			}
		}
		stack = append(stack, f)
	}
	return stack
}

// Implement Solution to Problem 1
func solve1(input string) string {
	list := parseInput(input)

	// now we need to convert this to a set of allergens and possibilities
	// the repeatedly remove the "know" mappings (i.e only 1 option) from the rest of the set.
	m := map[string]map[string]struct{}{}
	for _, f := range list {
		for _, allergen := range f.allergens {
			s := m[allergen]
			if s == nil {
				s = map[string]struct{}{}
				m[allergen] = s
			}
			for _, i := range f.ingredients {
				s[i] = struct{}{}
			}
		}
	}

	// now the recursive bit
	// the final map holds the mappings we have extracted
	// ingredient => allergen
	final := map[string]string{}
	for {
		found := false
		for all, set := range m {
			// remove anything in the final.
			for k := range final {
				delete(set, k)
			}
			if len(set) == 1 {
				found = true
				// delete from the current map.
				delete(m, all)
				// add to the final
				final[only(set)] = all
			}
		}
		if !found {
			// we should be done.
			break
		}
	}
	// now the result is the "number of times" the ingredients with NO allergens appear.
	// we have identified the ingredients WITH allergens so the remainder have NOw
	return ""
}

// get the single key from a set
func only(m map[string]struct{}) string {
	for k := range m {
		return k
	}
	return ""
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return "<unsolved>"
}
