package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 20, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	list := parseInput(input)
	// find the lowest unblocked ip.
	// they are sorted by start port.
	// so the highest blocked port is the high value
	// we iterate the list until the next low port
	// leaves a gap from the current highest port
	// /	fmt.Println(list)
	var high uint32
	for _, r := range list {
		if r[0] > high+1 {
			// we found it
			break
		}
		if high < r[1] {
			high = r[1]
		}
	}
	return fmt.Sprintf("%d", high+1)
}

var MaxIP = uint32(4294967295)

// Implement Solution to Problem 2
func solve2(input string) string {
	list := parseInput(input)
	allowed := uint64(0)
	var high uint32
	for _, r := range list {
		//fmt.Println("Range", r, "CurrentHigh", high)
		if r[0] > high && r[0]-high > 1 {
			// we found a gap.
			//	fmt.Println("Gap from", uint64(high), "to", uint64(r[0]), "which is", uint64(r[0])-uint64(high)-1, "addresses")
			allowed += uint64(r[0]) - uint64(high) - 1
		}
		if high < r[1] {
			high = r[1]
		}
	}
	// add any leftover at the end (no -1 because we want the final one)
	if high != MaxIP {
		allowed += uint64(MaxIP) - uint64(high)
	}
	return fmt.Sprintf("%d", allowed)
}

type List [][2]uint32

func (l List) Len() int           { return len(l) }
func (l List) Swap(i, j int)      { l[i], l[j] = l[j], l[i] }
func (l List) Less(i, j int) bool { return l[i][0] < l[j][0] }

func parseInput(input string) List {
	r := strings.NewReader(input)
	l := List{}
	var s, f uint32
	var err error
	for {
		if _, err = fmt.Fscanf(r, "%d-%d\n", &s, &f); err != nil {
			break
		}
		l = append(l, [2]uint32{s, f})
	}
	// this is important for our algorithm.
	// have the list sorted by start port.
	sort.Sort(l)
	return l
}
