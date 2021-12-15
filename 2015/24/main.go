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

type Group []int

func (g Group) QuantumEntanglement() int {
	q := 1
	for _, i := range g {
		q *= i
	}
	return q
}

func (g Group) Len() int {
	return len(g)
}

func makeGroups(l []int) []*Group {
	all := []*Group{}
	// they must all have the same weight and we must use them all.
	// so we just have to find groups of a specific wieght.
	sum := 0
	for _, i := range l {
		sum += i
	}
	// should be divisible by 3
	target := sum / 3
	fmt.Println("Sum is", sum, "target is", target)
	// now we create combinations that add up to target.
	// lets sort them to highest first.
	sort.Sort(sort.Reverse(sort.IntSlice(l)))
	// now let's try and make a
	//length := 1

	// for??

	return all
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// we need to find as many unique arrangements of
	// the packages into 3 groups.
	// then foreach group find the ones with the minimum number
	// of packages.
	// out of those, find the one with the smallest QE.

	packages := aoc.ToIntSlice(input, '\n')

	groups := makeGroups(packages)

	min := math.MaxInt64
	var minGroups []*Group
	for _, g := range groups {
		if min > g.Len() {
			min = g.Len()
			minGroups = []*Group{g}
		}
		if min == g.Len() {
			minGroups = append(minGroups, g)
		}
	}
	// sort by QE
	minQE := math.MaxInt64
	for _, g := range minGroups {
		qe := g.QuantumEntanglement()
		if minQE > qe {
			minQE = qe
		}
	}

	return fmt.Sprint(minQE)

}

// Implement Solution to Problem 2
func solve2(input string) string {
	return "<unsolved>"
}
