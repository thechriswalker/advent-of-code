package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2020, 12, solve1, solve2)
}

type action struct {
	kind  byte
	value int
}

func parseInput(in string) []action {
	lines := strings.Split(in, "\n")
	stack := make([]action, 0, len(lines))
	for _, line := range lines {
		if line == "" {
			continue
		}
		v, err := strconv.Atoi(line[1:])
		if err != nil {
			continue
		}
		stack = append(stack, action{kind: line[0], value: v})
	}
	return stack
}

type position [2]int

var (
	EAST  position = position{0, 1}
	WEST  position = position{0, -1}
	NORTH position = position{1, 0}
	SOUTH position = position{-1, 0}
)

func (p position) mahattanDistance() int {
	return abs(p[0]) + abs(p[1])
}

func rotate(from position, clockwise bool, degrees int) position {
	turns := degrees / 90
	y, x := from[0], from[1]
	for i := 0; i < turns; i++ {
		if clockwise {
			// we need to
			y, x = -1*x, y
		} else {
			y, x = x, -1*y
		}
	}
	return position{y, x}
}

func move(start, d position, distance int) position {
	y, x := start[0], start[1]
	for i := 0; i < distance; i++ {
		y, x = y+d[0], x+d[1]
	}
	return position{y, x}
}

// Implement Solution to Problem 1
func solve1(input string) string {
	facing := EAST
	pos := position{0, 0}
	actions := parseInput(input)

	for _, action := range actions {
		//fmt.Printf("before: facing:%v pos:%v action:%c %d\n", facing, pos, action.kind, action.value)
		switch action.kind {
		case 'F':
			// move in position.
			pos = move(pos, facing, action.value)
		case 'N':
			pos = move(pos, NORTH, action.value)
		case 'S':
			pos = move(pos, SOUTH, action.value)
		case 'E':
			pos = move(pos, EAST, action.value)
		case 'W':
			pos = move(pos, WEST, action.value)
		case 'L':
			facing = rotate(facing, false, action.value)
		case 'R':
			facing = rotate(facing, true, action.value)
		}
	}

	return fmt.Sprintf("%d", pos.mahattanDistance())
}

func abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

// Implement Solution to Problem 2
func solve2(input string) string {
	waypoint := position{1, 10} // relative position
	pos := position{0, 0}
	actions := parseInput(input)
	for _, action := range actions {
		switch action.kind {
		case 'F':
			// move to the waypoint x times.
			pos = move(pos, waypoint, action.value)
		case 'N':
			// move waypoint
			waypoint = move(waypoint, NORTH, action.value)
		case 'S':
			waypoint = move(waypoint, SOUTH, action.value)
		case 'E':
			waypoint = move(waypoint, EAST, action.value)
		case 'W':
			waypoint = move(waypoint, WEST, action.value)
		case 'L':
			waypoint = rotate(waypoint, false, action.value)
		case 'R':
			waypoint = rotate(waypoint, true, action.value)
		}
	}

	return fmt.Sprintf("%d", pos.mahattanDistance())
}
