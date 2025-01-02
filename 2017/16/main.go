package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

var LineSize = 16

func main() {
	aoc.Run(2017, 16, solve1, solve2)
}

// Better if we just tracked the "head", and made this a ring.
// then Spin is just a head move.
// makes Exchange require an "offset" to the head.
type Line struct {
	data []byte
	head int
}

func (l *Line) Spin(n int) {
	n = len(l.data) - n // reverse the spin
	l.head = (l.head + n) % len(l.data)
}

func (l *Line) Exchange(a, b int) {
	a = (a + l.head) % len(l.data)
	b = (b + l.head) % len(l.data)
	l.data[a], l.data[b] = l.data[b], l.data[a]
}

func (l *Line) Partner(a, b byte) {
	// find this indices of a and b and swap them (directly!)
	var ai, bi int
	var afound, bfound bool
	for i, c := range l.data {
		if c == a {
			ai = i
			afound = true
		}
		if c == b {
			bi = i
			bfound = true
		}
		if afound && bfound {
			break
		}
	}
	// these are absolute indexes, so just swap them directly
	l.data[ai], l.data[bi] = l.data[bi], l.data[ai]
}

func (l Line) String() string {
	return string(l.data[l.head:]) + string(l.data[:l.head])
}

// Implement Solution to Problem 1
func solve1(input string) string {
	line := Line{data: make([]byte, LineSize)}
	for i := range line.data {
		line.data[i] = byte('a' + i)
	}
	// parse the input into a list of moves
	moves := strings.Split(input, ",")

	var a, b int
	for _, move := range moves {
		switch move[0] {
		case 's':
			fmt.Sscanf(move, "s%d", &a)
			line.Spin(a)
		case 'x':
			fmt.Sscanf(move, "x%d/%d", &a, &b)
			line.Exchange(a, b)
		case 'p':
			fmt.Sscanf(move, "p%c/%c", &a, &b)
			line.Partner(byte(a), byte(b))
		}
	}

	return line.String()
}

// Implement Solution to Problem 2
func solve2(input string) string {

	// 1_000_000_000 iterations is a lot, so we need to find a pattern
	// we can speed up the whole process by finding the cycle length.
	// the problem is the Partner swap, which is not based on the indexes only
	// let's "look" at the first 1000 iterations and see if we can find a pattern
	// I might need to make the moves more efficient, but let's see if we can find a pattern first

	line := Line{data: make([]byte, LineSize)}
	for i := range line.data {
		line.data[i] = byte('a' + i)
	}
	// parse the input into a list of moves
	moves := strings.Split(input, ",")

	dance := func() {
		var a, b int
		for _, move := range moves {
			switch move[0] {
			case 's':
				fmt.Sscanf(move, "s%d", &a)
				line.Spin(a)
			case 'x':
				fmt.Sscanf(move, "x%d/%d", &a, &b)
				line.Exchange(a, b)
			case 'p':
				fmt.Sscanf(move, "p%c/%c", &a, &b)
				line.Partner(byte(a), byte(b))
			}
		}
	}
	i := 0
	cache := map[string]int{
		line.String(): i,
	}
	results := []string{line.String()}
	for {
		i++
		dance()
		s := line.String()
		results = append(results, s)
		if _, ok := cache[s]; ok {
			//fmt.Println("\ncycle found from", ii, "to", i)
			break
		}
		cache[s] = i
	}
	// now we know the cycle, we can just do the dance 1_000_000_000 % cycle times
	idx := 1_000_000_000 % i
	return results[idx]
}
