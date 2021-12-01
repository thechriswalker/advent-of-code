package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 9, solve1, solve2)
}

var (
	// this changes in testing.
	preambleLength = 25
)

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

func findBadNumber(nums []int) (int, bool) {
	for i := preambleLength; i < len(nums); i++ {
		// check that the number _is_ the sum of 2 different
		// numbers in the previous "preamble"
		// now the "sums" don't change much. we only lose those
		// from the dropped number and add those from the "added"
		// number. but unless this becomes prohibitive, let's do it
		// the naive way.
		if !isSumOfTwo(nums[i], nums[i-preambleLength:i]) {
			return nums[i], true
		}
	}
	return 0, false
}

// Implement Solution to Problem 1
func solve1(input string) string {
	nums := parseNumbers(input)

	// we only need to iterate from after the preamble
	// and sub-iterate over the previous "preamble"'s worth
	// of entries.
	if n, ok := findBadNumber(nums); ok {
		return fmt.Sprintf("%d", n)
	}
	return "<unsolved>"
}

// Implement Solution to Problem 2
func solve2(input string) string {
	nums := parseNumbers(input)
	target, ok := findBadNumber(nums)
	if !ok {
		panic("Could not solve part 1!")
	}

	// now we want to find a "contiguous" range of numbers that add up to the target.
	// I can think of two ways to do this.
	// 1. the naive way, iterate summing from the iteration point until it hits or exceeds the target.
	// 2. the clever way, do the first like 1. but then subtract the "first" number, and then substract
	//    repeatedly from the right until we hit or undercut. the add again and repeat. saving a bunch of
	//    addition. This will likely be faster, but more complicated. I'll try the naive way first...

	for i := 0; i < len(nums)-1; i++ {
		sum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			sum += nums[j]
			if sum == target {
				// we found it. sort and add the top and bottom
				//fmt.Printf("target:%d, i:%d, j:%d, slice:%v\n", target, i, j, nums[i:j+1])
				sort.Sort(sort.IntSlice(nums[i : j+1]))
				return fmt.Sprintf("%d", nums[i]+nums[j])
			}
			if sum > target {
				// didn't find it.
				break
			}
			// continue!
		}
	}
	return "<unsolved>"
}

func isSumOfTwo(target int, nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			//log.Printf("target:%d i:%d j:%d sum:%d + %d = %d", target, i, j, nums[i], nums[j], nums[i]+nums[j])
			if nums[i] != nums[j] && nums[i]+nums[j] == target {
				return true
			}
		}
	}
	return false
}
