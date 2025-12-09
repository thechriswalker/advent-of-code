package main

import (
	"fmt"
	"slices"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 8, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return solve1n(input, 1000)
}

func toVec3(input string) []aoc.V3 {
	out := []aoc.V3{}
	aoc.MapLines(input, func(line string) error {
		v := aoc.V3{}
		n, err := fmt.Sscanf(line, "%d,%d,%d", &v.X, &v.Y, &v.Z)
		if err != nil {
			panic(err)
		}
		if n != 3 {
			panic("expected three values in line: " + line)
		}
		out = append(out, v)
		return nil
	})
	return out
}

func solve1n(input string, n int) string {
	list := toVec3(input)
	type entry struct {
		a, b aoc.V3
		d    float64
	}

	sizes := make([]entry, 0, len(list)*len(list))

	// we have a big list, so finding the distance between all pairs is going to take a while...
	// but I guess we have to?
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			a, b := list[i], list[j]
			d := a.EuclideanDistance(b)
			//			aoc.Debug("a", a, "b", b, "d", d)
			sizes = append(sizes, entry{a, b, d})
		}
	}

	// now we sort the slice and get the top 3
	slices.SortFunc(sizes, func(a, b entry) int {
		if a.d < b.d {
			return -1
		}
		if a.d > b.d {
			return 1
		}
		return 0
	})

	type Circuit struct {
		nodes []aoc.V3
	}

	v3toC := map[aoc.V3]*Circuit{}
	cset := map[*Circuit]struct{}{}

	// now we need to connect the "n" shortest connections.
	for i := range n {
		aoc.Debugf("connecting shortest [%d]", i)
		shortest := sizes[i]
		ca, aok := v3toC[shortest.a]
		cb, bok := v3toC[shortest.b]

		if !aok && !bok {
			// no pre-existing, just add one.
			c := &Circuit{nodes: []aoc.V3{shortest.a, shortest.b}}
			cset[c] = struct{}{}
			v3toC[shortest.a] = c
			v3toC[shortest.b] = c
		}
		if aok && !bok {
			// merge into a
			ca.nodes = append(ca.nodes, shortest.b)
			v3toC[shortest.b] = ca
		}
		if bok && !aok {
			// merge into b
			cb.nodes = append(cb.nodes, shortest.a)
			v3toC[shortest.a] = cb
		}
		if aok && bok {
			// if they are the same, do nothing!
			if ca != cb {
				// merge together...
				ca.nodes = append(ca.nodes, cb.nodes...)
				v3toC[shortest.b] = ca
				for _, n := range cb.nodes {
					v3toC[n] = ca
				}
				delete(cset, cb)
			}
		}
	}

	// now siphon into a list, sort and take the top 3
	// but also deduplicate.
	circuits := make([]*Circuit, len(cset))
	x := 0
	for c := range cset {
		circuits[x] = c
		x++
	}

	slices.SortFunc(circuits, func(a, b *Circuit) int {
		aa, bb := len(a.nodes), len(b.nodes)
		return bb - aa
	})

	for i, c := range circuits {
		aoc.Debugf("Circuit %d size %d", i, len(c.nodes))
	}

	prod := len(circuits[0].nodes) * len(circuits[1].nodes) * len(circuits[2].nodes)

	return fmt.Sprint(prod)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	list := toVec3(input)
	type entry struct {
		a, b aoc.V3
		d    float64
	}

	sizes := make([]entry, 0, len(list)*len(list))

	// we have a big list, so finding the distance between all pairs is going to take a while...
	// but I guess we have to?
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			a, b := list[i], list[j]
			d := a.EuclideanDistance(b)
			//			aoc.Debug("a", a, "b", b, "d", d)
			sizes = append(sizes, entry{a, b, d})
		}
	}

	// now we sort the slice and get the top 3
	slices.SortFunc(sizes, func(a, b entry) int {
		if a.d < b.d {
			return -1
		}
		if a.d > b.d {
			return 1
		}
		return 0
	})

	type Circuit struct {
		nodes []aoc.V3
	}

	v3toC := map[aoc.V3]*Circuit{}
	cset := map[*Circuit]struct{}{}

	// now we need to connect the "n" shortest connections.
	for i := range len(sizes) {
		shortest := sizes[i]
		aoc.Debugf("connecting shortest [%d] %v and %v (distance: %0.2f)", i, shortest.a, shortest.b, shortest.d)
		ca, aok := v3toC[shortest.a]
		cb, bok := v3toC[shortest.b]

		if !aok && !bok {
			// no pre-existing, just add one.
			c := &Circuit{nodes: []aoc.V3{shortest.a, shortest.b}}
			cset[c] = struct{}{}
			v3toC[shortest.a] = c
			v3toC[shortest.b] = c
			continue
		}
		if aok && !bok {
			// merge into a
			ca.nodes = append(ca.nodes, shortest.b)
			v3toC[shortest.b] = ca
		}
		if bok && !aok {
			// merge into b
			cb.nodes = append(cb.nodes, shortest.a)
			v3toC[shortest.a] = cb
			// this to allow the check at the end to use this circuit instead of ca
			ca = cb
		}
		if aok && bok {
			// if they are the same, do nothing!
			if ca != cb {
				// merge together...
				ca.nodes = append(ca.nodes, cb.nodes...)
				v3toC[shortest.b] = ca
				for _, n := range cb.nodes {
					v3toC[n] = ca
				}
				delete(cset, cb)
			}
		}
		if len(ca.nodes) == len(list) {
			// ca has all the nodes.
			// we have the last pair!
			aoc.Debug("final pair", shortest.a, "and", shortest.b)
			prod := shortest.a.X * shortest.b.X
			return fmt.Sprint(prod)
		}
	}
	panic("not all connected!")
}
