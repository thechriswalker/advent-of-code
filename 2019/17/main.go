package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"

	"github.com/thechriswalker/advent-of-code/2019/intcode"
	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 17, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	m := explore(input)
	sum := 0
	for _, inter := range m.FindIntersections() {
		sum += inter[0] * inter[1]
		//fmt.Println(inter, "alignment", inter[0]*inter[1], "sum", sum)
	}
	//fmt.Print(m)

	return fmt.Sprint(sum)
}

func explore(code string) *Map {
	p := intcode.New(code)
	m := &Map{data: map[[2]int]byte{}}
	done := p.RunAsync()
	var x, y int
	for {
		select {
		case <-done:
			return m
		case c := <-p.Output:
			if byte(c) == '\n' {
				// new line!
				y++
				x = 0
			} else {
				switch byte(c) {
				case '^', 'v', '>', '<':
					m.robot = [2]int{x, y}
				}
				m.Set(x, y, byte(c))
				x++
			}
		}
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	/*
		m := explore(input)
		fmt.Print(m)
		p := m.WalkAll()
		fmt.Println(string(p))
	*/
	/*
		The output from the WalkAll was:

		R,10,R,10,R,6,R,4,R,10,R,10,L,4,R,10,R,10,R,6,R,4,R,4,L,4,L,10,L,10,R,10,R,10,R,6,R,4,R,10,R,10,L,4,R,4,L,4,L,10,L,10,R,10,R,10,L,4,R,4,L,4,L,10,L,10,R,10,R,10,L,4,

		I manually chopped up the string.
		It's some sort of maximum matching substring algorithm for the
		general solution, but it was easy enough to do by hand.

		R,10,R,10,R,6,R,4,  A
		R,10,R,10,L,4,      B
		R,10,R,10,R,6,R,4,  A
		R,4,L,4,L,10,L,10,  C
		R,10,R,10,R,6,R,4,  A
		R,10,R,10,L,4,      B
		R,4,L,4,L,10,L,10,  C
		R,10,R,10,L,4,      B
		R,4,L,4,L,10,L,10,  C
		R,10,R,10,L,4,      B
	*/

	A := []byte("R,10,R,10,R,6,R,4\n")
	B := []byte("R,10,R,10,L,4\n")
	C := []byte("R,4,L,4,L,10,L,10\n")
	R := []byte("A,B,A,C,A,B,C,B,C,B\n")

	ch := make(chan int64)
	go func() {
		// Main:
		for _, r := range R {
			ch <- int64(r)
		}
		// function A
		for _, r := range A {
			ch <- int64(r)
		}
		// function B
		for _, r := range B {
			ch <- int64(r)
		}
		// function C
		for _, r := range C {
			ch <- int64(r)
		}
		// continuous video feed?
		ch <- int64('n')
		ch <- int64('\n')
		// all input given
	}()

	feeder := func() int64 {
		return <-ch
	}

	pg := intcode.New(input)
	pg.Set(0, 2)
	done := pg.RunAsync()
	halt := false
	var result int64
	for {
		select {
		case pg.Input <- feeder:
			// fine keep feeding it
		case c := <-pg.Output:
			// the result! (maybe)
			result = c
		case <-done:
			halt = true
		}
		if halt {
			break
		}
	}

	return fmt.Sprint(result)
}

type Map struct {
	data       map[[2]int]byte
	xmax, ymax int
	robot      [2]int
}

func (m *Map) Set(x, y int, c byte) {
	m.data[[2]int{x, y}] = c
	if x > m.xmax {
		m.xmax = x
	}
	if y > m.ymax {
		m.ymax = y
	}
}

func (m *Map) IsScaffold(x, y int) bool {
	c := m.data[[2]int{x, y}]
	return c == '#' || c == 'O'
}

func (m *Map) IsIntersection(x, y int) bool {
	// it is an intersection if the 4 adjacent tiles are also scaffold
	c := m.data[[2]int{x, y}]
	if c == 'O' {
		return true
	}
	return m.IsScaffold(x, y) && m.IsScaffold(x+1, y) && m.IsScaffold(x-1, y) && m.IsScaffold(x, y+1) && m.IsScaffold(x, y-1)
}

func (m *Map) FindIntersections() [][2]int {
	intersections := [][2]int{}
	for y := 1; y < m.ymax; y++ {
		for x := 1; x < m.xmax; x++ {
			if m.IsIntersection(x, y) {
				m.Set(x, y, 'O')
				intersections = append(intersections, [2]int{x, y})
			}
		}
	}
	return intersections
}

func (m *Map) String() string {
	sb := strings.Builder{}
	m.WriteInto(&sb)
	return sb.String()
}

func (m *Map) WriteInto(w io.Writer) {
	for y := 0; y <= m.ymax; y++ {
		for x := 0; x <= m.xmax; x++ {
			c := m.data[[2]int{x, y}]
			switch c {
			// this is for color!
			case '.':
				w.Write([]byte("\x1b[1;90m"))
			case '#':
				w.Write([]byte("\x1b[1;93m"))
			case '^', '<', '>', 'v':
				w.Write([]byte("\x1b[1;96m"))
			case 'O':
				w.Write([]byte("\x1b[1;95m"))
			default:
				// nothing
				w.Write([]byte("\x1b[0m"))
			}
			w.Write([]byte{c})
		}
		// clear and add newline
		w.Write([]byte("\x1b[0m\n"))
	}
}

func (m *Map) WalkAll() []byte {
	r := m.robot
	// we will find the one direction to move
	var path bytes.Buffer
	dir := m.data[r]
	steps := 0
	// keep going in the direction until we hit an impass, count the steps.
	for {
		if n, ok := m.CanMove(r, dir); ok {
			steps++
			r = n
		} else {
			// cannot move.
			if steps > 0 {
				fmt.Fprint(&path, steps)
				path.WriteByte(',')
				steps = 0
			}
			if t, d, ok := m.NextDir(r, dir); !ok {
				// no more moves. the end!
				return path.Bytes()
			} else {
				// turn to d
				path.WriteByte(t)
				path.WriteByte(',')
				dir = d
			}
		}
	}
}

func (m *Map) CanMove(r [2]int, dir byte) ([2]int, bool) {
	//fmt.Print("m.CanMove(", r, ",", string(dir), ")")
	switch dir {
	case 'v':
		r[1]++
	case '^':
		r[1]--
	case '>':
		r[0]++
	case '<':
		r[0]--
	}
	c := m.data[r]
	//fmt.Print("{ n =", r, "; m.data[n] =", string(c), "}\n")
	return r, c == '#' || c == 'O'
}

func (m *Map) NextDir(r [2]int, dir byte) (byte, byte, bool) {
	// only 2 direction options left turn or right turn.
	var t1, t2 byte
	switch dir {
	case 'v', '^':
		t1, t2 = 'L', 'R'
		if dir == '^' {
			t1, t2 = 'R', 'L'
		}
		if _, ok := m.CanMove(r, '>'); ok {
			return t1, '>', true
		}
		if _, ok := m.CanMove(r, '<'); ok {
			return t2, '<', true
		}
	case '<', '>':
		t1, t2 = 'L', 'R'
		if dir == '>' {
			t1, t2 = 'R', 'L'
		}
		if _, ok := m.CanMove(r, 'v'); ok {
			return t1, 'v', true
		}
		if _, ok := m.CanMove(r, '^'); ok {
			return t2, '^', true
		}
	}
	return 'X', 'X', false
}
