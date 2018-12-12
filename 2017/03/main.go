package main

import (
	"fmt"

	"../../aoc"
)

func main() {
	aoc.Run(2017, 3, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var index int
	fmt.Sscanf(input, "%d", &index)
	x, y := coords(index)
	x = abs(x)
	y = abs(y)
	// mahattan distance to 0,0 is just the sum of abs values
	return fmt.Sprintf("%d", x+y)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// write values in a spiral until we reach a value larger than the input
	var target int
	fmt.Sscanf(input, "%d", &target)
	mem := SpiralMemory{}
	mem.Set(0, 0, 1) // set the initial value
	var x, y, value int
	index := 2
	for {
		x, y = coords(index)
		value = 0
		// add the surrounds
		value += mem.Get(x+1, y)
		value += mem.Get(x-1, y)
		value += mem.Get(x, y+1)
		value += mem.Get(x, y-1)
		value += mem.Get(x+1, y+1)
		value += mem.Get(x-1, y-1)
		value += mem.Get(x+1, y-1)
		value += mem.Get(x-1, y+1)

		if value > target {
			return fmt.Sprintf("%d", value)
		}
		// might as well bypass the set by coords and set directly
		// fmt.Println("Writing value", value, "at index", index, "(x:", x, "y:", y, ")")
		// fmt.Println(mem)
		// time.Sleep(time.Millisecond * 100)
		mem[index] = value
		index++
	}
}

// position 0,0 is the 1 index of the spiral pattern.
type SpiralMemory map[int]int

func (sm SpiralMemory) Get(x, y int) int {
	i := index(x, y)
	return sm[i]
}
func (sm SpiralMemory) Set(x, y, v int) {
	i := index(x, y)
	sm[i] = v
}

// this finds the coords on the plane starting 0,0 for 1, for a given index
func coords(index int) (x, y int) {
	if index == 1 {
		return 0, 0
	}
	// we work out the "square" of the spiral that we are on
	n := 3
	for ; n*n < index; n += 2 {
	}
	h := n - 1
	x, y = h/2, -1*h/2
	j := n*n - index // this is how many positions are on this square

	// 4 possibilities, depending on which side of the square we are on.
	switch {
	case j < h:
		// the bottom, y will be correct, we need to substract j from x though
		x -= j
	case j < 2*h:
		// on the left side
		x -= h
		y += j - h
	case j < 3*h:
		// the top side
		y += h
		x = x - h + (j - 2*h)
	default:
		// the right side
		y = y + h - (j - 3*h)
	}
	return
}

// this finds the index into the spiral for a given x,y
func index(x, y int) int {
	if x == 0 && y == 0 {
		return 1
	}
	// which square are we on squares start from the bottom right
	// so 2 <= n <= 9 is on square with bottom right 1,-1
	// a coord is on a the square with the maximum absolute coordinate value
	max := maxAbs(x, y)
	n := 1 + (max * 2)
	// we now have a lower and upper bound for the index
	// top is n*n 9,25,49,...
	// now we can work out the index by considering where this is on the square.
	switch {
	case y == -1*max:
		// the bottom row, include the final n numbers.
		// n*n is x = max and we work back to x = -max
		return n*n - max + x
	case x == -1*max:
		// we are on the left hand side (but not the bottom and the numbers run from
		// n*n - 2*max to n*n- 4*max
		return n*n - 3*max - y
	case y == max:
		// on the top, (but not the far left)
		// numbers are n*n - 5*max +/- max
		return n*n - 5*max - x
	case x == max:
		// on the right hand side (but not the top or bottom values)
		return n*n - 7*max + y
	default:
		panic("massive logic fail")
	}
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}
func maxAbs(x, y int) int {
	x = abs(x)
	y = abs(y)
	if x > y {
		return x
	}
	return y
}
