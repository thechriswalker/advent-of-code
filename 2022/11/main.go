package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 11, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return solveN(input, 20, true)
}

func parseUint64(s string) uint64 {
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	return uint64(n)
}

func getMonkeys(input string) []*Monkey {
	monkeys := []*Monkey{}

	// lets split by each monkey...
	chunks := strings.Split(input, "Monkey ")
	for _, c := range chunks {
		if c == "" {
			continue
		}
		lines := strings.Split(c, "\n")
		m := &Monkey{
			Index: len(monkeys),
		}
		// they are in order so this is OK.
		monkeys = append(monkeys, m)

		// line[1] contains the items
		m.Items = aoc.ToUint64Slice(lines[1][18:], ',')
		//fmt.Println(m.Items)
		// line[2] contains the operation.
		m.Operator = parseUint64(lines[2][24:])
		// a side effect of invalid parsing is 0
		m.Operation = lines[2][23]
		// divisor is on line[3]
		//fmt.Printf("lines[3][21:] %q\n", lines[3][19:])
		m.Divisor = parseUint64(lines[3][21:])
		//fmt.Printf("lines[4] %q\n", lines[4])
		m.TrueMonkey, _ = strconv.Atoi(strings.TrimSpace(lines[4][29:]))
		m.FalseMonkey, _ = strconv.Atoi(strings.TrimSpace(lines[5][30:]))

		//fmt.Printf("%#v\n", m)
	}
	return monkeys
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return solveN(input, 10000, false)
}

func solveN(input string, rounds int, withDiv3 bool) string {
	monkeys := getMonkeys(input)

	// given the divisors of all monkeys and multiply them, then
	// we can always drop "%" that.
	mod := uint64(1)
	for _, m := range monkeys {
		mod *= m.Divisor
	}

	for r := 0; r < rounds; r++ {
		// each monkey in turn
		for m := 0; m < len(monkeys); m++ {
			// monkey looks at each item. and passes it on
			mk := monkeys[m]
			// at the end this monkey will have nothing.
			for _, item := range mk.Items {
				switch mk.Operation {
				case '*':
					if mk.Operator == 0 {
						item = item * item
					} else {
						item = item * mk.Operator
					}
				case '+':
					if mk.Operator == 0 {
						item = item + item
					} else {
						item = item + mk.Operator
					}
				default:
					panic("unknown operator! `" + string([]byte{mk.Operation}) + "`")
				}
				mk.Inspections++
				if withDiv3 {
					item = item / 3
				}
				item = item % mod
				// check divisible
				if item%mk.Divisor == 0 {
					monkeys[mk.TrueMonkey].Items = append(monkeys[mk.TrueMonkey].Items, item)
				} else {
					monkeys[mk.FalseMonkey].Items = append(monkeys[mk.FalseMonkey].Items, item)
				}
			}
			// truncate our own items, we have passed them all on.
			mk.Items = mk.Items[0:0]
		}
		// if r == 0 || r == 19 {
		// 	fmt.Printf("Round %02d:\n", r+1)
		// 	for _, m := range monkeys {
		// 		fmt.Printf(" Monkey[%d] op:%c/%d test:%d/%d/%d interactions:%d items: %v\n", m.Index, m.Operation, m.Operator, m.Divisor, m.TrueMonkey, m.FalseMonkey, m.Inspections, m.Items)
		// 	}
		// 	fmt.Println()
		// }
	}

	counts := make([]uint64, len(monkeys))
	for i, m := range monkeys {
		counts[i] = m.Inspections
	}

	//fmt.Println(counts)

	sort.Slice(counts, func(i, j int) bool { return counts[i] < counts[j] })

	// top two!
	product := counts[len(counts)-1] * counts[len(counts)-2]

	return fmt.Sprint(product)
}

type Monkey struct {
	Index                   int
	Items                   []uint64
	Operation               byte   // '*' or '+'
	Operator                uint64 // 0 means self
	Divisor                 uint64
	TrueMonkey, FalseMonkey int
	Inspections             uint64
}
