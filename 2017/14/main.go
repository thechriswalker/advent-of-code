package main

import (
	"fmt"
	"math/bits"
	"strings"

	"github.com/thechriswalker/advent-of-code/2017/10/knothash"
	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 14, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	input = strings.TrimSpace(input)
	var b []uint8
	used := 0
	for i := 0; i < 128; i++ {
		b = knothash.Hash(fmt.Sprintf("%s-%d", input, i))
		for _, c := range b {
			used += bits.OnesCount8(c)
		}
	}
	// 8100 is too low.
	return fmt.Sprint(used)
}

// Implement Solution to Problem 2
//
// NB that I did this _after_ 2024-12 which needed a very similar algorithm for finding contiguous regions
// in a grid. That's the main reason I didn't bother with the bit-twiddling required to do this with the
// []uint8 from the knothash package and instead converted it to a grid pattern
func solve2(input string) string {
	input = strings.TrimSpace(input)
	// this time we need to count contiguous regions.
	// I cannot work out the bit-twiddling require to do this with the []uint8
	// but I can convert the bits into a grid of bools, and work with that...
	g := aoc.NewFixedByteGrid(128, 128, '.', nil)
	var b []uint8
	for i := 0; i < 128; i++ {
		b = knothash.Hash(fmt.Sprintf("%s-%d", input, i))
		for j, c := range b {
			// 8 bits. to be assigned to 8 cells.
			for k := 0; k < 8; k++ {
				if c&(1<<uint(7-k)) != 0 {
					g.Set(j*8+k, i, '#')
				}
			}
		}
	}

	// now we can walk this grid and find contiguous regions.
	m := &Map{
		grid:           g,
		regionsByPoint: map[aoc.V2]*Region{},
		pointsByRegion: map[*Region][]aoc.V2{},
	}
	merged := 0
	id := 0
	aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
		if b != '#' {
			return
		}
		var region *Region
		p := v.Add(aoc.North)
		if x, oob := g.Atv(p); !oob && x == b {
			// north we might have a region
			if r, ok := m.regionsByPoint[p]; ok {
				region = r
			}
		}
		// west could have a region as well.
		p = v.Add(aoc.West)
		if x, oob := g.Atv(p); !oob && x == b {
			if r, ok := m.regionsByPoint[p]; ok {
				// if this region is is not the same as the one we found to the north
				// we need to "merge" them.
				if region != nil && region != r {
					// merge the regions.
					// we will merge the smaller region into the larger region.
					if len(m.pointsByRegion[region]) < len(m.pointsByRegion[r]) {
						region, r = r, region
					}
					// merge r into region
					for _, p := range m.pointsByRegion[r] {
						m.regionsByPoint[p] = region
					}
					m.pointsByRegion[region] = append(m.pointsByRegion[region], m.pointsByRegion[r]...)
					delete(m.pointsByRegion, r)
					merged++
				} else {
					region = r
				}
			}
		}

		if region == nil {
			id++
			region = &Region{ID: id}
			m.pointsByRegion[region] = []aoc.V2{v}
			m.regionsByPoint[v] = region
		} else {
			m.pointsByRegion[region] = append(m.pointsByRegion[region], v)
			m.regionsByPoint[v] = region
		}
	})
	fmt.Println()

	// I'll leave this in, it is pretty
	m.Print()

	//fmt.Println("regions", id-merged)

	//1187 is too high - problem was the same as part 1, I forgot to "trim" the input.
	return fmt.Sprint(len(m.pointsByRegion))
}

type Region struct{ ID int }

type Map struct {
	grid aoc.ByteGrid
	// these will be updated together
	regionsByPoint map[aoc.V2]*Region
	pointsByRegion map[*Region][]aoc.V2
}

var colorlist = []aoc.Color{
	aoc.BoldCyan,
	aoc.BoldYellow,
	aoc.BoldGreen,
	aoc.BoldRed,
	aoc.BoldMagenta,
	aoc.BoldBlue,
	aoc.BoldBlack,
	aoc.BoldWhite,
}

func (g *Map) Print() {
	// need to color by region.
	regionColors := map[*Region]aoc.Color{}
	i := 0
	for r := range g.pointsByRegion {
		regionColors[r] = colorlist[i]
		i = (i + 1) % len(colorlist)
	}

	aoc.PrintByteGridFunc(g.grid, func(x, y int, b byte) aoc.Color {
		if r, ok := g.regionsByPoint[aoc.Vec2(x, y)]; ok {
			return regionColors[r]
		}
		return aoc.NoColor
	})
}
