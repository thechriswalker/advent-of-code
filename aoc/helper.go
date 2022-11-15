package aoc

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

func MapLines(input string, fn func(line string) error) error {
	sc := bufio.NewScanner(strings.NewReader(input))
	var err error
	for sc.Scan() {
		if err = fn(sc.Text()); err != nil {
			return err
		}
	}
	return sc.Err()
}

func ToIntSlice(input string, sep rune) []int {
	s := strings.FieldsFunc(strings.TrimSpace(input), func(r rune) bool { return r == sep })
	nn := make([]int, 0, len(s))
	for _, sn := range s {
		n, _ := strconv.Atoi(sn)
		nn = append(nn, n)
	}
	return nn
}

type ByteGrid interface {
	At(x, y int) (b byte, oob bool)
	Width() int
	Height() int
	Bounds() (x1, y1, x2, y2 int)
	Set(x, y int, z byte) bool
}

func SprintByteGrid(g ByteGrid, hilite map[byte]string) string {
	sb := strings.Builder{}
	FprintByteGrid(&sb, g, hilite)
	return sb.String()
}
func PrintByteGrid(g ByteGrid, hilite map[byte]string) {
	FprintByteGrid(os.Stdout, g, hilite)
}

func FprintByteGrid(w io.Writer, g ByteGrid, hilite map[byte]string) {
	x1, y1, x2, y2 := g.Bounds()
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			// shouldn't be oob
			b, _ := g.At(x, y)
			if h, ok := hilite[b]; ok {
				w.Write([]byte(`\x1b[`))
				w.Write([]byte(h))
				w.Write([]byte{'m', b})
				w.Write([]byte(`\x1b[0m`))
			} else {
				w.Write([]byte{b})
			}
		}
		w.Write([]byte{'\n'})
	}
}

func IterateByteGrid(g ByteGrid, f func(x, y int, b byte)) {
	x1, y1, x2, y2 := g.Bounds()
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			b, _ := g.At(x, y)
			f(x, y, b)
		}
	}
}

type FixedByteGrid struct {
	w, h    int
	data    []byte
	unknown byte
}

func (g *FixedByteGrid) At(x, y int) (byte, bool) {
	idx := GridIndex(x, y, g.w, g.h)
	if idx == -1 {
		return g.unknown, true
	}
	return g.data[idx], false
}
func (g *FixedByteGrid) Width() int  { return g.w }
func (g *FixedByteGrid) Height() int { return g.h }
func (g *FixedByteGrid) Bounds() (x1, y1, x2, y2 int) {
	return 0, 0, g.w - 1, g.h - 1
}
func (g *FixedByteGrid) Set(x, y int, b byte) bool {
	idx := GridIndex(x, y, g.w, g.h)
	if idx == -1 {
		return false
	}
	g.data[idx] = b
	return true
}

func CreateFixedByteGridFromString(input string, unknown byte) *FixedByteGrid {
	g := &FixedByteGrid{
		data:    make([]byte, 0, len(input)),
		unknown: unknown,
	}
	MapLines(input, func(line string) error {
		if g.w == 0 {
			g.w = len(line)
		}
		g.h++
		g.data = append(g.data, []byte(line)...)
		return nil
	})
	return g
}

type SparseByteGrid struct {
	xmin, xmax int
	ymin, ymax int
	data       map[[2]int]byte
	unknown    byte
}

func NewSparseByteGrid(unknown byte) *SparseByteGrid {
	return &SparseByteGrid{
		data:    map[[2]int]byte{},
		unknown: unknown,
	}
}

func (g *SparseByteGrid) Populate(other ByteGrid) {
	IterateByteGrid(other, func(x, y int, b byte) {
		g.Set(x, y, b)
	})
}

func (g *SparseByteGrid) At(x, y int) (byte, bool) {
	if x < g.xmin || x > g.xmax || y < g.ymin || y > g.ymax {
		return g.unknown, true
	}
	b, ok := g.data[[2]int{x, y}]
	if !ok {
		return g.unknown, false
	}
	return b, false
}
func (g *SparseByteGrid) Width() int  { return g.ymax - g.ymin }
func (g *SparseByteGrid) Height() int { return g.xmax - g.xmin }
func (g *SparseByteGrid) Bounds() (x1, y1, x2, y2 int) {
	return g.xmin, g.ymin, g.xmax, g.ymax
}
func (g *SparseByteGrid) Set(x, y int, b byte) bool {
	g.data[[2]int{x, y}] = b
	if x < g.xmin {
		g.xmin = x
	}
	if y < g.ymin {
		g.ymin = y
	}
	if x > g.xmax {
		g.xmax = x
	}
	if y > g.ymax {
		g.ymax = y
	}
	return true
}

func GridIndex(x, y, stride, height int) int {
	if x < 0 || x >= stride || y < 0 || y >= height {
		return -1
	}
	return y*stride + x
}

func GridCoords(idx, stride int) (x, y int) {
	x = idx / stride
	y = idx % stride
	return
}
