package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 8, solve1, solve2)
}

func parseNodesAndRoute(input string) ([]byte, map[string]*Node) {

	cache := map[string]*Node{}

	var route []byte

	aoc.MapLines(input, func(line string) error {
		if len(route) == 0 {
			route = []byte(line)
			return nil
		}
		if line == "" {
			return nil
		}
		parseNode(line, cache)
		return nil
	})

	return route, cache
}

// Implement Solution to Problem 1
func solve1(input string) string {
	route, cache := parseNodesAndRoute(input)
	curr := cache["AAA"]

	steps := 0
	for {

		switch route[steps%len(route)] {
		case 'L':
			curr = curr.left
		case 'R':
			curr = curr.right
		default:
			panic("Bad route! " + string(route))
		}
		steps++
		if curr.id == "ZZZ" {
			break
		}
	}

	return fmt.Sprint(steps)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	route, cache := parseNodesAndRoute(input)

	// find all start nodes
	nodes := []*Node{}
	for _, n := range cache {
		if n.id[2] == 'A' {
			nodes = append(nodes, n)
		}
	}

	// the naive way is a bad method - I really need to find the period of all
	// routes, then do the math...
	periods := make([]int, len(nodes))

	for i, n := range nodes {
		periods[i] = findPeriod(route, cache, n)
	}

	// now solve for LCM
	// LCM (a,b,c) = LCM(a, LCM(b,c))
	// so we can "reduce" our slice.
	x := lcm(periods[0], periods[1])
	for i := 2; i < len(periods); i++ {
		x = lcm(x, periods[i])
	}
	return fmt.Sprint(x)
}

// work out least common multiple using greatest common divisor
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// greatest common divisor using Euclid's Alogrithm
func gcd(a, b int) int {
	for {
		if b == 0 {
			return a
		}
		a, b = b, a%b
	}
}

func findPeriod(route []byte, cache map[string]*Node, start *Node) int {
	curr := start
	steps := 0
	for {
		switch route[steps%len(route)] {
		case 'L':
			curr = curr.left
		case 'R':
			curr = curr.right
		default:
			panic("Bad route! " + string(route))
		}
		steps++
		if curr.id[2] == 'Z' {
			return steps
		}
	}

}

func parseNode(line string, cache map[string]*Node) *Node {
	var x, l, r string
	_, err := fmt.Sscanf(line, "%3s = (%3s, %3s)", &x, &l, &r)
	if err != nil {
		fmt.Printf("BAD LINE: %q\n", line)
		panic(err)
	}
	n, ok := cache[x]
	if !ok {
		n = &Node{id: x}
		cache[x] = n
	}

	nl, ok := cache[l]
	if !ok {
		nl = &Node{id: l}
		cache[l] = nl
	}
	n.left = nl

	nr, ok := cache[r]
	if !ok {
		nr = &Node{id: r}
		cache[r] = nr
	}
	n.right = nr

	return n
}

type Node struct {
	id          string
	left, right *Node
}
