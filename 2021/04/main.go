package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 4, solve1, solve2)
}

func readNumbers(line string, sep rune) []int {
	sn := strings.FieldsFunc(line, func(r rune) bool {
		return r == sep
	})
	nn := make([]int, len(sn))
	for i, s := range sn {
		n, _ := strconv.Atoi(s)
		nn[i] = n
	}
	return nn
}

type BingoBoard struct {
	Nums   [25]int
	Marked [25]bool
	Won    bool
}

func (bb *BingoBoard) Mark(n int) {
	for i, m := range bb.Nums {
		if n == m {
			bb.Marked[i] = true
			bb.Check()
		}
	}
}

func (bb *BingoBoard) SumUnmarked() int {
	sum := 0
	for i, marked := range bb.Marked {
		if !marked {
			sum += bb.Nums[i]
		}
	}
	return sum
}

func (bb *BingoBoard) Check() {
	if bb.Won {
		return
	}
	// find 5 marked in a row
	// or 5 in a column.
	for i := 0; i < 5; i++ {
		//row is 5*i, 5*i + 1 ...
		row := 5 * i
		if bb.Marked[row] &&
			bb.Marked[row+1] &&
			bb.Marked[row+2] &&
			bb.Marked[row+3] &&
			bb.Marked[row+4] {
			bb.Won = true
			return
		}
		// column is i + 0*5, i+1*5,...
		if bb.Marked[i] &&
			bb.Marked[i+5] &&
			bb.Marked[i+10] &&
			bb.Marked[i+15] &&
			bb.Marked[i+20] {
			bb.Won = true
			return
		}
	}
	// nope
}

func (bb *BingoBoard) String() string {
	b := &strings.Builder{}
	// print in five rows
	// wrap "marked" elements in Bold
	for i := 0; i < 5; i++ {
		// each num
		for j := 0; j < 5; j++ {
			x := i*5 + j
			if bb.Marked[x] {
				fmt.Fprintf(b, "\x1b[93m%2d\x1b[0m ", bb.Nums[x])
			} else {
				fmt.Fprintf(b, "%2d ", bb.Nums[x])
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var caller []int
	boards := []*BingoBoard{}
	var curr *BingoBoard
	var row int
	aoc.MapLines(input, func(line string) error {
		if caller == nil {
			caller = readNumbers(line, ',')
			return nil
		}
		if line == "" {
			// next board
			if curr != nil {
				boards = append(boards, curr)
			}
			curr = &BingoBoard{}
			row = 0
			return nil
		}
		// a row of numbers.
		nums := readNumbers(line, ' ')
		copy(curr.Nums[row*5:], nums)
		row++
		return nil
	})

	if curr.Nums[0] != curr.Nums[1] {
		// we didn't add this board!
		boards = append(boards, curr)
	}

	fmt.Println("Found", len(boards), "boards")

	// OK we have the boards and the caller
	// we want to find the first board to "win"
	for _, n := range caller {
		for _, b := range boards {
			b.Mark(n)
			if b.Won {
				fmt.Println(b)
				// calculate sum of unmarked spaces,
				// multiply by n
				v := n * b.SumUnmarked()
				return fmt.Sprintf("%d", v)
			}
		}
	}
	return "<unsolved>"
}

// Implement Solution to Problem 2
func solve2(input string) string {
	var caller []int
	boards := []*BingoBoard{}
	var curr *BingoBoard
	var row int
	aoc.MapLines(input, func(line string) error {
		if caller == nil {
			caller = readNumbers(line, ',')
			return nil
		}
		if line == "" {
			// next board
			if curr != nil {
				boards = append(boards, curr)
			}
			curr = &BingoBoard{}
			row = 0
			return nil
		}
		// a row of numbers.
		nums := readNumbers(line, ' ')
		copy(curr.Nums[row*5:], nums)
		row++
		return nil
	})

	if curr.Nums[0] != curr.Nums[1] {
		// we didn't add this board!
		boards = append(boards, curr)
	}

	fmt.Println()
	// OK we have the boards and the caller
	// we want to find the last board to "win"
	var numWon int
	for _, n := range caller {
		for _, b := range boards {
			if b.Won {
				continue
			}
			b.Mark(n)
			if b.Won {
				numWon++
				if numWon == len(boards) {
					fmt.Println(b)
					// calculate sum of unmarked spaces,
					// multiply by n
					v := n * b.SumUnmarked()
					return fmt.Sprintf("%d", v)
				}
			}
		}
	}
	return "<unsolved>"
}
