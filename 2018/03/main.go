package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2018, 3, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	claims := parseClaims(input)
	//fmt.Println(claims)
	overlaps := map[[2]uint16]uint8{}
	for i := 0; i < len(claims); i++ {
		for j := i + 1; j < len(claims); j++ {
			for _, p := range claims[i].Intersection(claims[j]) {
				overlaps[p] = 1
			}
		}
	}
	return fmt.Sprintf("%d", len(overlaps))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	claims := parseClaims(input)
	//fmt.Println(claims)
	for i := 0; i < len(claims); i++ {
		hasIntersection := false
		for j := 0; j < len(claims); j++ {
			if i != j && len(claims[i].Intersection(claims[j])) != 0 {
				hasIntersection = true
				break
			}
		}
		if !hasIntersection {
			return fmt.Sprintf("%d", claims[i].Id)
		}
	}
	return "<unsolved>"
}

type claim struct {
	Id, Left, Right, Top, Bottom uint16
}

// the points here are 0,0 = square bounded by 0,0 -> 1,1
func (c claim) Intersection(d claim) [][2]uint16 {
	l := max(c.Left, d.Left)
	r := min(c.Right, d.Right)
	t := max(c.Top, d.Top)
	b := min(c.Bottom, d.Bottom)
	//	fmt.Println("intersection:", c, d, " is ", l, t, r, b)
	out := [][2]uint16{}
	for i := l; i < r; i++ {
		for j := t; j < b; j++ {
			//		fmt.Println("overlap", i, j)
			out = append(out, [2]uint16{i, j})
		}
	}
	return out
}

func min(a, b uint16) uint16 {
	if a < b {
		return a
	}
	return b
}

func max(a, b uint16) uint16 {
	if a > b {
		return a
	}
	return b
}

func parseClaims(input string) []claim {
	var id, left, top, width, height uint16
	var err error
	rd := strings.NewReader(input)
	out := []claim{}
	for {
		if _, err = fmt.Fscanf(rd, "#%d @ %d,%d: %dx%d\n", &id, &left, &top, &width, &height); err != nil {
			break
		}
		out = append(out, claim{Id: id, Left: left, Top: top, Right: left + width, Bottom: top + height})
	}
	return out
}
