package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 18, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	list := []Ins{}
	aoc.MapLines(input, func(line string) error {
		list = append(list, parseInstruction(line))
		return nil
	})

	// g1 := gridFromInstructions(list)

	// n1 := countInteriorAndBoundary(g1)

	// aoc.PrintByteGridFunc(g1, func(x, y int, b byte) aoc.Color {
	// 	if isPartOfLoop(b) {
	// 		return aoc.BoldGreen
	// 	}
	// 	if b == '#' {
	// 		return aoc.BoldCyan
	// 	}
	// 	return aoc.NoColor
	// })

	// should work for the simple method as well as the big one
	g := customGrid(list)
	g.Print(os.Stdout)
	n := g.CountInteriorAndBoundary()

	// 48503
	// if n != n1 {
	// 	fmt.Println("got", n, "expected", n1)
	// 	panic("regression!")
	// }
	return fmt.Sprint(n)
}

func isPartOfLoop(b byte) bool {
	switch b {
	case 'F', '7', 'J', 'L', '-', '|':
		return true
	}
	return false
}

// Implement Solution to Problem 2
func solve2(input string) string {
	list := []Ins{}
	aoc.MapLines(input, func(line string) error {
		list = append(list, parseInstruction2(line))
		return nil
	})

	//g := gridFromInstructions(list)

	// I don't think this will work. too many points to iterate.
	// we need another solution.
	// we have a massive polygon, and we need the area.
	// we could split it into chunks... but that would not help.
	// boundary counting solution could be improved to work sparsely in one direction
	// but I think we would need to work sparsely in both...
	//n := countInteriorAndBoundary(g)

	//fmt.Println(g.Bounds())
	// 0 0 1186328 1_186_328
	// hmm, it is a square. 1 million ops (optimizing row by row)
	// could be feasible (more than 1million ^2)
	// but we don't want to use our normal grid.
	// we need something specialised.
	// we will ignore `-` pieces so we can traverse the rows quickly.
	// we will need the `|` pieces to know where boundaries are
	g := customGrid(list)
	n := g.CountInteriorAndBoundary()
	return fmt.Sprint(n)
}

const (
	UP    = 'U'
	DOWN  = 'D'
	LEFT  = 'L'
	RIGHT = 'R'
)

type Ins struct {
	dir   byte
	count int
	color string
}

func parseInstruction(line string) Ins {
	i := Ins{}
	s := ""
	fmt.Sscanf(line, "%c %d %s", &(i.dir), &(i.count), &s)
	i.color = strings.Trim(s, "(#)")
	return i
}

func parseInstruction2(line string) Ins {
	i := Ins{}
	s := ""
	fmt.Sscanf(line, "%c %d %s", &(i.dir), &(i.count), &s)
	i.color = strings.Trim(s, "(#)")
	// now decode the color into distance and
	n, _ := strconv.ParseInt(i.color[0:5], 16, 64)
	i.count = int(n)
	switch i.color[5] {
	case '0':
		i.dir = RIGHT
	case '1':
		i.dir = DOWN
	case '2':
		i.dir = LEFT
	case '3':
		i.dir = UP

	}
	return i
}

func gridFromInstructions(list []Ins) *aoc.SparseByteGrid {
	g := aoc.NewSparseByteGrid('.')

	// current location
	x, y := 0, 0

	// set our initial location.
	// I intend to re-use the code from day 10,
	// so the byte we set will be the right "pipe" shape.
	// both our test grid and the real one start going right and the finish comes "up"
	// so this piece will be an 'F'

	g.Set(x, y, 'F')

	var dx, dy int
	var b byte
	final := len(list) - 1
	for i, ins := range list {
		// move that many in that direction.
		switch ins.dir {
		case UP:
			dx, dy = 0, -1
			b = '|'
		case RIGHT:
			dx, dy = 1, 0
			b = '-'
		case DOWN:
			dx, dy = 0, 1
			b = '|'
		case LEFT:
			dx, dy = -1, 0
			b = '-'
		}
		n := ins.count
		if i == final {
			n-- // skip the last piece on the final instruction
		}
		for j := 1; j <= n; j++ {
			if j == ins.count {
				// corner, need to check the "next direction"
				b = getCornerPiece(ins.dir, list[i+1].dir)
			}
			x, y = x+dx, y+dy
			g.Set(x, y, b)
		}
	}
	return g
}

var corners = map[byte]map[byte]byte{
	RIGHT: {
		RIGHT: '-',
		LEFT:  '-',
		UP:    'J',
		DOWN:  '7',
	},
	LEFT: {
		RIGHT: '-',
		LEFT:  '-',
		UP:    'L',
		DOWN:  'F',
	},
	UP: {
		RIGHT: 'F',
		LEFT:  '7',
		UP:    '|',
		DOWN:  '|',
	},
	DOWN: {
		RIGHT: 'L',
		LEFT:  'J',
		UP:    '|',
		DOWN:  '|',
	},
}

func getCornerPiece(curr, next byte) byte {
	return corners[curr][next]
}

func countInteriorAndBoundary(g aoc.ByteGrid) int {
	_, _, ymin, _ := g.Bounds()

	inside := 0
	row := ymin
	boundaryCount := 0
	boundaryOpener := byte(0x00)

	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		if row != y {
			row = y
			boundaryCount = 0
			boundaryOpener = 0x00
		}

		// is this a boundary cross? we are travelling west->east
		// a boundary is a pipe, or if the opener was a L then
		if isPartOfLoop(b) {
			// was this an opening or closing of the boundary.
			switch b {
			case '|':
				boundaryCount += 1 // definitely a boundary cross.
			case 'L', 'F': // start of a boundary
				boundaryOpener = b
			case '7': // end of boundary, the count
				// depends on the orientation.
				if boundaryOpener == 'L' {
					boundaryCount++
				}
				boundaryOpener = 0x00
			case 'J':
				if boundaryOpener == 'F' {
					boundaryCount++
				}
				boundaryOpener = 0x00
			}
			// but these count as part of the total
			inside++
		} else {
			if boundaryCount%2 == 1 {
				inside++
				//g.Set(x, y, '#')
			}
		}
	})
	return inside
}

type Boundary struct {
	x int
	b byte
}

func boundarySort(a, b Boundary) int {
	return a.x - b.x
}

type SparseRows struct {
	rows       map[int][]Boundary
	ymin, ymax int
	xmin       int
}

func (s *SparseRows) Print(w io.Writer) {
	var pos int
	for i := s.ymin; i <= s.ymax; i++ {
		// print this by skipping to each entry then printing.
		pos = s.xmin - 1
		for _, p := range s.rows[i] {
			for p.x > pos-1 {
				w.Write([]byte{'.'})
				pos++
			}
			w.Write([]byte{p.b})
			pos++
		}
		w.Write([]byte{'\n'})
	}
}

func (g *SparseRows) CountInteriorAndBoundary() int {
	// iterate the sparse rows, counting interior and boundary items.
	n := 0
	boundaryCount := 0
	boundaryOpener := Boundary{}
	lastBoundaryX := 0

	// doesn't matter what order we iterate the rows...
	for y, row := range g.rows {
		_ = y
		// row = []Boundary
		boundaryCount = 0
		boundaryOpener = Boundary{}
		lastBoundaryX = row[0].x - 1

		// now iterate the "vertical edges and corners"
		for _, piece := range row {
			// they are all "edge" pieces.
			//fmt.Printf("Handling piece at (%d,%d) '%c' [count=%d] boundaryCount = %d, if inside count += %d (%d - %d)\n", piece.x, y, piece.b, n, boundaryCount, piece.x-lastBoundaryX, piece.x, lastBoundaryX)
			switch piece.b {
			case '|':
				// a new boundary and "gulp" any enclosed pieces
				// if we were inside, then add all the interior pieces beforehand
				// otherwise, just this piece
				if boundaryCount%2 == 1 {
					// this means we _were_ inside a boundary.
					// so we need to add all ths interior pieces we missed.
					n += piece.x - lastBoundaryX
				} else {
					n++
				}
				boundaryCount++
				lastBoundaryX = piece.x
			case 'L', 'F':
				// boundary opener, we need to count the
				// if we _were_ inside the boundary, add
				// all the interior pieces between here and the lastBoundaryX
				if boundaryCount%2 == 1 {
					// yes, we were inside
					n += piece.x - lastBoundaryX
				} else {
					// otherwise just this piece
					n++
				}
				lastBoundaryX = piece.x
				boundaryOpener = piece
			case '7':
				// boundary closer
				// add the boundaries up to here.
				// i.e. since the opener
				n += piece.x - lastBoundaryX
				lastBoundaryX = piece.x
				// but whether this "counts" as a boundary increase depends on the opener
				if boundaryOpener.b == 'L' {
					// this is s-shape so yes, it counts
					boundaryCount++
				}
			case 'J':
				// boundary closer.
				// we need to add all the boundary pieces
				// including this one
				n += piece.x - lastBoundaryX
				lastBoundaryX = piece.x
				// but whether this "counts" as a boundary increase depends on the opener
				if boundaryOpener.b == 'F' {
					// this is s-shape so yes, it counts
					boundaryCount++
				}
			default:
				panic("unexpected boundary position!")
			}

		}

	}
	return n
}

func customGrid(list []Ins) *SparseRows {
	g := make(map[int][]Boundary)
	ymin, ymax := 0, 0
	xmin := 0

	// with an "F"
	set := func(x, y int, b byte) {
		g[y] = append(g[y], Boundary{x: x, b: b})
		if y < ymin {
			ymin = y
		}
		if y > ymax {
			ymax = y
		}
		if x < xmin {
			xmin = x
		}
	}

	set(0, 0, 'F')
	x, y := 0, 0
	for i := 0; i < len(list); i++ {
		ins := list[i]
		if ins.dir == UP || ins.dir == DOWN {
			// set the intermediate values.
			dy := 1
			if ins.dir == UP {
				dy = -1
			}
			for j := 0; j < ins.count-1; j++ {
				y += dy
				set(x, y, '|')
			}
			y += dy
		} else {
			// just skip to the end.
			// left and right
			dx := ins.count // right
			if ins.dir == LEFT {
				dx = -1 * ins.count
			}
			x += dx
		}
		// set the corner piece.
		// unless we are at the very end!
		if i < len(list)-1 {
			set(x, y, getCornerPiece(ins.dir, list[i+1].dir))
		}
	}

	for i := ymin; i <= ymax; i++ {
		slices.SortFunc(g[i], boundarySort)
	}

	return &SparseRows{
		rows: g,
		ymin: ymin,
		ymax: ymax,
		xmin: xmin,
	}
}
