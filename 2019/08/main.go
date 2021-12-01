package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 8, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	s := parseImage(input, 25, 6)
	c := s.LayerCount()
	var layer []uint8
	minZeros := 1000000
	for i := 0; i < c; i++ {
		l := s.GetLayer(i)
		// count the number of zeros in the layer.
		z := countDigitsInLayer(l, 0)
		if z < minZeros {
			layer = l
			minZeros = z
		}
	}

	return fmt.Sprintf("%d", countDigitsInLayer(layer, 1)*countDigitsInLayer(layer, 2))
}

func countDigitsInLayer(l []uint8, d uint8) int {
	c := 0
	for _, v := range l {
		if v == d {
			c++
		}
	}
	return c
}

// Implement Solution to Problem 2
func solve2(input string) string {
	s := parseImage(input, 25, 6)
	return "\n" + s.Flatten()
}

type SpaceImage struct {
	Width  int
	Height int
	Data   []uint8 // the raw data one row of one layer at a time.
}

func (s *SpaceImage) LayerCount() int {
	size := s.Width * s.Height
	return len(s.Data) / size
}

func (s *SpaceImage) GetLayer(n int) []uint8 {
	size := s.Width * s.Height
	offset := size * n
	return s.Data[offset : offset+size]
}

func parseImage(input string, w, h int) *SpaceImage {
	t := strings.TrimSpace(input)
	s := &SpaceImage{
		Width:  w,
		Height: h,
		Data:   make([]uint8, len(t)),
	}
	for i, r := range t {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			s.Data[i] = uint8(r) - '0'
		default:
			panic("bad image!")
		}
	}
	return s
}

const (
	white = '#'
	black = ' '
	trans = '.'
)

func (s *SpaceImage) GetPixel(p int) byte {
	size := s.Width * s.Height
	if p > size || p < 0 {
		panic("Pixel out of range!")
	}
	// iterate through the layers until we find a solid pixel
	l := 0
	for {
		v := size*l + p
		if v > len(s.Data) {
			return trans
		}
		switch s.Data[v] {
		case 0:
			return black
		case 1:
			return white
		default:
			l++
		}
	}
}

func (s *SpaceImage) Flatten() string {
	// add space for a \n on every row
	var data strings.Builder
	for i := 0; i < s.Width*s.Height; i++ {
		data.WriteByte(s.GetPixel(i))
		if (i+1)%s.Width == 0 {
			data.WriteByte('\n')
		}
	}
	return data.String()
}
