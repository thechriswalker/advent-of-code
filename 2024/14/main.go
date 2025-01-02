package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 14, solve1, solve2)
}

// available to change in tests
var GridSize = aoc.Vec2(101, 103)

// Implement Solution to Problem 1
func solve1(input string) string {
	// 4 quadrants. We count the number of points in each quadrant
	quadrants := []int{0, 0, 0, 0}

	ticks := 100

	//g := aoc.NewFixedByteGrid(GridSize.X, GridSize.Y, '.', nil)

	aoc.MapLines(input, func(line string) error {
		var px, py, vx, vy int
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		if err != nil {
			return err
		}
		//	fmt.Println("robot at", px, py, "moving", vx, vy)
		// after 100 ticks, see which quadrant the point is in.
		// we need to make these "positive" tho, so we may have to add the grid size to them
		px = (px + (vx * ticks)) % GridSize.X
		py = (py + (vy * ticks)) % GridSize.Y
		for px < 0 {
			px += GridSize.X
		}
		for py < 0 {
			py += GridSize.Y
		}
		//	fmt.Println("final position:", px, py)
		//	g.Set(px, py, '#')
		switch {
		case px < GridSize.X/2 && py < GridSize.Y/2:
			quadrants[0]++
		case px > GridSize.X/2 && py < GridSize.Y/2:
			quadrants[1]++
		case px < GridSize.X/2 && py > GridSize.Y/2:
			quadrants[2]++
		case px > GridSize.X/2 && py > GridSize.Y/2:
			quadrants[3]++
		case px == GridSize.X/2 || py == GridSize.Y/2:
			// ignore points on the axes
		default:
			// this should never happen
		}
		return nil
	})

	//aoc.PrintByteGrid(g, nil)

	//fmt.Println("quadrants:", quadrants)
	// multiply the numbers
	safety := quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]

	// 220827600 -- too low (had the grid x,y reversed)

	return fmt.Sprint(safety)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// robots should arrange themsolves into a christmas tree shape...
	// do we do this by simply watching? and narrowing in on the correct value?
	// robots := [][2]aoc.V2{}
	// aoc.MapLines(input, func(line string) error {
	// 	var px, py, vx, vy int
	// 	_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
	// 	if err != nil {
	// 		return err
	// 	}
	// 	robots = append(robots, [2]aoc.V2{{X: px, Y: py}, {X: vx, Y: vy}})
	// 	return nil
	// })

	// let's see if we find a full box at any point.
	// tick := 0
	// max := 7 // I know we are looking for eight and this cuts down on the output size.
	// for {
	// 	tick++
	// 	if c := findLongestLineAtTick(robots, tick, max); c >= 8 {
	// 		fmt.Println("possibility at", tick)
	// 		break
	// 	} else {
	// 		if max < c {
	// 			max = c
	// 			fmt.Println("max", max, "at", tick)
	// 		}
	// 	}
	// 	if tick%10000 == 0 {
	// 		fmt.Println("tick", tick)
	// 	}
	// }

	// tick := 7660
	// stdin := bufio.NewReader(os.Stdin)

	// for {
	// 	tick++
	// 	findLongestLineAtTick(robots, tick, 0)
	// 	// wait for input.
	// 	stdin.ReadBytes('\n')
	// }

	// wow. the image is pretty crap (or my grid drawing is really crap)
	// no idea why my render doesn't work...
	// it probably means that my algorithm is wrong for calculating the positions,
	// but somehow still gets the right answers for test and part1...
	// tick: 7672
	// OK I found it - I missed the `(x + dx * t) % X` instead doing `x + (dx * t) % X`
	// so, the error wasn't there in part 1.
	// of course it's still pretty cool that it got the right answer anyway... or horrifying.
	/*
		......................................................#..............................................
		.....................................................................................................
		.....................................................................................................
		.....................................................................................................
		......................................#.#.#######.....#...#.###...#..................................
		......................................#..............................................................
		....................................................................#................................
		......................................#.............................#................................
		......................................#..............................................................
		.....................................................................................................
		......................................#...................................#..........................
		.....................................................#..................#............................
		......................................#...........###...............#................................
		......................................#............##.#.............#................................
		......................................#.............###.............#................................
		...................................................#..#.............#....................#...........
		......................................#...........#.....#...........#......................#.........
		......................................#..........#.###..###.........#................................
		......................................#........#.#.##..###...........................................
		......................................#...........########..........#...............#.............#..
		......................................#..........#.#..#.#.#.........#................................
		..................................................#..#####..........#..............#.................
		......................................#........##.#.#.####.#........#.#..............................
		......................................#.......##.#......####.#......#................................
		......................................#.........########..##.........................................
		......................................#.......#..#.#......###...........#............................
		.................................................###.#..#...........#.......##.......................
		................................................##..##.#.#.##.#......................................
		......................................#.....###.####.###...#####..................................#..
		......................................#.............##..............#................................
		.#..................................................................#........................#.......
		......................................#..............##............................#.................
		....................................................................#................................
		..............................#.......#.............................#................................
		.....................................................................................................
		....................................................................#.......#........................
		#.....................................#.....#.#.##........#..#.#.##.#................................
		.....................................................................................................
		........................................................#.......................................#....
	*/

	// lets iterate our way to 7672
	// ok that worked... so I'll leave just this, because it is cool and draws the image.
	r := NewRobots(input)
	// for i := 0; i < 7672; i++ {
	// 	r.Tick(1)
	// }
	r.Tick(7672)
	fmt.Println("tick 7672")
	r.Print()

	return fmt.Sprint(r.ticks)
}

var _ = findLongestLineAtTick

func findLongestLineAtTick(robots [][2]aoc.V2, tick int, max int) int {
	// make a map of all the positions.
	m := map[aoc.V2]bool{}
	for _, r := range robots {
		p := aoc.Vec2(
			(r[0].X+(r[1].X*tick))%GridSize.X,
			(r[0].Y+(r[1].Y*tick))%GridSize.Y,
		)
		for p.X < 0 {
			p.X += GridSize.X
		}
		for p.Y < 0 {
			p.Y += GridSize.Y
		}
		m[p] = true
	}

	// we will do the print here as we have already done the work...
	g := aoc.NewFixedByteGrid(GridSize.X, GridSize.Y, '.', func(v aoc.V2) byte {
		if m[v] {
			return '#'
		}
		return '.'
	})

	// now we havea  grid, let's find the longest vertical line
	longest := 0
	aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
		if v.Y > GridSize.Y-longest {
			// we can't get longer..
			return
		}
		if b == '#' {
			// we have a point, let's see how long the line is
			l := 1
			for {
				v = v.Add(aoc.South)
				if b, _ := g.Atv(v); b == '#' {
					l++
				} else {
					break
				}
			}

			if l > longest {
				longest = l
			}
		}
	})

	if longest > max {
		aoc.PrintByteGrid(g, nil)
	}

	return longest
}

// lets produce an iterating version of the grid.
type Robots struct {
	counts     map[aoc.V2]int
	positions  []aoc.V2
	velocities []aoc.V2
	grid       aoc.ByteGrid
	ticks      int
}

func NewRobots(input string) *Robots {
	r := &Robots{}
	r.grid = aoc.NewFixedByteGrid(GridSize.X, GridSize.Y, '.', nil)
	aoc.MapLines(input, func(line string) error {
		var px, py, vx, vy int
		_, err := fmt.Sscanf(line, "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		if err != nil {
			return err
		}
		r.positions = append(r.positions, aoc.V2{X: px, Y: py})
		r.velocities = append(r.velocities, aoc.V2{X: vx, Y: vy})
		return nil
	})
	r.counts = make(map[aoc.V2]int, GridSize.X*GridSize.Y)
	for i := range r.positions {
		r.counts[r.positions[i]] = r.counts[r.positions[i]] + 1
	}
	for p := range r.counts {
		r.grid.Setv(p, '#')
	}
	return r
}

func (rs *Robots) Print() {
	aoc.PrintByteGridFunc(rs.grid, func(x, y int, b byte) aoc.Color {
		switch rs.counts[aoc.V2{X: x, Y: y}] {
		case 0:
			return aoc.NoColor
		case 1:
			return aoc.BoldGreen
		case 2:
			return aoc.BoldCyan
		default:
			return aoc.BoldRed
		}
	})
}

func (rs *Robots) Tick(t int) {
	rs.ticks += t
	var px, py int
	for i := range rs.positions {
		// remove from current position
		rs.counts[rs.positions[i]] = rs.counts[rs.positions[i]] - 1
		px = (rs.positions[i].X + (rs.velocities[i].X * t)) % GridSize.X
		for px < 0 {
			px += GridSize.X
		}
		py = (rs.positions[i].Y + (rs.velocities[i].Y * t)) % GridSize.Y
		for py < 0 {
			py += GridSize.Y
		}
		rs.positions[i] = aoc.V2{X: px, Y: py}
		// add to new position
		rs.counts[rs.positions[i]] = rs.counts[rs.positions[i]] + 1
	}
	for p := range rs.counts {
		if rs.counts[p] > 0 {
			rs.grid.Setv(p, '#')
		} else {
			rs.grid.Setv(p, '.')
		}
	}
}
