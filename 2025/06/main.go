package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2025, 6, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	sum := int64(0)
	var rows [][]int64
	var ops []string

	aoc.MapLines(input, func(line string) error {
		fields := strings.Fields(line)
		if fields[0] == "+" || fields[0] == "*" {
			ops = fields
		} else {
			if len(rows) == 0 {
				rows = make([][]int64, len(fields))
			}
			for i, s := range fields {
				n, _ := strconv.ParseInt(s, 10, 64)
				rows[i] = append(rows[i], n)
			}
		}
		return nil
	})

	for i := range ops {
		if ops[i] == "+" {
			sum += addRow(rows[i])
		} else {
			sum += multiplyRow(rows[i])
		}
	}
	return fmt.Sprint(sum)
}

func addRow(row []int64) int64 {
	sum := row[0]
	for _, n := range row[1:] {
		sum += n
	}
	return sum
}

func multiplyRow(row []int64) int64 {
	sum := row[0]
	for _, n := range row[1:] {
		sum *= n
	}
	return sum
}

// Implement Solution to Problem 2
func solve2(input string) string {

	// here its going to be easier to transpose the data first
	input = transpose(input)
	//aoc.Debug(input)
	// now each problem is separated by a blank line, and the first line of a problem contains the first number and the operator
	problems := strings.Split(input, "\n\n")

	sum := int64(0)
	for _, p := range problems {
		sum += solveProblem(p)
	}

	return fmt.Sprint(sum)
}

func transpose(input string) string {
	rows := bytes.Split([]byte(input), []byte{'\n'})
	// we know each line is the same length....
	l := len(rows[0])
	cols := make([][]byte, l)
	for i := 0; i < l; i++ {
		for j := 0; j < len(rows); j++ {
			if len(rows[j]) == 0 {
				// empy line at the end screwed this up...
				continue
			}
			cols[i] = append(cols[i], rows[j][i])
		}
	}
	for i := range cols {
		cols[i] = bytes.TrimSpace(cols[i])
	}
	return string(bytes.Join(cols, []byte("\n")))
}

func solveProblem(p string) int64 {
	lines := strings.Split(p, "\n")
	var op func(a, b int64) int64
	if lines[0][len(lines[0])-1] == '+' {
		op = func(a, b int64) int64 { return a + b }
	} else {
		op = func(a, b int64) int64 { return a * b }
	}
	sum, _ := strconv.ParseInt(strings.TrimSpace(lines[0][:len(lines[0])-1]), 10, 64)
	for _, s := range lines[1:] {
		n, _ := strconv.ParseInt(s, 10, 64)
		sum = op(sum, n)
	}
	return sum
}
