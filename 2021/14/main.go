package main

import (
	"fmt"
	"math"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 14, solve1, solve2)
}

// we want the rules as a map for easy lookup rather than
// iteration
type Rules map[[2]byte]byte

func iterate(tpl []byte, rules Rules) []byte {
	out := make([]byte, 0, len(tpl)*2)
	var match [2]byte

	// as we iterate, we put the "first" char into out
	// then if there is a match, we add the inner.
	// we don't add the final, we do that in the next iteration.
	for i := 0; i < len(tpl)-1; i++ {
		out = append(out, tpl[i])
		match[0], match[1] = tpl[i], tpl[i+1]
		if r, ok := rules[match]; ok {
			out = append(out, r)
		}
	}
	// now add the final element
	out = append(out, tpl[len(tpl)-1])
	return out
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// this was fine for 10 iterations...
	// but it'll be much faster with the "new" method
	//return solveN(input, 10)
	return solvePolymer(input, 10)
}
func solveN(input string, steps int) string {
	var template []byte
	rules := make(Rules)

	aoc.MapLines(input, func(line string) error {
		if template == nil {
			// first line is template.
			template = []byte(line)
		} else if line != "" {
			// a rule
			//  0123456
			// `AB -> C`
			rules[[2]byte{line[0], line[1]}] = line[6]
		}
		return nil
	})

	res := template
	// run iterations
	for i := 0; i < steps; i++ {
		res = iterate(res, rules)
	}

	// now we count the elements finding max and min
	count := map[byte]int{}
	for _, b := range res {
		count[b]++
	}
	max, min := 0, math.MaxInt64
	for _, c := range count {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	return fmt.Sprint(max - min)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// yeah, that's not going to work.
	//return solveN(input, 40)
	return solvePolymer(input, 40)
}

func solvePolymer(input string, steps int) string {
	var p *Polymer
	r := Rules{}
	aoc.MapLines(input, func(line string) error {
		if line == "" {
			// nothing
		} else if p == nil {
			// first line is template.
			p = PolymerFrom([]byte(line))
		} else {
			// a rule
			//  0123456
			// `AB -> C`
			r[[2]byte{line[0], line[1]}] = line[6]
		}
		return nil
	})

	// 40 iterations
	for i := 0; i < steps; i++ {
		p = p.Split(r)
	}

	return fmt.Sprint(p.MinMax())
}

// a better solution would be to create a list of pairs and counts.
// then on each iteration we "split" the pairs creating new pairs
type Polymer struct {
	First, Last byte
	Pairs       map[[2]byte]int
}

func PolymerFrom(b []byte) *Polymer {
	p := &Polymer{
		First: b[0],
		Last:  b[len(b)-1],
		Pairs: make(map[[2]byte]int),
	}
	for i := 0; i < len(b)-1; i++ {
		p.Pairs[[2]byte{b[i], b[i+1]}]++
	}
	return p
}

func (p Polymer) Split(rules Rules) *Polymer {
	next := &Polymer{
		First: p.First,
		Last:  p.Last,
		Pairs: make(map[[2]byte]int, len(p.Pairs)),
	}

	for m, c := range p.Pairs {
		if ins, ok := rules[m]; ok {
			// we need to split these into
			// 2 new ones.
			next.Pairs[[2]byte{m[0], ins}] += c
			next.Pairs[[2]byte{ins, m[1]}] += c
		} else {
			// just keep as is
			next.Pairs[m] += c
		}
	}
	return next
}

// every pair has has overlap with every other pair.
// except the ones at the ends.
// so we count them all and then half the results except
// we add back 1 each for the first and last.
func (p *Polymer) MinMax() int {
	count := map[byte]int{}
	for pair, c := range p.Pairs {
		a, b := pair[0], pair[1]
		count[a] += c
		count[b] += c
	}
	max, min := 0, math.MaxInt64
	for b, c := range count {
		c = c / 2
		if b == p.First {
			c++
		}
		if b == p.Last {
			c++
		}
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}
	return max - min
}
