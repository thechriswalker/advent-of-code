package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 8, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// I probably want a grid for this one, but I think we can just do the math on the each pair of nodes to find the antinode position and
	// whether it is out of bounds or not (which is all we care about in part 1)
	g := aoc.CreateFixedByteGridFromString(input, '.')

	// now we cant lists of the types of nodes and their positions.
	nodes := make(map[byte][]aoc.V2)

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		if b != '.' {
			nodes[b] = append(nodes[b], aoc.V2{X: x, Y: y})
		}
	})

	// iterate over node types and find the antinode positions
	// UNIQUE antinode positions
	antinodes := map[aoc.V2]struct{}{}
	for _, ns := range nodes {
		for i, n1 := range ns {
			for j, n2 := range ns {
				if i == j {
					continue
				}
				// there will be 2 antinodes
				a, b := findAntinodes(n1, n2)

				if _, oob := g.At(a.X, a.Y); !oob {
					antinodes[a] = struct{}{}
				}
				if _, oob := g.At(b.X, b.Y); !oob {
					antinodes[b] = struct{}{}
				}
			}
		}
	}

	return fmt.Sprint(len(antinodes))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '.')

	// now we cant lists of the types of nodes and their positions.
	nodes := make(map[byte][]aoc.V2)

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		if b != '.' {
			nodes[b] = append(nodes[b], aoc.V2{X: x, Y: y})
		}
	})

	// iterate over node types and find the antinode positions
	// UNIQUE antinode positions
	antinodes := map[aoc.V2]struct{}{}
	for _, ns := range nodes {
		for i, n1 := range ns {
			for j, n2 := range ns {
				if i == j {
					continue
				}
				// there will be an antinode at EACH antenna as well, possibly in the middle.
				// so for this pair. we start with the node itself and then iterate from that node in each direction.
				// this means some in the middle will also be found.
				// and we keep iterating until we are out of bounds.
				// I don't need a function for this.
				antinodes[n1] = struct{}{}
				dx, dy := reduceVector(n2.X-n1.X, n2.Y-n1.Y)
				v := n1
				var oob bool
				for {
					// up/left first.
					v = aoc.V2{X: v.X - dx, Y: v.Y - dy}
					if _, oob = g.At(v.X, v.Y); oob {
						break
					}
					g.Set(v.X, v.Y, '#')
					antinodes[v] = struct{}{}
				}
				// NB n1!
				v = n1
				for {
					v = aoc.V2{X: v.X + dx, Y: v.Y + dy}
					if _, oob = g.At(v.X, v.Y); oob {
						break
					}
					g.Set(v.X, v.Y, '#')
					antinodes[v] = struct{}{}
				}
			}
		}
	}

	// aoc.PrintByteGridFunc(g, func(x, y int, b byte) aoc.Color {
	// 	if b == '#' {
	// 		return aoc.BoldCyan
	// 	}
	// 	if b != '.' {
	// 		return aoc.BoldWhite
	// 	}
	// 	return aoc.NoColor
	// })

	return fmt.Sprint(len(antinodes))
}

func findAntinodes(n1, n2 aoc.V2) (aoc.V2, aoc.V2) {
	// we can assume the n1 is topper and lefter than n2, because of our iteration order.
	// we need to find the "difference" between the 2 vectors and add that to each one.
	dx, dy := n2.X-n1.X, n2.Y-n1.Y
	return aoc.V2{X: n1.X - dx, Y: n1.Y - dy}, aoc.V2{X: n2.X + dx, Y: n2.Y + dy}
}

func reduceVector(x1, y1 int) (x2, y2 int) {
	// lcm of x1 and y1 then divide
	// we can assume x1 and y1 are not 0
	gcd := aoc.GCD(x1, y1)
	return x1 / gcd, y1 / gcd
}
