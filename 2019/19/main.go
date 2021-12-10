package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/2019/intcode"
	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 19, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	p := intcode.New(input)
	count := int64(0)
	for x := int64(0); x < 50; x++ {
		for y := int64(0); y < 50; y++ {
			c := p.Copy()
			c.RunAsync()
			c.Input <- func() int64 { return x }
			c.Input <- func() int64 { return y }
			o := <-c.Output
			//fmt.Printf("Input: %d,%dOutput: %d\n", x, y, o)
			if o == 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			count += o
		}
		fmt.Print("\n")
	}

	return fmt.Sprint(count)
}

type v2 struct{ x, y int64 }

// Implement Solution to Problem 2
func solve2(input string) string {
	p := intcode.New(input)

	cache := map[v2]bool{{0, 0}: true}

	calculate := func(x, y int64) bool {
		c := p.Copy()
		c.RunAsync()
		c.Input <- func() int64 { return x }
		c.Input <- func() int64 { return y }
		return 1 == <-c.Output
	}

	getAt := func(x, y int64) bool {
		o, found := cache[v2{x, y}]
		if !found {
			o = calculate(x, y)
			cache[v2{x, y}] = o
		}
		return o
	}

	is100x100 := func(v v2) bool {
		for x := v.x; x < v.x+100; x++ {
			for y := v.y; y < v.y+100; y++ {
				if !getAt(x, y) {
					return false
				}
			}
		}
		return true
	}

	// we will start from 0,0
	// count up X until we are out of the beam
	// then increment.
	beamStartY := int64(0)
	for x := int64(100); ; x++ {
		inBeam := false
		for y := beamStartY; ; y++ {
			if getAt(x, y) {
				if !inBeam {
					beamStartY = y
					inBeam = true
				}
				if is100x100(v2{x, y}) {
					// found it! (probably)
					return fmt.Sprint(x*10000 + y)
				}
			} else {
				if inBeam {
					// we were in the beam, break out to
					// the next x
					break
				}
			}
		}
	}
}
