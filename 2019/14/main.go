package main

import (
	"bufio"
	"fmt"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2019, 14, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	reactions := parseInput(input)
	// find the reaction for FUEL
	// iterate until only fuel left!
	reqs := Inventory{fuel: 1}
	stock := Inventory{}
	totalOre := 0
	for {
		next, moreOre := nextRequirements(reactions, reqs, stock)
		totalOre += moreOre
		if len(next) == 0 {
			//	fmt.Printf("Finished - ore required: %d (stock left: %v)\n\n", totalOre, stock)
			return fmt.Sprint(totalOre)
		}
		//fmt.Printf("===== Need Another Step\nOre required: %d\nRequirements: %v\nStock: %v\n\n", totalOre, next, stock)
		reqs = next
	}
}

type Reactions map[Chemical]*Reaction

type Inventory map[Chemical]int

func nextRequirements(r Reactions, reqs, stock Inventory) (Inventory, int) {
	//how do we satisfy the current requirements?
	// iterate and find the minimum amounts to create the current reqs.
	// then we keep going until only ORE is required.
	next := Inventory{}
	oreReq := 0
	for chem, n := range reqs {
		// how do we make chem?
		reaction, ok := r[chem]
		if !ok {
			panic(fmt.Sprintf("failed to find reaction for %s", chem))
		}
		// how many reactions do we require?
		times := n / reaction.Output.amount
		if n%reaction.Output.amount != 0 {
			// need 1 more
			times++
			// add the superfluity to sotck
			stock[chem] += (reaction.Output.amount * times) - n
		}
		// so how much of the other elements do we need?
		for _, in := range reaction.Inputs {
			//	fmt.Printf("Producing %d %s requires %d %s (%d x %d) (stock has %d)\n", n, chem, in.amount*times, in.chemical, in.amount, times, stock[in.chemical])
			if in.chemical == ore {
				oreReq += in.amount * times
			} else {
				// how muc to we have?
				inStock := stock[in.chemical]
				needed := in.amount * times
				if inStock >= needed {
					stock[in.chemical] -= needed
					// no extra requirement.
				} else {
					/// we dont have enough in stock, so we need create some more
					delete(stock, in.chemical)
					needed -= inStock
					next[in.chemical] = next[in.chemical] + needed
				}
			}
		}
		// now we go back through the stock to see if we can reduce any requirements.
		// as we may not have processed everything in the optimal order
		for chem, n := range stock {
			if req := next[chem]; req > 0 {
				if req < n {
					// remove req
					stock[chem] -= req
					delete(next, chem)
				} else {
					// reduce req
					delete(stock, chem)
					next[chem] -= n
				}
			}

		}
	}
	return next, oreReq
}

func makeXFuel(reactions Reactions, stock Inventory, amount int) (Inventory, int) {
	reqs := Inventory{fuel: amount}
	// make a copy of stock so we can rollback
	newStock := Inventory{}
	for k, v := range stock {
		newStock[k] = v
	}
	totalOre := 0
	for {
		next, moreOre := nextRequirements(reactions, reqs, newStock)
		totalOre += moreOre
		if len(next) == 0 {
			//fmt.Printf("Finished - ore required: %d (stock left: %v)\n\n", totalOre, stock)
			return newStock, totalOre
		}
		//	fmt.Printf("===== Need Another Step\nOre required: %d\nRequirements: %v\nStock: %v\n\n", totalOre, next, stock)
		reqs = next
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	reactions := parseInput(input)
	// find the reaction for FUEL
	// iterate until only fuel left!
	stock := Inventory{}
	oreLimit := 1000000000000
	totalOre := 0
	fuelProduced := 0
	amount := 10000 // start by trying to produce 1000 at a time.
	for {
		nextStock, moreOre := makeXFuel(reactions, stock, amount)
		if totalOre+moreOre > oreLimit {
			if amount == 1 {
				return fmt.Sprint(fuelProduced)
			}
			amount /= 10
			// try making less
		} else {
			totalOre += moreOre
			stock = nextStock
			fuelProduced += amount
		}
		// loop
	}

	return "<unsolved>"
}

type Chemical string

const (
	ore  = Chemical("ORE")
	fuel = Chemical("FUEL")
)

type Produce struct {
	chemical Chemical
	amount   int
}

type Reaction struct {
	Inputs []Produce
	Output Produce
}

// we need to arrange these reactions in a dependency tree.
// we want to create 1 FUEL. So we must work backwards to see how much
// of each of the other chemicals we need. Then we can traverse the tree
// to make each one.
// each chemical can be produced in only one way, so we can make a map of the reactions
// and continue to build needs until we reach only needing ORE, then we are done.

func parseInput(input string) Reactions {
	scan := bufio.NewScanner(strings.NewReader(input))
	m := Reactions{}

	var name string
	var amount int

	parseProduce := func(in string) Produce {
		fmt.Sscanf(in, "%d %s", &amount, &name)
		return Produce{
			chemical: Chemical(name),
			amount:   amount,
		}
	}

	for scan.Scan() {
		text := strings.Split(scan.Text(), " => ")
		reaction := &Reaction{
			Inputs: []Produce{},
			Output: parseProduce(text[1]),
		}
		for _, s := range strings.Split(text[0], ", ") {
			reaction.Inputs = append(reaction.Inputs, parseProduce(s))
		}
		output := reaction.Output.chemical
		m[output] = reaction
	}
	return m
}
