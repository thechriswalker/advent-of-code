package main

import (
	"fmt"
	"math"
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 17, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var a, b, c int
	var p string
	fmt.Sscanf(input, `Register A: %d
Register B: %d
Register C: %d

Program: %s
`, &a, &b, &c, &p)

	program := aoc.ToIntSlice(p, ',')

	out := runProgram(program, a, b, c)

	// crunch into a comma-separated string
	var sb strings.Builder
	comma := ""
	for _, v := range out {
		fmt.Fprintf(&sb, "%s%d", comma, v)
		comma = ","
	}

	return sb.String()
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// var tmp int
	// var p string
	// fmt.Sscanf(input, `Register A: %d
	// Register B: %d
	// Register C: %d

	// Program: %s
	// `, &tmp, &tmp, &tmp, &p)

	// program := aoc.ToIntSlice(p, ',')

	// 	a = 0
	// 	for {
	// 		out := runProgram(program, a, b, c)
	// 		if equal(out, program) {
	// 			// check it matches
	// 			break
	// 		}
	// 		a++
	// 		if a%10_000_000 == 0 {
	// 			fmt.Println(a)
	// 		}
	// 	}
	// 	return fmt.Sprint(a)
	// unless my fast version is incorrect, it is still too slow.
	// I should check the fast version against part 1.
	// 48 bits.
	// so we have 16 octal digits.
	// but they interact with those close by so I will group in 3s.
	// basically, we check all possible "3-digit" groups for
	// the first number they output, and keep the ones that have the correct output.
	// the move to the next digit and repeat, keeping only the ones that keep the correct output.

	// 3 digits at a time
	bases := map[int]struct{}{
		0: {},
	}
	for n := 0; n < 16; n++ {
		// fmt.Println("digit", n, "bases", len(bases))
		d := int(math.Pow(8, float64(n)))
		options := map[int]struct{}{}
		for base := range bases {
			for a := range 8 * 8 * 8 {
				x := base + (a * d)
				// fast version is ~220ms but only works for my input
				if m, _ := fastProgram(x, nil); m >= n {
					// // slow version takes a full 2.8 seconds
					// out := runProgram(program, x, 0, 0)
					//if matches(out, program) >= n {
					options[x] = struct{}{}
				}
			}
		}
		if len(options) == 0 {
			panic("no options")
		}
		bases = options
	}

	b := make([]int, 0, len(bases))
	for base := range bases {
		b = append(b, base)
	}

	sort.Ints(b)

	// fmt.Println("program = ", []int{2, 4, 1, 1, 7, 5, 1, 5, 4, 3, 0, 3, 5, 5, 3, 0})
	// fmt.Println(" output = ", runProgram([]int{2, 4, 1, 1, 7, 5, 1, 5, 4, 3, 0, 3, 5, 5, 3, 0}, b[0], 0, 0))

	return fmt.Sprint(b[0])
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// how many entries in a match b
func matches(a, b []int) int {
	i := 0

	for i < len(a) && i < len(b) {
		if a[i] != b[i] {
			break
		}
		i++
	}
	return i
}

func runProgram(p []int, a, b, c int) (out []int) {
	pc := 0

	combo := func() int {
		v := p[pc]
		pc++
		switch v {
		case 0, 1, 2, 3:
			return v
		case 4:
			return a
		case 5:
			return b
		case 6:
			return c
		default:
			panic("invalid program")
		}
	}
	literal := func() int {
		v := p[pc]
		pc++
		return v
	}

	for pc < len(p) {
		// perform the operation
		op := p[pc]
		pc++
		switch op {
		case 0:
			// adv reg A / 2^combo
			a = a / (1 << combo())
		case 1:
			// bxl XOR B and literal
			b = b ^ literal()
		case 2:
			// bst combo() % 8  => b
			b = combo() % 8
		case 3:
			// jnz if A = 0 nothing
			if a == 0 {
				// ignore operand
				pc++
			} else {
				pc = literal()
			}
		case 4:
			// bxc XOR B and C
			b = b ^ c
			pc++ // ignore operand
		case 5:
			// out - combo %8 output.
			out = append(out, combo()%8)
		case 6:
			// bdv like adv but store in b
			b = a / (1 << combo())
		case 7:
			// cdv like adv but store in c
			c = a / (1 << combo())
		default:
			panic("invalid program")
		}
	}
	return
}

// Program: 2,4, 1,1, 7,5, 1,5, 4,3, 0,3, 5,5, 3,0
func fastProgram(a int, target []int) (int, bool) {
	var b, c int
	if target == nil {
		target = []int{2, 4, 1, 1, 7, 5, 1, 5, 4, 3, 0, 3, 5, 5, 3, 0}
	}
	outIdx := 0

	// dbg := func(m string) {
	// 	fmt.Printf("%s  \ta=%o, b=%o, c=%o\n", m, a, b, c)
	// }
	//dbg("start       ")

literal0:
	// first we 2,4
	// which is a mod 8 => b
	//b = a % 8

	// then we 1,1
	//b = b ^ 1 // i.e. b %2

	b = (a % 8) ^ 1 // a = 7, b = 1
	//dbg("b = (a % 8) ^ 1")
	// then we 7,5
	c = a / (1 << b) // c = 7 / 2^1 = 3
	//dbg("c = a / (1 << b)")

	// then we 1,5
	b = b ^ 5 // b = 1 ^ 5 = 4
	//dbg("b = b ^ 5")
	// then we 4,3
	b = b ^ c // b = 4 ^ 3 = 7
	//dbg("b = b ^ c")
	// then we 0,3
	a = a / 8 // a = 7 / 8 = 0
	//dbg("a = a / 8")
	// then 5,5
	if outIdx >= len(target) || target[outIdx] != b%8 { // b % 8 = 7
		//	dbg("returning false")
		return outIdx, false
	}
	//dbg(fmt.Sprintf("outputting: %d", b%8))
	outIdx++

	// then 3,0
	if a != 0 {
		//	dbg("a != 0 => jump to start")
		goto literal0
	}
	//dbg("halt")
	// halt.
	// if we matched all the outputs, we're done.
	return outIdx, outIdx == len(target)
}

// Can we work backwards? start with the registers at a certain value
// a is divided by 8 until it rounds down to 0, so the minimum possible value will be 8^8
// and the solution will have 16 * 3 bits (48 bits).
