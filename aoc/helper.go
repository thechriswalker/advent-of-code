package aoc

import (
	"bufio"
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
