package main

import (
	"fmt"
	"iter"
	"slices"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 12, solve1, solve2)
}

// all the presents are 3x3 in both test and input.
// this won't affect how I store them, but might affect the alogrithm.

var surrounds = [8]aoc.V2{
	aoc.NorthWest,
	aoc.North,
	aoc.NorthEast,
	aoc.East,
	aoc.SouthEast,
	aoc.South,
	aoc.SouthWest,
	aoc.West,
}

var presentColors = map[byte]aoc.Color{
	'0': aoc.BoldCyan,
	'1': aoc.BoldRed,
	'2': aoc.BoldGreen,
	'3': aoc.BoldYellow,
	'4': aoc.BoldMagenta,
	'5': aoc.BoldWhite,

	'6': aoc.BoldCyan,
	'7': aoc.BoldRed,
	'8': aoc.BoldGreen,
	'9': aoc.BoldYellow,
}

// on a whim, I an going to store my shapes as a uint8 for the 8 surrounding
// points and a bool for the center. the center never changes on a rotation
// or a flip.
// if I store the bits clockwise, then a rotation is a bitshift.
// a flip still needs a
type PresentShape struct {
	id       byte
	surround uint8
	center   bool
}

func (p PresentShape) Rotate90Clockwise() PresentShape {
	// top 5 move right, bottom 3 move left
	nextSurround := (p.surround >> 2) | (p.surround << 6)
	return PresentShape{id: p.id, surround: nextSurround, center: p.center}
}

func (p PresentShape) FlipY() PresentShape {
	n_s := ((p.surround & 0b01000000) >> 4) | ((p.surround & 0b00000100) << 4)
	ne_se := ((p.surround & 0b00100000) >> 2) | ((p.surround & 0b00001000) << 2)
	e_w := (p.surround & 0b00010001)
	nw_sw := ((p.surround & 0b10000000) >> 6) | ((p.surround & 0b00000010) << 6)

	next := n_s | ne_se | e_w | nw_sw

	return PresentShape{id: p.id, surround: next, center: p.center}
}

func (p PresentShape) FlipX() PresentShape {
	n_s := (p.surround & 0b01000100)
	ne_nw := ((p.surround & 0b00100000) << 2) | ((p.surround & 0b10000000) >> 2)
	e_w := ((p.surround & 0b00010000) >> 4) | ((p.surround & 0b000000001) << 4)
	se_sw := ((p.surround & 0b00001000) >> 2) | ((p.surround & 0b00000010) << 2)

	next := n_s | ne_nw | e_w | se_sw

	return PresentShape{id: p.id, surround: next, center: p.center}
}
func (p PresentShape) Points(center aoc.V2) iter.Seq[aoc.V2] {
	return func(yield func(aoc.V2) bool) {
		if p.center {
			if !yield(center) {
				return
			}
		}
		for i := 0; i < 8; i++ {
			mask := uint8(1 << (7 - i))
			if p.surround&mask == mask {
				if !yield(center.Add(surrounds[i])) {
					return
				}
			}
		}
	}
}

func (p PresentShape) Area() int {
	a := 0
	if p.center {
		a++
	}
	for i := 0; i < 8; i++ {
		mask := uint8(1 << (7 - i))
		if p.surround&mask == mask {
			a++
		}
	}
	return a
}

func (p PresentShape) CanDraw(g aoc.ByteGrid, at aoc.V2) bool {
	for p := range p.Points(at) {
		// no oob check as we stretch the grid
		b, _ := g.Atv(p)
		if b != '.' {
			return false
		}
	}
	return true
}

func (p PresentShape) Draw(g aoc.ByteGrid, at aoc.V2) {
	for x := range p.Points(at) {
		g.Setv(x, p.id)
	}
}

// ok I have loads of shitty code to manipulate the shape, but no idea about
// the right algorithm to use to find a solution.

// great, on to the parsing... then I'll have to deal with it...

func parseInput(input string) (presents []PresentShape, trees []Tree) {
	current := ""
	nextId := byte('0')
	aoc.MapLines(input, func(line string) error {
		// if empty, skip
		// if {idx}: skip.
		if len(line) > 0 && (line[0] == '.' || line[0] == '#') {
			// part of a shape.
			current += line[0:3]
			if len(current) == 9 {
				// we have a shape.
				presents = append(presents, parseShape([]byte(current), nextId))
				nextId++
				current = ""
			}
		} else {
			if strings.IndexByte(line, 'x') > 0 {
				// this is a line and there will be six slots in the tree,
				var w, h, a, b, c, d, e, f int
				fmt.Sscanf(line, "%dx%d: %d %d %d %d %d %d", &w, &h, &a, &b, &c, &d, &e, &f)
				trees = append(trees, Tree{w: w, h: h, list: []int{a, b, c, d, e, f}})
			}
		}
		return nil
	})
	return
}

func parseShape(s []byte, id byte) PresentShape {
	p := PresentShape{id: id}
	for i := 0; i < 9; i++ {
		if s[i] == '.' {
			continue
		}
		shift := -1
		switch i {
		case 0, 1, 2:
			// top row. in order
			shift = 7 - i
		case 3:
			// West, least significatn
			shift = 0
		case 4:
			p.center = true
			// leave shift < 0
		case 5:
			// East
			shift = 4
		case 6, 7, 8:
			// SW
			shift = i - 5
		}
		if shift >= 0 {
			p.surround |= 1 << shift
		}
	}
	aoc.Debugf("parseShape(%c): %s => %08b", id, s, p.surround)
	return p
}

// what about all the orientations?
// we would need to flip X and see if it is the same.
// flip Y and see if it is the same (symmetrical)
// for all the unique flip, then we rotate them 3 times
// of course, we only keep the unique values.
// in both cases there are only six shapes, so this should be OK.
func (p PresentShape) Permutations() []PresentShape {
	set := map[PresentShape]struct{}{
		p: {},
	}
	px := p.FlipX()
	set[px] = struct{}{}
	py := px.FlipY()
	set[py] = struct{}{}
	more := []PresentShape{}
	for p := range set {
		for range 3 {
			p = p.Rotate90Clockwise()
			more = append(more, p)
		}
	}
	for _, p := range more {
		set[p] = struct{}{}
	}
	more = more[0:0]
	for p := range set {
		more = append(more, p)
	}

	slices.SortFunc(more, func(a PresentShape, b PresentShape) int {
		return int(a.surround) - int(b.surround)
	})

	return more
}

// Implement Solution to Problem 1
func solve1(input string) int {
	shapes, trees := parseInput(input)
	g := aoc.NewSparseByteGrid('.')
	// shapes are 3x3 so start a move the center 4 each time.

	// this is the list of presetn shapes by index that are arrangements of
	// each specified shape.
	presents := [][]PresentShape{}

	for _, s := range shapes {
		presents = append(presents, s.Permutations())
	}

	bumpY := aoc.V2{Y: 4}

	for i, opts := range presents {
		p := aoc.V2{X: 4 * i}
		for _, opt := range opts {
			opt.Draw(g, p)
			p = p.Add(bumpY)
		}
	}

	if aoc.IsDebug() {
		aoc.PrintByteGridC(g, presentColors)
	}
	count := 0
	for _, t := range trees {
		if t.CanFitPresents(presents) {
			count++
		}
	}

	return count
}

type Tree struct {
	w, h int
	list []int
}

// not sure how I am going to cache these...
type State struct {
	grid      aoc.ByteGrid // the current state of the grid
	remaining []int        // number of each index left to place
}

type Choice struct {
	p   PresentShape
	idx int
	pos aoc.V2
}

// takes a slice of slices of permutations of shapes.
// uses the internal list and size to see if we can fit.
func (t *Tree) CanFitPresents(presents [][]PresentShape) bool {
	// wtf. I guess we try each shape in turn and see if we can fit it.
	// also I don't think we should care about the center - i.e. we use a sparse grid and grow as we fit presents.
	// every time, we can check the bounds of the grid haven't exceeded our shape.
	// again this feels like a massive depth first search...
	// lots of allocation and duplication... but it might work?

	// each iteration cloning a grid and finding every "adjacent" spot a  piece could go,
	// and removing the pieces from an array.
	// trim when we cannot place any pieces

	// success if any list of remaining pieces is len = 0
	// failure if any list of next options = 0

	// never leave any gaps, so a candidate poistion for a new piece must make no gap. how do we enforce that??

	// we could assume that there are no "tricks" and for each next shape to fit, the one with the smallest effect on the area is the only one to use.
	//?

	// there are a couple of conditions we know will pass.
	// first number of presents < number of 3x3 squares in the space = easy win
	// area of all presents is greater than area of space = easy fail.
	// resort to brute force.
	squares := (t.w / 3) * (t.h / 3)
	treeArea := t.w * t.h
	presentCount := 0
	presentArea := 0
	for i, n := range t.list {
		presentCount += n
		presentArea += n * presents[i][0].Area()
	}
	aoc.Debugf("tree space: %dx%d: area=%d (sections=%d); presentCount=%d, presentArea=%d", t.w, t.h, treeArea, squares, presentCount, presentArea)
	if presentCount <= squares {
		return true
	}
	if treeArea < presentArea {
		return false
	}

	// fuck it, lets just try brute force.
	// turns out this is only needed for the tests.
	// all the actual problem values just fit....
	// balls.
	// and I started by writing the area code, then dismissed the idea.... :facepalm:
	return t.canFitPresentsBrute(presents)
}

func (t *Tree) canFitPresentsBrute(presents [][]PresentShape) bool {
	// initial grid.
	initialGrid := aoc.NewSparseByteGrid('.')

	curr := []State{
		{grid: initialGrid, remaining: cp(t.list)},
	}
	var next []State
	for {
		aoc.Debugf("step: len(curr) = %d", len(curr))
		for _, c := range curr {
			for poss := range nextPossiblePieces(c, presents) {
				//aoc.Debugf("possible piece! id=%c, at=%v", poss.p.id, poss.pos)
				g := c.grid.Clone()
				poss.p.Draw(g, poss.pos)
				// check bounds
				if g.Width() > t.w || g.Height() > t.h {
					// nope
					continue
				}
				l := cp(c.remaining)
				l[poss.idx]--
				if allZero(l) {
					if aoc.IsDebug() {
						aoc.PrintByteGridC(g, presentColors)
					}
					return true
				}
				next = append(next, State{
					grid:      g,
					remaining: l,
				})
			}
		}
		curr = next
		if len(curr) == 0 {
			if aoc.IsDebug() {
				aoc.Debug(">> no more options")
			}
			return false
		}
		// cull all but the smallest?
		next = []State{}
	}
}

func allZero(l []int) bool {
	for _, n := range l {
		if n != 0 {
			return false
		}
	}
	return true
}

type Option struct {
	Choice
	Distance int
}

var origin = aoc.V2{}

func nextPossiblePieces(c State, presents [][]PresentShape) iter.Seq[Choice] {
	x1, y1, x2, y2 := c.grid.Bounds()
	// we will ignore attempt to expand the boundary if necessary
	//aoc.Debug("nextPossiblePiece, current bounds:", x1, y1, x2, y2)
	firstRun := x1|y1|x2|y2 == 0
	if !firstRun {
		// they are not all zero.
		// in this case we expand out iteration in 1 square in all directions.
		// the all zero exception is for the very start when there is only a single position
		x1--
		y1--
		x2++
		y2++
	}
	return func(yield func(Choice) bool) {
		// iterate through all piece orientations and all possible placements.
		// if "can place" yeild the piece
		for i, r := range c.remaining {
			if r == 0 {
				continue // no more of these are required.
			}
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					point := aoc.V2{X: x, Y: y}
					// find all the shapes that fit...
					// in the first run we only pick a single orientation
					l := len(presents[i]) - 1
					if firstRun {
						l = 1
					}
					for _, ss := range presents[i][0:l] {
						// wow quad nested loop... :(
						ok := ss.CanDraw(c.grid, point)
						if ok {
							if !yield(Choice{
								pos: point,
								p:   ss,
								idx: i,
							}) {
								return
							}
						}
					}
				}
			}
		}
	}
}

func cp(a []int) []int {
	b := make([]int, len(a))
	copy(b, a)
	return b
}

// Implement Solution to Problem 2
func solve2(input string) int {
	panic("unsolved")
}
