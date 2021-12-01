package main

import (
	"fmt"
	"io"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 9, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	rd := strings.NewReader(input)
	// final length
	length := int64(0)
	var c rune
	var err error
	// length and repeat variables for decompression sequences.
	var l, r int64
	for {
		c, _, err = rd.ReadRune()
		if err != nil {
			break
		}
		switch c {
		case '(':
			// start of decompression sequence,
			// so length is the length of the sequence * repeats,
			// and we seek forward just the length.
			_, err := fmt.Fscanf(rd, "%dx%d)", &l, &r)
			if err != nil {
				break
			}
			length += l * r
			rd.Seek(int64(l), io.SeekCurrent)
		case ' ', '\n', '\r', '\t', '\v':
			// ignore whitespace
		default:
			// increment length
			length++
		}
	}
	return fmt.Sprintf("%d", length)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return fmt.Sprintf("%d", ChunkLength(input))
}

// version 2 recurses
func ChunkLength(chunk string) int64 {
	rd := strings.NewReader(chunk)
	// final length
	length := int64(0)
	var c rune
	var err error
	// length and repeat variables for decompression sequences.
	var l, r int64
	for {
		c, _, err = rd.ReadRune()
		if err != nil {
			break
		}
		switch c {
		case '(':
			// start of decompression sequence,
			// so length is the length of the decompressed sub-sequence * repeats,
			// and we seek forward just the length.
			_, err := fmt.Fscanf(rd, "%dx%d)", &l, &r)
			if err != nil {
				break
			}
			pos, _ := rd.Seek(0, io.SeekCurrent)
			subchunk := chunk[pos : pos+l]
			length += r * ChunkLength(subchunk)
			rd.Seek(int64(l), io.SeekCurrent)
		case ' ', '\n', '\r', '\t', '\v':
			// ignore whitespace
		default:
			// increment length
			length++
		}
	}
	return length
}
