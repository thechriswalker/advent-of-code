package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 15, solve1, solve2)
}

type ByteGrid struct {
	w, h int
	data []byte
	oob  byte
}

func ToByteGrid(input string, oob byte) *ByteGrid {
	bg := &ByteGrid{
		oob:  oob,
		data: make([]byte, 0, len(input)),
	}
	err := aoc.MapLines(input, func(line string) error {
		if bg.w == 0 {
			bg.w = len(line)
		} else if bg.w != len(line) {
			return errors.New("not all lines the same length")
		}
		bg.h++
		bg.data = append(bg.data, []byte(line)...)
		return nil
	})
	if err != nil {
		panic(err)
	}
	return bg
}
func (bg *ByteGrid) Len() int {
	return len(bg.data)
}

func (bg *ByteGrid) Coords(idx int) (x, y int, oob bool) {
	x, y = aoc.GridCoords(idx, bg.w)
	oob = idx < 0 || idx >= len(bg.data)
	return
}
func (bg *ByteGrid) Index(x, y int) (idx int, oob bool) {
	idx = aoc.GridIndex(x, y, bg.w, bg.h)
	return idx, idx == -1
}

func (bg *ByteGrid) At(x, y int) (b byte, oob bool) {
	idx, _ := bg.Index(x, y)
	return bg.AtIndex(idx)
}
func (bg *ByteGrid) AtIndex(idx int) (b byte, oob bool) {
	if idx < 0 || idx >= len(bg.data) {
		return bg.oob, true
	}
	return bg.data[idx], false
}

func (bg *ByteGrid) Iterate(fn func(x, y int, b byte)) {
	for idx, b := range bg.data {
		x, y := aoc.GridCoords(idx, bg.w)
		fn(x, y, b)
	}
}

var cardinals = [][2]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func (bg *ByteGrid) CardinalsIndex(idx int) []int {
	x, y, _ := bg.Coords(idx)
	return bg.Cardinals(x, y)
}

// indices of valid cardinals.
func (bg *ByteGrid) Cardinals(x, y int) []int {
	c := []int{}
	for _, d := range cardinals {
		if i, oob := bg.Index(x+d[0], y+d[1]); !oob {
			c = append(c, i)
		}
	}
	return c
}

type Grid interface {
	Len() int
	//Cardinals(x, y int) []int
	CardinalsIndex(idx int) []int
	AtIndex(idx int) (b byte, oob bool)
	//At(x, y int) (b byte, oob bool)
	//Coords(idx int) (x, y int, oob bool)
	//Index(idx int) (x, y int, oob bool)
}

type Path struct {
	Pos  int
	Risk int
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := ToByteGrid(input, '9')
	return findLowestRisk(g, false)
}
func findLowestRisk(g Grid, debug bool) string {
	end := g.Len() - 1
	// cache of positions and risk at them
	cache := map[int]int{0: 0}
	// breadth first search for the "shortest" path
	curr := []*Path{{Pos: 0, Risk: 0}}
	var next map[int]int
	// lets make a naive path for our upper bound.
	// we will go down a side and then across the bottom.
	min := math.MaxInt64
	// g.Iterate(func(x, y int, b byte) {
	// 	if x == 0 && y == 0 {
	// 		// don't count the start
	// 		return
	// 	}
	// 	if y == 0 || x == g.h-1 {
	// 		min += int(b - '0')
	// 	}
	// })
	// fmt.Println("Upper bound is", min)
	for {
		next = map[int]int{}
		for _, p := range curr {
			// find possible routes
			if p.Pos == end {
				// done
				//fmt.Println("found a path with risk", p.Risk)
				if min > p.Risk {
					min = p.Risk
				}
			} else {
				for _, c := range g.CardinalsIndex(p.Pos) {
					v, _ := g.AtIndex(c)
					vr := p.Risk + int(v-'0')

					// should we take this path.
					// not if we have seen this entry before.
					// unless we are a better path.
					if r, ok := next[c]; ok && r < vr {
						// we already have a cheaper path
						continue
					} else {
						// we should also check the overall cache
						if r, ok := cache[c]; ok && r < vr {
							// we already have a cheaper path
							continue
						} else {
							// we are the cheaper path.
							next[c] = vr
							cache[c] = vr
						}
					}
				}
			}
		}
		if debug {
			fmt.Println("Tick, next options:", len(next))
		}
		if len(next) == 0 {
			// done!
			break
		}
		curr = make([]*Path, 0, len(next))
		for p, v := range next {
			curr = append(curr, &Path{Pos: p, Risk: v})
		}

	}

	return fmt.Sprint(min)
}

type FiveGrid struct {
	bg *ByteGrid
}

func (fg *FiveGrid) Len() int {
	// 25 time bigger!
	return 25 * fg.bg.Len()
}

func (fg *FiveGrid) CardinalsIndex(idx int) []int {
	// this is a bit more tricky we have to do it ourselves
	c := []int{}
	x, y, _ := fg.Coords(idx)
	for _, d := range cardinals {
		if i, oob := fg.Index(x+d[0], y+d[1]); !oob {
			c = append(c, i)
		}
	}
	return c
}

func (fg *FiveGrid) Coords(idx int) (x, y int, oob bool) {
	x, y = aoc.GridCoords(idx, fg.bg.w*5)
	oob = idx < 0 || idx >= fg.Len()
	return
}
func (fg *FiveGrid) Index(x, y int) (idx int, oob bool) {
	idx = aoc.GridIndex(x, y, fg.bg.w*5, fg.bg.h*5)
	return idx, idx == -1
}

func (fg *FiveGrid) AtIndex(idx int) (byte, bool) {
	// this is actually where we need to translate
	// back to grid coords
	x, y := aoc.GridCoords(idx, fg.bg.w*5)
	return fg.At(x, y)
}
func (fg *FiveGrid) At(x, y int) (byte, bool) {
	mx := x % fg.bg.h
	my := y % fg.bg.w

	// find the data at the original point
	v, oob := fg.bg.At(mx, my)
	if oob {
		return fg.bg.oob, true
	}
	// now turn to int, add the manhattan distance
	dx := (x - mx) / fg.bg.h
	dy := (y - my) / fg.bg.w
	d := dx + dy
	// levels above 9 wrap to 1 so
	// the available levels are 1-9
	// and wrap from 9 to 1.
	// but we can mod 9 it
	mv := int(v-'0') + (d % 9)
	// if the result is > 9 we subtract 9.
	// so 10 -> 1
	//    11 -> 2 etc...
	for mv > 9 {
		mv -= 9
	}

	// turn it back into a byte
	return byte('0' + mv), false

}

// Implement Solution to Problem 2
func solve2(input string) string {
	// we need to make a bigger grid.
	// we could fake it though.
	bg := ToByteGrid(input, '9')
	fg := &FiveGrid{bg: bg}

	// let's test a couple of numbers
	// for _, d := range []struct{ x, y int }{{0, 0}, {10, 10}, {49, 49}, {4, 23}} {
	// 	b, oob := fg.At(d.x, d.y)
	// 	fmt.Printf("At (%d,%d) = %c (oob? %v)\n", d.x, d.y, b, oob)
	// }
	//return ""
	return findLowestRisk(fg, false)
}
