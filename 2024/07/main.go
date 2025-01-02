package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 7, solve1, solve2)
}

type eq struct {
	target int
	nums   []int
}

// Implement Solution to Problem 1
func solve1(input string) string {
	eqs := []eq{}
	aoc.MapLines(input, func(s string) error {
		t, ns, _ := strings.Cut(s, ": ")
		target, _ := strconv.Atoi(t)
		nums := aoc.ToIntSlice(ns, ' ')

		eqs = append(eqs, eq{target, nums})
		return nil
	})

	sum := 0
	for _, e := range eqs {
		if checkEquation(e.target, e.nums, false) {
			sum += e.target
		}
	}

	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	eqs := []eq{}
	aoc.MapLines(input, func(s string) error {
		t, ns, _ := strings.Cut(s, ": ")
		target, _ := strconv.Atoi(t)
		nums := aoc.ToIntSlice(ns, ' ')

		eqs = append(eqs, eq{target, nums})
		return nil
	})

	sum := 0
	for _, e := range eqs {
		if checkEquation(e.target, e.nums, true) {
			sum += e.target
		}
	}

	return fmt.Sprint(sum)
}

func checkEquation(target int, nums []int, allowConcat bool) bool {
	// going left-to-right, we can "multiply" or "add" in order.
	// can we hit the target?
	// we have no zeros or negative numbers so exceeding the target at any point means we can't hit it.

	// we will do a breath first search of the tree of possibilities, pruning branches as we go.

	// possible current sums.
	curr := []int{nums[0]}
	// remaining numbers
	next := nums[1:]

	for {
		if len(next) == 0 {
			// do we have a "target" in curr
			for _, x := range curr {
				if x == target {
					return true
				}
			}
			return false
		}

		newCurr := []int{}

		for _, x := range curr {
			// attempt +
			if n := x + next[0]; n <= target {
				newCurr = append(newCurr, n)
			}
			if n := x * next[0]; n <= target {
				newCurr = append(newCurr, n)
			}
			if allowConcat {
				if n := concat(x, next[0]); n <= target {
					newCurr = append(newCurr, n)
				}
			}
		}

		if len(newCurr) == 0 {
			return false
		}
		curr = newCurr
		next = next[1:]
	}
}

func concat(l, r int) int {
	// there is probably a good mathsy way to do this, but I'm just going to convert to strings and back.
	c, _ := strconv.Atoi(fmt.Sprintf("%d%d", l, r))
	return c
}
