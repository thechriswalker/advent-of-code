package main

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 5, solve1, solve2)
}

// naive is not going to work..
// type Map map[int]int

// func (m Map) Get(k int) int {
// 	v, ok := m[k]
// 	if !ok {
// 		return k
// 	}
// 	return v
// }

// func (m Map) SetRange(dst, src, l int) {
// 	for x := 0; x < l; x++ {
// 		m[src+x] = dst + x
// 	}
//}

type ValueRange struct{ min, max int }

type MapRange struct{ dst, src, len int }

func (m MapRange) From(src int) int {
	x := src - m.src // this is the difference from the source
	return m.dst + x
}

type Map struct {
	s []MapRange
}

// sort the ranges into ascending order by src
func (m *Map) Sort() {
	slices.SortFunc(m.s, func(a, b MapRange) int {
		return a.src - b.src
	})
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// turn one range into a set of other ranges.
// assumes our ranges are sorted by src.
// this function needs tests!
func (m *Map) GetRanges(v ValueRange) []ValueRange {
	out := []ValueRange{}
	bottomOfRange := v.min
	// crap. we need to deal with the non-specified ranges as well.
	// so we start with the input and work through our ranges,
	// in order.
	for _, r := range m.s {
		rmin, rmax := r.src, r.src+r.len-1
		if bottomOfRange < rmin {
			// fill in an identity range.
			imax := min(v.max, rmin-1)
			out = append(out, ValueRange{min: bottomOfRange, max: imax})
			bottomOfRange = imax + 1
			if v.max < bottomOfRange {
				break
			}
		}
		// now if this range has values in our input range, create an output range
		if bottomOfRange >= rmin && bottomOfRange <= rmax {
			// we have some in this range.
			omin := bottomOfRange
			omax := min(rmax, v.max)

			dmin := r.From(omin)
			dmax := r.From(omax)

			// if dmin < 0 || dmax < 0 || dmin > dmax {
			// 	fmt.Println("Bad value range")
			// 	fmt.Println("value:", v)
			// 	fmt.Println("range:", r)
			// 	fmt.Println("source:", ValueRange{min: omin, max: omax})
			// 	fmt.Println("  dest:", ValueRange{min: dmin, max: dmax})
			// 	fmt.Println("out:", out)
			// 	panic("bad value range")
			// }

			out = append(out, ValueRange{min: dmin, max: dmax})
			bottomOfRange = omax + 1
			if v.max < bottomOfRange {
				break
			}
		}

	}
	// if we got to the end and there are any values remaining, we add them as identity mappings
	if bottomOfRange < v.max {
		out = append(out, ValueRange{min: bottomOfRange, max: v.max})
	}
	return out
}

func (m *Map) Get(x int) int {
	// find a matching entry
	for i := range m.s {
		min := m.s[i].src
		max := min + m.s[i].len
		if x >= min && x < max {
			return m.s[i].dst + (x - min)
		}
	}
	return x
}

func (m *Map) SetRange(x MapRange) {
	m.s = append(m.s, x)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	seeds := []int{}

	maps := []*Map{}
	var curr *Map
	// first we parse the seeds, then the maps
	aoc.MapLines(input, func(line string) error {
		// if seeds line, parse that.
		if len(seeds) == 0 && strings.HasPrefix(line, "seeds: ") {
			seeds = aoc.ToIntSlice(line[6:], ' ')
			return nil
		}
		// if empty ignore
		if strings.TrimSpace(line) == "" {
			return nil
		}
		// if "header:" start new map
		if strings.HasSuffix(line, ":") {
			// new map
			curr = &Map{s: []MapRange{}}
			maps = append(maps, curr)
			return nil
		}
		// otherwise add to current map
		ints := aoc.ToIntSlice(line, ' ')
		//fmt.Println("ints", ints, "line", line)
		curr.SetRange(MapRange{dst: ints[0], src: ints[1], len: ints[2]})
		return nil
	})

	// follow maps
	for _, m := range maps {
		for i := range seeds {
			x := m.Get(seeds[i])
			//fmt.Printf("Map[%d] %d => %d\n", j, seeds[i], x)
			seeds[i] = x
		}
	}
	// now find the minimum
	minLoc := math.MaxInt
	for i := range seeds {
		if seeds[i] < minLoc {
			minLoc = seeds[i]
		}
	}

	return fmt.Sprint(minLoc)
}

// create a map from the Source in m1, to the Dest in m2
// func (m1 Map) Combine(m2 Map) Map {
// 	// for each range in m1,
// 	// see if there is an overlapping range in m2
// 	// if so, split into the non-overlapping ranges and
//  // create a new entry...
// }

// Implement Solution to Problem 2
// in problem 2 we cannot do this the naive, or even the
// slightly less naive way that I did for part 1.
//
// Instead we might need a completely different approach.
// I think we might be able to dive through the maps, and
// instead of working on all of them, split them into
// ranges that all work the same, and whether the bottom or
// top of _that_ range will be minimum. then we only need to test
// each end of the range.
// then we can split the input seeds into those which fall at the end of
// a range.
// actually, I think we can collapse all the maps into one.
//
// Of course we may be able to cheat. we can find the minimum location
// then track it back to a seed, then check if that exists.
//
// no, we cannot work backwards, as these maps are sparse.
// so I think we need to collapse the ranges and work on whole
// ranges at once.
// i.e. we have a series of N "ranges" of seeds, which map to M ranges
// of soil, and so on. Don't work on the
func solve2(input string) string {
	seedRanges := []ValueRange{}

	maps := []*Map{}
	var currMap *Map
	// first we parse the seeds, then the maps
	aoc.MapLines(input, func(line string) error {
		// if seeds line, parse that.
		if len(seedRanges) == 0 && strings.HasPrefix(line, "seeds: ") {
			seeds := aoc.ToIntSlice(line[6:], ' ')
			for i := 0; i < len(seeds); i += 2 {
				seedRanges = append(seedRanges, ValueRange{min: seeds[i], max: seeds[i] + seeds[i+1] - 1})
			}
			return nil
		}
		// if empty ignore
		if strings.TrimSpace(line) == "" {
			return nil
		}
		// if "header:" start new map
		if strings.HasSuffix(line, ":") {
			// new map
			currMap = &Map{s: []MapRange{}}
			maps = append(maps, currMap)
			return nil
		}
		// otherwise add to current map
		ints := aoc.ToIntSlice(line, ' ')
		//fmt.Println("ints", ints, "line", line)
		currMap.SetRange(MapRange{dst: ints[0], src: ints[1], len: ints[2]})
		return nil
	})

	// find the minimum possible location. (the last map)
	slices.SortFunc(currMap.s, func(a, b MapRange) int {
		return a.dst - b.dst
	})
	// now working from the minimum possible, try and find the range of seeds that fit this.

	curr := seedRanges
	next := []ValueRange{}
	for _, m := range maps {
		m.Sort()
		for _, r := range curr {
			next = append(next, m.GetRanges(r)...)
		}
		// fmt.Println("From", curr)
		// fmt.Println("To", next)1701774694
		curr = next
		next = []ValueRange{}
	}
	// now we have all possible ranges, find the minimum minimum.
	minLoc := math.MaxInt
	for _, r := range curr {
		if r.min < minLoc {
			minLoc = r.min
		}
	}

	return fmt.Sprint(minLoc)
}
