package main

import (
	"fmt"
	"slices"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 14, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '#')

	sum := tiltNorth(g)

	return fmt.Sprint(sum)
}

var usePattern = false

// Implement Solution to Problem 2
func solve2(input string) string {
	if usePattern {
		return solve2findPattern(input)
	}
	// the cache state is "marginally" quicker. ~85ms => 60ms
	// if I had a better caching "value" (currently a _long_ string)
	// I am sure it would be faster.
	// but I don't know how to calculate that on the fly, so the string is the
	// best solution I have
	return solve2cacheState(input)
}

// try it with a hashmap
func solve2cacheState(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '#')
	i := 0
	tick := func() int {
		tiltNorth(g) // don't care about the number
		tiltWest(g)
		sum := tiltSouth(g)
		tiltEast(g)
		i++
		return sum
	}
	// give some entries to start, to save allocations later...
	cache := make(map[string]int, 300)
	weights := make(map[int]int, 300)
	offset := 0
	period := 0
	for {
		w := tick()
		v := g.Value()
		if prev, ok := cache[v]; ok {
			// cycle!
			period = i - prev
			offset = prev
			break
		}
		weights[i] = w
		cache[v] = i
	}

	// index into cycle is (1billion - offset) % period
	idx := (1_000_000_000 - offset) % period
	// which means index into cache = idx + offset
	//fmt.Println("weights:", weights, "offset:", offset, "period:", period, "idx:", idx+offset)
	return fmt.Sprint(weights[idx+offset])
}

func solve2findPattern(input string) string {
	g := aoc.CreateFixedByteGridFromString(input, '#')
	// we will need to find the period of the loop.
	// so a cache, but we only need to care about the weight
	// but also I will need to iterate the grid differently
	// for the different directions.
	// how do we check for a period?
	// we need to find an overlap of the previous values with the
	// current run. i.e. 1,2,3,4,2,3,4,2,3,4  would be 2,3,4
	// do we try and find 2 matching runs with the same pattern?
	// do I just print the pattern and see roughly when it repeats and use that as a starting point...
	i := 0
	tick := func() int {
		tiltNorth(g) // don't care about the number
		tiltWest(g)
		sum := tiltSouth(g)
		tiltEast(g)
		i++
		return sum
	}
	cache := make([]int, 0, 10000)
	// now look for a pattern.
	patternLength := 0
	patternOffset := 0
	for {
		cache = append(cache, tick())
		//fmt.Println(cache)

		// Pretty sure my implementation of this is sub-optimal
		o, x := findLongestRepeatingSubPattern(cache)

		// fmt.Println("cache", cache)
		// fmt.Println("longest repeating pattern:", x, "=>", cache[len(cache)-x:])
		if x > patternLength {
			patternLength = x
			patternOffset = o
		}
		if patternLength > 2 {
			// just use this value...
			break
		}
	}
	//fmt.Println("iterations", i)
	pattern := cache[patternOffset : patternOffset+patternLength]
	// fmt.Println("pattern", pattern)
	// fmt.Println("offset", patternOffset)
	// fmt.Println("cache", cache[:patternOffset+patternLength*2])

	// so our weight at 1_000_000_000 cycles is
	weightIndex := (1_000_000_000 - patternOffset - 1) % len(pattern)

	// 99522 too high
	// 99291
	return fmt.Sprint(pattern[weightIndex])
}

func findLongestRepeatingPattern(s []int) (offset, length int) {
	// start at the begining and see if the slice from there matches the
	for x := 0; x < len(s)/2; x++ {
		remaining := len(s) - x
		if remaining%2 == 1 {
			//skip
			continue
		}
		checkLen := remaining / 2
		testPattern := s[x : x+checkLen]
		if slices.Equal(testPattern, s[x+checkLen:x+checkLen*2]) {
			// now if that pattern has a repeating pattern... then we should use that instead
			// shorter, but the right answer.
			return x, checkLen
		}
	}
	return 0, 0
}

func findLongestRepeatingSubPattern(s []int) (int, int) {
	o, l := findLongestRepeatingPattern(s)
	return o, l
	// I thought we needed this, but we really didn't
	// only because of my input though...
	// i.e. I never got a pattern like `1,1,2, 1,1,2, 9, 1,1,2,1,1,2,9,...`
	// where there was a short repeating pattern earlier and a longer later.
	// the _first_ repeated pattern was the correct one.
	// I don't know whether that is a result of my input, or a property of the puzzle,
	// but it meant my "find shorter pattern"  code was not needed at all...

	// is there a shorter pattern?
	// i.e. is this all repeats?
	// is there a better way to test this?
	// we could find the factors of the length: l and see if
	// chunks of that length are all equal?
	// if o == 0 {
	// 	// not gonna happen
	// 	return 0, 0
	// }
	// return findMinimumRepeatingPattern(s[o:o+l], o, l)
}

func findMinimumRepeatingPattern(pattern []int, offset int, length int) (int, int) {
	// if this is minimal, return the original values.
	// check if every slice group is equal?
	for x := 2; x <= length/2; x++ {
		if length%x == 0 {
			// a factor. let's test it
			t := pattern[0:x] // first group
			fail := false
			for i := 1; i < length/x; i++ {
				o := i * x
				//fmt.Println("testing subpattern", t, "vs", pattern[o:o+x], " size=", x)
				if !slices.Equal(t, pattern[o:o+x]) {
					fail = true
					break
				}
			}
			if !fail {
				// we found the shortest repeating pattern in that pattern
				//fmt.Println("given", pattern, "\nshortest pattern is length", x, t)
				return offset, x
			}
		}
	}

	return offset, length
}

func isMultiple(a, b int) bool {
	if b == 0 {
		return true // we will consider this true
	}
	n := 1
	// max?
	for {
		x := b * n
		if x == a {
			return true
		}
		if x > a {
			return false
		}
		n++
	}
}

// tilt and return the weight
func tiltNorth(g aoc.ByteGrid) int {
	sum := 0
	_, ymin, _, ymax := g.Bounds()
	weightMax := 1 + ymax - ymin
	// to move north from the top slide everything up as far as it will go.
	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		if b == 'O' {
			// try to move it
			y1 := y
			for {

				if c, _ := g.At(x, y1-1); c != '.' {
					break
				}
				y1--
			}
			if y1 != y {
				// move
				g.Set(x, y1, 'O')
				g.Set(x, y, '.')
			}
			// now add the weight
			sum += weightMax - y1
		}
	})
	return sum
}

func tiltSouth(g aoc.ByteGrid) int {
	sum := 0
	xmin, ymin, xmax, ymax := g.Bounds()
	weightMax := 1 + ymax - ymin
	// iterate in rows from bottom to top
	for y := ymax; y >= ymin; y-- {
		for x := xmin; x <= xmax; x++ {
			if b, _ := g.At(x, y); b == 'O' {
				// try to move down
				y1 := y
				for {
					if c, _ := g.At(x, y1+1); c != '.' {
						break
					}
					y1++
				}
				if y1 != y {
					// move
					g.Set(x, y1, 'O')
					g.Set(x, y, '.')
				}
				// now add the weight
				sum += weightMax - y1
			}
		}
	}
	return sum
}

// weights will not change here.
func tiltEast(g aoc.ByteGrid) {
	xmin, ymin, xmax, ymax := g.Bounds()
	// iterate in cols from right to left, moving right
	for x := xmax; x >= xmin; x-- {
		for y := ymax; y >= ymin; y-- {
			if b, _ := g.At(x, y); b == 'O' {
				// try to move right
				x1 := x
				for {
					if c, _ := g.At(x1+1, y); c != '.' {
						break
					}
					x1++
				}
				if x1 != x {
					// move
					g.Set(x1, y, 'O')
					g.Set(x, y, '.')
				}

			}
		}
	}
}
func tiltWest(g aoc.ByteGrid) {
	xmin, ymin, xmax, ymax := g.Bounds()
	// iterate in cols from left to right, moving left
	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			if b, _ := g.At(x, y); b == 'O' {
				// try to move left
				x1 := x
				for {
					if c, oob := g.At(x1-1, y); oob || c != '.' {
						break
					}
					x1--
				}
				if x1 != x {
					// move
					g.Set(x1, y, 'O')
					g.Set(x, y, '.')
				}

			}
		}
	}
}
