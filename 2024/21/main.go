package main

import (
	"bytes"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 21, solve1, solve2)
}

var keypad = map[byte]aoc.V2{
	'7': {X: 0, Y: 0},
	'8': {X: 1, Y: 0},
	'9': {X: 2, Y: 0},
	'4': {X: 0, Y: 1},
	'5': {X: 1, Y: 1},
	'6': {X: 2, Y: 1},
	'1': {X: 0, Y: 2},
	'2': {X: 1, Y: 2},
	'3': {X: 2, Y: 2},
	'X': {X: 0, Y: 3},
	'0': {X: 1, Y: 3},
	'A': {X: 2, Y: 3},
}

var dpad = map[byte]aoc.V2{
	'X': {X: 0, Y: 0},
	'^': {X: 1, Y: 0},
	'A': {X: 2, Y: 0},
	'<': {X: 0, Y: 1},
	'v': {X: 1, Y: 1},
	'>': {X: 2, Y: 1},
}

func sequenceToDpadOptions(target map[byte]aoc.V2, sequence []byte, start byte) [][]byte {
	options := [][]byte{nil}
	blackhole := target['X']
	currentPos := target[start]

	for _, next := range sequence {
		moveHorizontalFirst := true
		moveVerticalFirst := true
		//l := len(moves)
		// move to the target
		nextPos := target[next]
		dx, dy := nextPos.X-currentPos.X, nextPos.Y-currentPos.Y

		if dx == 0 && dy == 0 {
			// we are already there.
			// do neither
			moveHorizontalFirst = false
			moveVerticalFirst = false
		} else

		// we will cross a black hole IIF
		// our curr.X == blackhole.X AND our next.Y = blackhole.Y
		// because that means if we move vertically first we will hit it.
		if currentPos.X == blackhole.X && nextPos.Y == blackhole.Y {
			// we will hit the backhole if we move vertically first.
			moveVerticalFirst = false
		} else if currentPos.Y == blackhole.Y && nextPos.X == blackhole.X {
			// we will hit the blackhole if we move horizontally first.
			moveHorizontalFirst = false
		}

		if moveHorizontalFirst && dx == 0 {
			moveHorizontalFirst = false
		}
		if moveVerticalFirst && dy == 0 {
			moveVerticalFirst = false
		}

		nextOptions := [][]byte{}

		if moveHorizontalFirst {
			// we move right first, then we will never hit the empty space.
			for _, moves := range options {
				m := make([]byte, 0, len(moves)*2)
				m = append(m, moves...)
				m = append(m, moveHorizontal(dx)...)
				m = append(m, moveVertical(dy)...)
				m = append(m, 'A')
				nextOptions = append(nextOptions, m)
			}
		}
		if moveVerticalFirst {
			// move up/down first, then left
			for _, moves := range options {
				m := make([]byte, 0, len(moves)*2)
				m = append(m, moves...)
				m = append(m, moveVertical(dy)...)
				m = append(m, moveHorizontal(dx)...)
				m = append(m, 'A')
				nextOptions = append(nextOptions, m)
			}
		}
		if !moveHorizontalFirst && !moveVerticalFirst {
			for _, moves := range options {
				moves = append(moves, 'A')
				nextOptions = append(nextOptions, moves)
			}
		}
		//	fmt.Println("to go from", currentPos, "to", nextPos, "we need to do", string(moves[l:]))
		//fmt.Println("from ", currentPos, "to", nextPos, "next", string(next), "dx", dx, "dy", dy, "moveHorizontalFirst", moveHorizontalFirst, "moveVerticalFirst", moveVerticalFirst, "nextOptions", len(nextOptions))
		currentPos = nextPos
		options = nextOptions
	}
	return options
}

func moveHorizontal(dx int) []byte {
	if dx == 0 {
		return nil
	}
	if dx < 0 {
		return bytes.Repeat([]byte{'<'}, -dx)
	}
	return bytes.Repeat([]byte{'>'}, dx)
}

func moveVertical(dy int) []byte {
	if dy == 0 {
		return nil
	}
	if dy < 0 {
		return bytes.Repeat([]byte{'^'}, -dy)
	}
	return bytes.Repeat([]byte{'v'}, dy)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return solveN(input, 2)
}

func solveN(input string, depth int) string {
	sum := 0
	cache := map[int]map[string]int{}
	aoc.MapLines(input, func(line string) error {

		code := []byte(strings.TrimSpace(line))
		n, _ := strconv.Atoi(line[:3]) //  the numeric bit
		l := minSequence(cache, string(code), 0, depth)
		complexity := l * n
		sum += complexity

		return nil
	})

	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {

	return solveN(input, 25)
}

// do we need to make a cache for the one step pairwise moves
// at each depth? rather than one cache with each?
// like a dynamic programming approach?
// so we have a [3]byte cache for "min length of sequence to get from a to b at depth c"
func minSequence(cache map[int]map[string]int, code string, depth, max int) (sum int) {
	if _, ok := cache[depth]; !ok {
		cache[depth] = map[string]int{}
	}
	if l, ok := cache[depth][code]; ok {
		return l
	}
	defer func() {
		cache[depth][code] = sum
	}()
	pad := dpad
	if depth == 0 {
		pad = keypad
	}
	// for each "pair" of characters in the code, we need to find the shortest sequence
	for i := 0; i < len(code); i++ {
		a, b := byte('A'), code[i]
		if i > 0 {
			a = code[i-1]
		}
		options := sequenceToDpadOptions(pad, []byte{b}, a)
		if depth == max {
			// sort and return the shortest
			slices.SortFunc(options, func(a, b []byte) int {
				return len(a) - len(b)
			})
			sum += len(options[0])
			continue
		}
		// otherwise recurse
		min := math.MaxInt
		for _, p := range options {
			m := minSequence(cache, string(p), depth+1, max)
			if m < min {
				min = m
			}
		}
		sum += min
	}
	return sum
}
