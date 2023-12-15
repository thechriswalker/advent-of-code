package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 15, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {

	sum := 0

	for _, s := range strings.Split(strings.TrimSpace(input), ",") {
		sum += h(s)
	}

	return fmt.Sprint(sum)
}

func h(s string) int {
	v := uint8(0)
	for i := range s {
		v += uint8(s[i])
		v *= 17
	}
	return int(v)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	ins := strings.Split(strings.TrimSpace(input), ",")

	hm := make([][]lens, 256)

	var label string
	var focal int
	var box int
	for _, x := range ins {
		if strings.HasSuffix(x, "-") {
			label = strings.TrimSuffix(x, "-")
			box = h(label)
			hm[box] = op_remove(hm[box], label)
		} else {
			s, n, _ := strings.Cut(x, "=")
			label = s
			box = h(label)
			focal, _ = strconv.Atoi(n)
			// insert
			hm[box] = op_insert(hm[box], label, focal)
		}
	}

	power := 0

	for i, b := range hm {
		for j, l := range b {
			power += (i + 1) * (j + 1) * l.focal
		}
	}

	return fmt.Sprint(power)
}

type lens struct {
	label string
	focal int
}

func op_insert(box []lens, lb string, fl int) []lens {
	newLens := lens{label: lb, focal: fl}
	// find label and replace. otherwise append

	for i := range box {
		if box[i].label == lb {
			// replace
			box[i] = newLens
			return box
		}
	}
	// append
	return append(box, newLens)
}

func op_remove(box []lens, lb string) []lens {
	// find label and remove if present
	for i := range box {
		if box[i].label == lb {
			copy(box[i:], box[i+1:])
			return box[:len(box)-1]
		}
	}
	return box
}
