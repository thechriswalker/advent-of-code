package main

import (
	"fmt"
	"sort"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2018, 13, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	track, prev := parseInput(input)
	next := prev.Copy()
	// track.Print(prev, true)
	for {
		collision, at, _ := tick(track, prev, next)
		if collision {
			return fmt.Sprintf("%d,%d", at[0], at[1])
		}
		// track.Print(next, true)
		// time.Sleep(time.Millisecond * 50)
		prev, next = next, prev
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	track, prev := parseInput(input)
	next := prev.Copy()
	// fmt.Println()
	//	track.Print(prev, true)
	var ticks int
	for {
		ticks++
		//fmt.Println("tick", ticks)
		_, _, n := tick(track, prev, next)
		if n <= 1 {
			// find the coords of the only remaining cart
			for _, c := range next {
				if !c.Destroyed {
					return fmt.Sprintf("%d,%d", c.X, c.Y)
				}
			}
			return fmt.Sprintf("All carts destroyed")
		}
		// track.Print(next, true)
		// time.Sleep(time.Millisecond * 500)
		prev, next = next, prev
	}
}

// ticks and writes the new positions into next based on current in prev.
// if there is a collision, we return true and the coordinates
// final result is the number of carts left
func tick(tracks *Tracks, prev, next CartList) (bool, [2]int, int) {
	sort.Sort(prev)
	collisionList := make([][2]int, len(prev))
	//put are current positions into the list
	var cartCount int
	offCanvas := [2]int{-1, -1}
	for i, c := range prev {
		if c.Destroyed {
			collisionList[i] = offCanvas
		} else {
			cartCount++
			collisionList[i] = [2]int{c.X, c.Y}
		}
		next[i].Id = c.Id
		next[i].Destroyed = c.Destroyed
		next[i].X = c.X
		next[i].Y = c.Y
		next[i].Direction = c.Direction
		next[i].NextTurn = c.NextTurn
	}

	var hasCollision bool
	var firstCollision [2]int // the x,y coords of the first collision
	for i := range next {
		if next[i].Destroyed {
			continue
		}
		for j := range collisionList {
			if i == j || next[j].Destroyed || collisionList[j] == offCanvas {
				continue
			}
			if collisionList[i] == collisionList[j] {
				// remove BOTH carts from the cart list.
				// but keep going
				// fmt.Printf("Carts Destroyed: %d and %d at (%d,%d)\n", next[i].Id, next[j].Id, collisionList[i][0], collisionList[j][1])
				// fmt.Println("i:", next[i], "j:", next[j])
				if !hasCollision {
					hasCollision = true
					firstCollision = collisionList[i]
				}
				collisionList[i] = offCanvas
				collisionList[j] = offCanvas
				next[i].Destroyed = true
				next[j].Destroyed = true
				cartCount -= 2
				if cartCount <= 1 {
					return hasCollision, firstCollision, cartCount
				}
				break
			}
		}
		if next[i].Destroyed {
			continue
		}
		// move the cart in the direction it is facing.
		// if the new track is a corner, or junction, change
		// direction.

		switch next[i].Direction {
		case NORTH:
			next[i].Y--
		case SOUTH:
			next[i].Y++
		case EAST:
			next[i].X++
		case WEST:
			next[i].X--
		}
		collisionList[i] = [2]int{next[i].X, next[i].Y}

		switch tracks.At(next[i].X, next[i].Y) {
		case JUNCTION:
			// turn in "next direction" and increment
			switch next[i].NextTurn {
			case STRAIGHT:
				// nothing to do
			case LEFT:
				// depends on current direction
				switch next[i].Direction {
				case NORTH:
					next[i].Direction = WEST
				case WEST:
					next[i].Direction = SOUTH
				case SOUTH:
					next[i].Direction = EAST
				case EAST:
					next[i].Direction = NORTH
				}
			case RIGHT:
				switch next[i].Direction {
				case NORTH:
					next[i].Direction = EAST
				case WEST:
					next[i].Direction = NORTH
				case SOUTH:
					next[i].Direction = WEST
				case EAST:
					next[i].Direction = SOUTH
				}
			}
			// update turn
			next[i].NextTurn = (next[i].NextTurn + 1) % 3

		case TURN_FORWARD:
			// turn / if we are moving NORTH turn to EAST
			// if WEST turn SOUTH, etc...
			switch next[i].Direction {
			case NORTH:
				next[i].Direction = EAST
			case WEST:
				next[i].Direction = SOUTH
			case SOUTH:
				next[i].Direction = WEST
			case EAST:
				next[i].Direction = NORTH
			}
		case TURN_BACK:
			// turn \ if we are moving NORTH turn WEST, etc...
			switch next[i].Direction {
			case NORTH:
				next[i].Direction = WEST
			case WEST:
				next[i].Direction = NORTH
			case SOUTH:
				next[i].Direction = EAST
			case EAST:
				next[i].Direction = SOUTH
			}
		}
	}

	return hasCollision, firstCollision, cartCount
}

func parseInput(input string) (*Tracks, CartList) {
	track := []byte{}
	carts := CartList{}

	rd := strings.NewReader(input)
	var b byte
	var err error
	var x, y int
	var width int
	var cart int
	for {
		b, err = rd.ReadByte()
		if err != nil {
			break
		}
		//		fmt.Printf("Handling byte '%c' (%d)\n", b, b)
		switch b {
		case EMPTY, JUNCTION, HORIZONTAL, VERTICAL, TURN_BACK, TURN_FORWARD:
			track = append(track, b)
			x++
		case '\n':
			// new line
			y++
			width = x
			x = 0
		case NORTH, SOUTH, EAST, WEST:
			// cart.
			carts = append(carts, &Cart{
				Id: cart,
				X:  x, Y: y,
				Direction: b, NextTurn: LEFT,
			})
			cart++
			// but also track
			switch b {
			case EAST, WEST:
				track = append(track, HORIZONTAL)
			case NORTH, SOUTH:
				track = append(track, VERTICAL)
			}
			x++
		}
	}

	return &Tracks{
		Track: track,
		Width: width,
	}, carts
}

type Tracks struct {
	Width int
	Track []byte
}

func (t *Tracks) Index(x, y int) int {
	if x < 0 || x >= t.Width || y < 0 || y >= len(t.Track)/t.Width {
		return -1
	}
	return y*t.Width + x
}

func (t *Tracks) At(x, y int) byte {
	i := t.Index(x, y)
	if i == -1 {
		return EMPTY
	}
	return t.Track[i]
}

const (
	EMPTY        = ' '
	JUNCTION     = '+'
	HORIZONTAL   = '-'
	VERTICAL     = '|'
	TURN_FORWARD = '/'
	TURN_BACK    = '\\'
)

type CartList []*Cart

func (c CartList) Len() int      { return len(c) }
func (c CartList) Swap(i, j int) { c[i], c[j] = c[j], c[i] }
func (c CartList) Less(i, j int) bool {
	// y first then x
	if c[i].Y == c[j].Y {
		return c[i].X < c[j].X
	}
	return c[i].Y < c[j].Y
}

func (cl CartList) Copy() CartList {
	cpy := make(CartList, len(cl))
	for i, c := range cl {
		cpy[i] = &Cart{
			Id:        c.Id,
			Destroyed: c.Destroyed,
			X:         c.X,
			Y:         c.Y,
			Direction: c.Direction,
			NextTurn:  c.NextTurn,
		}
	}
	return cpy
}

type Cart struct {
	Id        int
	Destroyed bool
	X, Y      int
	Direction byte
	NextTurn  Turn
}

const (
	NORTH = '^'
	SOUTH = 'v'
	EAST  = '>'
	WEST  = '<'
)

type Turn uint8

const (
	LEFT     Turn = 0
	STRAIGHT Turn = 1
	RIGHT    Turn = 2
)

func (t *Tracks) Print(c CartList, resetCursor bool) {
	// basically loop through and print all track except if there is a cart at the index.
	// so first we create a map of index -> cart direction
	carts := make(map[int]*Cart, len(c))
	for _, cart := range c {
		if !cart.Destroyed {
			i := t.Index(cart.X, cart.Y)
			carts[i] = cart
		}
	}
	var lines int
	// now the print loop
	for i, b := range t.Track {
		if cart, ok := carts[i]; ok {
			fmt.Printf("\x1b[1;3%dm%d\x1b[0m", cart.Id, cart.Id)
		} else {
			fmt.Printf("%c", b)
		}
		if i%t.Width == t.Width-1 {
			fmt.Print("\n")
			lines++
		}
	}
	if resetCursor {
		fmt.Printf("\x1b[%dF", lines)
	}
}
