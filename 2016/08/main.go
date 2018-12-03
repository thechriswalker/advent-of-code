package main

import (
	"fmt"
	"io"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 8, solve1, solve2)
}

// so I can change for tests
var (
	Width  = 50
	Height = 6
	Space  = ' '
)

// Implement Solution to Problem 1
func solve1(input string) string {
	s := NewSreen(Width, Height)
	RunOperations(input, s)
	return fmt.Sprintf("%d", s.CountLitPixel())
}

// Implement Solution to Problem 2
func solve2(input string) string {
	s := NewSreen(Width, Height)
	RunOperations(input, s)
	return s.String()
}

func NewSreen(width, height int) *Screen {
	return &Screen{
		Stride: width,
		Height: height,
		Pixels: make([]uint8, width*height),
	}
}

func RunOperations(input string, screen *Screen) {
	rd := strings.NewReader(input)
	var a, b, n int
	var err error
	buf := make([]byte, 2)

	for {
		_, err = rd.Read(buf)
		if err != nil {
			break
		}
		if buf[1] == 'e' {
			// rect AxB seek to A
			rd.Seek(3, io.SeekCurrent)
			n, _ = fmt.Fscanf(rd, "%dx%d\n", &a, &b)
			if n != 2 {
				break //panic?
			}
			screen.Rect(a, b)
			continue
		}
		// rotate
		rd.Seek(5, io.SeekCurrent)
		_, err = rd.Read(buf)
		if err != nil {
			break
		}
		if buf[0] == 'r' {
			// row
			rd.Seek(4, io.SeekCurrent)
			n, _ = fmt.Fscanf(rd, "%d by %d\n", &a, &b)
			if n != 2 {
				break //panic?
			}
			screen.RotateRow(a, b)
			continue
		}
		// column
		rd.Seek(7, io.SeekCurrent)
		n, _ = fmt.Fscanf(rd, "%d by %d\n", &a, &b)
		if n != 2 {
			break //panic?
		}
		screen.RotateColumn(a, b)
		continue
	}
}

type Screen struct {
	Stride int
	Height int
	Pixels []uint8 // W*D regular stride, 0 = off, 1 = on
}

// rect AxB turns on all of the pixels in a
// rectangle at the top-left of the screen which is A wide and B tall.
func (s *Screen) Rect(a, b int) {
	for i := 0; i < a; i++ {
		for j := 0; j < b; j++ {
			s.Pixels[j*s.Stride+i] = 1
		}
	}
}

// rotate row y=A by B shifts all of the pixels in row A (0 is the top row) right
// by B pixels. Pixels that would fall off the right end appear at the left end of the row.
func (s *Screen) RotateRow(row, shift int) {
	curr := make([]uint8, s.Stride)
	r := row * s.Stride
	for i := 0; i < s.Stride; i++ {
		curr[i] = s.Pixels[r+i]
	}
	// now shift
	for i := 0; i < s.Stride; i++ {
		j := (i + shift) % s.Stride
		s.Pixels[r+j] = curr[i]
	}
}

// rotate column x=A by B shifts all of the pixels in column A (0 is the left column)
// down by B pixels. Pixels that would fall off the bottom appear at the top of the column.
func (s *Screen) RotateColumn(col, shift int) {
	curr := make([]uint8, s.Height)
	for i := 0; i < s.Height; i++ {
		curr[i] = s.Pixels[i*s.Stride+col]
	}
	// now shift
	for i := 0; i < s.Height; i++ {
		j := (i + shift) % s.Height
		s.Pixels[j*s.Stride+col] = curr[i]
	}
}

func (s *Screen) CountLitPixel() int {
	sum := 0
	for _, lit := range s.Pixels {
		sum += int(lit)
	}
	return sum
}

func (s *Screen) String() string {
	out := strings.Builder{}
	for i, lit := range s.Pixels {
		if i%s.Stride == 0 {
			out.WriteRune('\n')
		}
		if lit == 1 {
			out.WriteRune('#')
		} else {
			out.WriteRune(Space)
		}
	}
	return out.String()
}
