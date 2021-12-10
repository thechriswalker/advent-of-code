package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 18, solve1, solve2)
}

// this will help keeping our set of keys collected easily comparable
type Keys uint32

func fromKeyRune(r rune) Keys {
	x := int(r - 'a')
	return 1 << x
}

func fromDoorRune(r rune) Keys {
	x := int(r - 'A')
	return 1 << x
}

func (k Keys) Add(n Keys) Keys {
	return k | n
}

func (k Keys) HasAll(n Keys) bool {
	return k&n == n
}

func (k Keys) Without(n Keys) Keys {
	return k & ^n
}
func (k Keys) String() string {
	// each bit represents a key.
	set := []rune{}
	for r := 'a'; r <= 'z'; r++ {
		if k.HasAll(fromKeyRune(r)) {
			set = append(set, r)
		}
	}
	return string(set)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// This is ugly, but it works
	// shortest path to all keys.
	// first model grid.
	g := createMaze(input)
	// I'm not sure if we should just do our generic maze solving
	// breadth first search, keeping track of keys collected and
	// so whether we should get all the distances up front then solve
	// based on that.
	// I think the latter may be easier.
	// create a cache of the distance between all key pairs AND
	// which doors are in the way.
	//fmt.Println("Calculating KeyPairs...")
	g.CacheKeyPairs()
	//fmt.Println(g)

	// then we find the distance from the origin to all freely available keys.
	score := walkMaze(g, g.Origin, 0)

	return fmt.Sprint(score)
}

func walkMaze(g *Maze, origin int, initialKeys Keys) int {
	current := g.FindInitialKeys(origin, initialKeys)
	var next []*Path

	// The key here is knowing which options to discard.
	// which means keeping a cache of the state.
	visited := map[State]int{}
	// we actually need to keep going until we get a minimum score.
	score := math.MaxInt64
	for {
		next = []*Path{}
		for _, p := range current {
			if g.IsComplete(p) {
				//fmt.Println("Found a solution with score:", p.Steps)
				if p.Steps < score {
					score = p.Steps
				}
			} else if p.Steps > score {
				// just give up on this route...
				// it is not optimal.
			} else {
				next = append(next, g.GetNextOptions(p, visited)...)
			}
		}
		if len(next) == 0 {
			break
		}

		// switch and continue
		current = next
	}
	return score
}

func isKey(r rune) bool {
	return r >= 'a' && r <= 'z'
}

func isDoor(r rune) bool {
	return r >= 'A' && r <= 'Z'
}

func createMaze(in string) *Maze {
	m := &Maze{
		Grid:     []rune{},
		Keys:     map[Keys]int{},
		KeyPairs: map[Keys]map[Keys]*Path{},
	}
	aoc.MapLines(in, func(line string) error {
		m.Height++
		m.Stride = len(line)
		m.Grid = append(m.Grid, []rune(line)...)
		return nil
	})
	// now find the index of all the keys and the origin
	for idx, r := range m.Grid {
		switch {
		case isKey(r):
			k := fromKeyRune(r)
			m.AllKeys = m.AllKeys.Add(k)
			m.Keys[k] = idx
			m.KeyPairs[k] = map[Keys]*Path{}
		case r == ORIGIN:
			m.Origin = idx
		}
	}
	return m
}

type Maze struct {
	// storage for the grid of stuff
	Grid           []rune
	Stride, Height int
	Keys           map[Keys]int
	AllKeys        Keys

	Origin   int                     // start
	KeyPairs map[Keys]map[Keys]*Path // distances and doors between each pair of keys
}

// comparable struct
type State struct {
	Position int
	Keys     Keys
}

func (m *Maze) GetNextOptions(p *Path, visited map[State]int) []*Path {
	// this means find all keys we can get to from here,
	options := []*Path{}
	// we should be at a key.
	r := m.Grid[p.Position]
	if !isKey(r) {
		panic("not at a key!")
	}
	key := fromKeyRune(r)

	for nextKey, route := range m.KeyPairs[key] {
		// do we need this key?
		if p.Keys.HasAll(nextKey) {
			// we don't need it
			//fmt.Printf("Not taking route from %s -> %s as we already have the key (we have: %s)\n", key, nextKey, p.Keys)
			continue
		}
		// can we get to this key? (do our keys cover all the doors?)
		if !p.Keys.HasAll(route.Doors) {
			// no, we can't get there
			//fmt.Printf("Not taking route from %s -> %s as we don't have the door keys (we have: %s, we need: %s)\n", key, nextKey, p.Keys, route.Doors)
			continue
		}
		// check there are no keys we don't have on the way
		if !p.Keys.HasAll(route.Keys) {
			// there are keys on the way we don't have
			//fmt.Printf("Not taking route from %s -> %s as we there are keys enroute we don't have (we have: %s, en route: %s)\n", key, nextKey, p.Keys, route.Keys)
			continue
		}
		// OK this is an option.
		// would the next state be one we have seen?
		state := State{Position: m.Keys[nextKey], Keys: p.Keys.Add(nextKey)}
		steps := p.Steps + route.Steps
		if v, ok := visited[state]; ok && v < steps {
			// we already visited this and the previous visit was less steps
			// lets just ignore this
			//fmt.Printf("Not taking route from %s -> %s as we already visited this position, with the same keys and less steps\n", key, nextKey)
			continue
		}
		visited[state] = steps
		options = append(options, &Path{
			Position: m.Keys[nextKey],
			Steps:    steps,
			Keys:     p.Keys.Add(nextKey),
		})
	}
	return options
}

func (m *Maze) IsComplete(p *Path) bool {
	return p.Keys.HasAll(m.AllKeys)
}

func (m *Maze) CacheKeyPairs() {
	// for each key, find all pairs and the keys required
	for key := range m.Keys {
		// we need to make m.KeyPairs[key] have entries for all other keys.
		m.CreateKeyPairs(key)
	}
	// now lets fix the routes so they
	for k1, kp := range m.KeyPairs {
		for k2, route := range kp {
			route.Keys = route.Keys.Without(k1.Add(k2))
		}
	}
}

func (m *Maze) CreateKeyPairs(key Keys) {
	kp := m.KeyPairs[key]
	curr := []*Path{{Position: m.Keys[key]}}
	cache := map[int]struct{}{m.Keys[key]: {}, -1: {}}
	var next []*Path
	for {
		next = []*Path{}
		for _, p := range curr {
			x, y := aoc.GridCoords(p.Position, m.Stride)
			// try to walk north south east west
			for _, d := range directions {
				di := aoc.GridIndex(x+d[0], y+d[1], m.Stride, m.Height)
				if _, seen := cache[di]; !seen {
					cache[di] = struct{}{}
					at := m.Grid[di]
					switch {
					case at == WALL:
						// dead end.
					case isDoor(at):
						// means we need the key for this door!
						// but also continue
						//fmt.Printf("found a door: %c, adding key %c to required keys: %s", at, keyFor(at), stringKeymap(p.Keys))
						next = append(next, &Path{
							Position: di,
							Steps:    p.Steps + 1,
							Doors:    p.Doors.Add(fromDoorRune(at)),
							Keys:     p.Keys,
						})
					default:
						// key or corridor
						path := &Path{
							Position: di,
							Steps:    p.Steps + 1,
							Keys:     p.Keys,
							Doors:    p.Doors,
						}
						if isKey(at) {
							k := fromKeyRune(at)
							// it's a key, mark that it is on the route
							path.Keys = path.Keys.Add(k)
							// find the paths. only update if shorter or new.
							if existing, ok := kp[k]; !ok || existing.Steps > path.Steps {
								// upsert.
								kp[k] = path
							}
							// the the other way around as well.
							if existing, ok := m.KeyPairs[k][key]; !ok || existing.Steps > path.Steps {
								// upsert.
								m.KeyPairs[k][key] = path
							}
						}
						// this becomes an output
						next = append(next, path)

					}
				}
			}
		}
		if len(next) == 0 {
			break
		}
		curr = next
	}

}

func (m *Maze) String() string {
	sb := strings.Builder{}

	for row := 0; row < m.Height; row++ {
		sb.WriteString(string(m.Grid[row*m.Stride : (row+1)*m.Stride]))
		sb.WriteByte('\n')
	}
	// Now some info:
	x, y := aoc.GridCoords(m.Origin, m.Stride)
	fmt.Fprintf(&sb, "Origin: %d,%d\n", x, y)

	return sb.String()
}

var directions = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

const (
	WALL     = '#'
	CORRIDOR = '.'
	ORIGIN   = '@'
)

// we want to reuse this function. The findAll parameter if true
// will traverse and find all keys "collecting", doors on the way
// if false, we will stop at doors and only collect available keys
func (m *Maze) FindInitialKeys(origin int, initialKeys Keys) []*Path {
	curr := []int{origin}
	// this is where we will store the useful paths
	// ie. those with keys!
	output := []*Path{}
	// a cache of indices we have visited
	// we start with the origin and the "invalid" position
	cache := map[int]struct{}{origin: {}, -1: {}}
	// without backtracking, follow the paths until we find a key
	var next []int
	steps := 0
	for {
		next = []int{}
		steps++
		for _, p := range curr {
			x, y := aoc.GridCoords(p, m.Stride)
			// try to walk north south east west
			for _, d := range directions {
				di := aoc.GridIndex(x+d[0], y+d[1], m.Stride, m.Height)
				if _, seen := cache[di]; !seen {
					cache[di] = struct{}{}
					at := m.Grid[di]
					switch {
					case at == WALL:
						// dead end.
					case isDoor(at):
						// dead end
						// do we have the key?
						if initialKeys.HasAll(fromDoorRune(at)) {
							// ok
							next = append(next, di)
						}
					case isKey(at):
						// this becomes an output
						//fmt.Printf("Found initial key: %c (steps:%d)\n", at, steps)
						output = append(output, &Path{
							Position: di,
							Steps:    steps,
							Keys:     fromKeyRune(at).Add(initialKeys),
						})
					default:
						// assume that is corridor
						next = append(next, di)
					}
				}
			}
		}
		if len(next) == 0 {
			// we are done, no more options
			break
		}
		// else switch current and next
		curr = next
	}
	return output
}

// in KeyPairs, this is the distance to a key and the keys/doors passed to get there.
// in the path finding it is the current number of steps and the keys we have.
// the position is the position we are in. That means nothing in KeyPairs
// doors is not used in the path finding.
type Path struct {
	Position int
	Steps    int
	Keys     Keys
	Doors    Keys
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// the key insight here is that each of the 4 grids can be solved by assuming that any doors for keys
	// NOT in the section have been unlocked already.
	// then we can just sum the minimum distances for each 4.
	// so we just do it 4 times. with different start points and "initial keys"
	g := createMaze(input)
	// now HACK IT
	g.Grid[g.Origin] = '#'
	x, y := aoc.GridCoords(g.Origin, g.Stride)
	for _, d := range directions {
		g.Grid[aoc.GridIndex(x+d[0], y+d[1], g.Stride, g.Height)] = '#'
	}
	g.CacheKeyPairs()
	// now we need the keys in each quadrant
	nw, ne, se, sw := g.KeysByQuadrant()

	// no for the robots
	// NW
	score := walkMaze(g, aoc.GridIndex(x-1, y-1, g.Stride, g.Height), ne.Add(sw).Add(se))
	// NE
	score += walkMaze(g, aoc.GridIndex(x-1, y+1, g.Stride, g.Height), nw.Add(sw).Add(se))
	// SE
	score += walkMaze(g, aoc.GridIndex(x+1, y+1, g.Stride, g.Height), nw.Add(sw).Add(ne))
	// SW
	score += walkMaze(g, aoc.GridIndex(x+1, y-1, g.Stride, g.Height), ne.Add(nw).Add(se))

	return fmt.Sprint(score)
}

func (m *Maze) KeysByQuadrant() (nw, ne, se, sw Keys) {
	for k, idx := range m.Keys {
		x, y := aoc.GridCoords(idx, m.Stride)
		north := x < m.Height/2
		south := !north
		west := y < m.Stride/2
		east := !west
		switch {
		case north && west:
			nw = nw.Add(k)
		case north && east:
			ne = ne.Add(k)
		case south && west:
			sw = sw.Add(k)
		case south && east:
			se = se.Add(k)
		}
	}
	return
}
