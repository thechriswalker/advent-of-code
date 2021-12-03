package main

import (
	"fmt"
	"strconv"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 3, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	bitcounts := []int{}
	bitlen := 0
	linecount := 0
	aoc.MapLines(input, func(line string) error {
		if bitlen == 0 {
			bitlen = len(line)
			bitcounts = make([]int, bitlen)
		}
		for i, bit := range line {
			if bit == '1' {
				bitcounts[i]++
			}
		}
		linecount++
		return nil
	})

	mostCommon := 0
	leastCommon := 0
	//fmt.Println("bitlen:", bitlen, "linecount:", linecount)
	// now if bitcounts[i] > linecount/2 build a binary number foreach option.
	for i := 0; i < bitlen; i++ {
		//	fmt.Println("pos:", i, "count:", bitcounts[i])
		if bitcounts[i] > linecount/2 {
			mostCommon |= 1 << (bitlen - i - 1)
		} else {
			leastCommon |= 1 << (bitlen - i - 1)
		}
	}
	//	fmt.Printf("gamma: %b, epsilon: %b", mostCommon, leastCommon)
	return fmt.Sprintf("%d", mostCommon*leastCommon)
}

type Diags struct {
	Readings  []string
	Bitcounts []int
}

func (d *Diags) Count() {
	if d.Bitcounts != nil {
		return
	}
	d.Bitcounts = make([]int, len(d.Readings[0]))
	for _, s := range d.Readings {
		for i, bit := range s {
			if bit == '1' {
				d.Bitcounts[i]++
			}
		}
	}
}

func (d *Diags) Filter(bit int, most bool) *Diags {
	lines := []string{}
	d.Count()

	isCommon := d.Bitcounts[bit] > len(d.Readings)/2

	isEqual := len(d.Readings)%2 == 0 && d.Bitcounts[bit] == len(d.Readings)/2
	for _, r := range d.Readings {
		// keep those that HAVE the correct bit
		if isCommon || isEqual {
			if most && r[bit] == '1' {
				lines = append(lines, r)
			}
			if !most && r[bit] == '0' {
				lines = append(lines, r)
			}
		} else {
			if most && r[bit] == '0' {
				lines = append(lines, r)
			}
			if !most && r[bit] == '1' {
				lines = append(lines, r)
			}
		}
	}
	return &Diags{Readings: lines}
}

func (d *Diags) Reduce(most bool, pos int) int64 {
	if len(d.Readings) == 1 {
		n, _ := strconv.ParseInt(d.Readings[0], 2, 64)
		return n
	}
	return d.Filter(pos, most).Reduce(most, pos+1)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	d := &Diags{Readings: []string{}}
	aoc.MapLines(input, func(s string) error {
		d.Readings = append(d.Readings, s)
		return nil
	})

	o2 := d.Reduce(true, 0)
	co2 := d.Reduce(false, 0)

	fmt.Printf("o2: %b (%d), co2: %b (%d)\n", o2, o2, co2, co2)

	return fmt.Sprintf("%d", o2*co2)
}
