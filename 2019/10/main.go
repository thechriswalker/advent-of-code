package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
	"time"

	"../../aoc"
)

func main() {
	aoc.Run(2019, 10, solve1, solve2)
}

const drawSolutionTwo = true

// Implement Solution to Problem 1
func solve1(input string) string {
	field := parseInput(input)
	max := 0
	var asteroid *Asteroid
	for i, a := range field.Asteroids {
		if a == nil {
			continue
		}
		if count := a.CountObservable(field.Asteroids[i+1:]); count > max {
			max = count
			asteroid = a
		}
	}
	log.Printf("\nBest position is %d,%d (%d)\n", asteroid.X, asteroid.Y, max)
	return fmt.Sprintf("%d", max)
}

func solve2(input string) string {
	// from problem one solution our station is at 23,19
	return solve2at(input, 23, 19, drawSolutionTwo)
}

// Implement Solution to Problem 2
func solve2at(input string, x, y int, draw bool) string {
	// to find the positions of all the other asteroids
	// in order of distance we need to store all the angle/distances
	// for all the asteroids, in relation to x,y
	field := parseInput(input)
	// we know this from round one.
	index := field.Width*y + x

	station := field.Asteroids[index]

	vectors := createDestroyList(station, field.Asteroids)
	// we have sorted by theta into buckets,
	// as created a doubly linked list, starting a vectors.
	// as we destroy an asteroid if no more exist in the bucket,
	// we can easily cut it out of the list. So we need only
	// 200 iterations.
	var destroyed *Asteroid

	if draw {
		// let's watch.
		fmt.Printf("\n%s", input)
		field.markAsteroid(station, "\x1b[1;96m#\x1b[0m")
	}
	for i := 0; i < 200; i++ {
		// if we destroyed one last tick, make it just red
		if destroyed != nil && draw {
			field.markAsteroid(destroyed, "\x1b[1;31m#\x1b[0m")
		}
		//destroy one element, move the list on the the next theta
		destroyed = vectors.DestroyNext()
		if draw {
			field.markAsteroid(destroyed, "\x1b[1;93m#\x1b[0m")
			time.Sleep(30 * time.Millisecond)
		}
	}
	// 200th destroyed asteroid is...
	log.Printf("200th destroyed asteroid is: (%d,%d)", destroyed.X, destroyed.Y)

	return fmt.Sprintf("%d", 100*destroyed.X+destroyed.Y)
}

// assuming we are on a terminal at the start (x = 0) of the line below the grid (y = height)
//
func (f *Field) markAsteroid(a *Asteroid, w string) {
	// move to asteroid X lines up
	// output string
	// move to start of line, the same number of lines down
	linesUp := f.Height - a.Y
	fmt.Fprintf(os.Stdout, "\x1b[%dF\x1b[%dC%s\x1b[%dE", linesUp, a.X, w, linesUp)
}

type DestroyList struct {
	current *DestroyBucket
}

type DestroyBucket struct {
	prev  *DestroyBucket
	next  *DestroyBucket
	roids []*Asteroid
}

type Pos struct {
	roid            *Asteroid
	theta, distance float64
}

type PosSlice []*Pos

var _ sort.Interface = PosSlice{}

func (ps PosSlice) Len() int      { return len(ps) }
func (ps PosSlice) Swap(i, j int) { ps[i], ps[j] = ps[j], ps[i] }
func (ps PosSlice) Less(i, j int) bool {
	if ps[i].theta == ps[j].theta {
		return ps[i].distance < ps[j].distance
	}
	return ps[i].theta < ps[j].theta
}

func createDestroyList(station *Asteroid, all []*Asteroid) *DestroyList {
	// first find all the vectors and distances to all the other asteroids
	// and sort by theta asc then distance asc we will bucket them afterwards.
	list := PosSlice{}
	for _, roid := range all {
		if roid != nil && roid != station {
			v, d := station.VectorTo(roid)
			list = append(list, &Pos{roid: roid, theta: v, distance: d})
		}
	}
	sort.Sort(list)
	// now we gather this list into the linked-list
	first := &DestroyBucket{
		prev:  nil,
		next:  nil,
		roids: []*Asteroid{},
	}
	curr := first
	lastTheta := list[0].theta
	for i, x := range list {
		_ = i
		//fmt.Printf("[%04d] (%02d, %02d) distance %02.3f, vector %02.3f\n", i, x.roid.X, x.roid.Y, x.distance, x.theta)
		if lastTheta != x.theta {
			// new link in the chain.
			curr = &DestroyBucket{
				prev:  curr,
				next:  nil,
				roids: []*Asteroid{x.roid},
			}
			// doubly linked!
			curr.prev.next = curr
			lastTheta = x.theta
		} else {
			curr.roids = append(curr.roids, x.roid)
		}
	}
	// now link it up!
	curr.next = first
	first.prev = curr

	return &DestroyList{
		current: first,
	}
}

func (d *DestroyList) DestroyNext() *Asteroid {
	// take one from the current list.
	// if current list empty, chop it out (so we skip it next time)
	for len(d.current.roids) == 0 {
		if d.current == d.current.next {
			return nil // list empty!
		}
		// cut out the current
		// set the prev.next to next.prev
		d.current.prev.next = d.current.next.prev
		// set the next's prev to current prv
		d.current.next.prev = d.current.prev.next
		d.current = d.current.next
	}
	// now slice out the first element of the list (shift)
	roid := d.current.roids[0]
	d.current.roids = d.current.roids[1:]
	// move on to the next list
	d.current = d.current.next
	return roid
}

// Field is the plane, and it has asteroids and blanks on it.
type Field struct {
	Input         string
	Width, Height int
	Asteroids     []*Asteroid // nil means nothing at that point, index is width * X + Y
}

// to count how many asteroids each other one can see, we could count
// both ways at once, as if A can see B then B can see A. But that
// is tricky to do, so maybe in part two, otherwise my alogithm is n^2
// which might be enough for the tests, but probably wont for the main
// to determine whether that path is blocked or not we don't store the
// nodes, but the reduced vector, i.e. just direction. if the paths
// are block the directions will match. So a set of the directions yield
// the number in it's cardinality.
type Asteroid struct {
	X, Y int
	// keyed on direction, value is distance
	Vectors     map[float64]float64
	Observables map[float64]*Asteroid // the closest seen asteroid on that vector
}

// zero angle is vertical up
func (a *Asteroid) VectorTo(b *Asteroid) (float64, float64) {
	// to make up 0
	//log.Printf("calculating angle from %#v to %#v\n", a, b)
	dy, dx := float64(b.Y-a.Y), float64(b.X-a.X)
	distance := math.Sqrt(math.Pow(dy, 2) + math.Pow(dx, 2))

	// this gives a value from -Pi to Pi, where -Pi is
	// left and 0 is right.
	// but we want 0 at the top

	// this took some trial and error and drawing pretty pictures...

	//theta := math.Atan2(dy, dx)
	//theta := math.Atan2(-1*dy, -1*dx)
	//theta := math.Atan2(-1*dy, dx)
	//theta := math.Atan2(dy, -1*dx)
	//theta := math.Atan2(dx, dy) // moves backwards!
	//theta := math.Atan2(-1*dx, -1*dy)
	theta := math.Atan2(-1*dx, dy)
	// theta := math.Atan2(dx, -1*dy)

	return theta, distance
}

func (a *Asteroid) CountObservable(others []*Asteroid) int {
	for _, b := range others {
		if b != nil && b != a {
			theta, d := a.VectorTo(b)
			if min, found := a.Vectors[theta]; !found || min > d {
				a.Vectors[theta] = d
				a.Observables[theta] = b
			}
			// mark b in return (we always progress through the set linearly)
			theta = reverse(theta)
			if min, found := b.Vectors[theta]; !found || min > d {
				b.Vectors[theta] = d
				b.Observables[theta] = a
			}
		}
	}
	return len(a.Observables)
}

func NewAsteroid(x, y int) *Asteroid {
	return &Asteroid{
		X: x, Y: y,
		Vectors:     map[float64]float64{},
		Observables: map[float64]*Asteroid{},
	}
}

func parseInput(input string) *Field {
	rd := bufio.NewScanner(strings.NewReader(input))
	a := []*Asteroid{}

	var x, y int
	for rd.Scan() {
		for i, c := range rd.Text() {
			if c == '#' {
				a = append(a, NewAsteroid(i, y))
			} else {
				a = append(a, nil)
			}
			if i > x {
				x = i
			}
		}
		y++
	}
	return &Field{
		Height:    y,
		Width:     x + 1,
		Asteroids: a,
	}
}

func reverse(angle float64) float64 {
	if angle < math.Pi {
		return angle + math.Pi
	}
	return angle - math.Pi
}
