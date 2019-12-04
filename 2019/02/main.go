package main

import (
	"fmt"
	"io"
	"log"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2019, 2, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	p := NewProgram(input)
	// so simulate the actual problem, not the tests
	if len(input) > 50 {
		p.Set(1, 12)
		p.Set(2, 2)
	}
	p.Run()
	return fmt.Sprintf("%d", p.Get(0))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	target := 19690720
	clean := NewProgram(input)
	// naive exhaustive search
	for n := 0; n < 100; n++ {
		for v := 0; v < 100; v++ {
			p := clean.Copy()
			p.Set(1, n)
			p.Set(2, v)
			if p.Run() {
				if p.Get(0) == target {
					// found answer
					return fmt.Sprintf("%d", 100*n+v)
				}
			}
		}
	}
	return "<unsolved>"
}

type Program struct {
	pc        int
	registers []int
	Failed    bool
}

const (
	OP_ADD  = 1
	OP_MUL  = 2
	OP_HALT = 99
)

func (p *Program) Debug() bool {
	p.pc = 0
	for p.Tick() {
		log.Println(p)
	}
	return !p.Failed

}

func (p *Program) Run() bool {
	p.pc = 0
	for p.Tick() {
	}
	return !p.Failed
}

func (p *Program) String() string {
	return fmt.Sprintf("{pc:%d, reg:%v}", p.pc, p.registers)
}

// returns true while the program should continue
func (p *Program) Tick() bool {
	op := p.Consume()
	switch op {
	case OP_ADD:
		a := p.Consume()
		b := p.Consume()
		c := p.Consume()
		p.Set(c, p.Get(a)+p.Get(b))
		return true
	case OP_MUL:
		a := p.Consume()
		b := p.Consume()
		c := p.Consume()
		p.Set(c, p.Get(a)*p.Get(b))
		return true
	case OP_HALT:
		return false
	default:
		p.Failed = true
		return false
	}
}

func (p *Program) Consume() int {
	v := p.registers[p.pc]
	p.pc++
	return v
}

func (p *Program) Get(register int) int {
	return p.registers[register]
}

func (p *Program) Set(register, value int) {
	p.registers[register] = value
}

func NewProgram(input string) *Program {
	rd := strings.NewReader(strings.TrimSpace(input))
	reg := []int{}
	var x int
	var err error
	for {
		_, err = fmt.Fscanf(rd, "%d", &x)
		if err != nil {
			break
		}
		rd.Seek(1, io.SeekCurrent)
		reg = append(reg, x)
	}
	return &Program{
		registers: reg,
	}
}

// create a copy (which we can then mutate)
func (p *Program) Copy() *Program {
	p2 := &Program{
		registers: make([]int, len(p.registers)),
	}
	copy(p2.registers, p.registers)
	return p2
}
