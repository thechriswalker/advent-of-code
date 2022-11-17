package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 24, solve1, solve2)
}

func quantumEntanglement(g []int) int {
	q := 1
	for _, i := range g {
		q *= i
	}
	return q
}

//func findGroupsOfWeight(l []int, x int) (groups [][]int, remaining []int)

// Implement Solution to Problem 1
func solve1(input string) string {
	return fmt.Sprint(solveN(input, 3))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return fmt.Sprint(solveN(input, 4))
}

func solveN(input string, groups int) int {
	packages := aoc.ToIntSlice(input, '\n')
	sort.Sort(sort.Reverse(sort.IntSlice(packages)))
	sum := 0
	for _, i := range packages {
		sum += i
	}
	// should be divisible by $groups
	target := sum / groups
	//fmt.Println("Sum is", sum, "target is", target)
	combos := CombinationsThatMatch(packages, target)
	//fmt.Println("combos:", combos[0])
	// find the "shortest" combos
	shortest := 10000000000
	sets := [][]int{}
	for _, c := range combos {
		if len(c) < shortest {
			shortest = len(c)
			sets = [][]int{c}
		} else if len(c) == shortest {
			sets = append(sets, c)
		}
	}
	// shortestSet is first
	l := len(sets[0])
	q := math.MaxInt
	for _, s := range sets {
		if len(s) > l {
			break
		}
		m := quantumEntanglement(s)
		if m < q {
			q = m
		}
	}
	// arrangements := makeArrangements(packages)

	// sort.Sort(arrangements)

	// q := quantumEntanglement(arrangements[0][0])

	return q
}

func CombinationsThatMatch(options []int, target int) [][]int {
	all := [][]int{}
	var recur func(a, b []int, sum int)
	recur = func(base, rem []int, sum int) {
		if sum == target {
			all = append(all, base)
			return
		}
		//log.Printf("base: %v, remaining: %v\n", base, rem)
		if len(base) == 6 {
			return
		}
		for i := range rem {
			nextSum := sum + rem[i]
			if nextSum > target {
				continue
			}
			nextBase := make([]int, len(base)+1)
			copy(nextBase, base)
			// pick i
			nextBase[len(base)] = rem[i]
			// create the next remainder.
			if len(rem) == 1 {
				recur(nextBase, []int{}, nextSum)
			} else {
				nextRem := make([]int, len(rem)-1)
				copy(nextRem, rem[:i])
				copy(nextRem[i:], rem[i+1:])
				recur(nextBase, nextRem, nextSum)
			}

		}
	}
	recur([]int{}, options, 0)
	return all
}
