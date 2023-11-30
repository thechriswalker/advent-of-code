package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 9, solve1, solve2)
}

func parseGroups(input string) (score, garbage int) {
	top := &Group{
		inner:  []*Group{},
		parent: nil,
	}
	curr := top
	// we will assume the first char is `{` and start a 1
	pc := 1
	inGarbage := false
	garbage = 0
outside:
	for {
		c := input[pc]
		pc++
		switch c {
		case '{': // start of subgroup unless in garbage
			if inGarbage {
				garbage++
			} else {
				sub := &Group{inner: []*Group{}, parent: curr}
				curr.inner = append(curr.inner, sub)
				curr = sub
			}
		case '}': // end of group (unless in garbage)
			if inGarbage {
				garbage++
			} else {
				curr = curr.parent
				if curr == nil {
					break outside
				}
			}
		case '!': // whatever the next char is, ignore it,
			pc++
		case '<': // entering garbage
			if inGarbage {
				garbage++
			} else {
				inGarbage = true
			}
		case '>': // leaving garbage
			inGarbage = false
		default:
			if inGarbage {
				garbage++
			}
		}
	}
	return top.Score(1), garbage
}

// Implement Solution to Problem 1
func solve1(input string) string {
	score, _ := parseGroups(input)
	return fmt.Sprint(score)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	_, garbage := parseGroups(input)
	return fmt.Sprint(garbage)
}

type Group struct {
	inner  []*Group
	parent *Group
}

func (g Group) Score(depth int) int {
	score := depth
	for _, gg := range g.inner {
		score += gg.Score(depth + 1)
	}
	return score
}
