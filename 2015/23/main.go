package main

import (
	"fmt"
	"strconv"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 23, solve1, solve2)
}

// ooh a computer. I like these.
type Computer struct {
	registers map[byte]uint
	counter   int
	program   []string
}

func signedInt(str string) int {
	s := 1
	if str[0] == '-' {
		s = -1
	}
	n, _ := strconv.Atoi(str[1:])
	return s * n
}

// really simple
func (pc *Computer) Tick() bool {
	if pc.counter >= len(pc.program) {
		return false
	}
	ins := pc.program[pc.counter]
	r := ins[4]
	//	fmt.Println("Handling instruction", pc.counter, " => ", ins, "Registers: a=", pc.registers['a'], "b=", pc.registers['b'])
	switch ins[:3] {
	case "hlf":
		pc.registers[r] /= 2
		pc.counter++
	case "tpl":
		pc.registers[r] *= 3
		pc.counter++
	case "inc":
		pc.registers[r] += 1
		pc.counter++
	case "jie":
		if pc.registers[r]%2 == 0 {
			pc.counter += signedInt(ins[6:])
		} else {
			pc.counter++
		}
	case "jio": // NB jump if ONE
		if pc.registers[r] == 1 {
			pc.counter += signedInt(ins[6:])
		} else {
			pc.counter++
		}
	case "jmp":
		j, _ := strconv.Atoi(ins[4:])
		pc.counter += j
	default:
		fmt.Println("UNEXPECTED INSTRUCTION:", ins)
		return false
	}
	return true
}

// Implement Solution to Problem 1
func solve1(input string) string {
	pc := &Computer{
		registers: map[byte]uint{},
		program:   []string{},
	}
	aoc.MapLines(input, func(s string) error {
		pc.program = append(pc.program, s)
		return nil
	})

	for pc.Tick() {
	}

	return fmt.Sprint(pc.registers['b'])
}

// Implement Solution to Problem 2
func solve2(input string) string {
	pc := &Computer{
		registers: map[byte]uint{'a': 1},
		program:   []string{},
	}
	aoc.MapLines(input, func(s string) error {
		pc.program = append(pc.program, s)
		return nil
	})

	for pc.Tick() {
	}

	return fmt.Sprint(pc.registers['b'])
}
