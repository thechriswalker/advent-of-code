package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/2017/10/knothash"
	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 10, solve1, solve2)
}

var ringSize = 256

// Implement Solution to Problem 1
func solve1(input string) string {
	lengths := aoc.ToIntSlice(input, ',')

	r := knothash.NewRing(ringSize)

	for _, l := range lengths {
		r.PinchAndTwist(l)
	}
	l0, l1 := r.List01()
	return fmt.Sprint(l0 * l1)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// now input is bytes
	// we should trim the string though
	return knothash.Hex(strings.TrimSpace(input))

}
