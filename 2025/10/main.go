package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 10, solve1, solve2)
}

type Machine struct {
	// indicator lights are never more than 16 bits, but bigger than 8,
	// but I'll just use int for ease
	InitTarget int

	// list of buttons and the lights they toggle
	ButtonEffects []int
	Buttons       [][]int

	// for part 2?
	Joltages [JOLT_SIZE]int
}

func parseMachines(input string) []Machine {
	machines := []Machine{}
	aoc.MapLines(input, func(line string) error {
		m := Machine{}
		fields := strings.Fields(line)
		for _, f := range fields {
			inner := f[1 : len(f)-1]
			switch f[0] {
			case '[':
				// indicator light target
				for i, b := range inner {
					if b == '#' {
						m.InitTarget ^= 1 << i
					}
				}
			case '(':
				// button
				s := aoc.ToIntSlice(inner, ',')
				effect := 0
				for _, b := range s {
					effect ^= 1 << b
				}
				m.ButtonEffects = append(m.ButtonEffects, effect)
				m.Buttons = append(m.Buttons, s)
			case '{':
				//joltage
				jolts := aoc.ToIntSlice(inner, ',')
				for i, j := range jolts {
					m.Joltages[i] = j
				}
			default:
				// ignore
			}
		}
		machines = append(machines, m)
		return nil
	})
	return machines
}

// Implement Solution to Problem 1
func solve1(input string) string {

	machines := parseMachines(input)
	sum := 0

	for _, m := range machines {
		sum += m.Initialise()
	}

	return fmt.Sprint(sum)
}

func (m Machine) Initialise() int {
	// minimum steps to initialisation.
	// definitely a breadth first search!
	// NB I totally missed that buttons can only be pressed 0 or 1 times as they xor.
	// would have cut down significantly on the iterations, but this still runs in ~2ms so I don't care.
	current := []int{0}
	next := []int{}
	steps := 0
	// a cache of "seen" states
	cache := map[int]struct{}{
		0: {},
	}
	for {
		steps++ // next step
		for _, x := range current {
			for _, b := range m.ButtonEffects {
				// press every thing!
				n := x ^ b
				if n == m.InitTarget {
					return steps
				}
				if _, ok := cache[n]; !ok {
					cache[n] = struct{}{}
					next = append(next, n)
				}
			}
		}
		if len(next) == 0 {
			panic("fail")
		}
		current = next
		next = []int{}
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// boo I can't bit twiddle for this one... or at least not easily with my skills!
	machines := parseMachines(input)

	sum := 0

	for _, m := range machines {
		// this was too slow.
		//sum += m.AdjustJoltage_BreadthFirstLikePart1()
		// what do we do instead?
		// we need a way of finding the shortest path to the target
		// ???
		// Apparently we need to use a SMT Solver like Z3
		// We model the system with each button press count as x0,x1,x2
		// and a series of constant values for the target amounts
		// i.e. j0 = x0 * (1 or 0 based on if button 0 affect j0) + ...
		// then solve for the x's and sum them.
		sum += m.AdjustJoltage_Solver()
	}

	return fmt.Sprint(sum)
}

// we use a fixed size array to allow using it as a key
const JOLT_SIZE = 10

// OK the exhaustive search didn't work on the actual input
// I guess too large values.
func (m Machine) AdjustJoltage_BreadthFirstLikePart1() int {
	// definitely a breadth first search!
	initial := [JOLT_SIZE]int{}
	current := [][JOLT_SIZE]int{initial}
	next := [][JOLT_SIZE]int{}
	steps := 0
	// cache := map[[JOLT_SIZE]int]struct{}{
	// 	initial: {},
	// }
	for {
		steps++ // next step
		for _, x := range current {
			for _, b := range m.Buttons {
				n := x
				// press every thing!
				bust := false
				for _, i := range b {
					n[i]++
					if n[i] > m.Joltages[i] {
						//nope.
						bust = true
						break
					}
				}
				if bust {
					break
				}
				if n == m.Joltages {
					return steps
				}
				// if _, ok := cache[n]; !ok {
				// 	cache[n] = struct{}{}
				next = append(next, n)
				// }
			}
		}
		if len(next) == 0 {
			panic("fail")
		}
		current = next
		next = [][JOLT_SIZE]int{}
	}
}

func (m Machine) AdjustJoltage_Solver() int {
	// need a Z3 solver here... which is a bummer.
	// I don't like having to use dependencies in this code...
	// I might just call it a fail for today an move on.
	return 0

}
