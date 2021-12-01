package main

import (
	"fmt"
	"math/bits"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 13, solve1, solve2)
}

var TargetX, TargetY uint = 31, 39

// Implement Solution to Problem 1
func solve1(input string) string {
	var plane Plane
	fmt.Sscanln(input, &plane)
	initialState := &State{X: 1, Y: 1}
	depth := 0
	var nextMoves []*State
	pendingMoves := []*State{initialState}
	for {
		nextMoves = []*State{}
		for _, move := range pendingMoves {
			if move.X == TargetX && move.Y == TargetY {
				return fmt.Sprintf("%d", depth)
			}
			// add more moves.
			nextMoves = append(nextMoves, move.NextMoves(plane)...)
		}
		if len(nextMoves) == 0 {
			return fmt.Sprintf("Halted after %d moves", depth)
		}
		depth++
		pendingMoves = nextMoves
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	var plane Plane
	fmt.Sscanln(input, &plane)
	initialState := &State{X: 1, Y: 1}
	empty := struct{}{}
	moveCache := map[[2]uint]struct{}{
		[2]uint{1, 1}: empty,
	}

	var nextMoves []*State
	pendingMoves := []*State{initialState}
	key := [2]uint{0, 0}
	for depth := 0; depth <= 50; depth++ {
		nextMoves = []*State{}
		for _, move := range pendingMoves {
			key[0], key[1] = move.X, move.Y
			moveCache[key] = empty
			// add more moves.
			nextMoves = append(nextMoves, move.NextMoves(plane)...)
		}
		pendingMoves = nextMoves
	}

	return fmt.Sprintf("%d", len(moveCache))
}

type Plane uint

func (p Plane) IsWall(x, y uint) bool {
	return !p.IsOpen(x, y)
}
func (p Plane) IsOpen(x, y uint) bool {
	n := uint(p) + x*x + 3*x + 2*x*y + y + y*y
	b := bits.OnesCount(n)
	return b%2 == 0
}
func (p Plane) String() string {
	return p.PrintSize(10)
}
func (p Plane) PrintSize(n uint) string {
	s := &strings.Builder{}
	s.WriteString("  0123456789")
	for y := uint(0); y < n; y++ {
		fmt.Fprintf(s, "\n%d ", y)
		for x := uint(0); x < n; x++ {
			if p.IsWall(x, y) {
				s.WriteRune('#')
			} else {
				s.WriteRune('.')
			}
		}
	}
	s.WriteRune('\n')
	return s.String()
}

type State struct {
	Prev *State
	X    uint
	Y    uint
}

func (s *State) IsBacktrack(x, y uint) bool {
	prev := s.Prev
	for prev != nil {
		if prev.X == x && prev.Y == y {
			return true
		}
		prev = prev.Prev
	}
	return false
}

func (s *State) NextMoves(p Plane) []*State {
	var x, y uint
	// at most 4 new states.
	moves := make([]*State, 0, 4)
	x, y = s.X, s.Y-1
	if s.Y != 0 && p.IsOpen(x, y) && !s.IsBacktrack(x, y) {
		// up is possible
		moves = append(moves, &State{Prev: s, X: x, Y: y})
	}
	x, y = s.X-1, s.Y
	if s.X != 0 && p.IsOpen(x, y) && !s.IsBacktrack(x, y) {
		// left is possible
		moves = append(moves, &State{Prev: s, X: x, Y: y})
	}
	x, y = s.X, s.Y+1
	if p.IsOpen(x, y) && !s.IsBacktrack(x, y) {
		// down is possible
		moves = append(moves, &State{Prev: s, X: x, Y: y})
	}
	x, y = s.X+1, s.Y
	if p.IsOpen(x, y) && !s.IsBacktrack(x, y) {
		// right is possible
		moves = append(moves, &State{Prev: s, X: x, Y: y})
	}
	return moves
}
