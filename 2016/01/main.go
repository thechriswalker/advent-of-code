package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 1, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(in string) string {
	instructions := parseInstructions(in)
	state := State{Dir: North}
	for _, i := range instructions {
		//fmt.Println("Instruction:", i)
		state = turn(state, i.Turn)
		state = move(state, i.Move)
		//fmt.Println("State:", state)
	}
	return fmt.Sprintf("%d", abs(state.X)+abs(state.Y))
}

// Implement Solution to Problem 2
func solve2(in string) string {
	instructions := parseInstructions(in)
	state := State{Dir: North}
	visitMap := map[Point]bool{}
	revisit := false
	var p Point
	for _, i := range instructions {
		//fmt.Println("Instruction:", i)
		state = turn(state, i.Turn)
		// now we move one step at a time

		for d := i.Move; d > 0; d-- {
			state = move(state, 1)
			p.X = state.X
			p.Y = state.Y
			if revisit = visitMap[p]; revisit {
				break
			} else {
				visitMap[p] = true
			}
		}
		if revisit {
			break
		}
		//fmt.Println("State:", state)
	}
	return fmt.Sprintf("%d", abs(state.X)+abs(state.Y))
}

type Point struct{ X, Y int }

type Dir int

const (
	North Dir = 0
	East      = 1
	South     = 2
	West      = 3
)

func (d Dir) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		panic("unknown direction")
	}
}

type State struct {
	Dir  Dir
	X, Y int
}

type Turn int

const (
	Right Turn = 1
	Left       = -1
)

func (t Turn) String() string {
	switch t {
	case Left:
		return "Left"
	case Right:
		return "Right"
	default:
		panic("unknown turn direction")
	}
}

type Instruction struct {
	Turn Turn
	Move int
}

func turn(s State, t Turn) State {
	s.Dir = Dir((int(s.Dir) + int(t) + 4) % 4)
	return s
}

func move(s State, d int) State {
	switch s.Dir {
	case North:
		s.Y += d
	case South:
		s.Y -= d
	case East:
		s.X += d
	case West:
		s.X -= d
	default:
		panic(fmt.Sprintf("Bad state, invalid direction: %d", s.Dir))
	}

	return s
}

func whichTurn(s string) Turn {
	switch s {
	case "R":
		return Right
	case "L":
		return Left
	default:
		panic(fmt.Sprintf("unknown turn direction: %s", s))
	}
}

func parseInstructions(in string) []Instruction {
	var turn string
	var dist int
	instructions := []Instruction{}
	r := strings.NewReader(in)
	for {
		n, _ := fmt.Fscanf(r, "%1s%d,", &turn, &dist)
		if n == 2 {
			instructions = append(instructions, Instruction{
				Turn: whichTurn(turn),
				Move: dist,
			})
		} else {
			break
		}
	}
	return instructions
}

func abs(d int) int {
	if d < 0 {
		return -1 * d
	}
	return d
}
