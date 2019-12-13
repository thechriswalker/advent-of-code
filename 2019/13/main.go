package main

import (
	"fmt"
	"io"
	"os"

	"../../aoc"
	"../intcode"
)

func main() {
	aoc.Run(2019, 13, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	pg := intcode.New(input)
	pg.RunAsync()
	screen := &Screen{pixels: map[[2]int64]int64{}}

	var x, y, t int64

	done := false
	for {
		// we wait for the program to halt collecting 3 outputs at a time.
		// or halt if program halts
		select {
		case <-pg.Halted:
			done = true
		case x = <-pg.Output:
		}
		if done {
			break
		}
		y = <-pg.Output
		t = <-pg.Output
		// draw the pixel
		screen.draw(x, y, t)
	}
	// now count the "block tiles"
	c := 0
	for _, t := range screen.pixels {
		if t == tile_block {
			c++
		}
	}
	return fmt.Sprint(c)

}

const (
	joy_left    = int64(-1)
	joy_neutral = 0
	joy_right   = 1
)

// Implement Solution to Problem 2
func solve2(input string) string {
	// for input we will provide interactive control via stdin
	// get := func() int64 {
	// 	rd := bufio.NewReaderSize(os.Stdin, 1)
	// 	in, _ := rd.ReadByte()
	// 	switch in {
	// 	case 'a':
	// 		return joy_left
	// 	case 's':
	// 		return joy_neutral
	// 	case 'd':
	// 		return joy_right
	// 	}
	// 	panic("bad input")
	// }

	// now we use AI to play the game.
	// always move towards the last "ball" X coordinate
	arcade := &Arcade{screen: &Screen{pixels: map[[2]int64]int64{}}}

	get := func() int64 {
		// move padd_x towards ball_x
		if arcade.padd_x == arcade.ball_x {
			return joy_neutral
		}
		if arcade.padd_x < arcade.ball_x {
			// paddle is left of ball
			return joy_right
		}
		return joy_left
	}

	arcade.Play(input, get, false)

	return fmt.Sprint(arcade.score)
}

type Arcade struct {
	screen *Screen
	score  int64
	ticks  int64
	ball_x int64
	padd_x int64
}

func (a *Arcade) Play(code string, input func() int64, draw bool) {
	pg := intcode.New(code)
	pg.Set(0, 2)
	pg.RunAsync()

	init := false

	redraw := func() {
		if !draw || !init {
			return
		}
		//	time.Sleep(10 * time.Millisecond)
		// now display the game!
		// wipe the terminal, move to home
		fmt.Print("\x1b[2J\x1b[H")
		// output Score line:
		fmt.Printf("[ARCADE] score: %06d [ticks:%06d]\n", a.score, a.ticks)
		// output screen
		a.screen.Print(os.Stdout)
	}

	var x, y, t int64
	for {
		redraw()
		// we wait for the program to halt collecting 3 outputs at a time.
		// or halt if program halts
		select {
		case <-pg.Halted:
			fmt.Println("Game Over!")
			return
		case pg.Input <- input:
			// first input is initialisation
			init = true
		case x = <-pg.Output:
			y = <-pg.Output
			t = <-pg.Output
			if x == -1 && y == 0 {
				inc := t - a.score
				fmt.Printf("Score increased by %d, to %d\n", inc, t)
				a.score = t
			} else {
				if t == tile_ball {
					a.ball_x = x
				}
				if t == tile_paddle {
					a.padd_x = x
				}
				a.screen.draw(x, y, t)
			}
		}
		a.ticks++
	}

}

// we don't know how big the screen is so keep track of the
// bounds, and use a sparse map
type Screen struct {
	pixels                 map[[2]int64]int64
	xmin, xmax, ymin, ymax int64
}

const (
	tile_empty  = int64(0)
	tile_wall   = 1
	tile_block  = 2
	tile_paddle = 3
	tile_ball   = 4
)

func (s *Screen) draw(x, y, tile int64) {
	s.pixels[[2]int64{x, y}] = tile
	if x < s.xmin {
		s.xmin = x
	}
	if x > s.xmax {
		s.xmax = x
	}
	if y < s.ymin {
		s.ymin = y
	}
	if y > s.ymax {
		s.ymax = y
	}
}

func (s *Screen) Print(w io.Writer) {
	for y := s.ymin - 1; y <= s.ymax+1; y++ {
		for x := s.xmin - 1; x <= s.xmax+1; x++ {
			tile, ok := s.pixels[[2]int64{x, y}]
			if !ok {
				w.Write([]byte{' '})
			} else {
				switch tile {
				case tile_empty:
					w.Write([]byte{' '})
				case tile_wall:
					w.Write([]byte("\x1b[1;90m#\x1b[0m"))
				case tile_block:
					w.Write([]byte("\x1b[1;97m#\x1b[0m"))
				case tile_ball:
					w.Write([]byte("\x1b[1;93mO\x1b[0m"))
				case tile_paddle:
					w.Write([]byte("\x1b[1;96m=\x1b[0m"))
				default:
					panic("unknown tile")
				}
			}
		}
		// add a newline
		w.Write([]byte{'\n'})
	}
}
