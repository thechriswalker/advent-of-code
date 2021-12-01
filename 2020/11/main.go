package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 11, solve1, solve2)
}

const (
	FLOOR = '.'
	EMPTY = 'L'
	TAKEN = '#'
)

type Seats struct {
	s      []rune // x=0,y=0,1,2,3,...,stride-1 x=1,y=0,1,2,...
	stride int
	height int
}

func (s *Seats) At(x, y int) rune {
	if x < 0 || y < 0 || y >= s.stride {
		return FLOOR
	}
	n := x*s.stride + y
	if n >= len(s.s) {
		return FLOOR
	}
	//log.Printf("Seat At(%d, %d) = %c", x, y, s.s[n])
	return s.s[n]
}

func (s *Seats) CountAdjacentOccupied(p int) int {
	x := p / s.stride
	y := p % s.stride
	count := 0
	// 8 adjacents. 3 with x-1, 3 with x+1 and 2 with x=x
	// x-1
	if s.At(x-1, y-1) == TAKEN {
		count++
	}
	if s.At(x-1, y) == TAKEN {
		count++
	}
	if s.At(x-1, y+1) == TAKEN {
		count++
	}
	// same x
	if s.At(x, y-1) == TAKEN {
		count++
	}
	if s.At(x, y+1) == TAKEN {
		count++
	}
	// x+1
	if s.At(x+1, y-1) == TAKEN {
		count++
	}
	if s.At(x+1, y) == TAKEN {
		count++
	}
	if s.At(x+1, y+1) == TAKEN {
		count++
	}
	//log.Printf("Adjacent to x:%d,y:%d = %d", x, y, count)
	return count
}

func (s *Seats) CountVisibleOccupied(p int) int {
	x := p / s.stride
	y := p % s.stride
	count := 0
	// 8 directions.
	// and we need to continue in each until we go off the map, or
	// hit a chair.
	// we use dx, dy for directions.
	count += s.VisibleInDirection(x, y, -1, 0)  // north
	count += s.VisibleInDirection(x, y, -1, 1)  // north-east
	count += s.VisibleInDirection(x, y, 0, 1)   // east
	count += s.VisibleInDirection(x, y, 1, 1)   // south-east
	count += s.VisibleInDirection(x, y, 1, 0)   // south
	count += s.VisibleInDirection(x, y, 1, -1)  // south-west
	count += s.VisibleInDirection(x, y, 0, -1)  // west
	count += s.VisibleInDirection(x, y, -1, -1) // north-west
	return count
}

func (s *Seats) VisibleInDirection(x, y, dx, dy int) int {
	for {
		x += dx
		y += dy
		// are we still on the map?
		if x < 0 || y < 0 || x >= s.height || y >= s.stride {
			// nope
			return 0
		}
		// is it a seat?
		switch s.At(x, y) {
		case TAKEN:
			return 1
		case EMPTY:
			return 0
		}
	}
}

func (s *Seats) String() string {
	builder := strings.Builder{}
	for i := 0; i < len(s.s); i++ {
		if i != 0 && i%s.stride == 0 {
			builder.WriteByte('\n')
		}
		builder.WriteByte(byte(s.s[i]))
	}
	builder.WriteByte('\n')
	return builder.String()
}

func (s *Seats) CountOccupied() int {
	count := 0
	for _, r := range s.s {
		if r == TAKEN {
			count++
		}
	}
	return count
}

func parseGrid(input string) *Seats {
	lines := strings.Split(input, "\n")
	seats := &Seats{
		s:      []rune{},
		stride: len(lines[0]),
		height: 0,
	}

	for _, line := range lines {
		for i, r := range line {
			seats.s = append(seats.s, r)
			if i == 0 {
				seats.height++
			}
		}

	}
	return seats
}

// Implement Solution to Problem 1
func solve1(input string) string {
	curr := parseGrid(input)
	next := parseGrid(input)
	tick := 0
	for {
		//log.Printf("tick:%d, occupied:%d\n%s", tick, curr.CountOccupied(), curr)
		tick++
		changed := false
		for i := 0; i < len(curr.s); i++ {
			r := curr.s[i]
			if r != FLOOR {
				n := curr.CountAdjacentOccupied(i)
				switch {
				case n == 0 && r == EMPTY:
					changed = true
					next.s[i] = TAKEN
				case n >= 4 && r == TAKEN:
					changed = true
					next.s[i] = EMPTY
				default:
					next.s[i] = r
				}
			}
		}
		if changed == false {
			return fmt.Sprintf("%d", next.CountOccupied())
		}
		// swap the buffers.
		curr, next = next, curr
		if tick > 1000000 {
			return "infinite loop"
		}
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	curr := parseGrid(input)
	next := parseGrid(input)
	tick := 0
	for {
		//log.Printf("tick:%d, occupied:%d\n%s", tick, curr.CountOccupied(), curr)
		tick++
		changed := false
		for i := 0; i < len(curr.s); i++ {
			r := curr.s[i]
			if r != FLOOR {
				n := curr.CountVisibleOccupied(i)
				switch {
				case n == 0 && r == EMPTY:
					changed = true
					next.s[i] = TAKEN
				case n >= 5 && r == TAKEN:
					changed = true
					next.s[i] = EMPTY
				default:
					next.s[i] = r
				}
			}
		}
		if changed == false {
			return fmt.Sprintf("%d", next.CountOccupied())
		}
		// swap the buffers.
		curr, next = next, curr
		if tick > 1000000 {
			return "infinite loop"
		}
	}
}
