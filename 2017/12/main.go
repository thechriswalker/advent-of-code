package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 12, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	m := map[int][]int{}
	aoc.MapLines(input, func(line string) error {
		id, links, _ := strings.Cut(line, " <-> ")
		pid, _ := strconv.Atoi(id)
		connected := aoc.ToIntSlice(links, ',')
		m[pid] = connected
		return nil
	})

	// how many programs are in the group that contains program ID 0?
	containsZero := map[int]bool{}
	containsZero[0] = true
	changed := true
	for changed {
		changed = false
		for pid, connected := range m {
			if containsZero[pid] {
				// we already know this one is in the group
				for _, c := range connected {
					if !containsZero[c] {
						containsZero[c] = true
						changed = true
					}
				}
			}
		}
	}

	return fmt.Sprint(len(containsZero))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	m := map[int][]int{}
	aoc.MapLines(input, func(line string) error {
		id, links, _ := strings.Cut(line, " <-> ")
		pid, _ := strconv.Atoi(id)
		connected := aoc.ToIntSlice(links, ',')
		m[pid] = connected
		return nil
	})

	// how many programs are in the group that contains program ID 0?
	groups := map[int]map[int]bool{}
	// pids we have already assigned to a group
	groupMembers := map[int]struct{}{}

	enumerateGroup := func(g int) {
		changed := true
		group := map[int]bool{
			g: true,
		}
		groupMembers[g] = struct{}{}
		groups[g] = group
		for changed {
			changed = false
			for pid, connected := range m {
				if group[pid] {
					// we already know this one is in the group
					for _, c := range connected {
						if !group[c] {
							group[c] = true
							groupMembers[c] = struct{}{}
							changed = true
						}
					}
				}
			}
		}
	}
	enumerateGroup(0)
	for pid := range m {
		if _, ok := groupMembers[pid]; !ok {
			enumerateGroup(pid)
		}
	}

	return fmt.Sprint(len(groups))

}
