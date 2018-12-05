package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 16, solve1, solve2)
}

var DiskSize = 272

// Implement Solution to Problem 1
func solve1(input string) string {
	data := createData(input)
	data = data.Fill(DiskSize)
	return fmt.Sprintf("%s", data.Checksum())
}

// Implement Solution to Problem 2
func solve2(input string) string {
	DiskSize = 35651584
	return solve1(input)
}

func createData(input string) Data {
	input = strings.TrimSpace(input)
	data := make(Data, len(input))
	for i, c := range input {
		switch c {
		case '1':
			data[i] = 1
		case '0':
			data[i] = 0
		default:
			panic("expecting 0 or 1!")
		}
	}
	return data
}

type Data []byte

func (d Data) String() string {
	s := &strings.Builder{}
	for _, b := range d {
		fmt.Fprintf(s, "%d", b)
	}
	return s.String()
}

func (d Data) Expand() Data {
	l := len(d)
	// this adds the zero and makes it the correct size
	d = append(d, 0)
	d = append(d, d[0:l]...)
	// now reverse the b into len(d) +
	end := len(d) - 1
	for i := 0; i < l; i++ {
		// first i goes in last place (flipped)
		if d[i] == 1 {
			d[end-i] = 0
		} else {
			d[end-i] = 1
		}
	}
	return d
}

func (d Data) Fill(size int) Data {
	for len(d) < size {
		d = d.Expand()
	}
	return d[0:size]
}

func (d Data) Checksum() Data {
	c := make(Data, 0, len(d)/2)
	for i := 1; i < len(d); i += 2 {
		if d[i] == d[i-1] {
			c = append(c, 1)
		} else {
			c = append(c, 0)
		}
	}
	if len(c)%2 == 0 {
		// even, repeat
		return c.Checksum()
	}
	return c
}
