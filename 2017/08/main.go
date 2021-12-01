package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 8, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	instructions := parseInput(input)
	cpu := &CPU{}
	cpu.Run(instructions...)
	max := 0
	for _, v := range cpu.registers {
		if v > max {
			max = v
		}
	}
	return fmt.Sprint(max)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	instructions := parseInput(input)
	cpu := &CPU{}
	cpu.Run(instructions...)
	return fmt.Sprint(cpu.max)
}

type instruction struct {
	targetReg string
	increment int
	testReg   string
	testCond  string
	testVal   int
}

type CPU struct {
	registers map[string]int
	max       int
}

func (c *CPU) Run(instructions ...instruction) {
	c.registers = make(map[string]int)
	c.max = 0
	for _, ins := range instructions {
		if c.condition(ins.testReg, ins.testCond, ins.testVal) {
			c.registers[ins.targetReg] += ins.increment
			if c.registers[ins.targetReg] > c.max {
				c.max = c.registers[ins.targetReg]
			}
		}
	}
}

func (c *CPU) condition(reg, cond string, test int) bool {
	v := c.registers[reg]
	switch cond {
	case ">":
		return v > test
	case "<":
		return v < test
	case "<=":
		return v <= test
	case ">=":
		return v >= test
	case "==":
		return v == test
	case "!=":
		return v != test
	}
	panic("unknown condition")
}

func parseInput(input string) []instruction {
	scanner := bufio.NewScanner(strings.NewReader(input))
	ins := []instruction{}
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		inc := 1
		if fields[1] == "dec" {
			inc = -1
		}
		v, _ := strconv.Atoi(fields[2])
		tv, _ := strconv.Atoi(fields[6])
		instruction := instruction{
			targetReg: fields[0],
			increment: inc * v,
			testReg:   fields[4],
			testCond:  fields[5],
			testVal:   tv,
		}
		ins = append(ins, instruction)
	}
	return ins
}
