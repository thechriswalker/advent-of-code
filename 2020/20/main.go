package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 20, solve1, solve2)
}

type rotation uint8

const (
	rotate0 rotation = iota
	rotate90
	rotate180
	rotate270
)

type direction uint8

const (
	north direction = iota
	east
	south
	west
)

func printEdge(e int) string {
	b := strings.Builder{}
	for i := 0; i < 10; i++ {
		if e&(1<<i) != 0 {
			b.WriteByte('#')
		} else {
			b.WriteByte('.')
		}
	}
	return b.String()
}

func opposite(dir direction) direction {
	switch dir {
	case north:
		return south
	case east:
		return west
	case south:
		return north
	case west:
		return east
	}
	panic("bad direction!")
}

var (
	flips     = []bool{false, true}
	rotations = []rotation{rotate0, rotate90, rotate180, rotate270}
)

func hasIntersection(a, b []int) bool {
	for _, an := range a {
		for _, bn := range b {
			if an == bn {
				return true
			}
		}
	}
	return false
}

type tile struct {
	id     int
	pixels []bool
	flip   bool
	rotate rotation
	edges  map[direction][]int
	conns  map[*tile]struct{}

	// for part 2
	adjacent map[direction]*tile
}

func (t *tile) Connect(other *tile, dir direction) bool {
	// we cannot re-orient THIS one, but we can re-orient the "OTHER"
	_, ok := t.conns[other]
	if !ok {
		return false // these DO NOT CONNECT
	}
	// we need to find a direction
	thisEdge := t.GetEdge(dir)
	otherDir := opposite(dir)
	other.flip = true
	other.rotate = rotate270
	for _, f := range flips {
		other.flip = f
		for _, r := range rotations {
			other.rotate = r
			//			fmt.Printf("Checking edge from:%d --[%d]-> %d, target:%d flip:%v, rotate:%v, edge:%d\n", t.id, dir, other.id, thisEdge, f, r, other.GetEdge(otherDir))
			if other.GetEdge(otherDir) == thisEdge {
				// got it.
				t.adjacent[dir] = other
				other.adjacent[otherDir] = t
				return true
			}
		}
	}
	//fmt.Printf("Failed to connect in direction %d, this:\n%s\nto this:\n%s", dir, t.Draw(), other.Draw())

	return false // cannot connect in this direction
}

func (t *tile) AttemptConnect(other *tile) {
	if _, ok := t.conns[other]; ok {
		return
	}
	if hasIntersection(t.edges[north], other.edges[south]) {
		t.conns[other] = struct{}{}
		other.conns[t] = struct{}{}
	}
	if hasIntersection(t.edges[south], other.edges[north]) {
		t.conns[other] = struct{}{}
		other.conns[t] = struct{}{}
	}
	if hasIntersection(t.edges[east], other.edges[west]) {
		t.conns[other] = struct{}{}
		other.conns[t] = struct{}{}
	}
	if hasIntersection(t.edges[west], other.edges[east]) {
		t.conns[other] = struct{}{}
		other.conns[t] = struct{}{}
	}
}

func (t *tile) String() string {
	return fmt.Sprintf("Tile<%d>", t.id)
}
func (t *tile) Draw() string {
	w := strings.Builder{}
	fmt.Fprintf(&w, "Tile: %d [f:%v, r:%v]\n", t.id, t.flip, t.rotate)
	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if t.At(x, y) {
				w.WriteByte('#')
			} else {
				w.WriteByte('.')
			}
		}
		w.WriteByte('\n')
	}
	e := t.NorthEdge()
	fmt.Fprintf(&w, "North: %s (%d)\n", printEdge(e), e)
	e = t.EastEdge()
	fmt.Fprintf(&w, " East: %s (%d)\n", printEdge(e), e)
	e = t.SouthEdge()
	fmt.Fprintf(&w, "South: %s (%d)\n", printEdge(e), e)
	e = t.WestEdge()
	fmt.Fprintf(&w, " West: %s (%d)\n", printEdge(e), e)
	return w.String()
}

func (t *tile) At(x, y int) bool {
	// get a pixel at the give point, taking rotation into account.
	switch t.rotate {
	case rotate0:
		// do nothing
	case rotate90:
		// 90 clockwise
		x, y = y, 9-x
	case rotate180:
		// 180 degrees
		x, y = 9-x, 9-y
	case rotate270:
		x, y = 9-y, x
	}
	// we only flip along a single axis
	if t.flip {
		x = 9 - x
	}
	return t.pixels[x*10+y]
}

func (t *tile) Edges() (int, int, int, int) {
	return t.NorthEdge(), t.EastEdge(), t.SouthEdge(), t.WestEdge()
}
func (t *tile) GetEdge(dir direction) int {
	switch dir {
	case north:
		return t.NorthEdge()
	case east:
		return t.EastEdge()
	case south:
		return t.SouthEdge()
	case west:
		return t.WestEdge()
	}
	panic("bad direction for edge")
}

func (t *tile) NorthEdge() int {
	// this is the edge at the top (x = 0)
	// left-to-right
	sum := 0
	for i := 0; i < 10; i++ {
		if t.At(0, i) {
			sum |= 1 << i
		}
	}
	return sum
}
func (t *tile) EastEdge() int {
	// this is the edge at the right. (y = 9)
	// top to bottom
	sum := 0
	for i := 0; i < 10; i++ {
		if t.At(i, 9) {
			sum |= 1 << i
		}
	}
	return sum
}
func (t *tile) SouthEdge() int {
	// this is the edge at the bottom (x = 9)
	// left-to-right
	sum := 0
	for i := 0; i < 10; i++ {
		if t.At(9, i) {
			sum |= 1 << i
		}
	}
	return sum
}
func (t *tile) WestEdge() int {
	// this is the edge at the left. (y = 0)
	// top to bottom
	sum := 0
	for i := 0; i < 10; i++ {
		if t.At(i, 0) {
			sum |= 1 << i
		}
	}
	return sum
}

func (t *tile) calculateEdges() {
	for _, f := range flips {
		t.flip = f
		for _, r := range rotations {
			t.rotate = r
			t.edges[north] = append(t.edges[north], t.NorthEdge())
			t.edges[east] = append(t.edges[east], t.EastEdge())
			t.edges[south] = append(t.edges[south], t.SouthEdge())
			t.edges[west] = append(t.edges[west], t.WestEdge())
		}
	}
}

func parseTiles(input string) []*tile {
	rd := strings.NewReader(input)
	stack := []*tile{}
	var id int

	for {
		if _, err := fmt.Fscanf(rd, "Tile %d:\n", &id); err != nil {
			break
		}
		t := &tile{
			id:       id,
			pixels:   make([]bool, 100),
			edges:    map[direction][]int{},
			conns:    map[*tile]struct{}{},
			adjacent: map[direction]*tile{},
		}
		// now read ten lines of ten pixels
		for x := 0; x < 10; x++ {
			for y := 0; y < 10; y++ {
				c, _ := rd.ReadByte()
				switch c {
				case '.', '#':
					t.pixels[x*10+y] = c == '#'
				default:
					panic("unexpected characeter! " + string(c))
				}
			}
			rd.ReadByte() // consume the newline
		}
		t.calculateEdges()
		stack = append(stack, t)
		rd.ReadByte() // consume the newline
	}
	return stack
}

// Implement Solution to Problem 1
func solve1(input string) string {
	tiles := parseTiles(input)

	// note that we actually don't know where our starting tile will be.
	// the idea is we **assume** that each tile only fits in one place.
	// this may be wrong, but we will run with it because it will be easy
	// if we count how many possible connections each tile has in each direction and see what the pattern is
	for i := 0; i < len(tiles)-1; i++ {
		for j := i + 1; j < len(tiles); j++ {
			// if these 2 can connect
			tiles[i].AttemptConnect(tiles[j])
		}
		//	fmt.Printf("Tile %d: connections=%d\n", tiles[i].id, len(tiles[i].conns))
	}
	//fmt.Printf("Tile %d: connections=%d\n", tiles[len(tiles)-1].id, len(tiles[len(tiles)-1].conns))
	// so now we just need to find the corners, which only have 2 connections.
	sum := 1
	for _, t := range tiles {
		if len(t.conns) == 2 {
			sum *= t.id
		}
	}
	return fmt.Sprintf("%d", sum)
}

type image struct {
	pixels   map[[2]int]bool
	monsters map[[2]int]bool
	size     int
	rotate   rotation
	flip     bool
}

func (im *image) At(x, y int) (bool, bool) {
	x, y = im.translate(x, y)
	ok := im.pixels[[2]int{x, y}]
	mon := im.monsters[[2]int{x, y}]

	return ok, mon
}

func (im *image) translate(x, y int) (int, int) {
	switch im.rotate {
	case rotate0:
		// do nothing
	case rotate90:
		// 90 clockwise
		x, y = y, im.size-x-1
	case rotate180:
		// 180 degrees
		x, y = im.size-x-1, im.size-y-1
	case rotate270:
		x, y = im.size-y-1, x
	}
	// we only flip along a single axis
	if im.flip {
		x = im.size - 1 - x
	}
	return x, y
}

func (im *image) SetMonster(x, y int) {
	xi, yi := im.translate(x, y)
	im.monsters[[2]int{xi, yi}] = true
	for _, pos := range monsterPositions {
		xx, yy := im.translate(x+pos[0], y+pos[1])
		im.monsters[[2]int{xx, yy}] = true
	}
}

// a seamonster looks like:
/*012345678901234567890123
                      #
	#    ##    ##    ###
 	 #  #  #  #  #
*/
// so we look for a head at y>=18 && y < width-1
// then we check for the "relative positions" of the other
// marks we want.
var monsterPositions = [][2]int{
	{1, 0}, {1, 1}, {1, -1}, {2, -2}, // the three under the head and the next
	{2, -5}, {1, -6}, {1, -7}, {2, -8}, // hump 1
	{2, -11}, {1, -12}, {1, -13}, {2, -14}, // hump 2
	{1, -18}, {2, -17}, // tail
}

func (im *image) FindSeaMonsters() int {
	countMonsters := 0
	im.monsters = map[[2]int]bool{}  // flatten the map
	for x := 0; x < im.size-3; x++ { // they need to be 3 high
		for y := 18; y < im.size-1; y++ { // they need to be 18 long (at the first row) but the can't be in the last 1
			if ok, _ := im.At(x, y); ok {
				found := true
				for i, pos := range monsterPositions {
					//	fmt.Printf("found %d pieces of seamonster start at %d,%d (checking %d,%d => %d,%d)\n", i+1, x, y, pos[0], pos[1], x+pos[0], y+pos[1])
					_ = i
					if ok, _ := im.At(x+pos[0], y+pos[1]); !ok {
						//not found.
						found = false
						break
					}
				}
				if found {
					//	fmt.Println("FOUND THE MONSTER!")
					im.SetMonster(x, y)
					countMonsters++
				}
			}
		}
	}
	return countMonsters
}

func buildImage(tilemap map[[2]int]*tile) *image {
	size := 12 // 144 tiles
	if len(tilemap) == 9 {
		size = 3
	}
	im := &image{
		pixels:   make(map[[2]int]bool, size*8),
		monsters: map[[2]int]bool{},
		size:     size * 8,
	}
	for p, t := range tilemap {
		for x := 0; x < 8; x++ {
			for y := 0; y < 8; y++ {
				if t.At(x+1, y+1) {
					im.pixels[[2]int{x + p[0]*8, y + p[1]*8}] = true
				}
			}
		}
	}
	return im
}

func (im *image) Draw() {
	fmt.Printf("Image(Rotation:%d, Flip:%v, Size:%d)\n", im.rotate, im.flip, im.size)
	fmt.Printf("     01234567890123456789012345678901234567890\n")
	for x := 0; x < im.size; x++ {
		fmt.Printf("%04d ", x)
		for y := 0; y < im.size; y++ {
			if ok, mon := im.At(x, y); ok {
				if mon {
					fmt.Print("\x1b[1;32m0")
				} else {
					fmt.Print("\x1b[1;31m#")
				}
			} else {
				fmt.Print("\x1b[1;30m.")
			}
		}
		fmt.Print("\x1b[0m\n")

	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	tiles := parseTiles(input)
	size := 3
	if len(tiles) == 144 {
		size = 12
	}
	// for this one we need to build the grid, so that means putting all the images together into a bigger picture.
	// we had 12x12 images of 10 pixels, so without the borders we will have 12 images of 8 pixels, so we will work in
	// 8pixel chunks (bytes)
	// but first we build the grid. this is done by orienting a corner and then finding the possible layout
	// there can be only one.
	// as we find tiles we will compress them into our image.
	// also we need to know the total number of "on" pixels.
	//fmt.Println(tiles)
	for i := 0; i < len(tiles)-1; i++ {
		for j := i + 1; j < len(tiles); j++ {
			// if these 2 can connect
			tiles[i].AttemptConnect(tiles[j])
		}

	}
	twoConnections := map[*tile]struct{}{}
	threeConnections := map[*tile]struct{}{}
	fourConnections := map[*tile]struct{}{}
	// need to find the starting tile that is in the top-left corner.
	var last *tile
	for _, t := range tiles {
		//fmt.Println(t)
		switch len(t.conns) {
		case 2:
			twoConnections[t] = struct{}{}
			last = t
		case 3:
			threeConnections[t] = struct{}{}
		case 4:
			fourConnections[t] = struct{}{}
		}
	}
	if last == nil {
		panic("could not find top-left tile")
	}

	// we need to find a 2-connector that is oriented as top left.
	// take the first corner and orient it so it is on top-left
	var found bool
	for _, f := range flips {
		last.flip = f
		//	fmt.Printf("update start flip: %v\n", f)
		for _, r := range rotations {
			last.rotate = r
			//		fmt.Printf("update start rotation: %v\n", r)
			//	fmt.Println(last.Draw())
			// try to connect to the 3 connections
			hasEast := false
			hasSouth := false
			for other := range threeConnections {
				if last.Connect(other, east) {
					hasEast = true
					continue
				}
				if last.Connect(other, south) {
					hasSouth = true
					continue
				}

			}
			if hasEast && hasSouth {
				//	fmt.Printf("piece is orientated: f:%v, r:%v\n%s", f, r, last.Draw())
				found = true
				break
			}

		}
		if found {
			break
		}
	}
	if !found {
		panic("could not orient starting piece")
	}

	//fmt.Printf("2 connections: %v\n3 connections: %v\n4 connections: %v\n", twoConnections, threeConnections, fourConnections)

	// start with the first corner.
	tilemap := map[[2]int]*tile{
		[2]int{0, 0}: last,
	}
	delete(twoConnections, last)
	dir := east // from last to next
	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			if x == 0 && y == 0 {
				continue
			}
			// we have moved one to the east.
			// so the last tile is fixed and
			// we must orient this one to fit next to
			// it on the west.
			var options map[*tile]struct{}
			if x == 0 && y == size-1 || x == size-1 && y == 0 || x == size-1 && y == size-1 {
				// a corner
				options = twoConnections
			} else if x == 0 || x == size-1 || y == 0 || y == size-1 {
				// an edge (but not corner)
				options = threeConnections
			} else {
				options = fourConnections
			}
			found := false
			for t := range options {
				if last.Connect(t, dir) {
					// good.
					found = true
					tilemap[[2]int{x, y}] = t
					delete(options, t)
					last = t
					//	fmt.Printf("Found tile for pos:%d,%d => %d\n", x, y, t.id)
					break
				}
			}
			if !found {
				// fmt.Println("placed", tilemap)
				// fmt.Println("options", options)
				// fmt.Printf("last:%d edges:%v conns:%v\n", last.id, last.edges[dir], last.conns)
				panic(fmt.Sprintf("could not find tile for position %d,%d", x, y))
			}
			dir = east
		}
		// we will add an x which means move down.
		// the "last" should be the tile directly above.
		// which is y=0 x=x
		last = tilemap[[2]int{x, 0}]
		dir = south
		if last == nil {
			panic("no first element of this row...")
		}
	}
	im := buildImage(tilemap)

	total := len(im.pixels)
	var monsters int
	for _, f := range flips {
		im.flip = f
		//	fmt.Printf("update start flip: %v\n", f)
		for _, r := range rotations {
			im.rotate = r

			im.flip = true
			im.rotate = rotate0
			monsters = im.FindSeaMonsters()
			//im.Draw()
			if monsters != 0 {
				break
			}
		}
		if monsters != 0 {
			break
		}
	}

	//fmt.Printf("Found %d monsters (*15) in %d waves = %d roughness\n", monsters, total, total-len(im.monsters))

	return fmt.Sprintf("%d", total-len(im.monsters))
}
