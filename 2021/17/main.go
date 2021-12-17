package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 17, solve1, solve2)
}

type Target struct {
	x1, x2, y1, y2 int
}

func (t *Target) Hit(x, y int) (hit bool, overshot bool) {
	undershot := x < t.x1 || y > t.y1
	overshot = x > t.x2 || y < t.y2

	return !(undershot || overshot), overshot
}

func sumBetweenInts(max, min int) int {
	return ((max - min) + 1) * (max + min) / 2
}

//The distance travelled decreases by 1 each time
// so the distance travelled in step X for a velocity V is the sum
// (max-min+1) * (min+max) /2
// (V - (V - X) + 1) * ((V-X)+V) / 2
// (X+1)*(2V-X) / 2
// but it doesn't keep increasing, it will hit delta-zero
// after V steps.
func xtravel(v int, steps int) int {
	i := steps - 1
	if i > v {
		i = v
	}
	return sumBetweenInts(v, v-i)
}

// We actually need to check "negative" y values.
func ytravel(v int, steps int) int {
	if v < 0 {
		// this is just the same as the next bit
		// without the "ignore" section.
		// but we will pretend the velocity
		// is positive
		v = -1 * v
	} else {
		ignore := (2 * v) + 1
		if steps <= ignore {
			// this is our sentinel value.
			// anything >0 means above the surface
			return 1
		}
		// now pretend we started at "steps"
		steps = steps - ignore
		// however the initial velocity of
		// the probe is now v (down)
		// so we leave it positive
		v = v + 1
	}
	// we have now fudged it so steps is the number
	// of steps AFTER we start going down.
	// Also we have made the velocity positive so
	// it INCREASES by 1 each step
	// we want the sum from v to v+steps,
	d := sumBetweenInts(v+steps-1, v)
	// now invert it again
	return -1 * d
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// The X and Y velocities are independent
	// So we can find all X where there is a chance of hitting the target
	// There is a bound, the steps where we might hit with X are
	// going to be X >= 0 and X <= the right hand bound of the target.
	// (that upper bound is probably wrong as it will be a Flat shot)

	// So we calculate all possible (Vx, steps) pairs which is actually unbounded
	// if steps > Vx, limit to steps <= Vx.

	// Now we look at Y and see if there is a Y at each "steps", that produces
	// a hit. If steps = Vx, we can go higher - this is where it is tricky?
	// how do we rule out going higher to produce a hit?

	// 1  2  3  4  5  6  7  8  9 10 11 12 13 14
	// 4  3  2  1  0 -1 -2 -3 -4 -5 -6 -7 -8 -9
	// 4, 7, 9 10 10  9  7  4  0 -5 -11 -18
	// OK as it is symmetrical the 2*steps + 1th value will be 0
	// and the NEXT value will be -V so V must be within range
	// of the bottom of the target.

	// OK with these to bounds, we can brute force a solution.
	t := Target{}
	// note that it is x1,x2, y2,y1 as a y values are reversed in both
	// the test and the input. we _could_ normalise it after...
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &(t.x1), &(t.x2), &(t.y2), &(t.y1))

	// there is a lower bound, but we might as well not bother calculating it
	// the upper bound on the x value is the far right of the target
	// the upper bound on the y value is the absolute distance to bottom of the target

	// we can find the answer more easily if we can find an X where
	// Dx == 0 in the range of the target.
	tx := 0
	for x := 1; x < t.x2; x++ {
		dx := xtravel(x, x) // this is the max distance we go.
		if dx > t.x1 {
			// boom.
			tx = x
			break
		}
	}
	// found optimal X
	if tx == 0 {
		panic("no easy solution")
	}
	// now max Y is Y = (-1 * t.y2) -1
	// and the height is actually the same calculation as the xtravel.
	ymax := (-1 * t.y2) - 1

	highest := xtravel(ymax, ymax)

	//fmt.Printf("Optimal solution is (%d,%d) with height %d\n", tx, ymax, highest)

	return fmt.Sprint(highest)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	t := Target{}
	// note that it is x1,x2, y2,y1 as a y values are reversed in both
	// the test and the input. we _could_ normalise it after...
	fmt.Sscanf(input, "target area: x=%d..%d, y=%d..%d", &(t.x1), &(t.x2), &(t.y2), &(t.y1))

	count := 0
	xmin := 1
	xmax := t.x2
	ymin := t.y2
	ymax := (-1 * t.y2) - 1
	//fmt.Printf("Xmin=%d, Xmax=%d, Ymin=%d, Ymax=%d\n", xmin, xmax, ymin, ymax)

	for x := xmin; x <= xmax; x++ {
		for y := ymin; y <= ymax; y++ {
			if isSolution(&t, x, y) {
				count++
			}
		}
	}

	return fmt.Sprint(count)
}

func isSolution(t *Target, vx, vy int) bool {
	var hit, over bool
	var x, y int
	for s := 1; ; s++ {
		x, y = xtravel(vx, s), ytravel(vy, s)
		hit, over = t.Hit(x, y)
		// fmt.Printf("Test initial V(%d,%d) after %d steps => (%d, %d) (hit:%v, over:%v)\n",
		// 	vx, vy, s, x, y, hit, over)
		if hit {

			return true
		}
		if over {
			return false
		}
	}
}
