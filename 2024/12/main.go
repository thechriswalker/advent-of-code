package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 12, solve1, solve2)
}

type Region struct {
	Type      byte
	Area      int
	Perimeter int
}

type GradenMap struct {
	grid aoc.ByteGrid
	// these will be updated together
	regionsByPoint map[aoc.V2]*Region
	pointsByRegion map[*Region][]aoc.V2
}

// Implement Solution to Problem 1
func createMap(input string) *GradenMap {
	// parse into a grid.
	// walk the grid, on each cell,  see if it is part of a contained region.
	// if not, search around it for it's region.
	// feels like we should be able to do it in one pass.
	// if we keep track regions by point, then  once we have all the regions we can walk them for their area.
	// or can we count the "area" as we go?
	// should do, if we add 1 area and 4 perimeter for each cell we add to a region, then substract the number of adjacent cells as perimeter.
	g := aoc.CreateFixedByteGridFromString(input, ' ')

	m := &GradenMap{
		grid:           g,
		regionsByPoint: map[aoc.V2]*Region{},
		pointsByRegion: map[*Region][]aoc.V2{},
	}

	aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
		// to keep it simple we will check all around the cell, to identify the perimeter addition.
		perimeter := 4
		var region *Region
		p := v.Add(aoc.North)
		if x, oob := g.Atv(p); !oob {
			if x == b {
				perimeter--
			}
			// north we might have a region
			if r, ok := m.regionsByPoint[p]; ok && r.Type == b {
				region = r
			}
		}
		// west could have a region as well.
		p = v.Add(aoc.West)
		if x, oob := g.Atv(p); !oob {
			if x == b {
				perimeter--
			}
			if r, ok := m.regionsByPoint[p]; ok && r.Type == b {
				// if this region is is not the same as the one we found to the north
				// we need to "merge" them.
				if region != nil && region != r {
					// merge the regions.
					// we will merge the smaller region into the larger region.
					if len(m.pointsByRegion[region]) < len(m.pointsByRegion[r]) {
						region, r = r, region
					}
					// merge r into region
					region.Area += r.Area
					region.Perimeter += r.Perimeter
					for _, p := range m.pointsByRegion[r] {
						m.regionsByPoint[p] = region
					}
					m.pointsByRegion[region] = append(m.pointsByRegion[region], m.pointsByRegion[r]...)
					delete(m.pointsByRegion, r)
				} else {
					region = r
				}
			}
		}

		// south and east, should not have regions as we move through the grid.
		p = v.Add(aoc.South)
		if x, oob := g.Atv(p); !oob && x == b {
			perimeter--
		}
		p = v.Add(aoc.East)
		if x, oob := g.Atv(p); !oob && x == b {
			perimeter--
		}

		if region == nil {
			region = &Region{Type: b, Perimeter: perimeter, Area: 1}
			m.pointsByRegion[region] = []aoc.V2{v}
			m.regionsByPoint[v] = region
		} else {
			region.Area++
			region.Perimeter += perimeter
			m.pointsByRegion[region] = append(m.pointsByRegion[region], v)
			m.regionsByPoint[v] = region
		}
	})

	return m
}
func solve1(input string) string {
	m := createMap(input)

	// now price up each region
	sum := 0
	for r := range m.pointsByRegion {
		price := r.Area * r.Perimeter
		sum += price
	}

	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	m := createMap(input)

	// for each region, iterate all the points.
	//
	// inner corners
	//
	// if a point has west and north, but not north west it is a corner - add 2 sides.
	// if a point has east and north, but not north east it is a corner - add 2 sides.
	// if a point has west and south, but not south west it is a corner - add 2 sides.
	// if a point has east and south, but not south east it is a corner - add 2 sides.
	//
	// outer corners
	//
	// if a point doesn't have west and north, then it is an outer corner - add 2 side.
	// if a point doesn't have east and north, then it is an outer corner - add 2 side.
	// if a point doesn't have west and south, then it is an outer corner - add 2 side.
	// if a point doesn't have east and south, then it is an outer corner - add 2 side.
	//
	// finally, we cut the number of corners in half as we will have doubled up (counted 2 for each corner)
	// this means number of corners = number of sides, so we just count corners.
	//
	// price is sum of sides * area
	price := 0

	for r, p := range m.pointsByRegion {
		corners := 0
		for _, v := range p {
			// inner corners
			hasNorth := m.regionsByPoint[v.Add(aoc.North)] == r
			hasWest := m.regionsByPoint[v.Add(aoc.West)] == r
			hasEast := m.regionsByPoint[v.Add(aoc.East)] == r
			hasSouth := m.regionsByPoint[v.Add(aoc.South)] == r
			if hasNorth && hasWest && m.regionsByPoint[v.Add(aoc.North).Add(aoc.West)] != r {
				corners += 1
			}
			if hasNorth && hasEast && m.regionsByPoint[v.Add(aoc.North).Add(aoc.East)] != r {
				corners += 1
			}
			if hasSouth && hasWest && m.regionsByPoint[v.Add(aoc.South).Add(aoc.West)] != r {
				corners += 1
			}
			if hasSouth && hasEast && m.regionsByPoint[v.Add(aoc.South).Add(aoc.East)] != r {
				corners += 1
			}
			// outer corners
			if !hasNorth && !hasWest {
				corners += 1
			}
			if !hasNorth && !hasEast {
				corners += 1
			}
			if !hasSouth && !hasWest {
				corners += 1
			}
			if !hasSouth && !hasEast {
				corners += 1
			}
		}
		price += r.Area * corners
	}

	return fmt.Sprint(price)
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

func (g *GradenMap) Print(grd aoc.ByteGrid) {
	// need to color by region.
	regionColors := map[*Region]aoc.Color{}
	i := 0
	for r := range g.pointsByRegion {
		regionColors[r] = colorlist[i]
		i = (i + 1) % len(colorlist)
	}

	aoc.PrintByteGridFunc(grd, func(x, y int, b byte) aoc.Color {
		if r, ok := g.regionsByPoint[aoc.Vec2(x, y)]; ok {
			return regionColors[r]
		}
		return aoc.NoColor
	})
}
