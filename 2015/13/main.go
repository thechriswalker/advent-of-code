package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 13, solve1, solve2)
}

type H map[string]int

func (h H) Key(p1, p2 string) string {
	if strings.Compare(p1, p2) < 0 {
		p1, p2 = p2, p1
	}
	return p1 + p2
}

func (h H) Add(p1, p2 string, dh int) {
	key := h.Key(p1, p2)
	h[key] += dh
}

func (h H) Get(p1, p2 string) int {
	return h[h.Key(p1, p2)]
}

type Path struct {
	Dh    int
	Order string
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return runSolution(input, false)
}
func runSolution(input string, withMe bool) string {
	// for each pair of names, we can create a
	// total happiness change for the pair to be together
	h := H{}
	nameSet := map[string]struct{}{}
	var p1, p2, gainOrLoss string
	var v int
	aoc.MapLines(input, func(line string) error {
		fmt.Sscanf(line[:len(line)-1], "%s would %s %d happiness units by sitting next to %s", &p1, &gainOrLoss, &v, &p2)
		if gainOrLoss == "lose" {
			v *= -1
		}
		h.Add(p1, p2, v)
		nameSet[p1] = struct{}{}
		nameSet[p2] = struct{}{}
		return nil
	})

	names := make([]string, 0, len(nameSet))
	for name := range nameSet {
		names = append(names, name)
	}
	sort.Strings(names)
	if withMe {
		names = append(names, "ME")
	}

	// now we need to sum all combinations. how many people do we have? can we just do it the naive way?
	// bnah, lets recurse so we don't have to calculate everything every time
	max := 0
	var recur func(p Path, rem []string, first string, prev string)
	recur = func(p Path, rem []string, first string, prev string) {
		//fmt.Println("recur:", p, rem)
		if len(rem) == 0 {
			// we have to add the "end" to "beginning" value
			p.Dh += h.Get(first, prev)
			p.Order += " -> " + first + " $"
			//fmt.Printf("Path[%d]: %s\n", p.Dh, p.Order)
			if max < p.Dh {
				max = p.Dh
			}
			return
		}
		// other iterate over each
		for i := range rem {
			next := rem[i]
			nextRem := make([]string, len(rem)-1)
			nextFirst := first
			nextPath := Path{Dh: p.Dh, Order: p.Order}
			copy(nextRem, rem[:i])
			copy(nextRem[i:], rem[i+1:])
			//fmt.Println("Picked", next, "left:", nextRem)
			if prev != "" {
				nextPath.Dh += h.Get(prev, next)
				nextPath.Order += " -> " + next
			} else {
				nextFirst = next
				nextPath.Order = next
			}
			recur(nextPath, nextRem, nextFirst, next)
		}
	}
	recur(Path{}, names, "", "")
	// max should have the value!
	return fmt.Sprintf("%d", max)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return runSolution(input, true)
}
