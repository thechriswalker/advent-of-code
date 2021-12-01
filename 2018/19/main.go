package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2018, 19, solve1, solve2)
}

// Implement Solution to Problem 2
// naive solution fails because the program effectively doesn't halt.
// so there must be a pattern. or at least lets look at what the code is doing:
/*
#ip 5 				    # PC is register 5
00: addi 5 16 5  	# set PC to 16 (then it increments to 17)
01: seti 1 1 2
02: seti 1 8 1
03: mulr 2 1 3
04: eqrr 3 4 3
05: addr 3 5 5
06: addi 5 1 5
07: addr 2 0 0
08: addi 1 1 1
09: gtrr 1 4 3
10: addr 5 3 5
11: seti 2 6 5
12: addi 2 1 2
13: gtrr 2 4 3
14: addr 3 5 5
15: seti 1 2 5
16: mulr 5 5 5
17: addi 4 2 4		# R4 = R4 + 2
18: mulr 4 4 4		# R4 = R4 * R4
19: mulr 5 4 4    # R4 = PC (19) * R4
20: muli 4 11 4   # R4 *= 11
21: addi 3 2 3    # R3 *= 2
22: mulr 3 5 3    # R3 = PC(22) * R3
23: addi 3 13 3   # R3 += 13
24: addr 4 3 4    # R4 += R3
25: addr 5 0 5    # PC = PC(25) + R0 (PC then increments again )
26: seti 0 8 5    # PC = 8 (then inc to 9)
27: setr 5 5 3    # R3 = PC(27)
28: mulr 3 5 3    # R3 = PC(28) * R3
29: addr 5 3 3		# R3 = R3 + PC(29)
30: mulr 5 3 3	  # R3 = PC(30) * R3
31: muli 3 14 3   # R3 = R3 * 14
32: mulr 3 5 3    #
33: addr 4 3 4
34: seti 0 9 0
35: seti 0 9 5
*/

func program_long_version() int {
	var r0, r1, r2, r3, r4 int
	r0 = 1
	// 00: addi 5 16 5  	# set PC to 16 (then it increments to 17)
	// 17: addi 4 2 4		# R4 = R4 + 2
	r4 += 2
	// 18: mulr 4 4 4		# R4 = R4 * R4
	r4 += r4
	// 19: mulr 5 4 4    # R4 = PC (19) * R4
	r4 *= 19
	// 20: muli 4 11 4   # R4 *= 11
	r4 *= 11
	// 21: addi 3 2 3    # R3 += 2
	r3 += 2
	// 22: mulr 3 5 3    # R3 = PC(22) * R3
	r3 *= 22
	// 23: addi 3 13 3   # R3 += 13
	r3 += 13
	// 24: addr 4 3 4    # R4 += R3
	r4 += r3
	// 25: addr 5 0 5    # PC = PC(25) + R0 (1) = 26, PC then increments again to 27
	// 27: setr 5 5 3    # R3 = PC(27)
	r3 = 27
	// 28: mulr 3 5 3    # R3 = PC(28) * R3
	r3 *= 28
	// 29: addr 5 3 3		# R3 = R3 + PC(29)
	r3 += 29
	// 30: mulr 5 3 3	  # R3 = PC(30) * R3
	r3 *= 30
	// 31: muli 3 14 3   # R3 = R3 * 14
	r3 *= 14
	// 32: mulr 3 5 3     # R3 = R3 * PC(32)
	r3 *= 32
	//33: addr 4 3 4 # R4 = R4 + R3
	r4 += r3
	// 34: seti 0 9 0 # R0 = 0
	r0 = 0
	// 35: seti 0 9 5 # PC = 0
	// 00: addi 5 16 5  	# set PC to 16 (then it increments to 17)
	// 17: addi 4 2 4		# R4 = R4 + 2
	r4 += 2
	// 18: mulr 4 4 4		# R4 = R4 * R4
	r4 += r4
	// 19: mulr 5 4 4    # R4 = PC (19) * R4
	r4 *= 19
	// 20: muli 4 11 4   # R4 *= 11
	r4 *= 11
	// 21: addi 3 2 3    # R3 += 2
	r3 += 2
	// 22: mulr 3 5 3    # R3 = PC(22) * R3
	r3 *= 22
	// 23: addi 3 13 3   # R3 += 13
	r3 += 13
	// 24: addr 4 3 4    # R4 += R3
	r4 += r3
	// 25: addr 5 0 5    # PC = PC(25) + R0 (0) = 26, PC then increments again to 26
	// 26: seti 0 8 5    # PC = 8 (then inc to 9)
	// this looks like the start of a loop
	for {
		// 09: gtrr 1 4 3    # r3 = r1 > r4 ? 1 : 0
		if r1 > r4 {
			// 10: addr 5 3 5    # PC = PC(10) + r3 (1) = 11 -> then inc to 12
			// 12: addi 2 1 2
			r2 += 1
			// 13: gtrr 2 4 3    # r3 = r2 > r4 ? 1 : 0
			if r2 > r4 {
				// 14: addr 3 5 5  # PC(14) += r3(1) => 15 +1 => 16
				// 16: mulr 5 5 5  # PC(16) = PC(16) * PC(16)
				// EXIT!!!!
				return r0
			} else {
				// 14: addr 3 5 5  # PC(14) += r3(0) => 14 +1 => 15
				// 15: seti 1 2 5  # PC(15) = 1 (then inc to)
				// 02: seti 1 8 1  # R1 = 1
				r1 = 1 // reset the loop
				// re-merge branch @3
			}
		} //else {
		// 10: addr 5 3 5    # PC = PC(10) + r3 (0) = 10 -> then inc to 11
		// 11: seti 2 6 5    # PC = 2 (inc to 3)
		//}
		// 03: mulr 2 1 3    # R3 = R1 * R2
		r3 = r1 * r2
		// 04: eqrr 3 4 3    # R3 = R3 > R4 ? 1 : 0 => 0
		if r3 == r4 {
			// 05: addr 3 5 5    # PC = PC(5) + R3 (1) = 6 (then inc => 7
			// 07: addr 2 0 0  # R0 = R0 + r2
			r0 += r2
			// branch merges after if statements at PC = 8
		} //else {
		// 05: addr 3 5 5    # PC = PC(5) + R3 (0) = 6 (then inc => 6
		// 06: addi 5 1 5 # PC(6) += 1 (then inc => 8)
		// branch merges after if statements at PC = 8
		//}

		// 08: addi 1 1 1 # r1++
		r1++
	}
}

func program_shorter(r0 int) int {
	var r1, r2, r3, r4 int
	// making no assumption about r0 other than it is zero or one
	// if it's one we might as well do this jump first, as we come back to it otherwise
	if r0 == 1 {
		// 25: addr 5 0 5    # PC = PC(25) + R0 (1) = 26, PC then increments again to 27
		// 27: setr 5 5 3    # R3 = PC(27)
		r3 = ((27 * 28) + 29) * 30 * 14
		r4 += r3
		// 34: seti 0 9 0 # R0 = 0
		r0 = 0
	}
	r4 = (r4 + 2) * 2 * 19 * 11
	// 21: addi 3 2 3    # R3 += 2
	r3 = ((r3 + 2) * 22) + 13
	r4 += r3
	// 25: addr 5 0 5    # PC = PC(25) + R0 (0) = 26, PC then increments again to 26
	// 26: seti 0 8 5    # PC = 8 (then inc to 9)
	// this looks like the start of a loop
	// at this point r4 is one of 2 values... both fairly high
	for {
		// 09: gtrr 1 4 3    # r3 = r1 > r4 ? 1 : 0
		if r1 > r4 {
			// 10: addr 5 3 5    # PC = PC(10) + r3 (1) = 11 -> then inc to 12
			// 12: addi 2 1 2
			r2++
			// 13: gtrr 2 4 3    # r3 = r2 > r4 ? 1 : 0
			if r2 > r4 {
				// 14: addr 3 5 5  # PC(14) += r3(1) => 15 +1 => 16
				// 16: mulr 5 5 5  # PC(16) = PC(16) * PC(16)
				// EXIT!!!!
				return r0
			} // else {
			r1 = 1
			// branch merge at instruction 3
			//}
		}
		// 10: addr 5 3 5    # PC = PC(10) + r3 (0) = 10 -> then inc to 11
		// 11: seti 2 6 5    # PC = 2 (then in to 3)
		// 03: mulr 2 1 3    # R3 = R1 * R2
		r3 = r1 * r2
		// 04: eqrr 3 4 3    # R3 = R3 == R4 ? 1 : 0
		if r3 == r4 {
			// 05: addr 3 5 5    # PC = PC(5) + R3 (1) = 6 (then inc => 7
			// 07: addr 2 0 0  # R0 = R0 + r2
			// r0 gets incremented by r2 if r1 * r2 = r4
			// i.e. r2 is a factor of r4 (twice) this is a double loop.
			r0 += r2
			// branch merges after if statements at PC = 8
		}
		//else {
		// 05: addr 3 5 5    # PC = PC(5) + R3 (0) = 6 (then inc => 6
		// 06: addi 5 1 5 # PC(6) += 1 (then inc => 8)
		// branch merges after if statements at PC = 8
		//}

		// 08: addi 1 1 1 # r1++
		r1++
	}
}

func program_shorter2(r0 int) int {
	var x int
	switch r0 {
	case 1:
		x = 10551293
	case 0:
		x = 893
	default:
		panic("Bad r0 value must be 1 or 0")
	}
	i, j := 1, 1
	for {
		if i > x {
			j++
			if j > x {
				return r0
			}
			i = 1
		}
		if i*j == x {
			fmt.Println("Added", j, "to", r0, "now", r0+j, "i was", i)
			r0 += j
		}
		i++
	}
}

func program_opt(r0 int) int {
	var x int
	switch r0 {
	case 1:
		x = 10551293
	case 0:
		x = 893
	default:
		panic("Bad r0 value must be 1 or 0")
	}
	r0 = 0
	// we only have to iterate once calculating divisors
	// this turns the O(n^2) loop into O(n) which is doable.
	for i := 1; i <= x; i++ {
		// is it a factor?
		if x%i == 0 {
			r0 += i
		}
	}
	return r0
}

func solve2(input string) string {
	return fmt.Sprintf("%d", program_opt(1))
}

// Implement Solution to Problem 1
// solved initially with the slow programs iterating code
// switched to the faster code to use as a test case
func solve1(input string) string {
	// pcReg, program := parseInput(input)

	// reg := [6]int{}
	// runProgram(pcReg, program, &reg, false)

	// return fmt.Sprintf("%d", reg[0])

	return fmt.Sprintf("%d", program_opt(0))
}

func runProgram(pcIndex int, prog []instruction, registers *[6]int, slow bool) {
	pc := 0
	//	fmt.Println(prog)
	tick := 0
	for {
		if pc < 0 || pc >= len(prog) {
			return
		}
		ins := prog[pc]
		// set pc into register pcIndex
		registers[pcIndex] = pc
		// execute instruction
		ins.Execute(registers)
		// read pc from registers and increment
		pc = registers[pcIndex] + 1
		if slow && tick > 6000 {
			fmt.Printf("T%06d | PC%05d | A%6d B%6d C%6d D%6d E%6d F%6d\n", tick, pc, registers[0], registers[1], registers[2], registers[3], registers[4], registers[5])
			time.Sleep(time.Millisecond * 10)
		}
		tick++
	}
}

type instruction struct {
	fn      func(r *[6]int, a, b, c int)
	a, b, c int
}

func (i instruction) Execute(r *[6]int) {
	i.fn(r, i.a, i.b, i.c)
}

func parseInput(input string) (pcReg int, program []instruction) {
	rd := strings.NewReader(input)
	// line one will be the instruction pointer register
	fmt.Fscanf(rd, "#ip %d\n", &pcReg)
	var name string
	var a, b, c int
	var err error
	for {
		// following lines will be program
		if _, err = fmt.Fscanf(rd, "%s %d %d %d\n", &name, &a, &b, &c); err != nil {
			break
		}
		program = append(program, instruction{
			fn: opMap[name],
			a:  a, b: b, c: c,
		})
	}
	return
}

// everything below here was copied from day 16 and the registers widened

/*
Addition:
*/
// addr (add register) stores into register C the result of adding register A and register B.
func addr(r *[6]int, a, b, c int) {
	r[c] = r[a] + r[b]
}

// addi (add immediate) stores into register C the result of adding register A and value B.
func addi(r *[6]int, a, b, c int) {
	r[c] = r[a] + b
}

//Multiplication:

// mulr (multiply register) stores into register C the result of multiplying register A and register B.
func mulr(r *[6]int, a, b, c int) {
	r[c] = r[a] * r[b]
}

// muli (multiply imm2ediate) stores into register C the result of multiplying register A and value B.
func muli(r *[6]int, a, b, c int) {
	r[c] = r[a] * b
}

//Bitwise AND:

//banr (bitwise AND register) stores into register C the result of the bitwise AND of register A and register B.
func banr(r *[6]int, a, b, c int) {
	r[c] = r[a] & r[b]
}

// bani (bitwise AND immediate) stores into register C the result of the bitwise AND of register A and value B.
func bani(r *[6]int, a, b, c int) {
	r[c] = r[a] & b
}

// Bitwise OR:

// borr (bitwise OR register) stores into register C the result of the bitwise OR of register A and register B.
func borr(r *[6]int, a, b, c int) {
	r[c] = r[a] | r[b]
}

// bori (bitwise OR immediate) stores into register C the result of the bitwise OR of register A and value B.
func bori(r *[6]int, a, b, c int) {
	r[c] = r[a] | b
}

// Assignment:

//setr (set register) copies the contents of register A into register C. (Input B is ignored.)
func setr(r *[6]int, a, b, c int) {
	r[c] = r[a]
}

//seti (set immediate) stores value A into register C. (Input B is ignored.)
func seti(r *[6]int, a, b, c int) {
	r[c] = a
}

//Greater-than testing:

//gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0.
func gtir(r *[6]int, a, b, c int) {
	if a > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

//gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
func gtri(r *[6]int, a, b, c int) {
	if r[a] > b {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

//gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0.
func gtrr(r *[6]int, a, b, c int) {
	if r[a] > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

//Equality testing:

//eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0.
func eqir(r *[6]int, a, b, c int) {
	if a == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

//eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
func eqri(r *[6]int, a, b, c int) {
	if r[a] == b {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

//eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0.
func eqrr(r *[6]int, a, b, c int) {
	if r[a] == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

type operation func(r *[6]int, a, b, c int)

var opMap = map[string]operation{
	"addr": addr,
	"addi": addi,
	"mulr": mulr,
	"muli": muli,
	"banr": banr,
	"bani": bani,
	"borr": borr,
	"bori": bori,
	"setr": setr,
	"seti": seti,
	"gtir": gtir,
	"gtri": gtri,
	"gtrr": gtrr,
	"eqir": eqir,
	"eqri": eqri,
	"eqrr": eqrr,
}
