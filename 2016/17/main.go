package main

import (
	"crypto/md5"
	"fmt"
	"hash"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 17, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	base := strings.TrimSpace(input)
	initialState := &State{X: 0, Y: 0}
	depth := 0
	var nextMoves []*State
	pendingMoves := []*State{initialState}
	h := md5.New()
	for {
		nextMoves = []*State{}
		for _, move := range pendingMoves {
			if move.X == MAX_ROOM && move.Y == MAX_ROOM {
				return move.Path
			}
			// add more moves.
			nextMoves = append(nextMoves, move.NextMoves(h, base)...)
		}
		if len(nextMoves) == 0 {
			return "<fail>"
		}
		depth++
		pendingMoves = nextMoves
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	base := strings.TrimSpace(input)
	initialState := &State{X: 0, Y: 0}
	depth := 0
	var nextMoves []*State
	pendingMoves := []*State{initialState}
	h := md5.New()
	max := 0
	for {
		nextMoves = []*State{}
		for _, move := range pendingMoves {
			if move.X == MAX_ROOM && move.Y == MAX_ROOM {
				if max < depth {
					max = depth
				}
				continue // just stop here
			}
			// add more moves.
			nextMoves = append(nextMoves, move.NextMoves(h, base)...)
		}
		if len(nextMoves) == 0 {
			return fmt.Sprintf("%d", max)
		}
		depth++
		pendingMoves = nextMoves
	}
}

// same breath first search on a plane we used for the 2016/13
// except we need the "path so far" as input to the NextMoves()

const (
	MIN_ROOM = 0
	MAX_ROOM = 3
)

type State struct {
	Prev *State
	X    uint
	Y    uint
	Path string
}

func (s *State) DoorsOpen(h hash.Hash, base string) (up bool, down bool, left bool, right bool) {
	h.Reset()
	fmt.Fprintf(h, "%s%s", base, s.Path)
	b := h.Sum(nil)
	// first 4 hex characters for up/down/left/right
	// each 2 hex chars are 1 byte
	// so up is b[0] >> 1 and down is b[0]
	up = b[0]>>4 > 0x0a
	down = b[0]&0x0f > 0x0a
	left = b[1]>>4 > 0x0a
	right = b[1]&0x0f > 0x0a
	return
}

func (s *State) NextMoves(h hash.Hash, base string) []*State {
	// at most 4 new states.
	moves := make([]*State, 0, 4)
	up, down, left, right := s.DoorsOpen(h, base)
	if up && s.Y != MIN_ROOM {
		moves = append(moves, &State{
			X:    s.X,
			Y:    s.Y - 1,
			Path: s.Path + "U",
		})
	}

	if left && s.X != MIN_ROOM {
		moves = append(moves, &State{
			X:    s.X - 1,
			Y:    s.Y,
			Path: s.Path + "L",
		})
	}

	if down && s.Y != MAX_ROOM {
		moves = append(moves, &State{
			X:    s.X,
			Y:    s.Y + 1,
			Path: s.Path + "D",
		})
	}
	if right && s.X != MAX_ROOM {
		moves = append(moves, &State{
			X:    s.X + 1,
			Y:    s.Y,
			Path: s.Path + "R",
		})
	}
	return moves
}
