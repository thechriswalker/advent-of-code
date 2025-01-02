package main

import (
	"fmt"
	"strconv"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 18, solve1, solve2)
}

type Op struct {
	Code  uint8
	X     uint8
	RawY  int64
	isReg bool // if Y value is a register reference
}

func (o Op) Y(regs map[uint8]int64) int64 {
	if o.isReg {
		return regs[uint8(o.RawY)]
	}
	return o.RawY
}

type Runtime struct {
	reg       map[uint8]int64
	pc        int
	lastSound int64
}

func (r *Runtime) RunUntilRcv(ops []Op) int64 {
	for r.pc >= 0 && r.pc < len(ops) {
		if r.Apply(ops[r.pc]) {
			return r.lastSound
		}
	}
	return -1
}

func (r *Runtime) Apply(o Op) (recovered bool) {
	r.pc++
	switch o.Code {
	case CodeSet:
		r.reg[o.X] = o.Y(r.reg)
	case CodeAdd:
		r.reg[o.X] += o.Y(r.reg)
	case CodeMul:
		r.reg[o.X] *= o.Y(r.reg)
	case CodeMod:
		r.reg[o.X] %= o.Y(r.reg)
	case CodeSnd:
		r.lastSound = r.reg[o.X]
	case CodeRcv:
		if r.reg[o.X] != 0 {
			recovered = true
		}
	case CodeJgz:
		if r.reg[o.X] > 0 {
			r.pc += int(o.Y(r.reg) - 1) // to counteract the pc++ in the switch
		}
	}
	return
}

const (
	CodeUnknown uint8 = iota
	CodeSet
	CodeAdd
	CodeMul
	CodeMod
	CodeSnd
	CodeRcv
	CodeJgz
)

func parseOps(input string) []Op {
	ops := make([]Op, 0)
	aoc.MapLines(input, func(line string) error {
		var op Op
		switch line[:3] {
		case "set":
			op.Code = CodeSet
		case "add":
			op.Code = CodeAdd
		case "mul":
			op.Code = CodeMul
		case "mod":
			op.Code = CodeMod
		case "snd":
			op.Code = CodeSnd
		case "rcv":
			op.Code = CodeRcv
		case "jgz":
			op.Code = CodeJgz
		default:
			panic("unknown op: " + line)
		}
		op.X = line[4]
		if op.Code == CodeRcv || op.Code == CodeSnd {
			// no Y value

		} else if line[6] >= 'a' && line[6] <= 'z' {
			// a register.
			op.isReg = true
			op.RawY = int64(line[6])
		} else {
			// a number
			op.RawY, _ = strconv.ParseInt(line[6:], 10, 64)
		}
		ops = append(ops, op)
		return nil
	})
	return ops
}

// Implement Solution to Problem 1
func solve1(input string) string {
	regs := make(map[uint8]int64)

	r := Runtime{regs, 0, 0}

	return fmt.Sprint(r.RunUntilRcv(parseOps(input)))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	ops := parseOps(input)

	// for i := range ops {
	// 	fmt.Println(ops[i])
	// }

	cpu0 := &CPU{
		ops:  ops,
		regs: map[uint8]int64{'p': 0},
	}
	cpu1 := &CPU{
		ops:  ops,
		regs: map[uint8]int64{'p': 1},
	}

	inbox0 := &queue{}
	inbox1 := &queue{}

	for {
		cpu0.Tick(inbox1, inbox0)
		cpu1.Tick(inbox0, inbox1)
		if cpu0.waiting && cpu1.waiting {
			break
		}
	}

	return strconv.Itoa(cpu1.sendCount)
}

type CPU struct {
	ops       []Op
	regs      map[uint8]int64
	pc        int
	sendCount int
	waiting   bool
}

func (cpu *CPU) Tick(inbox, outbox *queue) {
	// proces one instruction, return next PC
	cpu.waiting = false
	op := cpu.ops[cpu.pc]
	switch op.Code {
	case CodeUnknown:
		panic("unknown op")
	case CodeSet:
		cpu.regs[op.X] = op.Y(cpu.regs)
	case CodeAdd:
		cpu.regs[op.X] += op.Y(cpu.regs)
	case CodeMul:
		cpu.regs[op.X] *= op.Y(cpu.regs)
	case CodeMod:
		cpu.regs[op.X] %= op.Y(cpu.regs)
	case CodeSnd:
		outbox.Push(cpu.regs[op.X])
		cpu.sendCount++
	case CodeRcv:
		if inbox.Len() == 0 {
			cpu.waiting = true
			return // block
		} else {
			cpu.regs[op.X] = inbox.Shift()
		}
	case CodeJgz:
		if cpu.regs[op.X] > 0 {
			cpu.pc += int(op.Y(cpu.regs))
			return
		}
	}
	cpu.pc++
}

type queue struct {
	data []int64
}

func (q *queue) Len() int {
	return len(q.data)
}

func (q *queue) Push(v int64) {
	q.data = append(q.data, v)
}

func (q *queue) Shift() int64 {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
