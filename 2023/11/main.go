package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 11, solve1, solve2)
}

func solve1(input string) string {
	return solveD(input, 2)
}

// Implement Solution to Problem 1
func solveD(input string, expansion int) string {

	g := aoc.CreateFixedByteGridFromString(input, '.')

	countX, countY := map[int]int{}, map[int]int{}

	galaxies := [][2]int{}

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		if b == '#' {
			countX[x]++
			countY[y]++
			galaxies = append(galaxies, [2]int{x, y})
		}
	})

	emptyX, emptyY := findEmpties(g, countX, countY)

	// ok now manhattan distance is a bit funky due to expansion
	makeD := func(empties map[int]bool) func(a, b int) int {
		return func(a, b int) int {
			if a == b {
				return 0
			}
			if a > b {
				a, b = b, a
			}
			d := b - a
			for x := range empties {
				if x >= a && x <= b {
					d += expansion - 1
				}
			}
			return d
		}
	}

	dx := makeD(emptyX)
	dy := makeD(emptyY)

	// now iterate (~twice)
	sum := 0
	for i := 0; i < len(galaxies)-1; i++ {
		gi := galaxies[i]
		for j := i + 1; j < len(galaxies); j++ {
			gj := galaxies[j]
			d := dx(gi[0], gj[0]) + dy(gi[1], gj[1])
			sum += d
			//fmt.Printf("distance from %d->%d = %d\n", i+1, j+1, d)
		}
	}
	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return solveD(input, 1_000_000)
}

func findEmpties(g aoc.ByteGrid, countX, countY map[int]int) (ex, ey map[int]bool) {
	x1, y1, x2, y2 := g.Bounds()

	ex, ey = map[int]bool{}, map[int]bool{}

	for x := x1; x <= x2; x++ {
		if n := countX[x]; n == 0 {
			ex[x] = true
		}
	}
	for y := y1; y <= y2; y++ {
		if n := countY[y]; n == 0 {
			ey[y] = true
		}
	}
	return
}
