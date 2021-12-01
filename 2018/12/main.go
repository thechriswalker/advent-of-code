package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2018, 12, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	gens := 20
	state, rules := parseInput(input, gens) // space for 20 generations
	next := make(State, len(state))
	fmt.Println()
	printHeader(state, gens)
	printState(state, gens, 0)
	for i := 0; i < gens; i++ {
		state.Next(rules, next)
		next, state = state, next
		printState(state, gens, i+1)
	}

	// now count the living (summing indexes)
	var sum int
	for i, s := range state {
		if s == 1 {
			sum += i - gens
		}
	}
	return fmt.Sprintf("%d", sum)
}

func printHeader(state State, padding int) {
	fmt.Print("      ")
	l := len(state)
	for i := range state {
		j := i - padding
		if j >= 0 && j < l-2*padding {
			fmt.Printf("%d", j%10)
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}

func printState(state State, padding, gen int) {
	// the line with the values on it.
	fmt.Printf("%04d: ", gen)
	l := len(state)
	var format string
	dull := "\x1b[1;30m%c"
	alive := "\x1b[1;32m%c"
	dead := "\x1b[1;31m%c"
	for i, s := range state {
		j := i - padding
		format = dull
		if j >= 0 && j < l-2*padding {
			if s == 1 {
				format = alive
			} else {
				format = dead
			}
		}
		if s == 1 {
			fmt.Printf(format, '#')
		} else {
			fmt.Printf(format, '.')
		}
	}
	fmt.Print("\x1b[0m\n")
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// this probably only works for my input
	// I ran 500 generations and the output looks stable much sooner than that
	// in fact it stabalises at exactly generation 100
	// but it shifts to the right 1 every generation
	// which means the result is the sum at 100 + (50billion-100) * liveCount

	// lets generalise though and count the living on each pass
	// hoping that it stablises after 1000 generations at least
	gens := 1000
	prev, rules := parseInput(input, gens) // space for 20 generations
	next := make(State, len(prev))
	var alive, prevAlive, stableCount, shift, finalGen int
	prevAlive = prev.Alive()
	// fmt.Println()
	// printHeader(state, gens)
	// printState(state, gens, 0)
	for i := 0; i < gens; i++ {
		alive = prev.Next(rules, next)
		if alive == prevAlive {
			if bytes.Compare(prev.Representation(), next.Representation()) == 0 {
				stableCount++
			} else {
				stableCount = 0
			}
		} else {
			stableCount = 0
		}
		if stableCount > 25 {
			// we have a stable representation.
			// what is the shift?
			shift = bytes.IndexByte(next, 1) - bytes.IndexByte(prev, 1)
			finalGen = i + 1
			break
		}
		next, prev = prev, next
		prevAlive, alive = alive, prevAlive

		// printState(state, gens, i+1)
	}

	// now count the living (summing indexes and the count)
	// make sure it doesn't alternate so I count both
	var sum, count int
	for i, s := range next {
		if s == 1 {
			sum += i - gens
			count++
		}
	}

	// fmt.Println()
	// fmt.Printf("Stablised: Gen %d, Alive: %d, Sum: %d, Shift: %d\n", finalGen, count, sum, shift)
	// printState(next.Representation(), 0, finalGen)

	// now our final value is sum + (50e9-finalGen) * count * shift
	v := uint64(sum) + (50000000000-uint64(finalGen))*uint64(count)*uint64(shift)

	return fmt.Sprintf("%d", v)
}

func parseInput(input string, generations int) (State, Rules) {
	// just split on lines...
	lines := strings.Split(input, "\n")
	var state State
	rules := map[uint8]uint8{}
	var l1, l2, c, r1, r2, s, p uint8
	mapToBinary := func(c uint8) uint8 {
		switch c {
		case '.':
			return 0
		case '#':
			return 1
		}
		panic(fmt.Sprintf("unexpected character: %c", c))
	}
	for _, line := range lines {
		if len(strings.TrimSpace(line)) == 0 {
			continue
		}
		switch line[0] {
		case 'i':
			//initial state: #....
			state = make(State, 2*generations+len(line)-15)
			for i, c := range line[15:] {
				state[generations+i] = mapToBinary(uint8(c))
			}
		case '.', '#':
			// rule
			fmt.Sscanf(line, "%c%c%c%c%c => %c", &l1, &l2, &c, &r1, &r2, &s)
			p = mapToBinary(l1) + mapToBinary(l2)<<1 + mapToBinary(c)<<2 + mapToBinary(r1)<<3 + mapToBinary(r2)<<4
			rules[p] = mapToBinary(s)
		}
	}
	return state, rules
}

// The state could grow 20 elements bigger in each direction in 20 generations.
// so I will prepad the element array so we don't have to worry about it growing
type State []uint8

// these are the rules.
// they are the 5 bits for LLCRR
// shifted into a single integer, with the bool representing
// life or death for the next state
type Rules map[uint8]uint8

// write the next state into the given buffer
func (s State) Next(rules Rules, next State) int {
	var pattern uint8
	var alive int
	for i := 0; i < len(s); i++ {
		// assume nothing on either side of the buffer.
		pattern = 0
		switch {
		case i == 0:
			//left and left are dead
			pattern += s[i] << 2
			pattern += s[i+1] << 3
			pattern += s[i+2] << 4
		case i == 1:
			pattern += s[i-1] << 1
			pattern += s[i] << 2
			pattern += s[i+1] << 3
			pattern += s[i+2] << 4
		case i == len(s)-2:
			pattern += s[i-2]
			pattern += s[i-1] << 1
			pattern += s[i] << 2
			pattern += s[i+1] << 3
		case i == len(s)-1:
			pattern += s[i-2]
			pattern += s[i-1] << 1
			pattern += s[i] << 2
		default:
			// we have all the bits.
			pattern += s[i-2]
			pattern += s[i-1] << 1
			pattern += s[i] << 2
			pattern += s[i+1] << 3
			pattern += s[i+2] << 4
		}

		next[i] = rules[pattern]
		alive += int(rules[pattern])
	}
	return alive
}

func (s State) Alive() int {
	var alive int
	for i := 0; i < len(s); i++ {
		alive += int(s[i])
	}
	return alive
}

// index agnostic representation of the state.
// start from the first Live to the last live
func (s State) Representation() []byte {
	first := bytes.IndexByte([]byte(s), 1)
	last := bytes.LastIndexByte([]byte(s), 1)
	return s[first : last+1]
}
