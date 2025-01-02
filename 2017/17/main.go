package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 17, solve1, solve2)
}

// screw it lets us a linked list...
type Node struct {
	Value int
	Next  *Node
}

// Implement Solution to Problem 1
func solve1(input string) string {
	zero := &Node{Value: 0}
	zero.Next = zero

	current := zero
	steps, _ := strconv.Atoi(strings.TrimSpace(input))

	//fmt.Println("steps?", steps)

	for i := range 2017 {
		for j := 0; j < steps; j++ {
			current = current.Next
		}
		newNode := &Node{Value: i + 1, Next: current.Next}
		current.Next = newNode
		current = newNode
		// if i < 10 {
		// 	printLine(zero, current)
		// }
	}

	return strconv.Itoa(current.Next.Value)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// find the value after 0 after 50,000,000 insertions
	// I can probably do this the same way? the slow way...
	// OK the slow way takes 1.5 minutes to run... so I need to find a faster way.
	//return solve2slow(input)

	// let's think instead.
	// the value that is inserted after 0, means that ("current_position" + "step") % length == 0
	// so we only need to keep track of the current position and current length of list.
	// that's better only 190ms
	return solve2fast(input)
}

func solve2fast(input string) string {
	steps, _ := strconv.Atoi(strings.TrimSpace(input))
	current_position := 0
	after_zero := 0
	for i := range 50_000_000 {
		current_position = (current_position + steps) % (i + 1)
		if current_position == 0 {
			after_zero = i + 1
		}
		current_position++
	}
	return strconv.Itoa(after_zero)
}

var _ = solve2slow

func solve2slow(input string) string {
	zero := &Node{Value: 0}
	zero.Next = zero

	current := zero
	steps, _ := strconv.Atoi(strings.TrimSpace(input))

	//fmt.Println("steps?", steps)

	for i := range 50_000_000 {
		for j := 0; j < steps; j++ {
			current = current.Next
		}
		newNode := &Node{Value: i + 1, Next: current.Next}
		current.Next = newNode
		current = newNode
		// if i < 10 {
		// 	printLine(zero, current)
		// }
	}

	return strconv.Itoa(zero.Next.Value)
}

var _ = printLine

func printLine(zero, curr *Node) {
	for n := zero; n != nil; n = n.Next {
		if n == curr {
			fmt.Printf("(%d) ", n.Value)
		} else {
			fmt.Printf("%d ", n.Value)
		}
		if n.Next == zero {
			break
		}
	}
	fmt.Println()
}
