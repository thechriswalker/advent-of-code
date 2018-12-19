package main

import (
	"bytes"
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2018, 18, solve1, solve2)
}

// these are known outside of the puzzle.
// we make them globals here so they can be changed for the tests
var (
	Width = 50
)

const (
	TREE = '|'
	OPEN = '.'
	YARD = '#'
)

// Implement Solution to Problem 1
func solve1(input string) string {
	prev := parseInput(input)
	next := make(State, len(prev))
	//	fmt.Println(prev)
	for i := 0; i < 10; i++ {
		prev.Next(next)
		prev, next = next, prev
		//	fmt.Println(prev)
	}

	return fmt.Sprintf("%d", prev.ResourceValue())
}

// Implement Solution to Problem 2
// after watching the pattern for ~300 iterations it start to repeat.
// so we will capture the value at 500, then see what the period of repetition is
// then we can work out the value of the billion iterations as modulo the period
// then we can just work out that
func solve2(input string) string {
	prev := parseInput(input)
	next := make(State, len(prev))
	ticks := 0
	repetitionPoint := 500
	for {
		prev.Next(next)
		prev, next = next, prev
		ticks++
		if ticks == repetitionPoint {
			break
		}
	}

	// we are in the repetition now. lets find the period.
	checkpoint := make(State, len(prev))
	copy(checkpoint, prev)
	// fmt.Println("At tick", repetitionPoint)
	// fmt.Println(checkpoint)
	// fmt.Println()
	for {
		prev.Next(next)
		prev, next = next, prev
		ticks++
		//	lets print it anyway
		// fmt.Println("Tick", ticks)
		// fmt.Println(prev)
		if bytes.Compare(prev, checkpoint) == 0 {
			break
		}
		// fmt.Println("\x1b[53F")
		// time.Sleep(time.Millisecond * 100)
	}
	// we hit the same point again.
	period := ticks - repetitionPoint
	// fmt.Println("Repetition detected, period =", period)
	desired := (1000000000 - repetitionPoint) % period
	// run that many more ticks.
	for i := 0; i < desired; i++ {
		prev.Next(next)
		prev, next = next, prev
	}

	return fmt.Sprintf("%d", prev.ResourceValue())
}

func parseInput(input string) State {
	rd := strings.NewReader(input)
	state := make(State, 0, len(input))
	var b byte
	var err error
	for {
		b, err = rd.ReadByte()
		if err != nil {
			break
		}
		switch b {
		case TREE, OPEN, YARD:
			state = append(state, b)
		}
	}
	return state
}

type State []byte

func (s State) ResourceValue() int {
	var trees, yards int
	for _, b := range s {
		switch b {
		case TREE:
			trees++
		case YARD:
			yards++
		}
	}
	return trees * yards
}

func (s State) Next(next State) bool {
	var trees, yards int
	var change bool
	for i, b := range s {
		trees, _, yards = s.Adjacents(i)
		switch b {
		case OPEN:
			if trees >= 3 {
				next[i] = TREE
				change = true
			} else {
				next[i] = OPEN
			}
		case TREE:
			if yards >= 3 {
				next[i] = YARD
				change = true
			} else {
				next[i] = TREE
			}
		case YARD:
			if yards >= 1 && trees >= 1 {
				next[i] = YARD
			} else {
				next[i] = OPEN
				change = true
			}
		}
	}
	return change
}

func (s State) Adjacents(index int) (trees, opens, yards int) {
	// the eight adjacent squares.
	// we need to deal in X,Y values
	// to ensure we don't extend past a width.
	var tree, open, yard int
	x := index % Width
	y := index / Width
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 && j == 0 {
				continue
			}
			tree, open, yard = s.Counts(x+i, y+j)
			trees += tree
			opens += open
			yards += yard
		}
	}
	return
}

func (s State) Counts(x, y int) (trees, opens, yards int) {
	if x < 0 || x >= Width || y < 0 || y >= len(s)/Width {
		return
	}
	index := x + y*Width
	switch s[index] {
	case TREE:
		trees = 1
	case OPEN:
		opens = 1
	case YARD:
		yards = 1
	}
	return
}

func (s State) String() string {
	b := strings.Builder{}
	for i := range s {
		b.WriteByte(s[i])
		if i%Width == Width-1 {
			b.WriteByte('\n')
		}
	}
	return b.String()
}
