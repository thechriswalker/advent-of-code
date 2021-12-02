package intcode

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"
)

// Program represents the interpreter
type Program struct {
	pc           int64
	code         []int64
	memory       map[int64]int64 // to allow sparse memory
	relativeBase int64
	initialised  bool
	Failed       bool
	Input        chan func() int64
	Output       chan int64
	Halted       chan struct{}
	ID           string
	Debug        bool
}

// OpCodes
const (
	opADD       = 1
	opMUL       = 2
	opINPUT     = 3
	opOUTPUT    = 4
	opJUMPT     = 5
	opJUMPF     = 6
	opLESSTHAN  = 7
	opEQUALS    = 8
	opADJUSTREL = 9
	opHALT      = 99
)

// blocks until input is read
func (p *Program) NextInput(in int64) {
	p.Input <- func() int64 {
		return in
	}
}

// GetOutput waits for the next output to be emitted
func (p *Program) GetOutput() int64 {
	return <-p.Output
}

func (p *Program) Initialise() {
	p.pc = 0
	// we load the memory! (less need to copy the program now)
	// initialise with twice as much
	p.memory = make(map[int64]int64, len(p.code)*2)
	for i, c := range p.code {
		p.memory[int64(i)] = c
	}
	p.initialised = true
}

// Run executes the program returning true if no errors occured
func (p *Program) Run() bool {
	if !p.initialised {
		p.Initialise()
	}
	for p.Tick() {
	}
	return !p.Failed
}

// RunAsync returns a channel which returns the exit status when finished
func (p *Program) RunAsync() <-chan bool {
	ch := make(chan bool)
	go func() {
		ch <- p.Run()
	}()
	return ch
}

func (p *Program) String() string {
	return fmt.Sprintf("{pc:%d, reg:%v}", p.pc, p.memory)
}

func (p *Program) Log(a ...interface{}) {
	if p.Debug {
		log.Printf(`[%s] %s`, p.ID, fmt.Sprintln(a...))
	}
}

// Tick returns true while the program should continue
func (p *Program) Tick() bool {
	op := p.consume()
	switch op % 100 {
	case opADD:
		a, b, c := p.getTwoArgsAndAddress(op)
		p.Set(c, a+b)

	case opMUL:
		a, b, c := p.getTwoArgsAndAddress(op)
		p.Set(c, a*b)

	case opINPUT:
		// take an input, always a address address
		// a := p.getOneArg(op) definitely wrong
		a := p.getOneAddress(op)
		select {
		case in := <-p.Input:
			// immediate input
			i := in()
			p.Log("input...", i, "(sync)")
			p.Set(a, i)
		default:
			// deferred input
			select {
			case fi := <-p.Input:
				i := fi()
				p.Log("input...", i, "(async)")
				p.Set(a, i)
			case <-time.After(time.Second):
				p.Log("output... <TIMEOUT> (async)")
				p.Failed = true
				return false
			}
		}

	case opOUTPUT:
		// emit an output
		out := p.getOneArg(op)
		select {
		case p.Output <- out:
			// immediate send.
			p.Log("output...", out, "(sync)")
		default:
			// deferred send.
			select {
			case p.Output <- out:
				p.Log("output...", out, "(async)")
			case <-time.After(time.Second):
				p.Log("output... <TIMEOUT> (async)")
				p.Failed = true
				return false
			}
		}

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
		a, b, c := p.getTwoArgsAndAddress(op)
		if a < b {
			p.Set(c, 1)
		} else {
			p.Set(c, 0)
		}
	case opEQUALS:
		a, b, c := p.getTwoArgsAndAddress(op)
		if a == b {
			p.Set(c, 1)
		} else {
			p.Set(c, 0)
		}

	case opADJUSTREL:
		a := p.getOneArg(op)
		p.relativeBase += a

	case opHALT:
		close(p.Halted)
		return false

	default:
		p.Failed = true
		return false
	}
	return true
}

func (p *Program) consume() int64 {
	v := p.memory[p.pc]
	p.pc++
	return v
}

func (p *Program) getOneArg(op int64) int64 {
	a := p.consume()
	return p.get(a, op/100)
}

func (p *Program) getOneAddress(op int64) int64 {
	a := p.consume()
	return p.getAddress(a, op/100)
}

func (p *Program) getTwoArgs(op int64) (int64, int64) {
	a := p.consume()
	b := p.consume()
	return p.get(a, op/100), p.get(b, op/1000)
}

func (p *Program) getTwoArgsAndAddress(op int64) (int64, int64, int64) {
	a := p.consume()
	b := p.consume()
	c := p.consume()
	return p.get(a, op/100), p.get(b, op/1000), p.getAddress(c, op/10000)
}

// Get a value from a address
func (p *Program) Get(address int64) int64 {
	return p.get(address, 0) // positional, value at address.
}

func (p *Program) get(value, mode int64) int64 {
	// check the op and pos to see if we are in positional or immediate mode
	//log.Printf("Getting value from %v in mode %d\n", value, mode%10)
	switch mode % 10 {
	case 0:
		// positional, value is address
		return p.memory[value]
	case 1:
		// immediate, value is value
		return value
	case 2:
		// relative value, based on current relative-base
		return p.memory[p.relativeBase+value]
	default:
		panic("unknown memory access mode")
	}
}

func (p *Program) getAddress(value, mode int64) int64 {
	// no positional, but could be relative.
	switch mode % 10 {
	case 0, 1:
		// use value as address in either mode
		return value
	case 2:
		// relative value, based on current relative-base
		return p.relativeBase + value
	default:
		panic("unknown address reference mode")
	}
}

// Set a value at a address address
func (p *Program) Set(address, value int64) {
	//log.Printf("setting value %v at address %v", value, address)
	if address < 0 {
		log.Println(p)
		panic("Negative Memory Address Access")
	}
	if !p.initialised {
		p.Initialise()
	}
	p.memory[address] = value
}

// New creates a new program
func New(code string) *Program {
	rd := strings.NewReader(strings.TrimSpace(code))
	reg := []int64{}
	var x int64
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
		code:   reg,
		ID:     "intcode",
		Input:  make(chan func() int64),
		Output: make(chan int64),
		Halted: make(chan struct{}),
	}
}

// Copy creates a copy (which we can then mutate)
// it is created with the current memory state, so
// be sure to copy before running.
func (p *Program) Copy() *Program {
	p2 := &Program{
		code:   make([]int64, len(p.code)),
		ID:     p.ID,
		Input:  make(chan func() int64),
		Output: make(chan int64),
		Halted: make(chan struct{}),
	}
	copy(p2.code, p.code)
	return p2
}
