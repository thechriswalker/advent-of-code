package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 17, solve1, solve2)
}

// [3]int{x,y,z}
type pocket struct {
	data map[[3]int]struct{}
	// the maximum bounds we should consider.
	// min/max active cube in that plane
	min [3]int
	max [3]int
}

func (p *pocket) ActiveAt(x, y, z int) bool {
	_, ok := p.data[[3]int{x, y, z}]
	return ok
}

func (p *pocket) Activate(x, y, z int) {
	updateBound := func(n, i int) {
		if n < p.min[i] {
			p.min[i] = n
		}
		if n > p.max[i] {
			p.max[i] = n
		}
	}

	p.data[[3]int{x, y, z}] = struct{}{}
	updateBound(x, 0)
	updateBound(y, 1)
	updateBound(z, 2)

}

func newPocket() *pocket {
	return &pocket{
		min:  [3]int{math.MaxInt16, math.MaxInt16, math.MaxInt16},
		max:  [3]int{math.MinInt16, math.MinInt16, math.MinInt16},
		data: map[[3]int]struct{}{},
	}
}

func (p *pocket) CountActive() int {
	return len(p.data)
}

func (p *pocket) Tick() *pocket {
	// if there is nothing alive, do not iterate
	if len(p.data) == 0 {
		return p
	}
	next := newPocket()

	// but we iterate over the current bounds +/- 1
	for x := p.min[0] - 1; x <= p.max[0]+1; x++ {
		for y := p.min[1] - 1; y <= p.max[1]+1; y++ {
			for z := p.min[2] - 1; z <= p.max[2]+1; z++ {
				n := p.CountNeighboursActive(x, y, z)
				active := p.ActiveAt(x, y, z)
				if active {
					if n == 2 || n == 3 {
						// remains active.
						next.Activate(x, y, z)
					} else {
						// dies...
					}
				} else {
					if n == 3 {
						next.Activate(x, y, z)
					}
				}
			}
		}
	}
	return next
}

func (p *pocket) CountNeighboursActive(x, y, z int) (count int) {
	// in every direction.
	check := func(a, b, c int) {
		if p.ActiveAt(a, b, c) {
			count++
		}
	}

	//check(x, y, z)
	check(x, y, z+1)
	check(x, y, z-1)
	check(x, y+1, z)
	check(x, y+1, z+1)
	check(x, y+1, z-1)
	check(x, y-1, z)
	check(x, y-1, z+1)
	check(x, y-1, z-1)

	check(x+1, y, z)
	check(x+1, y, z+1)
	check(x+1, y, z-1)
	check(x+1, y+1, z)
	check(x+1, y+1, z+1)
	check(x+1, y+1, z-1)
	check(x+1, y-1, z)
	check(x+1, y-1, z+1)
	check(x+1, y-1, z-1)

	check(x-1, y, z)
	check(x-1, y, z+1)
	check(x-1, y, z-1)
	check(x-1, y+1, z)
	check(x-1, y+1, z+1)
	check(x-1, y+1, z-1)
	check(x-1, y-1, z)
	check(x-1, y-1, z+1)
	check(x-1, y-1, z-1)

	return
}

func (p *pocket) String() string {
	b := strings.Builder{}
	fmt.Fprintf(&b, "Active: %d Min:[x=%d,y=%d,z=%d] Max:[x=%d,y=%d,z=%d]\n", p.CountActive(), p.min[0], p.min[1], p.min[2], p.max[0], p.max[1], p.max[2])
	fmt.Fprintf(&b, "RAW: %v\n", p.data)
	for z := p.min[2] - 1; z <= p.max[2]+1; z++ {
		fmt.Fprintf(&b, "Z=%d\n", z)
		for y := p.min[1] - 1; y <= p.max[1]+1; y++ {
			for x := p.min[0] - 1; x <= p.max[0]+1; x++ {
				if p.ActiveAt(x, y, z) {
					b.Write([]byte{'#'})
				} else {
					b.Write([]byte{'.'})
				}
			}
			b.Write([]byte{'\n'})
		}
	}
	return b.String()
}

func parsePocket(in string) *pocket {
	p := newPocket()
	x, y := 0, 0
	for _, c := range in {
		switch c {
		case '.':
			//inactive!
			x++
		case '#':
			//active!
			p.Activate(x, y, 0)
			x++
		case '\n':
			// newline. update y, reset x
			y++
			x = 0
		}
	}
	return p
}

// Implement Solution to Problem 1
func solve1(input string) string {
	p := parsePocket(input)
	//fmt.Printf("INITIAL\n%s", p)
	ticks := 6
	for i := 0; i < ticks; i++ {
		p = p.Tick()
		//fmt.Printf("AFTER TICK %d\n%s", i+1, p)
	}
	return fmt.Sprintf("%d", p.CountActive())
}

// Implement Solution to Problem 2
func solve2(input string) string {
	p := parseHyperPocket(input)
	ticks := 6
	for i := 0; i < ticks; i++ {
		p = p.Tick()
		//fmt.Printf("TICK=%d ACTIVE=%d\n", i+1, p.CountActive())
	}
	return fmt.Sprintf("%d", p.CountActive())
}

func parseHyperPocket(in string) *hyperpocket {
	p := newHyperPocket()
	x, y := 0, 0
	for _, c := range in {
		switch c {
		case '.':
			//inactive!
			x++
		case '#':
			//active!
			p.Activate([4]int{x, y, 0, 0})
			x++
		case '\n':
			// newline. update y, reset x
			y++
			x = 0
		}
	}
	return p
}

// [3]int{x,y,z}
type hyperpocket struct {
	data map[[4]int]struct{}
}

func (p *hyperpocket) ActiveAt(c [4]int) bool {
	_, ok := p.data[c]
	return ok
}

func (p *hyperpocket) Activate(c [4]int) {
	p.data[c] = struct{}{}
}

func newHyperPocket() *hyperpocket {
	return &hyperpocket{
		data: map[[4]int]struct{}{},
	}
}

func (p *hyperpocket) CountActive() int {
	return len(p.data)
}

func (p *hyperpocket) Tick() *hyperpocket {
	// if there is nothing alive, do not iterate
	if len(p.data) == 0 {
		return p
	}
	next := newHyperPocket()

	// but we iterate over the current bounds +/- 1
	// OK this is pretty inefficient.
	// What we should do is iterate over the current ACTIVE
	// blocks and it's neighbours.
	// not repeating ones we have seen already.
	cache := map[[4]int]int{}
	getNeighbours := func(c [4]int) int {
		n, ok := cache[c]
		if !ok {
			n = p.CountNeighboursActive(c)
			cache[c] = n
		}
		return n
	}
	for coord := range p.data {
		// move all around
		for _, neighbour := range neighbours(coord) {
			n := getNeighbours(neighbour)
			if p.ActiveAt(neighbour) {
				if n == 2 || n == 3 {
					// remains active.
					next.Activate(neighbour)
				} else {
					// dies...
				}
			} else {
				if n == 3 {
					next.Activate(neighbour)
				}
			}
		}
		// this one is currently active.
		n := getNeighbours(coord)
		if n == 2 || n == 3 {
			// remains active.
			next.Activate(coord)
		} else {
			// dies...
		}
	}
	return next
}

func (p *hyperpocket) CountNeighboursActive(c [4]int) (count int) {
	for _, n := range neighbours(c) {
		if p.ActiveAt(n) {
			count++
		}
	}
	return
}

func neighbours(c [4]int) [][4]int {
	stack := make([][4]int, 0, 80)
	for dx := -1; dx < 2; dx++ {
		for dy := -1; dy < 2; dy++ {
			for dz := -1; dz < 2; dz++ {
				for dw := -1; dw < 2; dw++ {
					if dx == 0 && dy == 0 && dz == 0 && dw == 0 {
						continue
					}
					stack = append(stack, [4]int{c[0] + dx, c[1] + dy, c[2] + dz, c[3] + dw})
				}
			}
		}
	}
	return stack
}
