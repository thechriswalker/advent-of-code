package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 20, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	m := parseMaze(input)

	// aoc.PrintByteGrid(m.grid, nil)
	// fmt.Println(m.portals)
	return fmt.Sprint(m.Solve())
}

// Implement Solution to Problem 2
func solve2(input string) string {
	m := parseMaze(input)

	// aoc.PrintByteGrid(m.grid, nil)
	// fmt.Println(m.portals)
	return fmt.Sprint(m.SolveRecursive())
}

func parseMaze(input string) *DonutMaze {
	dm := &DonutMaze{
		grid:    aoc.CreateFixedByteGridFromString(input, ' '),
		portals: map[point]point{},
		names:   map[point]string{},
	}

	// find the portals?
	// we will keep a cache of portal locations as we find them
	// once we have found one end, we keep it here
	// then when we find the other end, we join them
	portalLocations := map[string]point{}

	isChar := func(b byte) bool {
		return b >= 'A' && b <= 'Z'
	}

	checkPortal := func(b byte, x1, y1 int) {
		// first find the "." there should be only one.
		var name string
		// we actually want the position of the dot
		var dot point
		if x, _ := dm.grid.At(x1, y1-1); x == '.' {
			// . is up, so second letter is down
			dot = point{x: x1, y: y1 - 1}
			b2, _ := dm.grid.At(x1, y1+1)
			name = string([]byte{b, b2})
		} else if x, _ := dm.grid.At(x1, y1+1); x == '.' {
			// . is down, so first letter is up
			dot = point{x: x1, y: y1 + 1}
			b2, _ := dm.grid.At(x1, y1-1)
			name = string([]byte{b2, b})
		} else if x, _ := dm.grid.At(x1-1, y1); x == '.' {
			// . is left, so second letter is right
			dot = point{x: x1 - 1, y: y1}
			b2, _ := dm.grid.At(x1+1, y1)
			name = string([]byte{b, b2})
		} else if x, _ := dm.grid.At(x1+1, y1); x == '.' {
			// . is right, so first letter is left
			dot = point{x: x1 + 1, y: y1}
			b2, _ := dm.grid.At(x1-1, y1)
			name = string([]byte{b2, b})
		} else {
			// nope.
			return
		}
		if name == "AA" {
			dm.start = dot
			return
		}
		if name == "ZZ" {
			dm.finish = dot
			return
		}

		if src, ok := portalLocations[name]; ok {
			// existing!
			dm.portals[src] = dot
			dm.portals[dot] = src
			dm.names[src] = name
			dm.names[dot] = name
		} else {
			// new!
			portalLocations[name] = dot
		}
	}

	w, h := dm.grid.Width()-1, dm.grid.Height()-1

	aoc.IterateByteGrid(dm.grid, func(x, y int, b byte) {
		if !isChar(b) {
			return
		}
		// where is this?
		switch {
		case x == 0, y == 0, x == w, y == h:
			// ignore
		default:
			// inside, we have to check around
			// we cannot tell where the donut is...
			checkPortal(b, x, y)
		}
	})

	return dm
}

// first we need not only the grid but the portal locations.
// we can import the whole thing though...
type DonutMaze struct {
	grid aoc.ByteGrid

	// AA and ZZ
	start, finish point

	// both directions src=>dst
	portals map[point]point

	// lets hold on to the portal names for convenience
	names map[point]string
}

func (d *DonutMaze) PortalName(p point) string {
	s, ok := d.names[p]
	if !ok {
		return "??"
	}
	return s
}

type point struct {
	x, y int
}

// breath-first search through the maze.
func (d *DonutMaze) Solve() int {

	// places we have been
	cache := map[point]struct{}{
		d.start: {},
	}

	// current points we are stood on
	curr := []point{d.start}
	var next []point
	steps := 0
	for {
		steps++
		next = []point{}
		for _, p := range curr {
			for _, option := range d.Options(p) {
				if option == d.finish {
					return steps
				} else if _, ok := cache[option]; !ok {
					// not been here before...
					cache[option] = struct{}{}
					next = append(next, option)
				}
			}
		}
		if len(next) == 0 {
			return -1
		}
		curr = next
	}
}

func (d *DonutMaze) Options(p point) []point {
	opts := []point{}
	if dst, ok := d.portals[p]; ok {
		opts = append(opts, dst)
	}
	if b, _ := d.grid.At(p.x, p.y-1); b == '.' {
		opts = append(opts, point{x: p.x, y: p.y - 1})
	}
	if b, _ := d.grid.At(p.x, p.y+1); b == '.' {
		opts = append(opts, point{x: p.x, y: p.y + 1})
	}
	if b, _ := d.grid.At(p.x-1, p.y); b == '.' {
		opts = append(opts, point{x: p.x - 1, y: p.y})
	}
	if b, _ := d.grid.At(p.x+1, p.y); b == '.' {
		opts = append(opts, point{x: p.x + 1, y: p.y})
	}
	return opts
}

type point3 struct {
	x, y  int
	level int
}

func (p point3) Point() point {
	return point{x: p.x, y: p.y}
}
func (p point3) EqualsPoint(p2 point) bool {
	return p.x == p2.x && p.y == p2.y
}

func portalIsOuter(g aoc.ByteGrid, p point3) bool {
	return p.x == 2 || p.y == 2 || p.x == g.Width()-3 || p.y == g.Height()-3
}

func (d *DonutMaze) Print() {
	hiliteFn := func(x, y int, b byte) aoc.Color {
		if _, ok := d.portals[point{x: x, y: y}]; ok {
			if portalIsOuter(d.grid, point3{x: x, y: y}) {
				return aoc.BoldGreen
			}
			return aoc.BoldRed
		}
		return aoc.NoColor
	}
	aoc.PrintByteGridFunc(d.grid, hiliteFn)
}

func (d *DonutMaze) IsOK3(p point3) (point3, bool) {
	if p.EqualsPoint(d.start) || p.EqualsPoint(d.finish) {
		// is on start or finish, they are only valid on level 0
		return p, p.level == 0
	}
	b, _ := d.grid.At(p.x, p.y)
	return p, b == '.'
}

func (d *DonutMaze) Options3(p point3) []point3 {
	opts := []point3{}
	if dst, ok := d.portals[p.Point()]; ok {
		// it is a portal.
		if portalIsOuter(d.grid, p) {
			// outer portals move UP
			// but you can't go that way from level 0
			if p.level != 0 {
				opts = append(opts, point3{x: dst.x, y: dst.y, level: p.level - 1})
			}
		} else {
			opts = append(opts, point3{x: dst.x, y: dst.y, level: p.level + 1})
		}
	}

	if x, ok := d.IsOK3(point3{x: p.x, y: p.y - 1, level: p.level}); ok {
		opts = append(opts, x)
	}
	if x, ok := d.IsOK3(point3{x: p.x, y: p.y + 1, level: p.level}); ok {
		opts = append(opts, x)
	}
	if x, ok := d.IsOK3(point3{x: p.x - 1, y: p.y, level: p.level}); ok {
		opts = append(opts, x)
	}
	if x, ok := d.IsOK3(point3{x: p.x + 1, y: p.y, level: p.level}); ok {
		opts = append(opts, x)
	}

	return opts
}

func (d *DonutMaze) SolveRecursive() int {
	// here our path and cache require a "level"
	cache := map[point3]struct{}{}
	curr := []*Path{{
		p: point3{x: d.start.x, y: d.start.y, level: 0},
	}}
	cache[curr[0].p] = struct{}{}
	var next []*Path
	steps := 0
	//fmt.Println("start:", d.start, "finish", d.finish)
	//d.Print()
	for {
		steps++
		//fmt.Printf("iterations: %03d:\n", steps)

		next = []*Path{}
		for _, p := range curr {
			//	fmt.Printf(" path: %s\n", p.Sprint(d))
			for _, option := range d.Options3(p.p) {
				if option.level == 0 && option.EqualsPoint(d.finish) {
					// d.PrintPath(&Path{
					// 	p:    option,
					// 	prev: p,
					// })
					return steps
				} else if _, ok := cache[option]; !ok {
					// not been here before...
					cache[option] = struct{}{}
					next = append(next, &Path{
						p:    option,
						prev: p,
					})
				}
			}
		}
		if len(next) == 0 {
			return -1
		}
		curr = next
	}
}

type Path struct {
	p    point3
	prev *Path
}

func (p *Path) Sprint(g *DonutMaze) string {
	sb := &strings.Builder{}
	p.buildString(sb, g)
	return sb.String()
}
func (p *Path) buildString(sb *strings.Builder, g *DonutMaze) {
	if p.prev != nil {
		p.prev.buildString(sb, g)
		if p.prev.p.level != p.p.level {
			fmt.Fprintf(sb, "\n\t[%s=>%s](%d=>%d)", g.PortalName(p.prev.p.Point()), g.PortalName(p.p.Point()), p.prev.p.level, p.p.level)
		}
	}
	//fmt.Fprintf(sb, "(x:%d,y:%d,l:%d)", p.p.x, p.p.y, p.p.level)
}

// iterate the path, caching each new level as a copy of the grid
// and setting the path on it. Then print each grid in order.
// with the path hilighted
func (d *DonutMaze) PrintPath(p *Path) {
	levels := []int{}
	levelGrids := map[int]aoc.ByteGrid{}

	getLevelGrid := func(lvl int) aoc.ByteGrid {
		g, ok := levelGrids[lvl]
		if !ok {
			levels = append(levels, lvl)
			g = d.grid.Clone()
			levelGrids[lvl] = g
		}
		return g
	}

	for p != nil {
		g := getLevelGrid(p.p.level)
		if _, ok := d.portals[p.p.Point()]; ok {
			g.Set(p.p.x, p.p.y, '@')
		} else {
			g.Set(p.p.x, p.p.y, '+')
		}
		p = p.prev
	}

	slices.Sort(levels)
	for _, lvl := range levels {
		fmt.Println("Level ", lvl)
		aoc.PrintByteGridFunc(getLevelGrid(lvl), func(x, y int, b byte) aoc.Color {
			if b == '+' {
				return aoc.BoldYellow
			}
			if _, ok := d.portals[point{x: x, y: y}]; ok {
				if portalIsOuter(d.grid, point3{x: x, y: y}) {
					return aoc.BoldGreen
				}
				return aoc.BoldRed
			}
			return aoc.NoColor
		})
	}

}
