package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 6, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var times, distances []int

	aoc.MapLines(input, func(line string) error {
		if strings.HasPrefix(line, "Time: ") {
			times = aoc.ToIntSlice(line[6:], ' ')
		}
		if strings.HasPrefix(line, "Distance: ") {
			distances = aoc.ToIntSlice(line[10:], ' ')
		}
		return nil
	})
	n := 1

	for i := range times {
		// how to win this game?
		n *= numWinningOptions(times[i], distances[i])
	}
	return fmt.Sprint(n)
}

func numWinningOptions(t, d int) int {
	//return numWinningOptionsNaive(t, d)
	return numWinningOptionsQuadratic(t, d)
}

func numWinningOptionsNaive(t, d int) int {
	n := 0
	for i := 1; i < t; i++ {
		di := i * (t - i)
		if di > d {
			n++
		}
	}
	//fmt.Println("t=", t, "d=", d, "n=", n)
	return n
}

// I think a more optimal solution would start in the middle
// and work outwards... or just solve i * (t-i) < d
// i*t - i*i < d
// -i^2 + t*i -d < 0
// solve for == 0 and we have the bounds!
// what was that quadratic formula again?
// (-b +/- sqrt( b^2 -4ac)) / 2a
//
// here a = -1, b = t, and c = -d
func numWinningOptionsQuadratic(t, d int) int {
	a := float64(-1)
	b := float64(t)
	c := -1 * float64(d)

	b2_4ac := (b * b) - (4 * a * c)
	if b2_4ac <= 0 {
		// no solutions
		return 0
	}

	sqrtb2_4ac := math.Sqrt(b2_4ac)

	// note that 2a = -2 and is constant

	s1 := (-1*b + sqrtb2_4ac) / -2
	s2 := (-1*b - sqrtb2_4ac) / -2

	// we want the _next_ integer or _prev_ integer
	if isInteger(s1) {
		s1 = s1 + 1
	} else {
		s1 = math.Ceil(s1)
	}
	if isInteger(s2) {
		s2 = s2 - 1
	} else {
		s2 = math.Floor(s2)
	}
	// add one to include the boundary numbers
	return int(1 + s2 - s1)
}

func isInteger(f float64) bool {
	return math.Floor(f) == f
}

// Implement Solution to Problem 2
func solve2(input string) string {
	var t, d int
	aoc.MapLines(input, func(line string) error {
		if strings.HasPrefix(line, "Time: ") {
			line := strings.ReplaceAll(line[6:], " ", "")
			t, _ = strconv.Atoi(line)
		}
		if strings.HasPrefix(line, "Distance: ") {
			line := strings.ReplaceAll(line[10:], " ", "")
			d, _ = strconv.Atoi(line)
		}
		return nil
	})
	n := numWinningOptions(t, d)
	return fmt.Sprint(n)
}
