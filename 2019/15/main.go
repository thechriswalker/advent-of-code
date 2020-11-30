package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"../../aoc"
	"../intcode"
)

func main() {
	aoc.Run(2019, 15, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	m := &Maze{
		data: map[[2]int]Tile{[2]int{0, 0}: start},
	}
	a, _ := os.LookupEnv("AOC_ANIMATE")
	ms, _ := strconv.Atoi(a)
	explore(m, input, ms)
	// now we have a complete maze, we can solve it for the shortest path.
	s := m.ShortestPath()
	return fmt.Sprint(s)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	m := &Maze{
		data: map[[2]int]Tile{[2]int{0, 0}: start},
	}
	explore(m, input, 0)
	// this gives us the whole maze.

	// now we need the longest possible path from oxygen.
	// very similar to the shortest path algorithm.
	a, _ := os.LookupEnv("AOC_ANIMATE")
	ms, _ := strconv.Atoi(a)
	s := m.LongestPath(ms)

	return fmt.Sprint(s)
}

// we will need to make a map. let's assume it is bounded and keep exploring.
// once we have the full map, we will

type Tile byte

const (
	unknown Tile = '.'
	wall         = '#'
	empty        = ' '
	start        = '*'
	oxygen       = 'O'
)

type Maze struct {
	// the explored area
	data map[[2]int]Tile
	// the bounds
	xmin, xmax int
	ymin, ymax int

	// some notable positions
	oxygen [2]int
}

func (m *Maze) WriteTo(w io.Writer) {
	for x := m.xmin - 1; x <= m.xmax+1; x++ {
		for y := m.ymin - 1; y <= m.ymax+1; y++ {
			c := m.data[[2]int{x, y}]
			switch c {
			// this is for color!
			case start:
				w.Write([]byte("\x1b[1;93m"))
			case wall:
				w.Write([]byte("\x1b[1;90m"))
			case oxygen:
				w.Write([]byte("\x1b[1;96m"))
			case empty:
				// nothing
				w.Write([]byte("\x1b[0m"))
			default:
				// override with unknown
				c = unknown
				w.Write([]byte("\x1b[90m"))
			}
			w.Write([]byte{byte(c)})
		}
		// clear and add newline
		w.Write([]byte("\x1b[0m\n"))
	}
}

func (m *Maze) ShortestPath() int {
	// start to oxygen
	var x1, y1, x2, y2 int
	x2, y2 = m.oxygen[0], m.oxygen[1]
	steps := 0
	path := &Step{X: x1, Y: y1}
	availablePaths := []*Step{path}
	visited := map[[2]int]struct{}{}

	for {
		steps++
		nextPaths := []*Step{}
		for _, step := range availablePaths {
			// walk a step
			for _, option := range step.Options(m, visited) {
				if option.X == x2 && option.Y == y2 {
					// done
					return steps
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

func (m *Maze) LongestPath(ms int) int {
	// start at oxygen, no target
	x1, y1 := m.oxygen[0], m.oxygen[1]
	steps := 0
	path := &Step{X: x1, Y: y1}
	availablePaths := []*Step{path}
	visited := map[[2]int]struct{}{}
	for {

		nextPaths := []*Step{}
		for _, step := range availablePaths {
			// walk a step
			m.Fill(step.X, step.Y, oxygen)
			for _, option := range step.Options(m, visited) {
				nextPaths = append(nextPaths, option)
			}
		}
		if ms > 0 {
			time.Sleep(time.Duration(ms) * time.Millisecond)
			fmt.Println("\x1b[2JStep", steps)
			fmt.Print(m)
		}
		availablePaths = nextPaths
		if len(availablePaths) == 0 {
			// we have finished!
			return steps
		}
		steps++
	}
}

//just a step on a path.
type Step struct {
	prev *Step
	X, Y int
}

func (s *Step) Options(m *Maze, visited map[[2]int]struct{}) []*Step {
	options := []*Step{}

	for _, o := range m.Options(s.X, s.Y) {
		if _, seen := visited[o]; !seen {
			options = append(options, &Step{
				prev: s,
				X:    o[0],
				Y:    o[1],
			})
			visited[o] = struct{}{}
		}
	}

	return options
}

type Direction int64

const (
	nowhere Direction = 0
	north             = 1 // negative x
	south             = 2 // positive x
	west              = 3 // negative y
	east              = 4 // postive y
)

func reverse(d Direction) Direction {
	switch d {
	case north:
		return south
	case west:
		return east
	case east:
		return west
	case south:
		return north
	}
	return nowhere
}

func (d Direction) String() string {
	switch d {
	case north:
		return "North"
	case west:
		return "West"
	case east:
		return "East"
	case south:
		return "South"
	}
	return "Nowhere"
}

type Response int64

const (
	hit_wall    Response = 0
	hit_nothing          = 1
	hit_oxygen           = 2
)

func (m *Maze) String() string {
	sb := strings.Builder{}
	m.WriteTo(&sb)
	return sb.String()
}

func (m *Maze) Options(x, y int) [][2]int {
	options := [][2]int{}
	var v [2]int

	v = [2]int{x, y + 1}
	if c := m.data[v]; c != wall {
		options = append(options, v)
	}

	v = [2]int{x, y - 1}
	if c := m.data[v]; c != wall {
		options = append(options, v)
	}
	v = [2]int{x + 1, y}
	if c := m.data[v]; c != wall {
		options = append(options, v)
	}
	v = [2]int{x - 1, y}
	if c := m.data[v]; c != wall {
		options = append(options, v)
	}
	return options
}

func (m *Maze) Unexplored(x, y int) Direction {
	// we must find the blocks on each side that are explored.
	if _, explored := m.data[[2]int{x, y + 1}]; !explored {
		return east
	}
	if _, explored := m.data[[2]int{x, y - 1}]; !explored {
		return west
	}
	if _, explored := m.data[[2]int{x + 1, y}]; !explored {
		return south
	}
	if _, explored := m.data[[2]int{x - 1, y}]; !explored {
		return north
	}
	return nowhere
}

func (m *Maze) Explore(x, y int, c Tile) {
	if _, explored := m.data[[2]int{x, y}]; explored {
		panic("cannot re-explore the same tile")
	}
	m.Fill(x, y, c)
}
func (m *Maze) Fill(x, y int, c Tile) {
	m.data[[2]int{x, y}] = c
	if x < m.xmin {
		m.xmin = x
	}
	if x > m.xmax {
		m.xmax = x
	}
	if y < m.ymin {
		m.ymin = y
	}
	if y > m.ymax {
		m.ymax = y
	}
}

type Path struct {
	prev            *Path
	returnDirection Direction
}

// we need to look around the entire maze until we have no empty space.
// but how do I know where that is.
// I guess I build a "path" linked list of where I have been.
// then I can go back and explore options I haven't explored until I
// get back to the beginning with no options left.
// also we never go to an explored square.
func explore(m *Maze, code string, ms int) {
	// the recursion is to keep going until you are blocked in all directions.
	// then backtrack until you get an option and if we run out of backtrack
	// and options, then we are done.
	pg := intcode.New(code)
	x, y := 0, 0

	path := &Path{
		prev:            nil,
		returnDirection: nowhere,
	}

	updatePosition := func(d Direction) {
		switch d {
		case nowhere:
		case north:
			x--
		case south:
			x++
		case west:
			y--
		case east:
			y++
		}
	}

	pg.RunAsync()

	for {
		if ms > 0 {
			time.Sleep(time.Duration(ms) * time.Millisecond)
			fmt.Println("\x1b[2JX", x, "Y", y, "D", m.Unexplored(x, y))
			fmt.Println(m)
		}
		nextDirection := m.Unexplored(x, y)
		if nextDirection == nowhere {
			// need to backtrack
			if path.prev == nil {
				// nowhere to go!
				return
			}
			// tell it to go back
			pg.Input <- func() int64 {
				return int64(path.returnDirection)
			}
			<-pg.Output // ignore output
			updatePosition(path.returnDirection)
			path = path.prev
			continue
		}
		// not a back track
		pg.Input <- func() int64 {
			return int64(nextDirection)
		}

		// fine
		// now read new position
		out := <-pg.Output
		back := reverse(nextDirection)
		updatePosition(nextDirection)
		switch Response(out) {
		case hit_nothing:
			// empty space
			m.Explore(x, y, empty)
			path = &Path{
				prev:            path,
				returnDirection: back,
			}
		case hit_wall:
			// don't move
			m.Explore(x, y, wall)
			updatePosition(back)
		case hit_oxygen:
			m.Explore(x, y, oxygen)
			m.oxygen = [2]int{x, y}
			path = &Path{
				prev:            path,
				returnDirection: back,
			}
		default:
			panic("bad output!")
		}
	}
}
