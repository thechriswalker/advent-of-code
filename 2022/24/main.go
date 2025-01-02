package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 24, solve1, solve2)
}

type Wind struct {
	pos aoc.V2
	dir aoc.V2
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// a grid, we always start at 1,0
	// and finish at max(x)-1, max(y)
	// but it will be easier to ignore the top and bottom rows, and the left and right columns.
	// then move 1 is always down and the final move is off the grid at the bottom.
	// we need to model all the wind and it's direction.
	// remember that the wind wraps, so the winds in each direction have a period of "width" or "height". This means that there might be periodicity?
	// anyway, we definitely need to model the wind by position and direction.
	// each step we can check multiple "pathways", but we need to "update" the wind onces per step (applies to all paths)
	zero := aoc.Vec2(0, 0)
	winds := []*Wind{}
	g := aoc.CreateFixedByteGridFromString(input, '#')
	aoc.IterateByteGrid(g, func(x, y int, b byte) {
		switch b {
		case '#':
			// wall pretend it is a wind that doesn't move
			winds = append(winds, &Wind{aoc.Vec2(x, y), zero})
		case '^':
			winds = append(winds, &Wind{aoc.Vec2(x, y), aoc.Vec2(0, -1)})
		case 'v':
			winds = append(winds, &Wind{aoc.Vec2(x, y), aoc.Vec2(0, 1)})
		case '<':
			winds = append(winds, &Wind{aoc.Vec2(x, y), aoc.Vec2(-1, 0)})
		case '>':
			winds = append(winds, &Wind{aoc.Vec2(x, y), aoc.Vec2(1, 0)})
		default:
			// do nothing
		}
	})

	x1, y1, x2, y2 := g.Bounds()
	target := aoc.Vec2(x2-1, y2)
	initialPos := aoc.Vec2(x1+1, y1)

	// maximum periodicity of the wind is lcm of x2-1 and y2-1
	// after this, they should be back where we were...
	period := aoc.LCM(x2-1, y2-1)
	// this means that if we end up in the same position this many steps ago, then
	// that is a dead end.
	// so we keep track of our steps

	// keep track of where the winds are.
	w := map[aoc.V2]int{}
	for _, wind := range winds {
		w[wind.pos] = w[wind.pos] + 1
	}

	updateWinds := func() {
		for _, wind := range winds {
			if wind.dir != zero {
				// remove the wind from the current position.
				w[wind.pos] = w[wind.pos] - 1
				// update wind position
				wind.pos.X += wind.dir.X
				if wind.dir.X == -1 && wind.pos.X == 0 {
					wind.pos.X = x2 - 1
				}
				if wind.dir.X == 1 && wind.pos.X == x2 {
					wind.pos.X = 1
				}
				wind.pos.Y += wind.dir.Y
				if wind.dir.Y == -1 && wind.pos.Y == 0 {
					wind.pos.Y = y2 - 1
				}
				if wind.dir.Y == 1 && wind.pos.Y == y2 {
					wind.pos.Y = 1
				}
				// add the wind to the new position
				w[wind.pos] = w[wind.pos] + 1
			}
		}
	}

	isSafe := func(p aoc.V2) bool {
		return w[p] <= 0
	}

	// step count
	step := 1
	positionsPerStep := []map[aoc.V2]struct{}{}

	// current positions
	curr := map[aoc.V2]struct{}{
		initialPos: {},
	}
	positionsPerStep = append(positionsPerStep, curr)
	next := make(map[aoc.V2]struct{}) // next possible positions

	maybeAddOption := func(t aoc.V2, options map[aoc.V2]struct{}) bool {
		if isSafe(t) {
			// were we here period ago?
			last := step - period
			if last >= 0 {
				if _, ok := positionsPerStep[last][t]; ok {
					// nothing!
					return false
				}
			}
			// add to set
			options[t] = struct{}{}
			return true
		}
		return false
	}

	for {
		updateWinds()
		for p := range curr {
			// we have 4 possible options
			// N, S, E, W and stay. if any are available, we should add them to the possibilities
			// stay still?
			maybeAddOption(p, next)
			// north
			maybeAddOption(aoc.Vec2(p.X, p.Y-1), next)
			// east
			maybeAddOption(aoc.Vec2(p.X+1, p.Y), next)
			// west
			maybeAddOption(aoc.Vec2(p.X-1, p.Y), next)
			// south
			if maybeAddOption(aoc.Vec2(p.X, p.Y+1), next) {
				if aoc.Vec2(p.X, p.Y+1) == target {
					// we are done!

					// 145 - too low? off by one, should we add the extra step now... as that is the "next" step...
					// nope 146 is also too low.
					// should I keep the actual route, so I can print it?
					// might be easier to see that is wrong...
					// stepping through this will be hellish with the number of options we have. stepping through the
					// single path I think is correct will still be difficult....

					return fmt.Sprint(step)
				}
			}
		}
		if len(next) == 0 {
			return "fail"
		}
		curr = next
		positionsPerStep = append(positionsPerStep, curr)
		next = map[aoc.V2]struct{}{}
		step++
		//	fmt.Println("step", step, "num options", len(curr))
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return "<unsolved>"
}
