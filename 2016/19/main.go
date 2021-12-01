package main

import (
	"fmt"
	"math/bits"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 19, solve1, solve2)
}

var ShouldPrint = false

// Implement Solution to Problem 1
func solve1(input string) string {
	var n int
	fmt.Sscanf(input, "%d", &n)
	return fmt.Sprintf("%d", formula(n))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	var n int
	fmt.Sscanf(input, "%d", &n)
	return fmt.Sprintf("%d", Solution3(n))
}

func formula(n int) int {
	// find the power of 2
	x := uint(62 - bits.LeadingZeros(uint(n)))
	return (n-(2<<x))*2 + 1
}

// let's print out some naive patterns to help work out what the pattern is
func NaiveSolution(n int, print bool) int {
	skipList := map[int]struct{}{}

	remaining := n
	iteration := 0
	// print initial line/ heading
	if print {
		fmt.Print("iter | ")
		for i := 1; i <= n; i++ {
			fmt.Print(i / 10)
		}
		fmt.Print(" | remaining\n")
	}
	var pendingSkip bool
	var final int
	for remaining > 0 {
		// print ln
		if print {
			fmt.Printf("%04d | ", iteration)
		}
		for i := 1; i <= n; i++ {
			if _, skip := skipList[i]; skip {
				if print {
					fmt.Print(" ")
				}
			} else {
				if print {
					fmt.Print(i % 10)
				}
				// should we skip this next time?
				if remaining == 1 || pendingSkip {
					// yes
					skipList[i] = struct{}{}
					pendingSkip = false
					if remaining == 1 {
						final = i
					}
				} else {
					// no
					pendingSkip = true
				}
			}
		}
		if print {
			fmt.Printf(" | %d\n", remaining)
		}
		iteration++
		remaining = n - len(skipList)
	}
	return final
}

func NaiveSolution2(n int) int {
	skipList := map[int]struct{}{}
	if n == 1 {
		return 1
	}
	remaining := n
	// print initial line/ heading
	index := 0
	for {
		// from this index who do we skip.
		// we skip index + remaining/2 elves, not in the skip list.
		stealmoves := index + remaining/2
		stealindex := index
		//fmt.Println("on index", index+1)
		for stealmoves > 0 {
			stealindex = (stealindex + 1) % n
			if _, ok := skipList[stealindex]; !ok {
				stealmoves--
			}
		}
		// steal this!
		//	fmt.Println("steal from", stealindex+1)
		skipList[stealindex] = struct{}{}
		remaining--
		//fmt.Println("remaining", remaining)
		if remaining == 1 {
			return index + 1
		}
		for {
			index = (index + 1) % n
			if _, ok := skipList[index]; !ok {
				// not on skip list
				break
			}
		}
	}
}

// better to make a circular linked list...
type Elf struct {
	Prev *Elf
	Next *Elf
	Pos  int
}

func Solution3(n int) int {
	start := &Elf{Pos: 1}
	prev := start
	var steal *Elf
	for x := 2; x <= n; x++ {
		curr := &Elf{
			Pos:  x,
			Prev: prev,
		}
		prev.Next = curr
		prev = curr
		if x == (n+1)/2 {
			// thisd is the first victim
			steal = curr
		}
	}
	// now link the head and tail
	prev.Next = start
	start.Prev = prev

	// now iterate until just one left.
	curr := start
	remaining := n

	for {
		// now steal from target
		//fmt.Println("current", curr.Pos, "steals from", steal.Pos)
		steal.Prev.Next, steal.Next.Prev = steal.Next, steal.Prev
		curr = curr.Next

		//	fmt.Println("remaining", remaining)
		// change victim
		steal = steal.Next
		if remaining%2 == 1 {
			// move again
			steal = steal.Next
		}
		remaining--
		if remaining == 1 {
			return curr.Pos
		}

	}
}
