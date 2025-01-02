package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 9, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// the ints are "fileIDs", -1 is empty space.
	data := make([]int, 0, len(input))
	fileID := 0
	for i, r := range input {
		n := int(r) - '0'
		if i%2 == 0 {
			// a file. length n
			for j := 0; j < n; j++ {
				data = append(data, fileID)
			}
			fileID++
		} else {
			// empty space
			for j := 0; j < n; j++ {
				data = append(data, -1)
			}
		}
	}

	// to compact this, we take the first non-empty piece from the right and move it to the first non-empty space on the left, until
	// our cursors cross.
	head, tail := 0, len(data)-1

	for head <= tail {
		// find the next gap.
		if data[head] != -1 {
			head++
			continue
		}
		// take a piece from the end to fill this gap.
		if data[tail] == -1 {
			tail--
			continue
		}
		// head points to gap, tail to piece
		data[head], data[tail] = data[tail], data[head]
		// move our pointers
		head++
		tail--
	}

	// now checksum.
	cs := 0
	for i := 0; i < len(data); i++ {
		if data[i] == -1 {
			break //
		}
		cs += data[i] * i
	}

	return fmt.Sprint(cs)
}

// Implement Solution to Problem 2
func solve2(input string) string {

	// we need MORE info than last time, as it will be useful to know what all the files are.
	// especially as we have to reverse traverse the data to find a file that will fit.
	// so we store the gaps in order, and the files in order (by ID)
	gaps := []aoc.V2{}  // X is start, Y is length, in order from left to right.
	files := []aoc.V2{} // index is "fileID", X is start, Y is length

	// the ints are "fileIDs", -1 is empty space.
	data := make([]int, 0, len(input))
	fileID := 0
	for i, r := range input {
		n := int(r) - '0'
		size := aoc.V2{len(data), n}
		if i%2 == 0 {
			// a file. length n
			files = append(files, size)
			for j := 0; j < n; j++ {
				data = append(data, fileID)
			}
			fileID++
		} else {
			// empty space
			gaps = append(gaps, size)
			for j := 0; j < n; j++ {
				data = append(data, -1)
			}
		}
	}

	printData(data)

	// now each time we find the first file big enough to fit the first gap.
	// this means we might need to splice gaps.
	// that is ok, because they will still be "in order" and we just use some of it, we remove the file after it is used.

	// bascially our loop now need to check if we were able to move anything each iteration and it not, we are done.
	for {
		moved := false
		// for each file, attempt to find the leftmost gap to fit it.
		// if we do, continue and iterate again.
	fileloop:
		for f := len(files) - 1; f >= 0; f-- {
			if files[f].Y == 0 { // we will set this to indicate a file that has been used.
				continue
			}
			// need a gap of at least f.Y
			for g := 0; g < len(gaps); g++ {
				// also we never move a file to the right.
				if gaps[g].X > files[f].X {
					continue
				}

				if gaps[g].Y >= files[f].Y {
					gapStart := gaps[g].X
					fileStart := files[f].X
					// we can fit this file.
					for i := 0; i < files[f].Y; i++ {
						// add file to gap.
						data[gapStart+i] = f
						data[fileStart+i] = -1
					}
					// close the used bit of the gap
					gaps[g].X += files[f].Y
					gaps[g].Y -= files[f].Y
					files[f].Y = 0 // mark as moved
					moved = true
					break fileloop
				}
			}
		}

		if !moved {
			break
		}
	}

	// now checksum, with no shortcut.
	cs := 0
	for i := 0; i < len(data); i++ {
		if data[i] != -1 {
			cs += data[i] * i
		}
	}
	printData(data)

	// 8705230292234 -- too high, what is the problem - test case passes, so we need to find a failing case...

	return fmt.Sprint(cs)
}

func printData(data []int) {
	// for i := 0; i < len(data); i++ {
	// 	if data[i] == -1 {
	// 		fmt.Print(".")
	// 	} else {
	// 		fmt.Print(data[i])
	// 	}
	// }
	// fmt.Println()
}
