package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 7, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := parseGraph(input)
	p := getBottom(g)
	return p.Name
}

func getBottom(g map[string]*Program) *Program {
	for _, p := range g {
		if p.HeldBy == nil {
			return p
		}
	}
	panic("unsolved!")
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := parseGraph(input)
	// find the bottom
	p := getBottom(g)
	// now find the weights of holding.
	for {
		//fmt.Println("-->", p.Name)
		next := p.FindUnbalanced()
		if next == nil {
			// the nodes above are balanced.
			// it should way the same as all
			// it's siblings.
			sibling := p.HeldBy.Holding[0]
			if sibling == p {
				sibling = p.HeldBy.Holding[1]
			}
			target := sibling.CombinedWeight()

			// we need to find all our children's weight
			for _, h := range p.Holding {
				target -= h.CombinedWeight()
			}
			return fmt.Sprint(target)
		}
		p = next
	}

}

type Program struct {
	Name    string
	Weight  int
	Holding []*Program
	HeldBy  *Program
}

func (p *Program) CombinedWeight() int {
	w := p.Weight
	for i := range p.Holding {
		w += p.Holding[i].CombinedWeight()
	}
	return w
}

func (p *Program) FindUnbalanced() *Program {
	// it is balanced if the programs is holding are
	// all balanced.
	weightsCount := make(map[int]int, len(p.Holding))
	weightsProg := make(map[int]*Program, len(p.Holding))
	for _, h := range p.Holding {
		w := h.CombinedWeight()
		weightsCount[w]++
		weightsProg[w] = h
	}
	for w, c := range weightsCount {
		if c == 1 {
			return weightsProg[w] // there's only one
		}
	}
	return nil
}

func parseGraph(input string) map[string]*Program {
	programs := map[string]*Program{}

	aoc.MapLines(input, func(line string) error {
		var id string
		var weight int
		_, err := fmt.Sscanf(line, "%s (%d)", &id, &weight)
		if err != nil {
			return err
		}
		p, ok := programs[id]
		if !ok {
			p = &Program{Name: id, Holding: []*Program{}}
			programs[id] = p
		}
		p.Weight = weight
		idx := strings.LastIndex(line, " -> ")
		if idx != -1 {
			held := strings.Split(line[idx+4:], ", ")
			for _, id := range held {
				h, ok := programs[id]
				if !ok {
					h = &Program{Name: id, Holding: []*Program{}}
					programs[id] = h
				}
				h.HeldBy = p
				p.Holding = append(p.Holding, h)
			}
		}
		return nil
	})

	return programs
}
