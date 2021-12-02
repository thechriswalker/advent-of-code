package aoc

import (
	"bufio"
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
