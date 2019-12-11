package main

import (
	"bufio"
	"fmt"
	"strings"

	"../../aoc"
)

const (
	CenterOfMass = "COM"
	Santa        = "SAN"
	You          = "YOU"
)

func main() {
	aoc.Run(2019, 6, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	dag := parseInput(input)[CenterOfMass]
	// find the sum of the depths of all nodes
	depths := dag.SumOfAllDepths(0)
	return fmt.Sprintf("%d", depths)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	dag := parseInput(input)
	curr := dag[You].parent     // what "you" are orbiting
	target := dag[Santa].parent // what "santa" is orbiting

	// we need to find the shortest path, sounds like a breadth first search!
	// AoC loves a bit of breadth first searching. but don't backtrack.
	// if we have tested a node, discard it next time.
	r := func() int {
		seen := map[*Orbital]struct{}{}
		distance := 0
		possibilities := []*Orbital{}
		nextPossibilities := append([]*Orbital{curr.parent}, curr.children...)
		for {
			distance++
			possibilities = nextPossibilities
			nextPossibilities = []*Orbital{}
			for _, p := range possibilities {
				if _, ok := seen[p]; ok {
					continue
				}
				seen[p] = struct{}{}
				if p == target {
					// found
					return distance
				}

				// not found add all the possibilities to the stack
				if p.parent != nil {
					nextPossibilities = append(nextPossibilities, p.parent)
				}
				nextPossibilities = append(nextPossibilities, p.children...)
			}
			if len(nextPossibilities) == 0 {
				return -1
			}
		}
	}()

	return fmt.Sprintf("%d", r)
}

// this will be our graph type
// the root node will be the one with id COM, but
// we can pick ANY node and follow the parents to the top.
// because this is a DAG
type Orbital struct {
	id       string
	parent   *Orbital
	children []*Orbital
}

func (o *Orbital) AddChild(c *Orbital) {
	o.children = append(o.children, c)
}

// finds the sum of the depths of all nodes.
func (o *Orbital) SumOfAllDepths(depth int) int {
	sum := depth // for this node.
	for _, c := range o.children {
		sum += c.SumOfAllDepths(depth + 1)
	}
	return sum
}

// parse the input and return the root orbital.
func parseInput(s string) map[string]*Orbital {
	lookup := map[string]*Orbital{}
	rd := strings.NewReader(s)
	sc := bufio.NewScanner(rd)
	var l, r string
	for sc.Scan() {
		line := strings.Split(sc.Text(), ")")
		l = line[0]
		r = line[1]
		left, ok := lookup[l]
		if !ok {
			left = &Orbital{
				id:       l,
				children: []*Orbital{},
			}
			lookup[l] = left
		}
		right, ok := lookup[r]
		if !ok {
			right = &Orbital{
				id:       r,
				children: []*Orbital{},
			}
			lookup[r] = right
		}
		// left is parent of right, right is child of left.
		right.parent = left
		left.AddChild(right)
	}
	return lookup
}
