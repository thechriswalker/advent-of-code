package main

import (
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2016, 12, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	ops := parseInput(input)
	//fmt.Println(ops)
	reg := runProgram(ops)
	return fmt.Sprintf("%d", reg[0])
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// add the instruction to initialize c to 1
	ops := parseInput("cpy 1 c\n" + input)
	reg := runProgram(ops)
	return fmt.Sprintf("%d", reg[0])
}

// fixed width op makes life easier
const OpLength = 3
const (
	CPV = iota
	CPR
	INC
	JNR
	JNV

	ADDR = 0
	ARG1 = 1
	ARG2 = 2
)

func parseInput(input string) []int {
	program := []int{}
	lines := strings.Split(input, "\n")
	var op, arg1, arg2 int
	for _, line := range lines {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		switch line[0] {
		case 'c':
			// cpv or cpr
			switch line[4] {
			case 'a', 'b', 'c', 'd':
				// register
				op = CPR
				arg1 = int(line[4] - 'a')
				arg2 = int(line[6] - 'a')
			default:
				//value
				op = CPV
				fmt.Sscanf(line[3:], "%d", &arg1)
				arg2 = int(line[len(line)-1] - 'a')
			}
		case 'i':
			// inc
			op = INC
			arg1 = int(line[4] - 'a')
			arg2 = 1
		case 'd':
			// dec
			op = INC
			arg1 = int(line[4] - 'a')
			arg2 = -1
		case 'j':
			// jnz
			switch line[4] {
			case 'a', 'b', 'c', 'd':
				// register
				op = JNR
				arg1 = int(line[4] - 'a')
				fmt.Sscanf(line[5:], "%d", &arg2)
			default:
				//value
				op = JNV
				fmt.Sscanf(line[3:], "%d %d", &arg1, &arg2)
			}
		}
		//	fmt.Println("Line:", line, "Op", op, arg1, arg2)
		program = append(program, op, arg1, arg2)
	}

	return program
}

func runProgram(prog []int) [4]int {
	//program counter
	pc := 0
	pl := len(prog) / 3
	// registers
	reg := [4]int{0, 0, 0, 0}
	loop := 0
	var op, arg1, arg2 int
	//fmt.Println("Prog:", prog, "\nPL:", pl)
	for pc < pl {
		addr := pc * 3
		loop++
		// if loop > 1000000 {
		// 	panic("not halting")
		// }
		op = prog[addr+ADDR]
		arg1 = prog[addr+ARG1]
		arg2 = prog[addr+ARG2]
		//	fmt.Printf("Step %06d | PC %03d | Reg A:%03d B:%03d C:%03d D:%03d | Op %d | Args: %d %d\n", loop, pc, reg[0], reg[1], reg[2], reg[3], op, arg1, arg2)
		switch op {
		case CPV:
			reg[arg2] = arg1
		case CPR:
			reg[arg2] = reg[arg1]
		case INC:
			reg[arg1] = reg[arg1] + arg2
		case JNV:
			if arg1 != 0 {
				pc += arg2 - 1
			}
		case JNR:
			//fmt.Printf("JNZ %c %d (RegValue: %d)\n", rune(arg1+'a'), arg2, reg[arg1])
			if reg[arg1] != 0 {
				pc += arg2 - 1
			}
		}
		pc++
	}
	return reg
}
