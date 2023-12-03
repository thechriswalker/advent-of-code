package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 3, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {

	// build a grid
	g := aoc.CreateFixedByteGridFromString(input, '.')
	// we will iterate the grid finding numbers.
	// for each "bit" of a number, we will check for adjacent symbols.
	// if we get to the end of a number without finding a symbol, it remains
	// invalid
	current := []int{}
	valid := false
	// as we find a valid number, we will add it to the sum
	sum := 0
	// keep track of the row we are on
	row := 0

	addNumber := func(c []int) {
		//fmt.Println("PART NUMBER: ", c)
		x := 1
		for i := len(c) - 1; i >= 0; i-- {
			sum += x * c[i]
			x *= 10
		}
	}

	isSymbol := func(x, y int) bool {
		b, _ := g.At(x, y)
		switch b {
		// these are the only "known good" elements.
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.':
			return false
		}
		return true
	}

	hasSurroundingSymbol := func(x, y int) bool {
		return isSymbol(x-1, y-1) || isSymbol(x, y-1) || isSymbol(x+1, y-1) ||
			isSymbol(x-1, y) || isSymbol(x+1, y) ||
			isSymbol(x-1, y+1) || isSymbol(x, y+1) || isSymbol(x+1, y+1)
	}

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		// if on a new row, we may need to add the last number.
		if row != y {
			if valid && len(current) > 0 {
				addNumber(current)
			}
			current = []int{}
			valid = false
			row = y
		}
		switch b {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			// add to the stack and check for part number
			if !valid && hasSurroundingSymbol(x, y) {
				valid = true
			}
			// add to stack
			current = append(current, int(b-'0'))
		default:
			// this could be after a number?
			if len(current) > 0 && valid {
				addNumber(current)
			}
			current = []int{}
			valid = false
		}
	})
	// 5776988 too high
	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// build a grid
	g := aoc.CreateFixedByteGridFromString(input, '.')
	// this time we will find the gears first.

	isNum := func(x, y int) bool {
		c, _ := g.At(x, y)
		return c >= '0' && c <= '9'
	}

	findNumAround := func(x, y int) int {
		// move to far right of number.
		for {
			c, oob := g.At(x+1, y)
			if oob {
				break
			}
			if c < '0' || c > '9' {
				break
			}
			x++
		}
		// now find work backwards adding numbers until we hit a non-number;
		n := 0
		m := 1
		for {
			c, _ := g.At(x, y)
			if c < '0' || c > '9' {
				break
			}
			// other add the number
			n += m * int(c-'0')
			m *= 10
			x--
		}
		return n
	}

	checkGear := func(x, y int) (ok bool, ratio int) {
		// first check the gear
		if c, _ := g.At(x, y); c != '*' {
			return false, 0
		}
		// OK, it might be a gear.
		// we can rule out some easily by
		// checking left and right first
		surrounding := 0
		nums := []int{}

		hasLeft := isNum(x-1, y)
		if hasLeft {
			surrounding++
			nums = append(nums, findNumAround(x-1, y))
		}
		hasRight := isNum(x+1, y)
		if hasRight {
			surrounding++
			nums = append(nums, findNumAround(x+1, y))
		}
		// there could be 2 above, or just one.
		hasAboveLeft := isNum(x-1, y-1)
		if hasAboveLeft {
			surrounding++
			if surrounding == 3 {
				return false, 0
			}
			nums = append(nums, findNumAround(x-1, y-1))
		}
		hasAboveCenter := isNum(x, y-1)
		if hasAboveCenter && !hasAboveLeft {
			surrounding++
			if surrounding == 3 {
				return false, 0
			}
			nums = append(nums, findNumAround(x, y-1))
		}
		hasAboveRight := isNum(x+1, y-1)
		if hasAboveRight && !hasAboveCenter {
			surrounding++
			if surrounding == 3 {
				return false, 0
			}
			nums = append(nums, findNumAround(x+1, y-1))
		}

		// bottom
		hasBelowLeft := isNum(x-1, y+1)
		if hasBelowLeft {
			surrounding++
			if surrounding == 3 {
				return false, 0
			}
			nums = append(nums, findNumAround(x-1, y+1))
		}
		hasBelowCenter := isNum(x, y+1)
		if hasBelowCenter && !hasBelowLeft {
			surrounding++
			if surrounding == 3 {
				return false, 0
			}
			nums = append(nums, findNumAround(x, y+1))
		}
		hasBelowRight := isNum(x+1, y+1)
		if hasBelowRight && !hasBelowCenter {
			surrounding++
			if surrounding == 3 {
				return false, 0
			}
			nums = append(nums, findNumAround(x+1, y+1))
		}
		// OK, there are at most 2, and they should be in `nums`
		if len(nums) != 2 {
			return false, 0
		}
		return true, nums[0] * nums[1]
	}

	sum := 0

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		if ok, r := checkGear(x, y); ok {
			sum += r
		}
	})
	return fmt.Sprint(sum)
}
