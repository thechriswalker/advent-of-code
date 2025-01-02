package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 2, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	safe := 0

	aoc.MapLines(input, func(line string) error {
		ints := aoc.ToIntSlice(line, ' ')

		isIncreasing := ints[0] < ints[1]
		for i := 0; i < len(ints)-1; i++ {
			a, b := ints[i], ints[i+1]

			diff := b - a

			if diff == 0 {
				// stop early if we find a duplicate, not-increasing or decreasing
				return nil
			}

			if isIncreasing && diff < 0 {
				// not increasing when it should be
				return nil
			}
			if !isIncreasing && diff > 0 {
				// not decreasing when it should be
				return nil
			}
			if diff > 3 || diff < -3 {
				// difference is too big
				return nil
			}
		}
		// all good
		safe++
		return nil
	})

	return fmt.Sprint(safe)
}

func isSafe(a, b int, isIncreasing bool) bool {
	diff := b - a
	if diff == 0 {
		return false
	}
	if isIncreasing && diff < 0 {
		return false
	}
	if !isIncreasing && diff > 0 {
		return false
	}
	if diff > 3 || diff < -3 {
		return false
	}
	return true
}

// Implement Solution to Problem 2
func solve2(input string) string {
	safe := 0

	// var checkInts func(ints []int, dampened bool) bool
	// checkInts = func(ints []int, dampened bool) bool {
	// 	isIncreasing := ints[0] < ints[1]
	// 	for i := 0; i < len(ints)-1; i++ {
	// 		a, b := ints[i], ints[i+1]
	// 		if !isSafe(a, b, isIncreasing) {
	// 			if !dampened {
	// 				if i == len(ints)-1 {
	// 					// we haven't dampened and it's only the final element left.
	// 					// so we are OK without that one element. just do it!
	// 					return true
	// 				}
	// 				// 2 options, either we remove a or b
	// 				// to remove b we can continue with just a slice
	// 				// to remove a, we need a new slice.
	// 				// but it's probably easier (if not as effcient) to just rebuild the slices.
	// 				option1 := make([]int, len(ints)-1)
	// 				// skip a (i.e. i)
	// 				copy(option1, ints[:i])
	// 				copy(option1[i:], ints[i+1:])
	// 				// option2 is is the b removed
	// 				option2 := make([]int, len(ints)-1)
	// 				// skip b (i.e. i+1)
	// 				copy(option2, ints[:i+1])
	// 				copy(option2[i+1:], ints[i+2:])

	// 				return checkInts(option1, true) || checkInts(option2, true)
	// 			}
	// 			return false
	// 		}
	// 	}
	// 	return true
	// }

	// screw it there's something wrong with that logic. lets just brute force it. our inputs are small enough.
	// we will do a simple check, then one with a skipped element, until we run out or find a safe solution...
	checkWithSkip := func(ints []int, skip int) bool {
		if skip > -1 && skip <= len(ints) {
			slice := make([]int, len(ints)-1)
			copy(slice, ints[:skip])
			copy(slice[skip:], ints[skip+1:])
			ints = slice
		}
		//fmt.Println("check with skip", skip, ints)

		isIncreasing := ints[0] < ints[1]
		for i := 0; i < len(ints)-1; i++ {
			if !isSafe(ints[i], ints[i+1], isIncreasing) {
				return false
			}
		}
		return true
	}
	checkInts := func(ints []int, dampened bool) bool {
		for i := -1; i < len(ints); i++ {
			if checkWithSkip(ints, i) {
				return true
			}
		}
		return false
	}

	aoc.MapLines(input, func(line string) error {
		ints := aoc.ToIntSlice(line, ' ')
		if checkInts(ints, false) {
			// all good
			safe++
			// 344 too low
		}
		return nil
	})

	return fmt.Sprint(safe)
}
