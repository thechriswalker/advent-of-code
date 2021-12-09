package main

import (
	"fmt"
	"math"
	"sort"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 17, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return solveSane(input, 150)
}

// type twoInts [][2]int

// var _ sort.Interface = (twoInts)(nil)

// func (ti twoInts) Len() int           { return len(ti) }
// func (ti twoInts) Less(i, j int) bool { return ti[i][0] < ti[j][0] }
// func (ti twoInts) Swap(i, j int)      { ti[i], ti[j] = ti[j], ti[i] }

// type Path struct {
// 	weight   int
// 	elements twoInts
// }

// func solveFor(input string, target int) string {
// 	ints := aoc.ToIntSlice(input, '\n')
// 	// sort them biggest first.
// 	sort.Sort(sort.Reverse(sort.IntSlice(ints)))

// 	// now we need to make a more explicit type for these ints.
// 	// with their INDEX included.
// 	containers := make(twoInts, len(ints))
// 	for i, n := range ints {
// 		containers[i] = [2]int{i, n}
// 	}

// 	fmt.Println(containers)
// 	// all combinations which add up to target
// 	combos := map[string]struct{}{}

// 	// we can do this the recursive breadth first way.
// 	// but we need to be able to uniquely identify the containers.
// 	// and so eliminate duplicates. They need an unique "ordering"
// 	// then we can do combinations, which is much less then permutations.
// 	var recur func(p *Path, rem twoInts)
// 	recur = func(p *Path, rem twoInts) {
// 		//fmt.Println("recur:", p, rem)
// 		if p.weight == target {
// 			// add to the list,
// 			sort.Sort(p.elements)
// 			combos[fmt.Sprint(p.elements)] = struct{}{}
// 			return
// 		}
// 		if p.weight > target {
// 			// nope
// 			return
// 		}
// 		// else try all the next possibilities
// 		for i, next := range rem {
// 			nextRem := make(twoInts, len(rem)-1)
// 			copy(nextRem, rem[:i])
// 			copy(nextRem[i:], rem[i+1:])

// 			nextPath := &Path{
// 				weight:   p.weight + next[1],
// 				elements: make(twoInts, len(p.elements)+1),
// 			}
// 			copy(nextPath.elements, p.elements)
// 			nextPath.elements[len(p.elements)] = next
// 			recur(nextPath, nextRem)
// 		}
// 	}
// 	recur(&Path{weight: 0, elements: twoInts{}}, containers)

// 	fmt.Println(combos)

// 	return fmt.Sprintf("%d", len(combos))
// }

// OK, so the naive method doesn't work. the permutations are too large.
// how do we make the problem space smaller.
// we only want combinations, so we can iterate only the "later" values,
// rather the the breadth first recursive search.
func solveSane(input string, target int) string {
	ints := aoc.ToIntSlice(input, '\n')
	// sort them biggest first.
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))

	var inner func(list []int, t int) int
	inner = func(list []int, t int) int {
		sum := 0
		for i, n := range list {
			switch {
			case n == t:
				sum++
			case n > t:
				// ignore
			case n < t:
				// recurse
				sum += inner(list[i+1:], t-n)
			}
		}
		return sum
	}

	return fmt.Sprintf("%d", inner(ints, target))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return solveSane2(input, 150)
}

func solveSane2(input string, target int) string {
	ints := aoc.ToIntSlice(input, '\n')
	// sort them biggest first.
	sort.Sort(sort.Reverse(sort.IntSlice(ints)))

	// we need to keep track of the number of containers required
	// in each solution, and the minimum number
	min := math.MaxInt64
	counts := map[int]int{}

	var inner func(d int, list []int, t int)
	inner = func(d int, list []int, t int) {
		d++
		for i, n := range list {
			switch {
			case n == t:
				//update the count
				counts[d]++
				// and maybe the min
				if d < min {
					min = d
				}
			case n > t:
				// ignore
			case n < t:
				// recurse
				inner(d, list[i+1:], t-n)
			}
		}
	}

	inner(0, ints, target)

	return fmt.Sprintf("%d", counts[min])
}
