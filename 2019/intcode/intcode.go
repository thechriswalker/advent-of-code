package intcode

import (
	"fmt"
	"io"
	"log"
	"strings"
)

// Program represents the interpreter
type Program struct {
	pc        int
	registers []int
	Failed    bool
	Input     chan int
	Output    chan int
}

// OpCodes
const (
	opADD      = 1
	opMUL      = 2
	opINPUT    = 3
	opOUTPUT   = 4
	opJUMPT    = 5
	opJUMPF    = 6
	opLESSTHAN = 7
	opEQUALS   = 8
	opHALT     = 99
)

// EnqueueInput asynchronously queues values for input
func (p *Program) EnqueueInput(in ...int) {
	go func() {
		for _, n := range in {
			p.Input <- n
		}
	}()
}

// GetOutput waits for the next output to be emitted
func (p *Program) GetOutput() int {
	return <-p.Output
}

// Debug is like Run, but dumps the program at each tick
func (p *Program) Debug() bool {
	p.pc = 0
	for p.Tick() {
		log.Println(p)
	}
	return !p.Failed
}

// Run executes the program returning true if no errors occured
func (p *Program) Run() bool {
	p.pc = 0
	for p.Tick() {
	}
	return !p.Failed
}

// RunAsync returns a channel which returns the exit status when finished
func (p *Program) RunAsync(debug bool) <-chan bool {
	ch := make(chan bool)
	go func() {
		if debug {
			ch <- p.Debug()
		} else {
			ch <- p.Run()
		}
	}()
	return ch
}

func (p *Program) String() string {
	return fmt.Sprintf("{pc:%d, reg:%v}", p.pc, p.registers)
}

// Tick returns true while the program should continue
func (p *Program) Tick() bool {
	op := p.consume()
	switch op % 100 {
	case opADD:
		a, b := p.getTwoArgs(op)
		c := p.consume() // c is a raw address
		p.Set(c, a+b)

	case opMUL:
		a, b := p.getTwoArgs(op)
		c := p.consume() // c is a raw address
		p.Set(c, a*b)

	case opINPUT:
		// take an input, always a register address
		a := p.consume()
		p.Set(a, <-p.Input)

	case opOUTPUT:
		// emit an output
		p.Output <- p.getOneArg(op)

	case opJUMPT:
		a, b := p.getTwoArgs(op)
		if a != 0 {
			p.pc = b
		}
	case opJUMPF:
		a, b := p.getTwoArgs(op)
		if a == 0 {
			p.pc = b
		}

	case opLESSTHAN:
		a, b := p.getTwoArgs(op)
		c := p.consume() // raw address
		if a < b {
			p.Set(c, 1)
		} else {
			p.Set(c, 0)
		}
	case opEQUALS:
		a, b := p.getTwoArgs(op)
		c := p.consume() // raw address
		if a == b {
			p.Set(c, 1)
		} else {
			p.Set(c, 0)
		}

	case opHALT:
		return false

	default:
		p.Failed = true
		return false
	}
	return true
}

func (p *Program) consume() int {
	v := p.registers[p.pc]
	p.pc++
	return v
}

func (p *Program) getOneArg(op int) int {
	a := p.consume()
	return p.get(a, op/100)
}

func (p *Program) getTwoArgs(op int) (int, int) {
	a := p.consume()
	b := p.consume()
	return p.get(a, op/100), p.get(b, op/1000)
}

func (p *Program) getThreeArgs(op int) (int, int, int) {
	a := p.consume()
	b := p.consume()
	c := p.consume()
	return p.get(a, op/100), p.get(b, op/1000), p.get(c, op/10000)
}

// Get a value from a register
func (p *Program) Get(register int) int {
	return p.get(register, 0) // positional, value at register.
}

func (p *Program) get(value, mode int) int {
	// check the op and pos to see if we are in positional or immediate mode
	if mode%2 == 0 {
		// positional, value is register
		return p.registers[value]
	}
	// immediate, value is value
	return value
}

// Set a value at a register address
func (p *Program) Set(register, value int) {
	p.registers[register] = value
}

// New creates a new program
func New(code string) *Program {
	rd := strings.NewReader(strings.TrimSpace(code))
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
		Input:     make(chan int),
		Output:    make(chan int),
	}
}

// Copy creates a copy (which we can then mutate)
// it is created with the current memory state, so
// be sure to copy before running.
func (p *Program) Copy() *Program {
	p2 := &Program{
		registers: make([]int, len(p.registers)),
		Input:     make(chan int),
		Output:    make(chan int),
	}
	copy(p2.registers, p.registers)
	return p2
}
