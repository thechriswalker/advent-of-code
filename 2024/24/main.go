package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 24, solve1, solve2)
}

type Gate struct {
	in1, in2 string
	op       string
	out      string
}

// Implement Solution to Problem 1
func solve1(input string) string {
	wires := map[string]uint8{}

	wireDefs, gateDefs, _ := strings.Cut(input, "\n\n")
	var name string
	var value uint8
	aoc.MapLines(wireDefs, func(line string) error {
		fmt.Sscanf(line, "%3s: %d", &name, &value)
		wires[name] = value
		return nil
	})

	gates := []Gate{}

	aoc.MapLines(gateDefs, func(line string) error {
		gate := Gate{}
		fmt.Sscanf(line, "%s %s %s -> %s", &gate.in1, &gate.op, &gate.in2, &gate.out)
		gates = append(gates, gate)
		return nil
	})

	n := processWires(wires, gates)

	return fmt.Sprint(n)
}

func numberToWires(prefix string, n int64, w map[string]uint8) {
	binString := fmt.Sprintf("%044b", n)
	// go backwards through the string
	for i := 43; i >= 0; i-- {
		w[prefix+strconv.Itoa(i)] = binString[43-i] - '0'
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	wireDefs, gateDefs, _ := strings.Cut(input, "\n\n")
	gates := []Gate{}

	_ = wireDefs

	gatesByOutput := map[string]Gate{}
	gatesByInput := map[string][]Gate{}

	aoc.MapLines(gateDefs, func(line string) error {
		gate := Gate{}
		fmt.Sscanf(line, "%s %s %s -> %s", &gate.in1, &gate.op, &gate.in2, &gate.out)
		if strings.Compare(gate.in1, gate.in2) > 0 {
			gate.in1, gate.in2 = gate.in2, gate.in1
		}
		gates = append(gates, gate)
		gatesByOutput[gate.out] = gate
		gatesByInput[gate.in1] = append(gatesByInput[gate.in1], gate)
		gatesByInput[gate.in2] = append(gatesByInput[gate.in2], gate)
		return nil
	})

	// we need to swap 4 gates to get to the right answer.
	// there must be a trick.
	// otherwise I need to find all combinations of 4 swaps, i.e. all combinations of a subset of 4 elements from the set of combinations of pairs of gates.
	// then we would need to iterate them and discard any for which our process wires function doesn't return the right answer.
	// that's a lot of work.

	// the actual solution to this (which I had to look up, then investigate) is that the circuit is a ripple carry adder.
	// the gates are all XOR, AND, OR.
	// a 44 bit ripple-carry adder.
	// so we have a half adder then a bunch of full adders.
	// we can follow the logic if we index the gates by their inputs and outputs.

	// easier said than done, but we need should have:
	// x00 XOR y00 = z00
	// x00 AND y00 = c00 (we will need to keep track of the aliases for the carry bits)
	// that was a half adder.
	// now the full
	// x01 XOR y01 = tmps01 (again an alias)
	// tmp01 XOR c00 = z01
	// x01 AND y01 = tmpc01 (again an alias)
	// tmps01 AND c00 = tmpx01
	// tmpc01 OR tmpx01 = c01

	// and generally for Nth bit N > 0 we have

	// xNN XOR yNN = tmpaNN
	// tmpaNN XOR c(N-1) = zNN
	// xNN AND yNN = tmpbNN
	// tmpaNN AND c(N-1) = tmpcNN
	// tmpbNN OR tmpcNN = cNN

	// so if we keep track of the aliases for the bits, we should be able to solve where this is wrong.
	// sounds complicated though...

	return "<unsolved>"
}

func processWires(wires map[string]uint8, gates []Gate) int64 {
	deferrals := map[string][]Gate{}
	var processGate func(gate Gate)
	processGate = func(gate Gate) {
		deferralNeeded := false
		a, ok := wires[gate.in1]
		if !ok {
			deferrals[gate.in1] = append(deferrals[gate.in1], gate)
			deferralNeeded = true
		}
		b, ok := wires[gate.in2]
		if !ok {
			deferrals[gate.in2] = append(deferrals[gate.in2], gate)
			deferralNeeded = true
		}
		if deferralNeeded {
			return
		}
		switch gate.op {
		case "AND":
			wires[gate.out] = a & b
		case "OR":
			wires[gate.out] = a | b
		case "XOR":
			wires[gate.out] = a ^ b
		default:
			panic("unknown op")
		}
		// now process any deferrals.

		if defers, ok := deferrals[gate.out]; ok {
			delete(deferrals, gate.out)
			for _, deferral := range defers {
				processGate(deferral)
			}
		}
	}

	for _, gate := range gates {
		processGate(gate)
	}

	if len(deferrals) > 0 {
		// fmt.Println("wires:", wires)
		// fmt.Println("deferrals:", deferrals)
		return -1
	}

	zWires := []string{}
	for k := range wires {
		if strings.HasPrefix(k, "z") {
			zWires = append(zWires, k)
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(zWires)))
	binString := make([]byte, len(zWires))
	for i, z := range zWires {
		if wires[z] == 1 {
			binString[i] = '1'
		} else {
			binString[i] = '0'
		}
	}

	n, _ := strconv.ParseInt(string(binString), 2, 64)
	return n
}
