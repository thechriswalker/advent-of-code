package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 13, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	machines := []Machine{}
	for _, s := range strings.Split(input, "\n\n") {
		machines = append(machines, parseMachine(s, aoc.Vec2(0, 0)))
	}

	total := 0
	for _, m := range machines {
		cost, won := m.PlayWithMath()
		if won {
			total += cost
		}
	}

	return fmt.Sprint(total)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	machines := []Machine{}
	for _, s := range strings.Split(input, "\n\n") {
		machines = append(machines, parseMachine(s, aoc.Vec2(10_000_000_000_000, 10_000_000_000_000)))
	}

	total := 0
	for _, m := range machines {
		cost, won := m.PlayWithMath()
		if won {
			total += cost
		}
	}

	// 150481754826599 is too high

	return fmt.Sprint(total)
}

type Machine struct {
	Prize aoc.V2
	A, B  aoc.V2
}

const (
	costA = 3
	costB = 1
)

// either X or Y is greater than the other
func gt(a, b aoc.V2) bool {
	return a.X > b.X || a.Y > b.Y
}

func parseMachine(s string, prizeOffset aoc.V2) Machine {
	var a, b, prize aoc.V2
	fmt.Sscanf(s, `Button A: X+%d, Y+%d
Button B: X+%d, Y+%d
Prize: X=%d, Y=%d`, &a.X, &a.Y, &b.X, &b.Y, &prize.X, &prize.Y)
	return Machine{
		Prize: prize.Add(prizeOffset),
		A:     a,
		B:     b,
	}
}

// this is not going to cut it for part 2, it sounded like this would be a math problem...
func (m *Machine) Play() (cost int, won bool) {
	// max moves to get prize is 100.
	// seen holds the number of tokens needed to get to a given position
	// if we get there cheaper, we update the value, but we don't repeat the move.
	origin := aoc.Vec2(0, 0)
	seen := map[aoc.V2]int{
		origin: 0,
	}
	// start at the origin
	curr := []aoc.V2{origin}
	next := []aoc.V2{}
	//fmt.Println("playing machine", m)
	for len(curr) > 0 {
		for _, c := range curr {
			cost := seen[c]
			a := c.Add(m.A)
			// if we haven't excceded the prize check this
			if !gt(a, m.Prize) {
				ca, ok := seen[a]
				if !ok || ca > cost+costA {
					seen[a] = cost + costA
					next = append(next, a)
				}
			}
			b := c.Add(m.B)
			if !gt(b, m.Prize) {
				cb, ok := seen[b]
				if !ok || cb > cost+costB {
					seen[b] = cost + costB
					next = append(next, b)
				}
			}
		}
		curr = next
		next = []aoc.V2{}
	}

	cost, won = seen[m.Prize]
	return
}

/*
We have some simultaneous equations to solve here.

Pa and Pb are the button presses for A and B respectively.

Prize.X = Pa * a.X + Pb * b.X
Prize.Y = Pa * a.Y + Pb * b.Y

solving to Pb we get:

Pb = (a.X*Prize.Y - a.Y*Prize.X) / (a.X*b.Y - a.Y*b.X)

so a solution exists if a.X*b.Y - a.Y*b.X != 0

then we can work out the Pa and the cost

but this doesn't get us the right answer, so I guess we have to solve the other way as well?

must be quadratic or something.
*/
func (m *Machine) SolveWithB() (pa, pb int, won bool) {
	det := m.A.X*m.B.Y - m.A.Y*m.B.X
	if det == 0 {
		return 0, 0, false
	}

	Pb := (m.A.X*m.Prize.Y - m.A.Y*m.Prize.X) / det
	Pa := (m.Prize.X - Pb*m.B.X) / m.A.X
	if Pa < 0 || Pb < 0 {
		return 0, 0, false
	}
	// not only do we need to check that Pa and Pb are positive, but we also need to check that they are integers
	// they _are_ integers becuase of the type system, so we need to double check our results.
	if Pa*m.A.X+Pb*m.B.X != m.Prize.X || Pa*m.A.Y+Pb*m.B.Y != m.Prize.Y {
		return 0, 0, false
	}
	//fmt.Println("solved with B:", Pa, Pb, "cost:", Pa*costA+Pb*costB)
	return Pa, Pb, true
}

func (m *Machine) SolveWithA() (pa, pb int, won bool) {
	det := m.A.X*m.B.Y - m.A.Y*m.B.X
	if det == 0 {
		return 0, 0, false
	}

	Pa := (m.B.X*m.Prize.Y - m.B.Y*m.Prize.X) / det
	Pb := (m.Prize.X - Pa*m.A.X) / m.B.X

	if Pa < 0 || Pb < 0 {
		return 0, 0, false
	}
	// not only do we need to check that Pa and Pb are positive, but we also need to check that they are integers
	// they _are_ integers becuase of the type system, so we need to double check our results.
	if Pa*m.A.X+Pb*m.B.X != m.Prize.X || Pa*m.A.Y+Pb*m.B.Y != m.Prize.Y {
		return 0, 0, false
	}

	//fmt.Println("solved with A:", Pa, Pb, "cost:", Pa*costA+Pb*costB)

	return Pa, Pb, true
}

func (m *Machine) PlayWithMath() (cost int, won bool) {
	//	fmt.Println("solving machine:", m)
	aa, ab, aw := m.SolveWithA()
	ba, bb, bw := m.SolveWithB()
	if !aw && !bw {
		return 0, false
	}
	if aw && bw {
		ac, bc := aa*costA+ab*costB, ba*costA+bb*costB
		if ac < bc {
			return ac, true
		}
		return bc, true
	}
	if aw {
		return aa*costA + ab*costB, true
	}
	if bw {
		return ba*costA + bb*costB, true
	}
	return 0, false
}
