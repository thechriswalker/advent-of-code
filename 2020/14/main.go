package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2020, 14, solve1, solve2)
}

type bitmask struct {
	str   string
	ones  uint64
	zeros uint64
}

func (bm *bitmask) Apply(x uint64) uint64 {
	// the ones are a mask like `000000100` and we OR so
	// the 0s in the mask dont change anything and the 1s
	// always leave a 1.
	a := bm.ones | x

	// the zeros are a mask like `111111011` and we AND so
	// the 1s in the mask do nothing, but the 0s always leave 0.
	b := bm.zeros & a

	if false {
		fmt.Printf(`----\n
  MASK: %s
 INPUT: %036b (%d)
  ONES: %036b
MIDDLE: %036b (%d)
 ZEROS: %036b
OUTPUT: %036b (%d)
`, bm.str, x, x, bm.ones, a, a, bm.zeros, b, b)
	}
	return b
}

const mask36bits = 2<<35 - 1

func parseMask(line string) *bitmask {
	// mask = XXXXXXX10
	var bit uint64
	s := strings.Split(line, " = ")
	bm := &bitmask{
		str:   s[1],
		ones:  0,
		zeros: mask36bits,
	}
	for i, m := range s[1] {
		bit = 1 << (35 - i)
		switch m {
		case 'X':
			// do nothing
		case '1':
			// this goes in the ones.
			bm.ones = bm.ones | bit
		case '0':
			// zero the bit
			bm.zeros = bm.zeros & (mask36bits ^ bit)
		default:
			panic("unexpeceted character in line: " + s[1])
		}
	}
	return bm
}

// mem[X] = <binary> => X, <decimal>
func parseLine(line string) (uint64, uint64) {
	var x uint64
	var b uint64
	if _, err := fmt.Sscanf(line, "mem[%d] = %d", &x, &b); err != nil {
		panic(err.Error() + "[" + line + "]")
	}
	return x, b
}

// Implement Solution to Problem 1
// 13114787670125 is too high...
// 13105044880745
func solve1(input string) string {
	lines := strings.Split(input, "\n")
	mask := parseMask(lines[0])
	mem := map[uint64]uint64{}
	for _, line := range lines[1:] {
		if line == "" {
			break
		}
		if line[0:4] == "mask" {
			mask = parseMask(line)
			continue
		}
		x, b := parseLine(line)
		mem[x] = mask.Apply(b)
	}
	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	lines := strings.Split(input, "\n")
	mask := parseMask2(lines[0])
	fmt.Printf("MASK:\n%s", mask)
	mem := map[uint64]uint64{}
	for _, line := range lines[1:] {
		if line == "" {
			break
		}
		if line[0:4] == "mask" {
			mask = parseMask2(line)
			//fmt.Printf("MASK:\n%s", mask)
			continue
		}
		x, b := parseLine(line)
		//	fmt.Printf("LINE mem[%d] = %d\n", x, b)
		addrs := mask.Apply(x)

		for _, addr := range addrs {
			//	fmt.Printf("    ADDR: %036b (%d)\n", addr, addr)
			mem[addr] = b
		}
	}
	sum := uint64(0)
	for _, v := range mem {
		sum += v
	}
	return fmt.Sprintf("%d", sum)

}

type bitmask2 struct {
	str      string
	floating []uint64
	ones     uint64
	zeros    uint64
}

func (bm *bitmask2) String() string {
	builder := strings.Builder{}
	fmt.Fprintf(&builder, `INPUT: %s
 ONES: %036b
ZEROS: %036b
`, bm.str, bm.ones, bm.zeros)
	for _, f := range bm.floating {
		fmt.Fprintf(&builder, "FLOAT: %036b\n", f)
	}
	return builder.String()
}

// this time we return all the possible values
func (bm *bitmask2) Apply(x uint64) []uint64 {
	// the ones are a mask like `000000100` and we OR so
	// the 0s in the mask dont change anything and the 1s
	// always leave a 1.
	a := bm.ones | x

	// the zeros are a mask like `111111011` and we AND so
	// the 1s in the mask do nothing, but the 0s always leave 0.
	b := bm.zeros & a
	// fmt.Printf("UNMASKED: %036b (%d)\n", x, x)
	// fmt.Printf("  MASKED: %036b (%d)\n", b, b)

	// now we work out the possibilities.
	return addRecursive(b, make([]uint64, 0, 1<<len(bm.floating)), bm.floating)
}

func addRecursive(base uint64, stack []uint64, floating []uint64) []uint64 {
	if len(floating) == 0 {
		return append(stack, base)
	}
	if len(floating) == 1 {
		return append(stack, base, base|floating[0])
	}
	for i := 0; i < len(floating[1:]); i++ {
		// with the zero value is as is.
		stack = addRecursive(base, stack, floating[1:])
		// and with the 1 value
		stack = addRecursive(base|floating[0], stack, floating[1:])
	}
	return stack
}

func parseMask2(line string) *bitmask2 {
	// mask = XXXXXXX10
	var bit uint64
	s := strings.Split(line, " = ")
	bm := &bitmask2{
		str:      s[1],
		floating: []uint64{},
		ones:     0,
		zeros:    mask36bits,
	}
	//fmt.Printf("PREPARSE:\n%s", bm)
	for i, m := range s[1] {
		bit = 1 << (35 - i)
		switch m {
		case 'X':
			// make a zero and and to floating
			bm.zeros = bm.zeros & (mask36bits ^ bit)
			bm.floating = append(bm.floating, bit)
		case '1':
			// this goes in the ones.
			bm.ones = bm.ones | bit
		case '0':
			// don't change it!
		default:
			panic("unexpeceted character in line: " + s[1])
		}
	}
	return bm
}
