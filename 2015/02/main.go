package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 2, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	list := parseInput(input)
	sum := 0
	for _, p := range list {
		sum += p.RequiredPaper()
	}
	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	list := parseInput(input)
	sum := 0
	for _, p := range list {
		sum += p.RequiredRibbon()
	}
	return fmt.Sprint(sum)
}

type Present struct {
	w, h, l int
}

func (p *Present) RequiredPaper() int {
	x := p.w * p.h
	y := p.w * p.l
	z := p.h * p.l

	return 2*x + 2*y + 2*z + min(x, y, z)
}

func (p *Present) RequiredRibbon() int {
	bow := p.w * p.h * p.l
	x := min(p.w, p.h, p.l)
	y := mid(p.w, p.h, p.l)

	return 2*x + 2*y + bow
}

func min(x, y, z int) int {
	s := sort.IntSlice{x, y, z}
	sort.Sort(s)
	return s[0]
}

func mid(x, y, z int) int {
	s := sort.IntSlice{x, y, z}
	sort.Sort(s)
	return s[1]
}

func parseInput(input string) []Present {
	sc := bufio.NewScanner(strings.NewReader(input))
	var w, h, l int
	list := []Present{}
	for sc.Scan() {
		if n, _ := fmt.Sscanf(sc.Text(), "%dx%dx%d", &w, &h, &l); n == 3 {
			list = append(list, Present{w, h, l})
		}
	}
	return list
}
