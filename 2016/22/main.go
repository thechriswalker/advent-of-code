package main

import (
	"fmt"
	"sort"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 22, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	nodes := parseInput(input)
	//fmt.Println(nodes)
	sum := 0
	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {
			a, b := nodes[i], nodes[j]
			if a.Used != 0 && a.Used <= b.Avail {
				sum++
			}
			// and the other way around
			if b.Used != 0 && b.Used <= a.Avail {
				sum++
			}
		}
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	nodes := parseInput(input)
	sort.Sort(Nodes(nodes))
	// now create the first grid.
	// assume that all the node positions are filled...
	grid := &Grid{
		State: make([][2]int, len(nodes)),
	}
	var empty int
	for i, n := range nodes {
		grid.State[i] = [2]int{n.Used, n.Avail}
		if n.Used == 0 {
			empty = i
		}
		if n.X > grid.Width {
			grid.Width = n.X
		}
	}
	grid.Width++ // nodes are zerobased

	// now printing this shows us something.
	// there is a single empty node
	// all others cannot swap.
	// so we must move the empty node to position and then solve
	// also the "wall" of large disks abuts the right hand side, so we must move
	// around it to the left.
	// this is the number of steps and the X position.
	steps, x := grid.MoveToTop(empty)
	// now we move the to the right-1
	// width is actually the
	steps += grid.Width - 2 - x
	// now we have to move the data left "width-1" times.
	steps += grid.Width - 1
	// each one must be followed by 4 moves to put the "empty" in
	// the right place again. we will have to do this width-2 times.
	steps += (grid.Width - 2) * 4

	return fmt.Sprintf("%d", steps)
}

// this doesn't run anywhere need fast enough...
func solveGeneralCase(input string) string {
	nodes := parseInput(input)
	sort.Sort(Nodes(nodes))
	// now create the first grid.
	// assume that all the node positions are filled...
	grid := &Grid{
		State: make([][2]int, len(nodes)),
	}
	for i, n := range nodes {
		grid.State[i] = [2]int{n.Used, n.Avail}
		if n.X > grid.Width {
			grid.Width = n.X
		}
	}
	grid.Interested = [2]int{grid.Width, 0}

	cache := GridCache{}
	cache.HasNotSeen(grid) // cache initial state
	// breath first search
	steps := 0
	availableNextMoves := []*Seq{{grid: grid}}
	for {
		steps++
		nextMoves := []*Seq{}
		for _, move := range availableNextMoves {
			for _, option := range move.grid.AvailableMoves() {
				if option.IsDone() {
					fmt.Println(move)
					fmt.Println(option)
					return fmt.Sprintf("%d", steps)
				}
				if cache.HasNotSeen(option) {
					nextMoves = append(nextMoves, &Seq{
						prev: move,
						grid: option,
					})
				}
			}
		}
		availableNextMoves = nextMoves
		if len(availableNextMoves) == 0 {
			panic("no more moves")
		}
	}
}

type Nodes []*Node

func (n Nodes) Len() int      { return len(n) }
func (n Nodes) Swap(i, j int) { n[i], n[j] = n[j], n[i] }
func (n Nodes) Less(i, j int) bool {
	if n[i].Y == n[j].Y {
		return n[i].X < n[j].X
	}
	return n[i].Y < n[j].Y
}

type Node struct {
	X, Y        int
	Avail, Used int
}

//Filesystem              Size  Used  Avail  Use%
///dev/grid/node-x0-y6     85T   67T    18T   78%

func parseInput(input string) []*Node {
	rd := strings.NewReader(input)
	var x, y, s, u, a, p int
	var err error
	// skip the first 2 lines
	var b byte
	for {
		b, err = rd.ReadByte()
		if err != nil {
			break
		}
		if b == '\n' {
			x++
			if x == 2 {
				break
			}
		}
	}
	// now create nodes
	nodes := []*Node{}
	for {
		_, err = fmt.Fscanf(rd, "/dev/grid/node-x%d-y%d ", &x, &y)
		if err != nil {
			break
		}
		_, err = fmt.Fscanf(rd, "%dT", &s)
		if err != nil {
			break
		}
		_, err = fmt.Fscanf(rd, "%dT", &u)
		if err != nil {
			break
		}
		_, err = fmt.Fscanf(rd, "%dT", &a)
		if err != nil {
			break
		}
		_, err = fmt.Fscanf(rd, "%d%%\n", &p)
		if err != nil {
			break
		}
		nodes = append(nodes, &Node{X: x, Y: y, Avail: a, Used: u})
	}
	return nodes
}

// linked list of states for debugging
type Seq struct {
	prev *Seq
	grid *Grid
}

func (s *Seq) String() string {
	grids := []*Grid{}
	curr := s
	for {
		grids = append(grids, curr.grid)
		if curr.prev == nil {
			break
		}
		curr = curr.prev
	}
	b := strings.Builder{}
	// now print in reverse order
	for i := len(grids) - 1; i >= 0; i-- {
		b.WriteString(grids[i].String())
		b.WriteByte('\n')
	}
	return b.String()
}

// prints the grid like in the readme
func (g *Grid) String() string {
	s := &strings.Builder{}

	var x, y int
	for i := 0; i < len(g.State); i++ {
		// each node will be printed like with Free/Total
		// square brackets around the target data.
		x, y = i%g.Width, i/g.Width
		if g.Interested[0] == x && g.Interested[1] == y {
			fmt.Fprintf(s, "[%3d/%3d]", g.State[i][USED], g.State[i][USED]+g.State[i][FREE])
		} else {
			fmt.Fprintf(s, "(%3d/%3d)", g.State[i][USED], g.State[i][USED]+g.State[i][FREE])
		}
		if x == g.Width-1 {
			s.WriteByte('\n')
		} else {
			s.WriteByte(' ')
		}
	}
	return s.String()
}

type GridCache map[string]struct{}

func (gc GridCache) HasNotSeen(g *Grid) bool {
	s := g.String()
	if _, ok := gc[s]; !ok {
		gc[s] = struct{}{}
		return true
	}
	return false
}

// for part 2 we need to work on a plane.
// we calculate the width/height after we
// have calculated the list of nodes.
// then we need to sort the nodes into
// order so they are at index [x + Width * y]
type Grid struct {
	Width      int
	State      [][2]int
	Interested [2]int // the (current) x,y of the wanted data
}

func (g *Grid) IsDone() bool {
	return g.Interested == [2]int{0, 0}
}

const (
	USED = 0
	FREE = 1
)

// how can we make this into an immutable state.
// once we have the list of nodes, we can simplify the
// problem into [][2]int{free,used}

// so then we can go from one grid to another via all available moves.
func (g *Grid) AvailableMoves() []*Grid {
	moves := []*Grid{}
	var used, x, y int
	var isInteresting bool
	for i := 0; i < len(g.State); i++ {
		// find all possible moves.
		// only to the adjacent nodes IF they have enough space.
		// we do need to keep track of the data from 0,H if moved
		// the dataAt func return 0,0 for nodes outside the grid,
		// so we are OK not to bounds check here.
		// we only consider moving data FROM this node, the other way happens in
		// the other node.
		// NEVER move the important data to a node where any adjacent nodes would
		// not be able to move the data on again (or larger than any node on the route
		// to the goal!)
		//
		used = g.State[i][USED]
		if used == 0 {
			// nothing to move
			continue
		}
		y = i / g.Width
		x = i % g.Width

		isInteresting = g.Interested[0] == x && g.Interested[1] == y

		// left
		if g.DataFreeAt(x-1, y) >= used {
			if !isInteresting || g.OkToMoveInteresingDataTo(used, x-1, y) {
				moves = append(moves, g.Move(x, y, x-1, y))
			}
		}

		// right
		if g.DataFreeAt(x+1, y) >= used {
			if !isInteresting || g.OkToMoveInteresingDataTo(used, x+1, y) {
				moves = append(moves, g.Move(x, y, x+1, y))
			}
		}

		// up
		if g.DataFreeAt(x, y-1) >= used {
			if !isInteresting || g.OkToMoveInteresingDataTo(used, x, y-1) {
				moves = append(moves, g.Move(x, y, x, y-1))
			}
		}

		// down
		if g.DataFreeAt(x, y+1) >= used {
			if !isInteresting || g.OkToMoveInteresingDataTo(used, x, y+1) {
				moves = append(moves, g.Move(x, y, x, y+1))
			}
		}
	}

	return moves
}

func (g *Grid) OkToMoveInteresingDataTo(amount, x, y int) bool {
	maxData := g.DataUsedAt(x, y) + amount
	if g.DataCapacityAt(x+1, y) >= maxData {
		return true
	}
	if g.DataCapacityAt(x-1, y) >= maxData {
		return true
	}
	if g.DataCapacityAt(x, y+1) >= maxData {
		return true
	}
	if g.DataCapacityAt(x, y-1) >= maxData {
		return true
	}
	return false

}

func (g *Grid) Move(srcX, srcY, dstX, dstY int) *Grid {

	next := &Grid{
		Width:      g.Width,
		State:      make([][2]int, len(g.State)),
		Interested: g.Interested,
	}
	copy(next.State, g.State)

	if srcX == g.Interested[0] && srcY == g.Interested[1] {
		// interested data has moved
		next.Interested = [2]int{dstX, dstY}
	}
	src := srcX + srcY*g.Width
	dst := dstX + dstY*g.Width

	// fmt.Printf("Move Data From Node (%2d,%2d) [%2d] -> (%2d,%2d) [%2d]\n", srcX, srcY, src, dstX, dstY, dst)
	// fmt.Printf("BEFORE        (free %2d, used %2d) -> (free %2d, used %2d)\n",
	// 	next.State[src][FREE],
	// 	next.State[src][USED],
	// 	next.State[dst][FREE],
	// 	next.State[dst][USED])

	// add used on src to used on dst
	// add used on src to free on src
	// subtract used on src from free on dst
	// set used on src to 0
	next.State[dst][USED] += next.State[src][USED]
	next.State[src][FREE] += next.State[src][USED]
	next.State[dst][FREE] -= next.State[src][USED]
	next.State[src][USED] = 0

	// fmt.Printf("AFTER         (free %2d, used %2d) -> (free %2d, used %2d)\n",
	// 	next.State[src][FREE],
	// 	next.State[src][USED],
	// 	next.State[dst][FREE],
	// 	next.State[dst][USED])

	return next
}

func (g *Grid) DataFreeAt(x, y int) int {
	// bounds check
	if x < 0 || x >= g.Width {
		return 0
	}
	i := x + g.Width*y
	if i >= 0 && i < len(g.State) {
		return g.State[i][FREE]
	}
	return 0
}

func (g *Grid) DataCapacityAt(x, y int) int {
	// bounds check
	if x < 0 || x >= g.Width {
		return 0
	}
	i := x + g.Width*y
	if i >= 0 && i < len(g.State) {
		return g.State[i][FREE] + g.State[i][USED]
	}
	return 0
}

func (g *Grid) DataUsedAt(x, y int) int {
	// bounds check
	if x < 0 || x >= g.Width {
		return 0
	}
	i := x + g.Width*y
	if i >= 0 && i < len(g.State) {
		return g.State[i][USED]
	}
	return 0
}

func (g *Grid) IsWall(x, y int) bool {
	i := x + g.Width*y
	// this is an assumption too
	return g.State[i][USED] > 400
}

func (g *Grid) MoveToTop(p int) (steps int, x int) {
	x = p % g.Width
	y := p / g.Width

	for y > 0 {
		steps++
		// assume data will always fit, we are using the empty node
		// try up first
		if g.IsWall(x, y-1) {
			// gotta go left (assume walls are horizontal)
			x--
		} else {
			y--
		}
	}
	return
}
