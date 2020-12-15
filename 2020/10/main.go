package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2020, 10, solve1, solve2)
}

func parseNumbers(in string) []int {
	lines := strings.Split(in, "\n")
	stack := make([]int, 0, len(lines))
	for _, line := range lines {
		if n, err := strconv.Atoi(line); err == nil {
			stack = append(stack, n)
		}
	}
	return stack
}

// Implement Solution to Problem 1
func solve1(input string) string {
	nums := parseNumbers(input)
	sort.Sort(sort.IntSlice(nums))
	// starting from zero, iterate
	// and add up the "differences"
	diffs := map[int]int{}
	last := 0
	for _, n := range nums {
		diffs[n-last]++
		last = n
	}
	// now the final adapter last+3;
	diffs[3]++

	//log.Printf("%#v", diffs)

	return fmt.Sprintf("%d", diffs[1]*diffs[3])
}

// Implement Solution to Problem 2
func solve2(input string) string {
	nums := parseNumbers(input)
	sort.Sort(sort.IntSlice(nums))

	// the number of combinations is directly related to the "runs" of consecutive 1-gaps
	// the 3-gaps "must" be as is. but the 1-gaps have valid arrangements.
	// e.g.
	// let's think about how we can arrange gaps in our consecutive runs.
	// the beginning and end must have the adapters.

	// I think the pattern is (in binary for ease) with consecutive numbers:

	// n = length of run
	// x = number of combinations
	//
	// n = 1, x = 1 (1)
	// n = 2, x = 1 (only 11) as the first and last must be 1
	// n = 3, x = 2 (111, 101)
	// n = 4, x = 4 (1[n3] (2), 10[n2] (1), 100[n1] (1) => 4)
	// n = 5, x = n4 + n3 + n2 = 4+2+1 = 7
	// n = 6, x = n5 + n4 + n3 = 7 + 4 + 2 = 13
	//
	// 111111 1 + n5
	// 110111
	// 111011
	// 111101
	// 110011
	// 111001
	// 110101
	// 101111 10 + n4
	// 101011
	// 101101
	// 101001
	// 100111 100 + n3
	// 100101

	// n = k, x = n(k-1) + n(k-2) + n(k-3)
	// inductive proof!

	// so we can work out all the combinations of the runs, then we can have them in any combination
	// this basically means multiplying all the combinations.
	combos := uint64(1)

	currentRun := 1
	last := 0
	for _, n := range nums {
		if n == last+1 {
			currentRun++
		} else {
			// if currentRun <= i {
			// 	log.Printf("run of %d, %v\n", currentRun, nums[i-currentRun:i])
			// }
			// run finished!
			combos *= getCombosForSequence(currentRun)
			currentRun = 1
		}
		last = n
	}
	// the last is the max +3 so never a full run, so if we had a sequence we should
	// add it. i.e. 1111 (final).
	if currentRun > 0 {
		combos *= getCombosForSequence(currentRun)
	}
	return fmt.Sprintf("%d", combos)
}

// must have 3!
var seqCache = map[int]uint64{
	1: 1,
	2: 1,
	3: 2,
}

func getCombosForSequence(l int) uint64 {
	if c, ok := seqCache[l]; ok {
		return c
	}
	n := getCombosForSequence(l-1) + getCombosForSequence(l-2) + getCombosForSequence(l-3)
	//log.Printf("caching seq for n=%d combos=%d\n", l, n)
	seqCache[l] = n
	return n
}
