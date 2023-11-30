package main

import (
	"fmt"
	"math/big"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 25, solve1, solve2)
}

var (
	seed       = big.NewInt(20151125)
	multiplier = big.NewInt(252533)
	modulo     = big.NewInt(33554393)
)

// the starting number in a row is calculable
// we can memoize this if needs be
func startOfRowX(row int) int {
	if row <= 0 {
		panic("bad row")
	}
	if row == 1 {
		return 1
	}
	return startOfRowX(row-1) + row - 1
}

// work out the number of the code in row X col Y
func rowXColY(row, col int) int {
	sum := startOfRowX(row)
	inc := row + 1
	fmt.Printf("start of row %d is %d, increments start at %d\n", row, sum, inc)
	for x := 2; x <= col; x++ {
		sum += inc
		inc++
	}
	return sum
}

// naive calculation of the code...
func calcCode(n int) uint64 {
	c := &big.Int{}
	c.Set(seed)
	for i := 1; i < n; i++ {
		c.Mul(c, multiplier)
		c.Mod(c, modulo)
	}
	return c.Uint64()
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// we need to parse out the row/col values
	var row, col int
	fmt.Sscanf(input, "To continue, please consult the code grid in the manual.  Enter the code at row %d, column %d.", &row, &col)

	fmt.Printf("finding row=%d col=%d\n", row, col)
	code_number := rowXColY(row, col)

	code := calcCode(code_number)

	return fmt.Sprintf("%d", code)
	// 27398364 too high
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return "<unsolved>"
}
