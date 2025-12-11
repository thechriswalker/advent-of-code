package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 11, solve1, solve2)
}

type node struct {
	id      string
	outputs []string
}

// Implement Solution to Problem 1
func solve1(input string) string {

	m := map[string]*node{}

	aoc.MapLines(input, func(line string) error {
		id := line[0:3]
		outputs := strings.Fields(line[5:])
		m[id] = &node{id: id, outputs: outputs}
		return nil
	})

	paths := 0

	curr := []string{"you"}
	next := []string{}
	for {
		for _, s := range curr {
			for _, o := range m[s].outputs {
				if o == "out" {
					paths++
					continue
				}
				next = append(next, o)
			}
		}
		if len(next) == 0 {
			break
		}
		curr = next
		next = []string{}
	}

	return fmt.Sprint(paths)
}

// Implement Solution to Problem 2

// as usual the simple update to the solution for 1 doesn't work
// first I thought it might loop, but even with a cache in the state
// it doesn't finish in time.
//
// so a different approach is needed.
// all paths need to go from through dac AND fft
// so I can find:
// - all paths svr->dac that don't pass fft
// - all paths svr->fft that don't pass dac
// - all paths from dac->fft
// - all paths from fft->dac
// - all paths from dac->out
// - all paths from fft->ou

// then we have:
//
//	  sum(svr->dac)*sum(dac->fft)*sum(fft->out)
//	+ sum(svr->fft)*sum(fft->dac)*sum(dac->out)
//
// ?
// so we need a function to count paths from a to b excluding a set.
//
// OK, that didn't work. same issue - the search for ALL paths srv->dac doesn't finsh
// because it doesn't know when to stop!
// I guess we have len(m) nodes, so that is our MAX DEPTH?

// OK that doesn't help either.

// what else do we know? if there are loops we could remove them.
// how to we shorten the searh space.
// sounds like a dynamic programming problem and AoC loves them....

// a comparable key for the cache
type key struct {
	id             string
	hasFFT, hasDAC bool
}

func dyn(cache map[key]int, m map[string]*node, curr *node, hasFFT, hasDAC bool) int {
	count := 0
	k := key{id: curr.id, hasFFT: hasFFT, hasDAC: hasDAC}
	if c, ok := cache[k]; ok {
		return c
	}
	for _, id := range curr.outputs {
		if id == "out" {
			if hasDAC && hasFFT {
				count++
			}
			continue
		}
		n, ok := m[id]
		if !ok {
			panic("cannot find node: " + id)
		}
		count += dyn(cache, m, n, hasFFT || id == "fft", hasDAC || id == "dac")
	}
	cache[k] = count
	return count
}

func solve2(input string) string {
	m := map[string]*node{}

	aoc.MapLines(input, func(line string) error {
		id := line[0:3]
		outputs := strings.Fields(line[5:])
		m[id] = &node{id: id, outputs: outputs}
		return nil
	})

	return fmt.Sprint(dyn(map[key]int{}, m, m["svr"], false, false))
}
