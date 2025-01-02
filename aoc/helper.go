package aoc

import (
	"bufio"
	"fmt"
	"io"
	"maps"
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
		n, _ := strconv.Atoi(strings.TrimSpace(sn))
		nn = append(nn, n)
	}
	return nn
}

func ToUint8Slice(input string, sep rune) []uint8 {
	s := strings.FieldsFunc(strings.TrimSpace(input), func(r rune) bool { return r == sep })
	nn := make([]uint8, 0, len(s))
	for _, sn := range s {
		n, _ := strconv.Atoi(strings.TrimSpace(sn))
		nn = append(nn, uint8(n))
	}
	return nn
}
func ToUint64Slice(input string, sep rune) []uint64 {
	s := strings.FieldsFunc(strings.TrimSpace(input), func(r rune) bool { return r == sep })
	nn := make([]uint64, 0, len(s))
	for _, sn := range s {
		n, _ := strconv.Atoi(strings.TrimSpace(sn))
		nn = append(nn, uint64(n))
	}
	return nn
}

// ByteGrid represents a grid of data represented by a single byte.
// out of bounds access returns the "unknown" byte and a bool oob=true
// the grid is laid out as left-to-right is increasing x, up-to-down is increasing y
type ByteGrid interface {
	At(x, y int) (b byte, oob bool)
	Atv(v V2) (b byte, oob bool)
	Width() int
	Height() int
	Bounds() (x1, y1, x2, y2 int)
	Set(x, y int, z byte) bool
	Setv(v V2, z byte) bool
	Clone() ByteGrid
}

func SprintByteGrid(g ByteGrid, hilite map[byte]string) string {
	sb := strings.Builder{}
	FprintByteGrid(&sb, g, hilite)
	return sb.String()
}

func SprintByteGridC(g ByteGrid, hilite map[byte]Color) string {
	sb := strings.Builder{}
	FprintByteGridC(&sb, g, hilite)
	return sb.String()
}

func PrintByteGrid(g ByteGrid, hilite map[byte]string) {
	FprintByteGrid(os.Stdout, g, hilite)
}

func PrintByteGridC(g ByteGrid, hilite map[byte]Color) {
	FprintByteGridC(os.Stdout, g, hilite)
}

func OOB(g ByteGrid, x, y int) (oob bool) {
	_, oob = g.At(x, y)
	return
}

type Color [2]uint8

var (
	NoColor = Color{0, 0}

	Black   = Color{0, 30}
	Red     = Color{0, 31}
	Green   = Color{0, 32}
	Yellow  = Color{0, 33}
	Blue    = Color{0, 34}
	Magenta = Color{0, 35}
	Cyan    = Color{0, 36}
	White   = Color{0, 37}

	BoldBlack   = Color{1, 30}
	BoldRed     = Color{1, 31}
	BoldGreen   = Color{1, 32}
	BoldYellow  = Color{1, 33}
	BoldBlue    = Color{1, 34}
	BoldMagenta = Color{1, 35}
	BoldCyan    = Color{1, 36}
	BoldWhite   = Color{1, 37}
)

func SprintByteGridFunc(g ByteGrid, hiliteFn func(x, y int, b byte) Color) string {
	sb := strings.Builder{}
	FprintByteGridFunc(&sb, g, hiliteFn)
	return sb.String()
}
func PrintByteGridFunc(g ByteGrid, hiliteFn func(x, y int, b byte) Color) {
	FprintByteGridFunc(os.Stdout, g, hiliteFn)
}

func FprintByteGridFunc(w io.Writer, g ByteGrid, hiliteFn func(x, y int, b byte) Color) {
	if hiliteFn == nil {
		hiliteFn = func(x, y int, b byte) Color { return NoColor }
	}
	x1, y1, x2, y2 := g.Bounds()
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			// shouldn't be oob
			b, _ := g.At(x, y)
			if c := hiliteFn(x, y, b); c != NoColor {
				w.Write([]byte("\x1b["))
				if c[0] == 1 {
					w.Write([]byte("1;"))
				}
				fmt.Fprintf(w, "%dm%c\x1b[0m", c[1], b)
			} else {
				w.Write([]byte{b})
			}
		}
		w.Write([]byte{'\n'})
	}
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

func FprintByteGridC(w io.Writer, g ByteGrid, hilite map[byte]Color) {
	x1, y1, x2, y2 := g.Bounds()
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			// shouldn't be oob
			b, _ := g.At(x, y)
			if c, ok := hilite[b]; ok && c != NoColor {
				w.Write([]byte("\x1b["))
				if c[0] == 1 {
					w.Write([]byte("1;"))
				}
				fmt.Fprintf(w, "%dm%c\x1b[0m", c[1], b)
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

func IterateByteGridv(g ByteGrid, f func(v V2, b byte)) {
	x1, y1, x2, y2 := g.Bounds()
	for y := y1; y <= y2; y++ {
		for x := x1; x <= x2; x++ {
			b, _ := g.At(x, y)
			f(V2{x, y}, b)
		}
	}
}

type FixedByteGrid struct {
	w, h    int
	data    []byte
	unknown byte
}

func (g *FixedByteGrid) Clone() ByteGrid {
	data := make([]byte, len(g.data))
	copy(data, g.data)
	return &FixedByteGrid{
		w:       g.w,
		h:       g.h,
		unknown: g.unknown,
		data:    data,
	}
}

func (g *FixedByteGrid) Value() string {
	return string(g.data)
}
func (g *FixedByteGrid) Atv(v V2) (byte, bool) {
	return g.At(v.X, v.Y)
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
func (g *FixedByteGrid) Setv(v V2, b byte) bool {
	return g.Set(v.X, v.Y, b)
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

func (g *SparseByteGrid) Clone() ByteGrid {
	return &SparseByteGrid{
		xmin:    g.xmin,
		xmax:    g.xmax,
		ymin:    g.ymin,
		ymax:    g.ymax,
		unknown: g.unknown,
		data:    maps.Clone(g.data),
	}
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
func (g *SparseByteGrid) Atv(v V2) (byte, bool) {
	return g.At(v.X, v.Y)
}
func (g *SparseByteGrid) Width() int  { return g.ymax - g.ymin }
func (g *SparseByteGrid) Height() int { return g.xmax - g.xmin }
func (g *SparseByteGrid) Bounds() (x1, y1, x2, y2 int) {
	return g.xmin, g.ymin, g.xmax, g.ymax
}
func (g *SparseByteGrid) Set(x, y int, b byte) bool {
	g.data[[2]int{x, y}] = b
	if len(g.data) == 1 {
		// first insert, set bounds around this point
		g.xmin, g.ymin = x, y
		g.xmax, g.ymax = x, y
	} else {
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
	}
	return true
}
func (g *SparseByteGrid) Setv(v V2, b byte) bool {
	return g.Set(v.X, v.Y, b)
}

func GridIndex(x, y, stride, height int) int {
	if x < 0 || x >= stride || y < 0 || y >= height {
		return -1
	}
	return y*stride + x
}

func GridCoords(idx, stride int) (x, y int) {
	x = idx % stride
	y = idx / stride
	return
}

type V2 struct {
	X, Y int
}

func (v V2) Add(o V2) V2 {
	return V2{v.X + o.X, v.Y + o.Y}
}

func Vec2(x, y int) V2 {
	return V2{x, y}
}

type V3 struct {
	X, Y, Z int
}

func Vec3(x, y, z int) V3 {
	return V3{x, y, z}
}

var (
	North = V2{0, -1}
	South = V2{0, 1}
	East  = V2{1, 0}
	West  = V2{-1, 0}
)

func GCD[T ~int | ~int8 | ~int16 | ~int32 | ~int64](a, b T) T {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM[T ~int | ~int8 | ~int16 | ~int32 | ~int64](a, b T) T {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	return a * b / GCD(a, b)
}
