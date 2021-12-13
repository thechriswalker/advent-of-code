package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 12, solve1, solve2)
}

type Cave struct {
	ID       string
	Big      bool
	Connects []*Cave
}

type Path struct {
	SmallCaveDoubleVisit bool
	Current              *Cave
	Previous             *Path
}

func (p *Path) NextPaths() []*Path {
	next := []*Path{}

	//rules for next path. Cannot be back to a Non-Big cave
	for _, c := range p.Current.Connects {
		if c.ID != "start" && (c.Big || p.NotVisited(c)) {
			next = append(next, &Path{Current: c, Previous: p})
		}
	}
	return next
}

func (p *Path) NextPaths2() []*Path {
	next := []*Path{}

	//rules for next path. Cannot be back to a Non-Big cave
	for _, c := range p.Current.Connects {
		if c.ID == "start" {
			continue
		}
		if c.ID == "end" || c.Big {
			next = append(next, &Path{
				SmallCaveDoubleVisit: p.SmallCaveDoubleVisit,
				Current:              c,
				Previous:             p,
			})
			continue
		}
		// now for a small cave, if we have done a double visit
		if p.NotVisited(c) {
			next = append(next, &Path{
				SmallCaveDoubleVisit: p.SmallCaveDoubleVisit,
				Current:              c,
				Previous:             p,
			})
			continue
		}

		if !p.SmallCaveDoubleVisit {
			// we haven't done a double visit yet, do it now.
			next = append(next, &Path{
				SmallCaveDoubleVisit: true,
				Current:              c,
				Previous:             p,
			})
			continue
		}
		// nope
	}
	return next
}

func (p *Path) NotVisited(c *Cave) bool {
	n := p
	for n != nil {
		if n.Current == c {
			// we have visited it.
			return false
		}
		n = n.Previous
	}
	return true
}

func (p *Path) Print() {
	ids := []string{}
	n := p
	for n != nil {
		ids = append(ids, n.Current.ID)
		n = n.Previous
	}
	// reverse the slice
	for i := 0; i < len(ids)/2; i++ {
		ids[i], ids[len(ids)-1-i] = ids[len(ids)-1-i], ids[i]
	}
	fmt.Println(ids)
}

// Implement Solution to Problem 1
func solve1(input string) string {

	system := map[string]*Cave{}

	getCave := func(id string) *Cave {
		c, ok := system[id]
		if !ok {
			c = &Cave{
				ID:       id,
				Big:      id[0] >= 'A' && id[0] <= 'Z',
				Connects: []*Cave{},
			}
			system[id] = c
		}
		return c
	}

	aoc.MapLines(input, func(line string) error {
		parts := strings.Split(line, "-")
		left, right := getCave(parts[0]), getCave(parts[1])

		// they connect to each other.
		left.Connects = append(left.Connects, right)
		right.Connects = append(right.Connects, left)
		return nil
	})

	// find ALL paths from "start" to "end", but don't cross
	// and non-Big caves more than once.
	// it would be good to describe paths as a comparable

	// breadth first search, start with the start cave.
	curr := []*Path{{
		Current: getCave("start"),
	}}
	var next []*Path
	count := 0
	for {
		next = []*Path{}

		for _, o := range curr {
			if o.Current.ID == "end" {
				// done.
				count++
				//o.Print()
			} else {
				next = append(next, o.NextPaths()...)
			}
		}
		if len(next) == 0 {
			return fmt.Sprint(count)
		}
		curr = next
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {

	system := map[string]*Cave{}

	getCave := func(id string) *Cave {
		c, ok := system[id]
		if !ok {
			c = &Cave{
				ID:       id,
				Big:      id[0] >= 'A' && id[0] <= 'Z',
				Connects: []*Cave{},
			}
			system[id] = c
		}
		return c
	}

	aoc.MapLines(input, func(line string) error {
		parts := strings.Split(line, "-")
		left, right := getCave(parts[0]), getCave(parts[1])

		// they connect to each other.
		left.Connects = append(left.Connects, right)
		right.Connects = append(right.Connects, left)
		return nil
	})

	// find ALL paths from "start" to "end", but don't cross
	// and non-Big caves more than once.
	// it would be good to describe paths as a comparable

	// breadth first search, start with the start cave.
	curr := []*Path{{
		Current: getCave("start"),
	}}
	var next []*Path
	count := 0
	for {
		next = []*Path{}

		for _, o := range curr {
			if o.Current.ID == "end" {
				// done.
				count++
				//o.Print()
			} else {
				next = append(next, o.NextPaths2()...)
			}
		}
		if len(next) == 0 {
			return fmt.Sprint(count)
		}
		curr = next
	}
}
