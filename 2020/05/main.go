package main

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 5, solve1, solve2)
}

// FBFFBFF RLL
type boardingPass struct {
	row, col int
	id       int
}

func parseBoardingPass(s string) *boardingPass {
	if len(s) < 10 {
		return nil
	}
	rmin, rmax := 0, 127
	cmin, cmax := 0, 7

	for _, c := range s[0:7] {
		h := (1 + rmax - rmin) / 2
		switch c {
		case 'F':
			// forward, lower half
			rmax = rmin + h - 1
		case 'B':
			// backward, upper half
			rmin = rmin + h
		default:
			log.Fatalln("Bad character in row:", s)
		}
		//log.Printf("max:%d, min:%d, c:%c\n", rmax, rmin, c)
	}
	for _, c := range s[7:10] {
		h := (1 + cmax - cmin) / 2
		switch c {
		case 'L':
			// forward, lower half
			cmax = cmin + h - 1
		case 'R':
			// backward, upper half
			cmin = cmin + h
		default:
			log.Fatalln("Bad character in column:", s)
		}
		//log.Printf("max:%d, min:%d, c:%c\n", cmax, cmin, c)
	}
	id := rmax*8 + cmax
	//log.Printf("pass:%s row:%d col:%d id:%d\n", s, rmax, cmax, id)
	return &boardingPass{
		row: rmax,
		col: cmax,
		id:  id,
	}
}

func parsePasses(in string) []*boardingPass {
	stack := make([]*boardingPass, 0, 500)
	lines := strings.Split(in, "\n")
	for _, line := range lines {
		b := parseBoardingPass(line)
		if b != nil {
			stack = append(stack, b)
		}
	}
	return stack
}

// Implement Solution to Problem 1
func solve1(input string) string {
	passes := parsePasses(input)
	max := 0
	for _, bp := range passes {
		if max < bp.id {
			max = bp.id
		}
	}
	return fmt.Sprintf("%d", max)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	passes := parsePasses(input)
	// we should sort then discard row 0
	// ID < 8 and row 127, ID >= 127*8
	// then we need to find a missing id.
	sort.Sort(sorter(passes))
	for i := 0; i < len(passes)-1; i++ {
		p := passes[i]
		if p.id < 8 || p.id > 127*8 {
			continue
		}
		if passes[i+1].id == p.id+2 {
			// ours is the middle
			return fmt.Sprintf("%d", p.id+1)
		}
	}
	return "<unsolved>"
}

type sorter []*boardingPass

var _ sort.Interface = sorter(nil)

func (s sorter) Len() int           { return len(s) }
func (s sorter) Less(i, j int) bool { return s[i].id < s[j].id }
func (s sorter) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
