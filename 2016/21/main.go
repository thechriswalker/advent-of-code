package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 21, solve1, solve2)
}

var Password1 = "abcdefgh"
var Password2 = "fbgdceah"

// Implement Solution to Problem 1
func solve1(input string) string {
	// we could implement an interpreter, but this is easy enough
	// first we make two buffers to swap between.
	prev := []byte(Password1)
	next := make([]byte, len(prev))
	copy(next, prev)
	var n1, n2 int
	var c1, c2 byte
	var err error
	p := func(s string) { panic(s) }
	for _, line := range strings.Split(input, "\n") {

		// work out what to do.
		if len(line) < 8 {
			// not a real line
			continue
		}
		//fmt.Printf("# %s\n", line)
		switch line[0] {
		case 'm': //move position X to position Y
			if _, err = fmt.Sscanf(line, "move position %d to position %d", &n1, &n2); err != nil {
				p(line)
			}
			movePosition(prev, next, n1, n2)
		case 's':
			switch line[5] {
			case 'l': // swap letter
				if _, err = fmt.Sscanf(line, "swap letter %c with letter %c", &c1, &c2); err != nil {
					p(line)
				}
				swapLetter(prev, next, c1, c2)
			case 'p': // swap position
				if _, err = fmt.Sscanf(line, "swap position %d with position %d", &n1, &n2); err != nil {
					p(line)
				}
				swapPosition(prev, next, n1, n2)
			}
		case 'r':
			switch line[7] {
			case ' ': // reverse positions
				if _, err = fmt.Sscanf(line, "reverse positions %d through %d", &n1, &n2); err != nil {
					p(line)
				}
				reversePositions(prev, next, n1, n2)
			case 'b': // rotate based on position
				if _, err = fmt.Sscanf(line, "rotate based on position of letter %c", &c1); err != nil {
					p(line)
				}
				rotateFromLetter(prev, next, c1)
			case 'l': //rotate left
				if _, err = fmt.Sscanf(line, "rotate left %d step", &n1); err != nil {
					p(line)
				}
				rotate(prev, next, n1*-1) //left is negative shift
			case 'r': //rotate right
				if _, err = fmt.Sscanf(line, "rotate right %d step", &n1); err != nil {
					break
				}
				rotate(prev, next, n1) //right is positive shift
			default:
				p(line)
			}
		default:
			p(line)
		}
		// make them the same with next as the template
		//	fmt.Printf("%s => %s\n", prev, next)
		// quick sanity check.
		//sanityCheck(next)
		copy(prev, next)
	}
	return string(prev)
}

func sanityCheck(next []byte) {
	// no duplicate letters.
	m := map[byte]struct{}{}
	for _, c := range next {
		if _, ok := m[c]; ok {
			panic("Duplicate letter found")
		}
		m[c] = struct{}{}
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	prev := []byte(Password2)
	next := make([]byte, len(prev))
	copy(next, prev)
	var n1, n2 int
	var c1, c2 byte
	var err error
	p := func(s string) { panic(s) }
	lines := strings.Split(input, "\n")
	// run backwards
	for i := len(lines) - 1; i > -1; i-- {
		line := lines[i]
		// work out what to do.
		if len(line) < 8 {
			// not a real line
			continue
		}
		//	fmt.Printf("# %s\n", line)
		switch line[0] {
		case 'm': //move position X to position Y
			if _, err = fmt.Sscanf(line, "move position %d to position %d", &n1, &n2); err != nil {
				p(line)
			}
			movePosition(prev, next, n2, n1)
		case 's':
			switch line[5] {
			case 'l': // swap letter
				if _, err = fmt.Sscanf(line, "swap letter %c with letter %c", &c1, &c2); err != nil {
					p(line)
				}
				swapLetter(prev, next, c2, c1)
			case 'p': // swap position
				if _, err = fmt.Sscanf(line, "swap position %d with position %d", &n1, &n2); err != nil {
					p(line)
				}
				swapPosition(prev, next, n2, n1)
			}
		case 'r':
			switch line[7] {
			case ' ': // reverse positions
				if _, err = fmt.Sscanf(line, "reverse positions %d through %d", &n1, &n2); err != nil {
					p(line)
				}
				reversePositions(prev, next, n1, n2) // nb this is it's own reverse function
			case 'b': // rotate based on position
				if _, err = fmt.Sscanf(line, "rotate based on position of letter %c", &c1); err != nil {
					p(line)
				}
				// how do we do this...
				reverseRotateFromLetter(prev, next, c1)
			case 'l': //rotate left
				if _, err = fmt.Sscanf(line, "rotate left %d step", &n1); err != nil {
					p(line)
				}
				rotate(prev, next, n1) //reversed to positive
			case 'r': //rotate right
				if _, err = fmt.Sscanf(line, "rotate right %d step", &n1); err != nil {
					break
				}
				rotate(prev, next, n1*-1) //reverse to negative
			default:
				p(line)
			}
		default:
			p(line)
		}
		// make them the same with next as the template
		//fmt.Printf("%s => %s\n", prev, next)
		// quick sanity check.
		//	sanityCheck(next)
		copy(prev, next)
	}
	return string(prev)
}

func swapPosition(prev, next []byte, x, y int) {
	next[x], next[y] = prev[y], prev[x]
}

func swapLetter(prev, next []byte, x, y byte) {
	swapPosition(prev, next, bytes.IndexByte(prev, x), bytes.IndexByte(prev, y))
}

func movePosition(prev, next []byte, x, y int) {
	shift := 0
	before := y < x
	for i := range prev {
		switch i {
		case x:
			if before {
				next[i] = prev[i+shift]
			}
			shift++
			if !before {
				next[i] = prev[i+shift]
			}
		case y:
			shift--
			next[i] = prev[x]
		default:
			next[i] = prev[i+shift]
		}
	}
}

func reversePositions(prev, next []byte, x, y int) {
	for i := 0; i <= y-x; i++ {
		next[x+i] = prev[y-i]
	}
}

func rotate(prev, next []byte, n int) {
	l := len(prev)
	for i := 0; i < l; i++ {
		y := (i + l + n) % l
		next[y] = prev[i]
	}
}
func rotateFromLetter(prev, next []byte, x byte) {
	// this is a right rotation based on position of x +1 or +2
	n := bytes.IndexByte(prev, x) + 1
	if n > 4 {
		n++
	}
	rotate(prev, next, n)
}

func reverseRotateFromLetter(prev, next []byte, x byte) {
	i := bytes.IndexByte(prev, x)
	n := i
	// now what brute force, we know the lengths.
	switch len(prev) {
	case 5:
		if n == 2 {
			panic("5 letter word with invalid shift")
		}
		n = map5[n]
	case 8:

		n = map8[n]
	}

	// now shift left! that much
	//fmt.Printf("reverseRotateFromLetter: %c current index %d shift %d\n", x, i, n)
	rotate(prev, next, n)
}

var map5 = map[int]int{
	0: -1, 1: -1, 2: 1, 3: -2,
}

var map8 = map[int]int{
	0: -1, 1: -1, 2: -6, 3: -2, 4: -7, 5: -3, 6: 0, 7: -4,
}
