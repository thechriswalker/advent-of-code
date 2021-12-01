package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 25, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	ops := parseInput(input)
	//fmt.Println(ops)
	// find the value that produces the stable alternating signal.
	var attempt int
	for !runProgram(ops, attempt, 0, 0, 0) {
		attempt++
	}
	return fmt.Sprintf("%d", attempt)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return "<unsolved>"
}

// fixed width op makes life easier
// op argc value_or_reg arg1 value_or_reg arg2
const OpLength = 6
const (
	CPY_VV = iota // copy value value (invalid)
	CPY_VR        // copy value register
	CPY_RR        // copy register register
	CPY_RV        // copy register value (invalid)
	INC_V         // inc value (noop)
	INC_R         // inc register
	DEC_V         // dec value (noop)
	DEC_R         // dec regsiter
	JNZ_VV        // jump-not-zero value value
	JNZ_RV        // jump-not-zero register value
	JNZ_VR        // jump-not-zero value register
	JNZ_RR        // jump-not-zero register register
	TGL_V         // toggle value
	TGL_R         // toggle register
	OUT_V         // output value
	OUT_R         // output register

	ADDR = 0
	ARG1 = 1
	ARG2 = 2
)

type Op int

func (o Op) String() string {
	switch int(o) {
	case CPY_VV:
		return "CPY_VV"
	case CPY_VR:
		return "CPY_VR"
	case CPY_RR:
		return "CPY_RR"
	case CPY_RV:
		return "CPY_RV"
	case INC_V:
		return "INC_V"
	case INC_R:
		return "INC_R"
	case DEC_V:
		return "DEC_V"
	case DEC_R:
		return "DEC_R"
	case JNZ_VV:
		return "JNZ_VV"
	case JNZ_RV:
		return "JNZ_RR"
	case JNZ_VR:
		return "JNZ_VR"
	case JNZ_RR:
		return "JNZ_RR"
	case TGL_V:
		return "TGL_V"
	case TGL_R:
		return "TGL_R"
	case OUT_V:
		return "OUT_V"
	case OUT_R:
		return "OUT_R"
	default:
		return fmt.Sprintf("?%d", int(o))
	}
}

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
			// cpy type
			switch line[4] {
			case 'a', 'b', 'c', 'd':
				// register
				op = CPY_RR
				arg1 = int(line[4] - 'a')
				arg2 = int(line[6] - 'a')
			default:
				//value
				op = CPY_VR
				fmt.Sscanf(line[3:], "%d", &arg1)
				arg2 = int(line[len(line)-1] - 'a')
			}
		case 'i':
			// inc
			op = INC_R
			arg1 = int(line[4] - 'a')
			arg2 = 1
		case 'd':
			// dec
			op = DEC_R
			arg1 = int(line[4] - 'a')
			arg2 = 1
		case 'j':
			// jnz
			switch line[4] {
			case 'a', 'b', 'c', 'd':
				// register
				arg1 = int(line[4] - 'a')
				switch line[6] {
				case 'a', 'b', 'c', 'd':
					// register
					op = JNZ_RR
					arg2 = int(line[6] - 'a')
				default:
					//value
					op = JNZ_RV
					fmt.Sscanf(line[5:], "%d", &arg2)
				}
			default:
				// value first
				switch line[len(line)-1] {
				case 'a', 'b', 'c', 'd':
					op = JNZ_VR
					arg2 = int(line[len(line)-1] - 'a')
					fmt.Sscanf(line[3:], "%d ", &arg1)
				default:
					op = JNZ_VV
					fmt.Sscanf(line[3:], "%d %d", &arg1, &arg2)
				}
			}
		case 't':
			// tgl
			switch line[4] {
			case 'a', 'b', 'c', 'd':
				// register
				op = TGL_R
				arg1 = int(line[4] - 'a')
			default:
				// value
				op = TGL_V
				fmt.Sscanf(line[3:], "%d", &arg1)
			}
		case 'o':
			//out
			switch line[4] {
			case 'a', 'b', 'c', 'd':
				// register
				op = OUT_R
				arg1 = int(line[4] - 'a')
			default:
				// value
				op = OUT_V
				fmt.Sscanf(line[3:], "%d", &arg1)
			}
		}
		//	fmt.Println("Line:", line, "Op", op, arg1, arg2)
		program = append(program, op, arg1, arg2)
	}

	return program
}

func runProgram(prog []int, a, b, c, d int) bool {
	//program counter
	pc := 0
	pl := len(prog) / 3
	// registers
	reg := [4]int{a, b, c, d}
	loop := 0
	var op, arg1, arg2 int

	var clockExpected, clockTicks int

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
		//fmt.Printf("Step %06d | PC %03d | Reg A:%03d B:%03d C:%03d D:%03d | Op %6s | Args: %d %d\n", loop, pc, reg[0], reg[1], reg[2], reg[3], Op(op), arg1, arg2)
		switch op {
		case CPY_VR:
			reg[arg2] = arg1
		case CPY_RR:
			reg[arg2] = reg[arg1]
		case INC_R:
			reg[arg1] = reg[arg1] + 1
		case DEC_R:
			reg[arg1] = reg[arg1] - 1
		case JNZ_VR:
			if arg1 != 0 {
				pc += reg[arg2] - 1
			}
		case JNZ_VV:
			if arg1 != 0 {
				pc += arg2 - 1
			}
		case JNZ_RV:
			if reg[arg1] != 0 {
				pc += arg2 - 1
			}
		case JNZ_RR:
			if reg[arg1] != 0 {
				pc += reg[arg2] - 1
			}
		case TGL_R:
			i := addr + reg[arg1]*3 + ADDR
			if i < 0 || i >= len(prog) {
				// If an attempt is made to toggle an instruction outside
				// the program, nothing happens.
				pc++
				continue
			}
			prog[i] = tglMap[prog[i]]
		case TGL_V:
			i := addr + arg1*3 + ADDR
			if i < 0 || i >= len(prog) {
				// If an attempt is made to toggle an instruction outside
				// the program, nothing happens.
				pc++
				continue
			}
			prog[i] = tglMap[prog[i]]
		case OUT_V:
			if clockExpected != arg1 {
				return false
			}
			clockExpected = -1 * (clockExpected - 1)
			clockTicks++
		case OUT_R:
			if clockExpected != reg[arg1] {
				return false
			}
			clockExpected = -1 * (clockExpected - 1)
			clockTicks++
		default:
			//skip invalid instruction
		}
		//time.Sleep(time.Millisecond * 50)
		pc++
		if clockTicks == 10000 {
			// gonna call that stable.
			return true
		}
	}
	return false
}

var tglMap = map[int]int{
	INC_R:  DEC_R,
	INC_V:  DEC_V,
	TGL_R:  INC_R,
	DEC_R:  INC_R,
	TGL_V:  INC_V,
	DEC_V:  INC_V,
	OUT_V:  INC_V,
	OUT_R:  INC_R,
	JNZ_RR: CPY_RR,
	JNZ_RV: CPY_RV,
	JNZ_VV: CPY_VV,
	JNZ_VR: CPY_VR,
	CPY_RR: JNZ_RR,
	CPY_RV: JNZ_RV,
	CPY_VV: JNZ_VV,
	CPY_VR: JNZ_VR,
}
