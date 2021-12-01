package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2018, 2, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	rd := strings.NewReader(input)
	twos := 0
	threes := 0
	var s string
	for {
		if n, _ := fmt.Fscanln(rd, &s); n != 1 {
			break
		}
		r := NewRunes(s)
		has2 := false
		has3 := false
		for _, count := range r.Map {
			if count == 2 {
				has2 = true
			}
			if count == 3 {
				has3 = true
			}
		}
		if has2 {
			twos++
		}
		if has3 {
			threes++
		}
	}
	return fmt.Sprintf("%d", twos*threes)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	rd := strings.NewReader(input)
	list := []string{}
	var s string
	for {
		if n, _ := fmt.Fscanln(rd, &s); n != 1 {
			break
		}
		list = append(list, s)
	}
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			diffs := 0
			//fmt.Printf("checking:\n>>> %s\n>>> %s\n", list[i], list[j])
			for k := range list[i] {
				if list[i][k] != list[j][k] {
					diffs++
					if diffs > 1 {
						break
					}
				}
			}
			//fmt.Printf("differences: %d\n", diffs)
			if diffs == 1 {
				// found i,j make the string with the common letter.
				b := strings.Builder{}
				for k, c := range list[i] {
					if c == rune(list[j][k]) {
						b.WriteRune(c)
					}
				}
				return b.String()
			}
		}
	}

	return "<unsolved>"
}

func NewRunes(s string) *Runes {
	r := &Runes{
		List: []rune{},
		Map:  map[rune]int{},
	}
	for _, c := range s {
		r.Add(c)
	}
	return r
}

type Runes struct {
	List []rune
	Map  map[rune]int
}

func (r *Runes) Add(c rune) {
	n, ok := r.Map[c]
	r.Map[c] = n + 1
	if !ok {
		r.List = append(r.List, c)
	}

}
