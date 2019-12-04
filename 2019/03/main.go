package main

import (
	"fmt"
	"io"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2019, 3, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := makeGrid(input)
	d := g.GetDistanceToClosestIntersection()
	return fmt.Sprintf("%d", d)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := makeGrid(input)
	d := g.GetDelayToMinimalDelayIntersection()
	return fmt.Sprintf("%d", d)
}

func makeGrid(input string) *Grid {
	// parse the wires by finding newlines in the input.

	p := 0
	s := strings.TrimSpace(input)
	g := &Grid{}
	for {
		//log.Println("RemainingString:", s[p:])
		idx := strings.IndexRune(s[p:], '\n')
		if idx == -1 {
			// last one
			g.AddWire(parseWire(s[p:]))
			break
		} else {
			// just a single one
			g.AddWire(parseWire(s[p : p+idx]))
			p += idx + 1
		}
	}
	return g
}

// best way to do this is to create a grid and then fill in the squares with the
// wires on them, as 2 different wires co-exist on a square we know there was a cross
// We will use powers of 2 for the wires and OR the values, so like wires do not count
// as intersection.
// then for any square on the grid where we find a value that is not `0` and not the
// same value as our current wire, it is an intersection.
// we will gather these intersections during the drawing process and then sort by
// distance.

// we don't know how big this grid will be, so we have to make it sparse.
// we use a map type for the points, the zero value is the empty point.
// 8 bits for the value gives us 8 wires max.
// int for the keys gives a massive grid size.

type Grid struct {
	points                 map[int]map[int]uint8 // the grid
	stepSums               map[int]map[int]int   // keeping track of the distance from each wire to this point
	intersections          [][2]int              // keep track of intersection points
	xmax, xmin, ymax, ymin int
	wires                  []Wire // all the wires
}

func abs(n int) int {
	if n < 0 {
		return -1 * n
	}
	return n
}

func (g *Grid) GetDistanceToClosestIntersection() int {
	var min int
	for i, p := range g.intersections {
		d := abs(p[0]) + abs(p[1])
		if i == 0 || d < min {
			min = d
		}
	}
	return min
}

func (g *Grid) GetDelayToMinimalDelayIntersection() int {
	var min int
	for i, p := range g.intersections {
		d := g.getStepSumAt(p[0], p[1])
		if i == 0 || d < min {
			min = d
		}
	}
	return min
}

func (g *Grid) setWireAt(x, y, d int, id uint8) {
	curr := g.getValueAt(x, y)
	// either no wires or not-our-wire
	// this catches multiple hits to the same intersection
	if curr&id != id {
		// this means this is the first time we have reached this point
		g.addStepSum(x, y, d)
		if curr != 0 {
			// intersection
			g.intersections = append(g.intersections, [2]int{x, y})
		}
	}
	// or the values together
	g.setValueAt(x, y, curr|id)
	// update bounds
	if x < g.xmin {
		g.xmin = x
	}
	if x > g.xmax {
		g.xmax = x
	}
	if y < g.ymin {
		g.ymin = y
	}
	if y > g.ymax {
		g.ymax = y
	}
}

func (g *Grid) setValueAt(x, y int, v uint8) {
	if g.points == nil {
		g.points = map[int]map[int]uint8{}
	}
	xm, ok := g.points[x]
	if !ok {
		xm = map[int]uint8{}
		g.points[x] = xm
	}
	xm[y] = v
}

func (g *Grid) addStepSum(x, y, v int) {
	if g.stepSums == nil {
		g.stepSums = map[int]map[int]int{}
	}
	xm, ok := g.stepSums[x]
	if !ok {
		xm = map[int]int{}
		g.stepSums[x] = xm
	}
	xm[y] += v
}

func (g *Grid) getValueAt(x, y int) uint8 {
	if g.points == nil {
		return 0
	}
	xm, ok := g.points[x]
	if !ok {
		return 0
	}
	return xm[y]
}

func (g *Grid) getStepSumAt(x, y int) int {
	if g.stepSums == nil {
		return 0
	}
	xm, ok := g.stepSums[x]
	if !ok {
		return 0
	}
	return xm[y]
}

type Wire [][2]int // series of Vectors, n then e. i.e. [1,0] is U1, and [0,-3] is L3

func (g *Grid) AddWire(wire Wire) {
	g.wires = append(g.wires, wire)
	wireId := uint8(len(g.wires))
	// start at the origin and follow the vectors point by point.
	// if we hit a point _with_ a different wire add point to intersections,
	// add wire to point, continue
	// current position.
	px, py := 0, 0
	d := 0 // distance travelled so far
	for _, v := range wire {
		// one of the values in the vect8or should be zero, so we can just process both
		// and the ordering doesn't matter.
		x, y := v[0], v[1]
		for x != 0 {
			// follow the x value
			inc := 1
			if x < 0 {
				inc = -1
			}
			// move and inc
			px += inc
			x -= inc
			d++
			g.setWireAt(px, py, d, wireId)
		}
		for y != 0 {
			// follow the y value
			// follow the x value
			inc := 1
			if y < 0 {
				inc = -1
			}
			// move and inc
			py += inc
			y -= inc
			d++
			g.setWireAt(px, py, d, wireId)
		}
		// continue
	}
}

func parseWire(s string) Wire {
	w := Wire{}
	rd := strings.NewReader(strings.TrimSpace(s))

	var x int
	var d byte
	var err error
	for {
		d, err = rd.ReadByte()
		if err != nil {
			break
		}
		switch d {
		case 'R', 'U', 'D', 'L':
			// ok
		default:
			panic("bad input")
		}
		// now read an integer
		_, err = fmt.Fscanf(rd, "%d", &x)
		if err != nil {
			break
		}
		switch d {
		case 'R':
			w = append(w, [2]int{x, 0})
		case 'L':
			w = append(w, [2]int{-1 * x, 0})
		case 'U':
			w = append(w, [2]int{0, x})
		case 'D':
			w = append(w, [2]int{0, -1 * x})
		}
		// then a comma or EOF
		rd.Seek(1, io.SeekCurrent)
	}
	//log.Println("Parsed Wire:", w)
	return w
}
