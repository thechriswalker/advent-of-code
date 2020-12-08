package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2020, 8, solve1, solve2)
}

type program struct {
	code []*instruction
}

type op uint8

const (
	oNOP op = iota
	oACC
	oJMP
)

type instruction struct {
	kind op
	arg  int
}

func (p *program) Run() (int, bool) {
	acc, pc := 0, 0
	tracker := map[int]struct{}{}
	for {
		if _, hit := tracker[pc]; hit {
			// hit a instruction we already used
			return acc, false
		}
		tracker[pc] = struct{}{}
		if pc < 0 || pc >= len(p.code) {
			// this counts as a termination!
			return acc, true
		}
		// now handle the op.
		ins := p.code[pc]
		switch ins.kind {
		case oNOP:
			pc++
		case oJMP:
			pc += ins.arg
		case oACC:
			acc += ins.arg
			pc++
		default:
			panic("unknown operation")
		}
	}
}

func parseProgram(input string) *program {
	lines := strings.Split(input, "\n")
	p := &program{code: make([]*instruction, 0, len(lines))}
	var ins op
	var arg int
	for _, line := range lines {
		if line == "" {
			break
		}
		arg, _ = strconv.Atoi(line[4:])
		switch line[:3] {
		case "nop":
			ins = oNOP
		case "jmp":
			ins = oJMP
		case "acc":
			ins = oACC
		}
		p.code = append(p.code, &instruction{kind: ins, arg: arg})
	}
	return p
}

// Implement Solution to Problem 1
func solve1(input string) string {
	p := parseProgram(input)
	acc, _ := p.Run()
	return fmt.Sprintf("%d", acc)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	p := parseProgram(input)
	var restore func()
	for i := 0; i < len(p.code); i++ {
		ins := p.code[i]
		switch ins.kind {
		case oACC:
			continue // does this skip the loop?
		case oNOP:
			ins.kind = oJMP
			restore = func() {
				ins.kind = oNOP
			}
		case oJMP:
			ins.kind = oNOP
			restore = func() {
				ins.kind = oJMP
			}
		}
		if acc, ok := p.Run(); ok {
			return fmt.Sprintf("%d", acc)
			// } else {
			// 	log.Printf("acc:%d, terminated:%v, modified:%#v", acc, ok, ins)
		}
		restore()
	}
	return "<unsolved>"
}
