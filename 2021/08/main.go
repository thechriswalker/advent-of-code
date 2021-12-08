package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 8, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	sum := 0
	aoc.MapLines(input, func(line string) error {
		outputs := strings.Fields(line)[11:]
		// 0-9 are inputs, then 10 is pipe, the 11-14 are outputs.
		for _, s := range outputs {
			switch len(s) {
			case 2, 3, 4, 7:
				sum += 1
			}
		}
		return nil
	})
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// more tricky, we need to use extrapolate the
	// actual signals from the ones we have
	// num signals for each digit:
	// 0 -> 6
	// 1 -> 2
	// 2 -> 5
	// 3 -> 5
	// 4 -> 4
	// 5 -> 5
	// 6 -> 6
	// 7 -> 3
	// 8 -> 7
	// 9 -> 6

	// the other way around:
	// 2 -> 1
	// 3 -> 7
	// 4 -> 4
	// 5 -> 2,3,5
	// 6 -> 6, 9, 0
	// 7 -> 8

	// so we need to work out which of the 2,3,5,6,9
	// are which.
	// 6 does not have both segments from 1 so we have 6
	// now 6 segments is 9, 0
	// 3 has both segments from 1 so we have 3
	// now 5 segments is 2, 5

	// the difference betwen 1 and 4:
	// - should be in 9 but not 0
	// - should be in 5 but not 2

	// boom.

	// add number to sum
	displays := []Display{}
	aoc.MapLines(input, func(line string) error {
		d := Display{}
		d.Unmarshal(line)
		displays = append(displays, d)
		return nil
	})

	sum := 0
	for _, d := range displays {
		sum += d.Solve()
	}

	return fmt.Sprintf("%d", sum)
}

type Display [15]string

func (d *Display) Unmarshal(l string) {
	fields := strings.Fields(l)
	copy(d[:], fields)
}
func (d *Display) Signals() []string {
	return d[:10]
}
func (d *Display) Output() []string {
	return d[11:]
}

// so algorithm:
// find 1,4,7,8,
// find 3 by comparing the 5 digit ones with 1
// find 6 by comparing the 6 digit ones with 1
// find the difference between 1 and 4
// use that to compare against the 6 digits to find 6 and 9
// use again against the 5 digits to find 2 and 5
// use the map to find the digits of the display
func (d *Display) Solve() int {
	mapping := map[string]int{}
	var one, four []rune
	// first pass to pick out 1478
	sigs := d.Signals()
	for _, s := range sigs {

		switch len(s) {
		case 2:
			s = sortString(s)
			mapping[s] = 1
			one = []rune(s)
		case 3:
			s = sortString(s)
			mapping[s] = 7
		case 4:
			s = sortString(s)
			mapping[s] = 4
			four = []rune(s)
		case 7:
			s = sortString(s)
			mapping[s] = 8
		}
	}
	// diff between 4 and 1.
	diff := difference(one, four)

	for _, s := range sigs {
		s = sortString(s)
		switch len(s) {
		case 5:
			s = sortString(s)
			// if has all the elements of 1 -> 3
			// else if has the diff -> 5 else -> 2
			if hasAll(s, one) {
				mapping[s] = 3
			} else if hasAll(s, diff) {
				mapping[s] = 5
			} else {
				mapping[s] = 2
			}
		case 6:
			s = sortString(s)
			if !hasAll(s, one) {
				mapping[s] = 6
			} else if hasAll(s, diff) {
				mapping[s] = 9
			} else {
				mapping[s] = 0
			}
		}
	}
	//fmt.Println("mapping", mapping)
	// now we have all the mappings lets get the number
	sum := 0
	m := 1000
	for _, s := range d.Output() {
		s = sortString(s)
		//fmt.Println("digit:", s, "mapped:", mapping[s], "mag:", m)
		sum += m * mapping[s]
		m /= 10
	}
	return sum
}

type RuneSorter []rune

var _ sort.Interface = (RuneSorter)(nil)

func (rs RuneSorter) Len() int           { return len(rs) }
func (rs RuneSorter) Less(i, j int) bool { return rs[i] < rs[j] }
func (rs RuneSorter) Swap(i, j int)      { rs[i], rs[j] = rs[j], rs[i] }

func sortString(s string) string {
	r := RuneSorter(s)
	sort.Sort(r)

	return string(r)
}

func difference(s1, s2 []rune) []rune {
	m := map[rune]int{}
	for _, r := range s1 {
		m[r]++
	}
	for _, r := range s2 {
		m[r]--
	}
	out := []rune{}
	for r, i := range m {
		if i != 0 {
			out = append(out, r)
		}
	}
	return out
}

func hasAll(s string, rr []rune) bool {
	for _, r := range rr {
		if !strings.ContainsRune(s, r) {
			return false
		}
	}
	return true
}
