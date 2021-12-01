package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2018, 16, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	tests := parseInputOne(input)
	var n int
	//	fmt.Println(tests)
	for _, t := range tests {
		if opcodePossibilityCheck(t.before, t.after, t.a, t.b, t.c) > 2 {
			n++
		}
	}
	return fmt.Sprintf("%d", n)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	tests := parseInputOne(input)
	//	fmt.Println(tests)
	opcodePossibilities := map[int][]string{}
	for _, t := range tests {
		list := opcodePossibilityList(t.before, t.after, t.a, t.b, t.c)
		//	fmt.Println(t.code, list)
		if prev, ok := opcodePossibilities[t.code]; ok {
			// intersect list.
			opcodePossibilities[t.code] = intersectOpList(prev, list)
		} else {
			opcodePossibilities[t.code] = list
		}
	}
	// hopefully we have 1 for each now.
	// for i := 0; i < 16; i++ {
	// 	fmt.Println("Code:", i, "Ops:", opcodePossibilities[i])
	// }
	// but we don't so we have to take the ones that DO have one
	// and use those removing the  the others until we have no 1 length codes left
	ops := [16]operation{}
	var found int
	removals := map[string]struct{}{}
	for found < 16 {
		//fmt.Println("list", opcodeList(opcodePossibilities))
		// find all one length lists.
		for code, list := range opcodePossibilities {
			if len(list) == 1 {
				//	fmt.Println("found code", code, "=>", list[0])
				ops[code] = opMap[list[0]]
				found++
				removals[list[0]] = struct{}{}
				delete(opcodePossibilities, code)
			}
		}
		if found == 16 {
			break
		}
		removeCodes(opcodePossibilities, removals)
		removals = map[string]struct{}{}
		//time.Sleep(time.Second)
	}

	// got them all, now run the program.
	p := parseProgram(input)
	registers := runProgram(ops, p, [4]int{})
	return fmt.Sprintf("%d", registers[0])
}

type opcodeList map[int][]string

func (o opcodeList) String() string {
	s := &strings.Builder{}
	for i := 0; i < 16; i++ {
		if len(o[i]) >= 1 {
			fmt.Fprintf(s, "Code: %2d, Options: %v\n", i, o[i])
		}
	}
	return s.String()
}

func removeCodes(codes map[int][]string, removals map[string]struct{}) {
	for i, list := range codes {
		next := []string{}
		for _, s := range list {
			if _, ok := removals[s]; !ok {
				next = append(next, s)
			}
		}
		codes[i] = next
	}
}

func intersectOpList(a, b []string) []string {
	s := map[string]int{}
	for i := range a {
		s[a[i]] = 1
	}
	for i := range b {
		s[b[i]] = s[b[i]] - 1
	}
	intersection := []string{}
	for k, v := range s {
		if v == 0 {
			intersection = append(intersection, k)
		}
	}
	sort.Sort(sort.StringSlice(intersection))
	return intersection
}

func parseProgram(input string) []int {
	index := strings.Index(input, "\n\n\n\n")
	rd := strings.NewReader(input[index+4:])
	var op, a, b, c int
	program := []int{}
	for {
		if _, err := fmt.Fscanf(rd, "%d %d %d %d\n", &op, &a, &b, &c); err != nil {
			break
		}
		program = append(program, op, a, b, c)
	}
	return program
}

func runProgram(opcodes [16]operation, program []int, registers [4]int) [4]int {
	for i := 0; i < len(program); i += 4 {
		//	fmt.Println(registers)
		opcodes[program[i]](&registers, program[i+1], program[i+2], program[i+3])
	}
	//fmt.Println(registers)
	return registers
}

func parseInputOne(input string) []Test {
	// Before: [3, 2, 1, 1]
	// 9 2 1 2
	// After:  [3, 2, 2, 1]
	tests := []Test{}
	rd := strings.NewReader(input)
	var i1, i2, i3, i4, op, a, b, c, o1, o2, o3, o4 int
	for {
		if _, err := fmt.Fscanf(rd, "Before: [%d, %d, %d, %d]\n%d %d %d %d\nAfter: [%d, %d, %d, %d]\n\n", &i1, &i2, &i3, &i4, &op, &a, &b, &c, &o1, &o2, &o3, &o4); err != nil {
			break
		}
		tests = append(tests, Test{
			before: [4]int{i1, i2, i3, i4},
			after:  [4]int{o1, o2, o3, o4},
			code:   op,
			a:      a,
			b:      b,
			c:      c,
		})
	}
	return tests
}

type Test struct {
	before, after [4]int
	code          int
	a, b, c       int
}

func opcodePossibilityCheck(before, after [4]int, a, b, c int) int {
	test := [4]int{0, 0, 0, 0}
	var count int
	for _, op := range opMap {
		copy(test[0:4], before[0:4])
		//	fmt.Printf("Testing op '%s' with input a:%d, b:%d, c:%d on %v => %v\n", name, a, b, c, test, after)
		op(&test, a, b, c)
		//	fmt.Printf("Actual after: %v\n", test)
		if test == after {
			count++
		}
	}
	return count
}

var codes = []string{"addr", "addi", "mulr", "muli", "banr", "bani", "borr", "bori", "setr", "seti", "gtrr", "gtri", "gtir", "eqrr", "eqri", "eqir"}

func init() {
	sort.Sort(sort.StringSlice(codes))
	fmt.Println(codes)
}

func opcodePossibilityList(before, after [4]int, a, b, c int) []string {
	test := [4]int{0, 0, 0, 0}
	list := []string{}

	for _, name := range codes {
		copy(test[0:4], before[0:4])
		//	fmt.Printf("Testing op '%s' with input a:%d, b:%d, c:%d on %v => %v\n", name, a, b, c, test, after)
		opMap[name](&test, a, b, c)
		//	fmt.Printf("Actual after: %v\n", test)
		if test == after {
			list = append(list, name)
		}
	}
	return list
}

/*
Addition:
*/
// addr (add register) stores into register C the result of adding register A and register B.
func addr(r *[4]int, a, b, c int) {
	r[c] = r[a] + r[b]
}

// addi (add immediate) stores into register C the result of adding register A and value B.
func addi(r *[4]int, a, b, c int) {
	r[c] = r[a] + b
}

//Multiplication:

// mulr (multiply register) stores into register C the result of multiplying register A and register B.
func mulr(r *[4]int, a, b, c int) {
	r[c] = r[a] * r[b]
}

// muli (multiply imm2ediate) stores into register C the result of multiplying register A and value B.
func muli(r *[4]int, a, b, c int) {
	r[c] = r[a] * b
}

//Bitwise AND:

//banr (bitwise AND register) stores into register C the result of the bitwise AND of register A and register B.
func banr(r *[4]int, a, b, c int) {
	r[c] = r[a] & r[b]
}

// bani (bitwise AND immediate) stores into register C the result of the bitwise AND of register A and value B.
func bani(r *[4]int, a, b, c int) {
	r[c] = r[a] & b
}

// Bitwise OR:

// borr (bitwise OR register) stores into register C the result of the bitwise OR of register A and register B.
func borr(r *[4]int, a, b, c int) {
	r[c] = r[a] | r[b]
}

// bori (bitwise OR immediate) stores into register C the result of the bitwise OR of register A and value B.
func bori(r *[4]int, a, b, c int) {
	r[c] = r[a] | b
}

// Assignment:

//setr (set register) copies the contents of register A into register C. (Input B is ignored.)
func setr(r *[4]int, a, b, c int) {
	r[c] = r[a]
}

//seti (set immediate) stores value A into register C. (Input B is ignored.)
func seti(r *[4]int, a, b, c int) {
	r[c] = a
}

//Greater-than testing:

//gtir (greater-than immediate/register) sets register C to 1 if value A is greater than register B. Otherwise, register C is set to 0.
func gtir(r *[4]int, a, b, c int) {
	if a > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

//gtri (greater-than register/immediate) sets register C to 1 if register A is greater than value B. Otherwise, register C is set to 0.
func gtri(r *[4]int, a, b, c int) {
	if r[a] > b {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

//gtrr (greater-than register/register) sets register C to 1 if register A is greater than register B. Otherwise, register C is set to 0.
func gtrr(r *[4]int, a, b, c int) {
	if r[a] > r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

//Equality testing:

//eqir (equal immediate/register) sets register C to 1 if value A is equal to register B. Otherwise, register C is set to 0.
func eqir(r *[4]int, a, b, c int) {
	if a == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

//eqri (equal register/immediate) sets register C to 1 if register A is equal to value B. Otherwise, register C is set to 0.
func eqri(r *[4]int, a, b, c int) {
	if r[a] == b {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

//eqrr (equal register/register) sets register C to 1 if register A is equal to register B. Otherwise, register C is set to 0.
func eqrr(r *[4]int, a, b, c int) {
	if r[a] == r[b] {
		r[c] = 1
	} else {
		r[c] = 0
	}
}

type operation func(r *[4]int, a, b, c int)

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
