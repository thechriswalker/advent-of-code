package main

import (
	"fmt"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 2, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	keypad := &Keypad{
		Size: 3,
		Pad: [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		},
	}
	state := State{1, 1}
	return getCode(keypad, state, input)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	keypad := &Keypad{
		Size: 5,
		Pad: [][]int{
			{0, 0, 1, 0, 0},
			{0, 2, 3, 4, 0},
			{5, 6, 7, 8, 9},
			{0, 10, 11, 12, 0},
			{0, 0, 13, 0, 0},
		},
	}
	// start on the 5
	state := State{2, 0}

	return getCode(keypad, state, input)
}

type Keypad struct {
	Size int
	Pad  [][]int // will be [Size][Size]int
}

type State struct {
	X, Y int
}

func (k *Keypad) KeyAt(s State) string {
	v := k.Pad[s.X][s.Y]
	return fmt.Sprintf("%X", v)
}
func (k *Keypad) ValidState(s State) bool {
	if s.X < 0 || s.X >= k.Size || s.Y < 0 || s.Y >= k.Size {
		return false
	}
	return k.Pad[s.X][s.Y] > 0
}

type Dir uint

const (
	Up Dir = iota
	Down
	Left
	Right
)

type Sequence []Dir

func move(s State, dir Dir) State {
	switch dir {
	case Up:
		s.X--
	case Down:
		s.X++
	case Left:
		s.Y--
	case Right:
		s.Y++
	default:
		panic("bad direction")
	}
	return s
}

func runSequence(s State, k *Keypad, seq Sequence) State {
	for _, d := range seq {
		//fmt.Printf("move: %v\n", d)
		if next := move(s, d); k.ValidState(next) {
			s = next
		}
		//fmt.Printf("new state: %v\n", s)
	}
	return s
}

func getCode(pad *Keypad, state State, in string) string {
	seqs := parseSequences(in)
	code := ""
	for _, s := range seqs {
		state = runSequence(state, pad, s)
		code += pad.KeyAt(state)
	}
	return code
}

func parseSequences(in string) []Sequence {
	seqs := []Sequence{}
	curr := Sequence{}

	for _, tok := range in {
		switch tok {
		case '\n':
			// next sequence
			seqs = append(seqs, curr)
			curr = Sequence{}
		case 'U':
			curr = append(curr, Up)
		case 'D':
			curr = append(curr, Down)
		case 'L':
			curr = append(curr, Left)
		case 'R':
			curr = append(curr, Right)
		default:
			panic(fmt.Sprintf("unexpected character in input: %c (%d)", tok, tok))
		}
	}
	if len(curr) > 0 {
		seqs = append(seqs, curr)
	}
	//fmt.Printf("Input: %s\n\nSequences: %v", in, seqs)
	return seqs
}
