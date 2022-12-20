package main

import (
	"fmt"
	"math"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 15, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return solve1x(input, 2000000)
}

// because the test and the real use different values for `row“
func solve1x(input string, row int) string {
	sensors := parseSensorsAndBeacons(input)
	//fmt.Println(sensors)

	// naive solution failed for the bigger picture.
	// which was to plot all points...
	// there's a trick somewhere.
	// how about we find the bounds of possibility along the row we care about.
	// then iterate along that, find if there are any sensors in range, or a beacon present
	// for a start we can discount any sensors whose closest beacon is row +- md.

	findInRange := func(x int) []int {
		var s []int
		for i := range sensors {
			d := sensors[i].Position.MD(pos{x: x, y: row})
			if d <= sensors[i].MD {
				s = append(s, i)
			}
		}
		return s
	}

	max := 0
	left, right := math.MaxInt, math.MinInt
	for _, s := range sensors {
		if s.MD > max {
			max = s.MD
		}
		if s.Position.x > right {
			right = s.Position.x
		}
		if s.Position.x < left {
			left = s.Position.x
		}
	}
	sum := 0

	//	fmt.Println("max:", max, "left:", left, "right:", right)

	for x := left - max; x <= right+max; x++ {
		// how far is the closest sensor.
		p := pos{x: x, y: row}
		// and is it further that that sensor's closest beacon.
		s := findInRange(x)
		for _, i := range s {
			if sensors[i].ClosestBeacon == p || sensors[i].Position == p {
				// it is at the beacon or the sensor
				// not this time.
				break
			}
			d := sensors[i].Position.MD(p)
			//fmt.Println("pos:", p, "closest:", sensors[i].Position, "range:", sensors[i].MD, "d:", d, "excluded?", d <= sensors[i].MD)
			// if this sensor covers it, then great.
			if d <= sensors[i].MD {
				sum++
				break
			}
			// nope, try the next sensor.
		}
	}

	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return solve2x(input, 4000000)
}

func solve2x(input string, xymax int) string {
	sensors := parseSensorsAndBeacons(input)
	// find the only possible location for the beacon
	// within X,Y between 0,xymax
	// we have just found how to find the positions that it could be in a single row.
	// now in a square.
	// what do we know.
	// - we know that the missing spot will be md+1 from a sensor.
	// so we can find sensors in range of all points md+1 from each sensor
	// that are also within our box.
	// ???
	// sounds slow.
	checked := map[pos]struct{}{}
	for _, s := range sensors {
		// find all point along the border that are just out of it's range. (and we have tested yet)
		pts := s.FindPointsOutOfRange(xymax)

		// for all those points,
		for _, p := range pts {
			if _, ok := checked[p]; ok {
				continue
			}
			checked[p] = struct{}{}
			// is there another sensor in range?
			if p.OutOfRange(sensors) {
				return fmt.Sprint(p.x*4000000 + p.y)
			}
		}
	}
	return "<unsolved>"
}

type pos struct{ x, y int }

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (a pos) MD(b pos) int {
	x, y := abs(a.x-b.x), abs(a.y-b.y)
	return x + y
}

func (a pos) OutOfRange(sensors []Sensor) bool {
	for _, s := range sensors {
		if a.MD(s.Position) <= s.MD {
			return false
		}
	}
	return true
}

type Sensor struct {
	Position      pos
	ClosestBeacon pos
	MD            int
}

func parseSensorsAndBeacons(input string) []Sensor {
	var sensors []Sensor

	aoc.MapLines(input, func(line string) error {
		var s Sensor
		fmt.Sscanf(line, "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &(s.Position.x), &(s.Position.y), &(s.ClosestBeacon.x), &(s.ClosestBeacon.y))
		s.MD = s.Position.MD(s.ClosestBeacon)
		sensors = append(sensors, s)
		return nil
	})

	return sensors
}

func (s Sensor) FindPointsOutOfRange(xymax int) []pos {
	var pts []pos
	// all of the diamond points around s.Position by 1 extra pixe;
	p := s.Position
	for i := 0; i <= s.MD+1; i++ {
		d := s.MD + 1 - i
		if p.x+i <= xymax {
			if p.y+d <= xymax {
				pts = append(pts, pos{x: p.x + i, y: p.y + d})
			}
			if p.y-d > 0 {
				pts = append(pts, pos{x: p.x + i, y: p.y - d})
			}
		}
		if p.x-i > 0 {
			if p.y+d <= xymax {
				pts = append(pts, pos{x: p.x - i, y: p.y + d})
			}
			if p.y-d > 0 {
				pts = append(pts, pos{x: p.x - i, y: p.y - d})
			}
		}
	}
	return pts
}
