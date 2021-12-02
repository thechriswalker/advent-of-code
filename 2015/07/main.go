package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 7, solve1, solve2)
}

// we need a tree, which we can iteratively solve, once we have all the values.
// each wire has a "dependency" on the wires that precede it,
// then at the end we can say wire["a"].Signal() and it will backtrack to find
// the Signal.
type Wire struct {
	Name   string
	signal uint16
	Input  Signaller
}

func (w *Wire) Signal() uint16 {
	if w.Input == nil {
		panic("wire " + w.Name + " has no input!")
	}
	if w.signal == 0 {
		w.signal = w.Input.Signal()
	}
	return w.signal
}

type Signaller interface {
	Signal() uint16
}

type fixedSignal uint16

func (f fixedSignal) Signal() uint16 {
	return uint16(f)
}

type binaryGate struct {
	Op   string
	A, B Signaller
}

func (bg *binaryGate) Signal() uint16 {
	a, b := bg.A.Signal(), bg.B.Signal()
	switch bg.Op {
	case "AND":
		return a & b
	case "OR":
		return a | b
	case "LSHIFT":
		return a << b
	case "RSHIFT":
		return a >> b
	}
	panic("unreachable!")
}

type notGate struct {
	A Signaller
}

func (n *notGate) Signal() uint16 {
	return ^n.A.Signal()
}

func mustGetFixedSignal(s string) fixedSignal {
	if n, err := strconv.Atoi(s); err == nil {
		return fixedSignal(n)
	}
	panic("Bad fixed signal: " + s)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	w := getWires(input)
	// hopefully we have enough info to work this out!
	return fmt.Sprintf("%d", w["a"].Signal())
}
func getWires(input string) map[string]*Wire {
	// wires by id and signal
	wires := map[string]*Wire{}
	var getWire = func(id string) *Wire {
		if strings.ContainsAny(id, "0123456789") {
			panic("bad wire name! " + id)
		}
		w, ok := wires[id]
		if !ok {
			// create a wire.
			w = &Wire{Name: id}
			wires[id] = w
		}
		return w
	}
	var getWireOrFixed = func(id string) Signaller {
		if strings.ContainsAny(id, "0123456789") {
			return mustGetFixedSignal(id)
		}
		return getWire(id)
	}
	aoc.MapLines(input, func(line string) error {
		//fmt.Println(line)
		bits := strings.Split(line, " -> ")
		src, dst := bits[0], bits[1]
		// dst is always a wire.
		output := getWire(dst)
		if output.Input != nil {
			panic("already connected something to wire: " + dst)
		}
		// otherwise work out what the input is.
		inputFields := strings.Fields(src)
		switch len(inputFields) {
		case 1:
			// fixed source, but this _could_ be another wire...
			output.Input = getWireOrFixed(inputFields[0])
		case 2:
			// NOT gate
			output.Input = &notGate{A: getWireOrFixed(inputFields[1])}
		case 3:
			// binary gate
			// the L/R SHIFT gates have fixed inputs.
			switch inputFields[1] {
			case "LSHIFT", "RSHIFT":
				output.Input = &binaryGate{
					A:  getWireOrFixed(inputFields[0]),
					Op: inputFields[1],
					B:  getWireOrFixed(inputFields[2]),
				}
			case "AND", "OR":
				// two wires
				output.Input = &binaryGate{
					A:  getWireOrFixed(inputFields[0]),
					Op: inputFields[1],
					B:  getWireOrFixed(inputFields[2]),
				}
			default:
				panic("unknown gate: " + inputFields[1])
			}
		default:
			panic("bad src: " + src)
		}
		return nil
	})
	return wires
}

// Implement Solution to Problem 2
func solve2(input string) string {
	w := getWires(input)
	// override b with the output from part 1
	w["b"].Input = fixedSignal(3176)
	return fmt.Sprintf("%d", w["a"].Signal())
}
