package main

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 15, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	gridStr, insStr, _ := strings.Cut(input, "\n\n")

	g := aoc.CreateFixedByteGridFromString(gridStr, '#')
	ins := bytes.TrimSpace([]byte(insStr))

	// find the positions of the robot.
	var robot aoc.V2

	aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
		if b == '@' {
			robot = v
		}
	})

	// follow the instructions,
	// if the space we want to move into is empty, move there.
	// if it is a wall, do nothing.
	// if it is a box, keep checking in the same direction until:
	// - we hit a wall - do nothing
	// - we hit a space - move the first box there, move the robot to the first box position.
	// when moving remember to leave an empty space.
	dirs := map[byte]aoc.V2{
		'^': aoc.North,
		'v': aoc.South,
		'<': aoc.West,
		'>': aoc.East,
	}
	for _, i := range ins {
		dir := dirs[i]
		b, _ := g.Atv(robot.Add(dir))
		switch b {
		case '.':
			// move into the space.
			g.Setv(robot, '.')
			robot = robot.Add(dir)
			g.Setv(robot, '@')
		case '#':
			// do nothing
		case 'O':
			// keep going
			initialBox := robot.Add(dir)
			notBox := initialBox.Add(dir)
			for {
				b, _ := g.Atv(notBox)
				if b == '.' {
					// move the box here
					g.Setv(initialBox, '.')
					g.Setv(notBox, 'O')
					// move the robot to the initial box position
					g.Setv(robot, '.')
					robot = initialBox
					g.Setv(robot, '@')
					break
				}
				if b == '#' {
					// nothing moves
					break
				}
				// keep going
				notBox = notBox.Add(dir)
			}
		}

	}

	sum := 0
	aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
		// positions of boxes
		if b == 'O' {
			sum += v.X + 100*v.Y
		}
	})

	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// change the map
	gridStr, insStr, _ := strings.Cut(input, "\n\n")
	gridStr = strings.NewReplacer("@", "@.", "O", "[]", "#", "##", ".", "..").Replace(gridStr)
	g := aoc.CreateFixedByteGridFromString(gridStr, '#')

	ins := bytes.TrimSpace([]byte(insStr))

	var robot aoc.V2
	aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
		if b == '@' {
			robot = v
		}
	})

	// similar to part 1, but now we have to keep track of the boxes which are wider.
	// so moving east/west is different to moving north/south.
	dirs := map[byte]aoc.V2{
		'^': aoc.North,
		'v': aoc.South,
		'<': aoc.West,
		'>': aoc.East,
	}
	t := 1
	for _, i := range ins {
		dir := dirs[i]
		b, _ := g.Atv(robot.Add(dir))
		switch b {
		case '.':
			// move into the space.
			g.Setv(robot, '.')
			robot = robot.Add(dir)
			g.Setv(robot, '@')
		case '#':
			// do nothing
		case '[', ']':
			switch dir {
			case aoc.North, aoc.South:
				if moveBoxesWideVertical(g, robot.Add(dir), dir) {
					// move the robot
					g.Setv(robot, '.')
					robot = robot.Add(dir)
					g.Setv(robot, '@')
				}
			case aoc.East, aoc.West:
				if moveBoxesWideHorizontal(g, robot.Add(dir), dir) {
					// move the robot
					g.Setv(robot, '.')
					robot = robot.Add(dir)
					g.Setv(robot, '@')
				}
			}
		}
		// fmt.Println("tick", t)
		// aoc.PrintByteGridC(g, map[byte]aoc.Color{
		// 	'#': aoc.BoldWhite,
		// 	'[': aoc.BoldYellow,
		// 	']': aoc.BoldYellow,
		// 	'@': aoc.BoldCyan,
		// })
		t++
	}

	// find the positions of the `[` part of the boxes
	sum := 0
	aoc.IterateByteGridv(g, func(v aoc.V2, b byte) {
		if b == '[' {
			sum += v.X + 100*v.Y
		}
	})

	return fmt.Sprint(sum)
}

// update the grid in place, just moving the boxes.
// return true if the robot can move into the space.
func moveBoxesWideHorizontal(g aoc.ByteGrid, initialBox, dir aoc.V2) bool {
	// horizontal movement, from the side we are pushing, the "next interesting space" is 2 away.
	next := initialBox.Add(dir).Add(dir)
	moves := 1
	for {
		b, _ := g.Atv(next)
		switch b {
		case '.':
			// move the box here, but we will only have shift the row "half" a box
			// so all the [] need to turn to ][
			boxOrder := []byte{'[', ']'}
			if dir == aoc.West {
				boxOrder = []byte{']', '['}
			}
			g.Setv(initialBox, '.') //reset the initial box postiion
			p := initialBox.Add(dir)
			for i := 0; i <= moves; i++ {
				g.Setv(p, boxOrder[i%2])
				p = p.Add(dir)
			}
			return true
		case '#':
			// cannot move anything
			return false
		}
		// keep going
		next = next.Add(dir)
		moves++
	}
}

func moveBoxesWideVertical(g aoc.ByteGrid, initialBox, dir aoc.V2) bool {
	// in this case a single box moving down could push 2 boxes in the next row (and 3 in the row after...)
	// so we need to keep track of how many boxes we are moving.
	front := []int{initialBox.X} // x coordinates of the moving boxes.
	if b, _ := g.At(initialBox.X, initialBox.Y); b == '[' {
		front = append(front, initialBox.X+1)
	} else {
		front = append(front, initialBox.X-1)
	}
	dy := 1
	if dir == aoc.North {
		dy = -1
	}
	initialRow := initialBox.Y
	boxesPerRow := [][]int{front}
	// check spaces in the next row:
	currentRow := initialRow
	for {
		currentRow += dy
		newBoxes := map[int]struct{}{}
		// fmt.Println("current row", currentRow)
		// fmt.Println("front", front)
		// fmt.Println("boxes per row", boxesPerRow)
		for _, x := range front {
			b, _ := g.At(x, currentRow)
			//fmt.Println("at", x, currentRow, "we have:", string(b))
			switch b {
			case '[', ']':
				newBoxes[x] = struct{}{}
				// also the one to the left or right
				if b == '[' {
					// the ] is to the right
					newBoxes[x+1] = struct{}{}
				} else {
					// the [ is to the left
					newBoxes[x-1] = struct{}{}
				}
			case '#':
				return false // block anywhere prevents motion
			}
		}

		if len(newBoxes) == 0 {
			// space to move all the boxes...

			// we have to work backwards
			// so we can move the boxes in the correct order.
			for j := len(boxesPerRow) - 1; j >= 0; j-- {
				rowBoxes := boxesPerRow[j]
				currentRow := initialRow + dy*j
				nextRow := initialRow + dy*(j+1)
				//fmt.Println("moving boxes", rowBoxes, "from", currentRow, "to", nextRow)
				for _, x := range rowBoxes {
					c, _ := g.At(x, currentRow)
					g.Set(x, currentRow, '.')
					g.Set(x, nextRow, c)
				}
			}
			return true
		}

		// more boxes to move.
		front = make([]int, 0, len(newBoxes))
		for x := range newBoxes {
			front = append(front, x)
		}
		boxesPerRow = append(boxesPerRow, front)
	}
}
