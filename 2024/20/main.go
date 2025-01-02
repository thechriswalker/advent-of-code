package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 20, solve1, solve2)
}

var SaveThreshold = 100

// Implement Solution to Problem 1
func solve1(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '#')
	var start, finish aoc.V2
	aoc.IterateByteGridv(g, func(p aoc.V2, b byte) {
		if b == 'S' {
			start = p
		}
		if b == 'E' {
			finish = p
		}
	})

	nonCheatRoutes := aoc.GetShortestPaths(g, start, finish, '#')

	if len(nonCheatRoutes) != 1 {
		panic("expected exactly one route")
	}

	// basically we need the actual path, and
	// then we walk it, seeing how much each possible cheat would save.
	path := nonCheatRoutes[0]

	pathCost := map[aoc.V2]int{}
	for i, p := range path {
		pathCost[p] = i
	}
	found := 0
	for i, p := range path {
		// see if we can cheat in any direction
		for _, d := range []aoc.V2{aoc.North, aoc.East, aoc.South, aoc.West} {
			n := p.Add(d)
			b, _ := g.Atv(n)
			if b == '#' {
				// we could cheat here
				normalCost := pathCost[n.Add(d)]
				cheatCost := i + 2
				diff := normalCost - cheatCost
				if diff >= SaveThreshold {
					found++
				}
			}
		}
	}

	// 1366 is too low. probably an off by one error in the diff > SaveThreshold check (yes! should have been >=)
	return fmt.Sprint(found)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '#')
	var start, finish aoc.V2
	aoc.IterateByteGridv(g, func(p aoc.V2, b byte) {
		if b == 'S' {
			start = p
		}
		if b == 'E' {
			finish = p
		}
	})

	nonCheatRoutes := aoc.GetShortestPaths(g, start, finish, '#')

	if len(nonCheatRoutes) != 1 {
		panic("expected exactly one route")
	}

	// basically we need the actual path, and
	// then we walk it, seeing how much each possible cheat would save.
	path := nonCheatRoutes[0]
	//fmt.Println("normal path length", len(path))

	// for i, p := range path {
	// 	g.Setv(p, '0'+byte(i%10))
	// }
	// aoc.PrintByteGridFunc(g, func(x, y int, b byte) aoc.Color {
	// 	if b >= '0' && b <= '9' {
	// 		return aoc.BoldCyan
	// 	}
	// 	return aoc.NoColor
	// })

	found := 0
	for i := 0; i < len(path)-1; i++ {
		// we need to find all possible cheats.
		// cheats will:
		// - start at p
		// - finish at a point on p[i+1:]
		// - have a manhattan distance of 20 or less
		// - not be a duplicate of an existing cheat
		found += possibleCheats(g, i, path)

	}

	return fmt.Sprint(found)
}

type Cheat struct {
	Start, Finish aoc.V2
}

func possibleCheats(g aoc.ByteGrid, pos int, path []aoc.V2) int {
	// we cannot save enough
	// for every other point on the path, we need to check if we can reach it in 20 or less steps
	found := 0
	p := path[pos]
	for j := 0; j < len(path); j++ {
		d := p.ManhattanDistance(path[j])
		if d <= 20 {
			// OK it is less that equ 20 steps
			// j is the number of steps it would normally take
			// pos + d is the number of steps if we cheat
			// so we save j - (pos + d) steps
			saving := j - (pos + d)
			if saving >= SaveThreshold {
				// we can reach here in 20 or less steps
				//fmt.Println("possible cheat", pos, path[pos], "saving", saving)
				found++
			}
		}
	}

	return found

}

// for each of those cheats, which can actually be created by the final step being "coming out of a wall"
