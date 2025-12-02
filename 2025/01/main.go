package main

import (
	"fmt"
	"strconv"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 1, solve1, solve2)
}

type Rot struct {
	left bool
	n    int
}

// Implement Solution to Problem 1
func solve1(input string) string {
	seq := []Rot{}

	aoc.MapLines(input, func(line string) error {
		r := Rot{left: line[0] == 'L'}
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			return err
		}
		r.n = n
		seq = append(seq, r)
		return nil
	})

	count0 := 0
	safe := 50
	for _, r := range seq {
		if r.left {
			safe -= r.n
		} else {
			safe += r.n
		}
		safe = safe % 100
		if safe == 0 {
			count0++
		}
	}

	return fmt.Sprint(count0)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	seq := []Rot{}

	aoc.MapLines(input, func(line string) error {
		r := Rot{left: line[0] == 'L'}
		n, err := strconv.Atoi(line[1:])
		if err != nil {
			return err
		}
		r.n = n
		seq = append(seq, r)
		return nil
	})

	count0 := 0
	safe := 50
	//	fmt.Println("-----")
	for _, r := range seq {
		//	before := safe
		ticks := 0
		if r.left {
			i := r.n
			for i >= 100 {
				ticks++
				i -= 100
			}
			if safe == 0 {
				safe = 100
			}
			safe -= i
			if safe < 0 {
				ticks++
				safe += 100
			}
		} else {
			i := r.n
			for i >= 100 {
				ticks++
				i -= 100
			}
			safe += i
			if safe > 100 {
				ticks++
				safe -= 100
			}
		}
		safe = safe % 100
		if safe == 0 {
			ticks++
		}
		count0 += ticks
		//	fmt.Println("before", before, "rot", r, "after", safe, "ticks", ticks, "total", count0)
	}

	//6795 too high
	return fmt.Sprint(count0)
}
