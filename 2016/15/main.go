package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 15, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	discs := parseInput(input)
	time := 0
	for {
		if Attempt(discs, time) {
			return fmt.Sprintf("%d", time)
		}
		time++
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	discs := parseInput(input)
	// add the extra disc
	discs = append(discs, Disc{11, 0})
	time := 0
	for {
		if Attempt(discs, time) {
			return fmt.Sprintf("%d", time)
		}
		time++
	}
}

func parseInput(input string) []Disc {
	s := strings.NewReader(input)
	var err error
	var id, count, pos int
	discs := []Disc{}
	for {
		_, err = fmt.Fscanf(s, "Disc #%d has %d positions; at time=0, it is at position %d.\n", &id, &count, &pos)
		if err != nil {
			break
		}
		discs = append(discs, Disc{count, pos})
	}
	return discs
}

// 0 is number of positions, 1 is initial position
type Disc [2]int

func (d Disc) InRightPosition(t int) bool {
	return (d[1]+t)%d[0] == 0
}

func Attempt(discs []Disc, start int) bool {
	for i := 0; i < len(discs); i++ {
		if !discs[i].InRightPosition(start + i + 1) {
			return false
		}
	}
	return true
}
