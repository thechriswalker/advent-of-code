package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 24, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	grid := parseInput(input)
	fmt.Println(grid)

	// looking at the grid I guessed the order.
	// var ProblemOneOrder = []byte{'0', '4', '5', '6', '7', '3', '1', '2'}
	// but maybe we could solve generally by running permutations
	// then solving in order.
	//
	// turns out my guess was wrong, so we'll have to.
	// lets work out the min distance from each point to each other.
	distances := map[byte]map[byte]PathAndDistance{}
	var a, b byte
	remaining := make([]byte, len(grid.Markers))
	for i := 0; i < len(grid.Markers); i++ {
		distances['0'+byte(i)] = map[byte]PathAndDistance{}
		remaining[i] = '0' + byte(i)
	}
	for i := 0; i < len(grid.Markers)-1; i++ {
		for j := i + 1; j < len(grid.Markers); j++ {
			a, b = '0'+byte(i), '0'+byte(j)
			p, distance := grid.ShortestPath(a, b)
			pd := PathAndDistance{P: p, D: distance}
			distances[a][b] = pd
			distances[b][a] = pd
		}
	}

	// we have found all the distances, now start at 1 lets make all the permutations and find the smallest
	// one.
	// we have to start at 0.
	// lets start a 0 and work forwards

	permutations := []Permutation{Permutation{
		Current:   '0',
		Remaining: without('0', remaining),
		Path:      &Path{X: grid.Markers['0'][0], Y: grid.Markers['0'][1]},
	}}
	for {
		nextPermutations := []Permutation{}
		for _, p := range permutations {
			if len(p.Remaining) == 0 {
				break
			}
			for _, r := range p.Remaining {
				// add the distance
				nextPermutations = append(nextPermutations, Permutation{
					Current:   r,
					Remaining: without(r, p.Remaining),
					Distance:  p.Distance + distances[p.Current][r].D,
					Path:      join(p.Path, distances[p.Current][r].P),
				})
			}
		}
		if len(nextPermutations) == 0 {
			break
		}
		permutations = nextPermutations
	}

	// now find the smallest
	var min = permutations[0]
	for i := 1; i < len(permutations); i++ {
		if permutations[i].Distance < min.Distance {
			min = permutations[i]
		}
	}
	grid.AddPath(min.Path)
	fmt.Println(grid)

	return fmt.Sprintf("%d", min.Distance)
}

func without(c byte, b []byte) []byte {
	o := make([]byte, 0, len(b)-1)
	for i := 0; i < len(b); i++ {
		if b[i] != c {
			o = append(o, b[i])
		}
	}
	return o
}

func join(a, b *Path) *Path {
	// returns a new path which is both of these combined
	joined := &Path{X: a.X, Y: a.Y}
	src := joined
	curr := a.prev
	for curr != nil {
		src.prev = &Path{X: curr.X, Y: curr.Y}
		src = src.prev
		curr = curr.prev
	}
	curr = b
	for curr != nil {
		src.prev = &Path{X: curr.X, Y: curr.Y}
		src = src.prev
		curr = curr.prev
	}
	return joined
}

type PathAndDistance struct {
	P *Path
	D int
}

type Permutation struct {
	Current   byte
	Remaining []byte
	Distance  int
	Path      *Path
}

// Implement Solution to Problem 2
func solve2(input string) string {
	grid := parseInput(input)
	distances := map[byte]map[byte]PathAndDistance{}
	var a, b byte
	remaining := make([]byte, len(grid.Markers))
	for i := 0; i < len(grid.Markers); i++ {
		distances['0'+byte(i)] = map[byte]PathAndDistance{}
		remaining[i] = '0' + byte(i)
	}
	for i := 0; i < len(grid.Markers)-1; i++ {
		for j := i + 1; j < len(grid.Markers); j++ {
			a, b = '0'+byte(i), '0'+byte(j)
			p, distance := grid.ShortestPath(a, b)
			pd := PathAndDistance{P: p, D: distance}
			distances[a][b] = pd
			distances[b][a] = pd
		}
	}

	// almost identically to part 1, except we add the trip to 0 to the end of every permutation

	permutations := []Permutation{Permutation{
		Current:   '0',
		Remaining: without('0', remaining),
		Path:      &Path{X: grid.Markers['0'][0], Y: grid.Markers['0'][1]},
	}}
	for {
		nextPermutations := []Permutation{}
		for _, p := range permutations {
			if len(p.Remaining) == 0 {
				break
			}
			for _, r := range p.Remaining {
				// add the distance
				nextPermutations = append(nextPermutations, Permutation{
					Current:   r,
					Remaining: without(r, p.Remaining),
					Distance:  p.Distance + distances[p.Current][r].D,
					Path:      join(p.Path, distances[p.Current][r].P),
				})
			}
		}
		if len(nextPermutations) == 0 {
			break
		}
		permutations = nextPermutations
	}
	// add the trip back to zero
	for i := range permutations {
		permutations[i].Distance = permutations[i].Distance + distances[permutations[i].Current]['0'].D
		permutations[i].Path = join(permutations[i].Path, distances[permutations[i].Current]['0'].P)
	}

	// now find the smallest
	var min = permutations[0]
	for i := 1; i < len(permutations); i++ {
		if permutations[i].Distance < min.Distance {
			min = permutations[i]
		}
	}
	grid.AddPath(min.Path)
	fmt.Println(grid)

	return fmt.Sprintf("%d", min.Distance)
}

// just parse this into a grid.
type Grid struct {
	Width       int
	Height      int
	Data        []byte
	Start       [2]int
	StartIndex  int
	Markers     map[byte][2]int
	MarkerIndex map[int]byte
}

func (g *Grid) AddPath(p *Path) {
	for p != nil {
		i := g.Index(p.X, p.Y)
		if g.Data[i] == OPEN {
			g.Data[i] = PATH
		}

		p = p.prev
	}
}

func (g *Grid) XY(i int) (x, y int) {
	x = i % g.Width
	y = i / g.Width
	return
}

func (g *Grid) Index(x, y int) int {
	if x < 0 || x >= g.Width || y < 0 || y >= g.Height {
		return -1
	}
	return y*g.Width + x
}

func (g *Grid) SetByteAt(x, y int, b byte) {
	i := g.Index(x, y)
	if i != -1 {
		g.Data[i] = b
	}
}

func (g *Grid) ByteAt(x, y int) byte {
	i := g.Index(x, y)
	if i == -1 {
		return WALL
	}
	return g.Data[i]
}

const (
	WALL byte = '#'
	OPEN byte = '.'
	MARK byte = '*'
	PATH byte = '+'
)

func parseInput(input string) *Grid {
	rd := strings.NewReader(input)
	var b byte
	var err error
	grid := &Grid{
		Data:    make([]byte, 0, len(input)),
		Markers: make(map[byte][2]int, 10),
	}
	var x, y int
	for {
		b, err = rd.ReadByte()
		if err != nil {
			break
		}
		if b == '0' {
			// start
			grid.Start = [2]int{x, y}
			grid.StartIndex = len(grid.Data)
		}
		if b == '\n' {
			y++
			if grid.Width == 0 {
				grid.Width = x
			}
			x = 0
		} else {
			if b >= '0' && b <= '9' {
				grid.Markers[b] = [2]int{x, y}
				grid.Data = append(grid.Data, b)
			} else {
				grid.Data = append(grid.Data, b)
			}
			x++
		}
	}
	grid.Height = len(grid.Data) / grid.Width
	grid.MarkerIndex = make(map[int]byte, len(grid.Markers))
	for m, p := range grid.Markers {
		grid.MarkerIndex[grid.Index(p[0], p[1])] = m
	}
	return grid
}

func (g *Grid) String() string {
	s := &strings.Builder{}
	fmt.Fprintf(s, "Grid[%dx%d](%d,%d):", g.Width, g.Height, g.Start[0], g.Start[1])
	for i, c := range g.Data {
		if i%g.Width == 0 {
			fmt.Fprint(s, "\x1b[0m\n")
		}

		switch c {
		case WALL:
			fmt.Fprint(s, "\x1b[1;30m#")
		case OPEN:
			fmt.Fprint(s, "\x1b[1;30m.")
		case PATH:
			fmt.Fprint(s, "\x1b[1;35m+")
		default:
			if i == g.StartIndex {
				fmt.Fprint(s, "\x1b[1;32m0")
			} else {
				fmt.Fprintf(s, "\x1b[1;34m%c", c)
			}
		}
	}
	fmt.Fprint(s, "\x1b[0m\n")
	for m := 0; m < len(g.Markers); m++ {
		c := '0' + byte(m)
		p := g.Markers[c]
		fmt.Fprintf(s, "Marker '%c' at (%d,%d)\n", c, p[0], p[1])
	}
	return s.String()
}

type Dir uint8

const (
	U Dir = iota
	D
	L
	R
)

func (g *Grid) Fill(x, y int) {
	i := g.Index(x, y)
	if i != -1 {
		g.Data[i] = WALL
	}
}

func (g *Grid) Options(x, y int) []Dir {
	up := g.ByteAt(x, y-1) != WALL
	down := g.ByteAt(x, y+1) != WALL
	left := g.ByteAt(x-1, y) != WALL
	right := g.ByteAt(x+1, y) != WALL
	options := make([]Dir, 0, 4)
	if up {
		options = append(options, U)
	}
	if down {
		options = append(options, D)
	}
	if left {
		options = append(options, L)
	}
	if right {
		options = append(options, R)
	}
	return options
}

func (g *Grid) ShortestPath(a, b byte) (*Path, int) {
	var x1, y1, x2, y2 int
	x1 = g.Markers[a][0]
	y1 = g.Markers[a][1]
	x2 = g.Markers[b][0]
	y2 = g.Markers[b][1]
	steps := 0
	path := &Path{
		X: x1,
		Y: y1,
	}
	availablePaths := []*Path{path}
	visited := map[[2]int]struct{}{}
	for {
		steps++
		nextPaths := []*Path{}
		for _, path := range availablePaths {
			// walk a step
			for _, option := range path.Options(g, visited) {
				if option.X == x2 && option.Y == y2 {
					// done
					return option, steps
				}
				nextPaths = append(nextPaths, option)
			}
		}
		availablePaths = nextPaths
		if len(availablePaths) == 0 {

			panic("no more options!")
		}
	}
}

// just a step on a path.
type Path struct {
	prev *Path
	X, Y int
}

func (p *Path) Options(g *Grid, cache map[[2]int]struct{}) []*Path {
	options := []*Path{}
	possible := g.Options(p.X, p.Y)
	// up/down/left/right.
	for _, dir := range possible {
		switch dir {
		case U:
			if _, ok := cache[[2]int{p.X, p.Y - 1}]; !ok {
				options = append(options, &Path{prev: p, X: p.X, Y: p.Y - 1})
				cache[[2]int{p.X, p.Y - 1}] = struct{}{}
			}
		case D:
			if _, ok := cache[[2]int{p.X, p.Y + 1}]; !ok {
				options = append(options, &Path{prev: p, X: p.X, Y: p.Y + 1})
				cache[[2]int{p.X, p.Y + 1}] = struct{}{}
			}
		case L:
			if _, ok := cache[[2]int{p.X - 1, p.Y}]; !ok {
				options = append(options, &Path{prev: p, X: p.X - 1, Y: p.Y})
				cache[[2]int{p.X - 1, p.Y}] = struct{}{}
			}
		case R:
			if _, ok := cache[[2]int{p.X + 1, p.Y}]; !ok {
				options = append(options, &Path{prev: p, X: p.X + 1, Y: p.Y})
				cache[[2]int{p.X, p.Y - 1}] = struct{}{}
			}
		}
	}
	return options
}
