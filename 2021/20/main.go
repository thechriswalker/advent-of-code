package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 20, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// create the enchancment
	enhance := parseEnhance(input[0:512])

	g := aoc.CreateFixedByteGridFromString(input[514:], '.')
	//fmt.Println("original")
	//aoc.PrintByteGrid(g, nil)

	g1 := runEnhance(enhance, g)
	//fmt.Println("\n\nfirst pass")
	//aoc.PrintByteGrid(g1, nil)

	g2 := runEnhance(enhance, g1)
	//fmt.Println("\n\nsecond pass")
	//aoc.PrintByteGrid(g2, nil)

	lit := 0
	aoc.IterateByteGrid(g2, func(_, _ int, b byte) {
		if b == '#' {
			lit++
		}
	})

	return fmt.Sprint(lit)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// 50 enhancements!
	enhance := parseEnhance(input[0:512])
	var current aoc.ByteGrid
	current = aoc.CreateFixedByteGridFromString(input[514:], '.')

	for i := 0; i < 50; i++ {
		current = runEnhance(enhance, current)
	}
	lit := 0
	aoc.IterateByteGrid(current, func(_, _ int, b byte) {
		if b == '#' {
			lit++
		}
	})

	return fmt.Sprint(lit)
}

type EnhanceMap map[int]struct{}

func (em EnhanceMap) Lit(i int) bool {
	_, ok := em[i]
	return ok
}

func parseEnhance(input string) EnhanceMap {
	// we set the bits the other way around.
	// e.g. index 0 in the string is bit 511
	if len(input) != 512 {
		panic("invalid length for enhancement string")
	}
	em := make(EnhanceMap, 512)
	for i := 0; i < 512; i++ {
		switch input[i] {
		case '#':
			em[i] = struct{}{}
		}
	}
	return em
}

func runEnhance(enhance EnhanceMap, g aoc.ByteGrid) aoc.ByteGrid {
	// we will make a Sparse one.
	var next aoc.ByteGrid
	// the infinite image starts dark, so the sum
	// of all 9 pixels in the void is 0.
	// if the enhance 0 in our input is 1, so after 1
	// enhancement all infinite pixels are lit around our little image.
	// let's pick an unknow pixel to test.
	unknown, _ := g.At(1000000, 1000000)
	if unknown == '.' {
		// unknown is dark
		if enhance.Lit(0) {
			// enhancing dark gives light
			next = aoc.NewSparseByteGrid('#')
		} else {
			// enhancing dark gives dark
			next = aoc.NewSparseByteGrid('.')
		}
	} else {
		// currently the infinite void is lit.
		if enhance.Lit(511) {
			// enchancing gives dark
			next = aoc.NewSparseByteGrid('#')
		} else {
			next = aoc.NewSparseByteGrid('.')
		}
	}

	// we actually need to iterate from "outside" the bounds...
	// because we want all squares that have at least 1 pixel inside the image
	// which means 1 extra pixel each side.
	x1, y1, x2, y2 := g.Bounds()
	x1 -= 1
	y1 -= 1
	x2 += 1
	y2 += 1
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			// we need the nine pixels around this one.
			code := find9(g, x, y)
			if enhance.Lit(code) {
				next.Set(x, y, '#')
			} else {
				next.Set(x, y, '.')
			}
		}
	}
	return next
}

func find9(g aoc.ByteGrid, x, y int) int {
	//debug := x == 3 && y == 3
	n := 8
	v := 0
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			xx, yy := x+j, y+i
			b, _ := g.At(xx, yy)

			// if debug {
			// 	fmt.Println("grid at ", xx, ",", yy, "is", string([]byte{b}), "position:", n)
			// }
			if b == '#' {
				v |= 1 << n
			}
			n--

		}
	}
	// if debug {
	// 	fmt.Println("x", x, "y", y, "n", n, "v", v)
	// }
	return v
}
