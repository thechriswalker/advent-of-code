package main

import (
	"fmt"
	"strconv"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2020, 16, solve1, solve2)
}

type ticket []int

type rule func(x int) bool

type notes struct {
	fields map[string]rule
	mine   ticket
	nearby []ticket
}

func parseNotes(input string) *notes {
	n := &notes{
		fields: map[string]rule{},
		mine:   ticket{},
		nearby: []ticket{},
	}
	lines := strings.Split(input, "\n")
	phase := "fields"
	for _, line := range lines {
		if line == "" {
			continue
		}
		switch phase {
		case "fields":
			if strings.HasPrefix(line, "your ticket:") {
				// transition
				phase = "mine"
				continue
			}
			addFieldRule(n.fields, line)
		case "mine":
			if strings.HasPrefix(line, "nearby tickets:") {
				// transition
				phase = "nearby"
				continue
			}
			n.mine = parseTicket(line)
		case "nearby":
			n.nearby = append(n.nearby, parseTicket(line))
		}
	}
	return n
}

func addFieldRule(fields map[string]rule, line string) {
	s := strings.Split(line, ": ")
	key := s[0]
	v := strings.Split(s[1], " or ")
	ranges := make([][2]int, len(v))
	var a, b int
	for _, r := range v {
		fmt.Sscanf(r, "%d-%d", &a, &b)
		ranges = append(ranges, [2]int{a, b})
	}
	ruleFn := func(x int) bool {
		for _, r := range ranges {
			if x >= r[0] && x <= r[1] {
				return true
			}
		}
		return false
	}
	fields[key] = ruleFn
}

func parseTicket(line string) ticket {
	nums := strings.Split(line, ",")
	stack := make(ticket, 0, len(nums))
	for _, s := range nums {
		n, err := strconv.Atoi(s)
		if err == nil {
			stack = append(stack, n)
		}
	}
	return stack
}

func (t ticket) FindInvalidFieldSum(fields map[string]rule) int {
	sum := 0
	for _, i := range t {
		valid := false
		for _, r := range fields {
			if r(i) {
				// valid for a field
				valid = true
				break
			}
		}
		if !valid {
			sum += i
		}
	}
	return sum
}

// Implement Solution to Problem 1
func solve1(input string) string {
	n := parseNotes(input)
	sum := 0
	for _, t := range n.nearby {
		sum += t.FindInvalidFieldSum(n.fields)
	}
	return fmt.Sprintf("%d", sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	n := parseNotes(input)

	validTickets := []ticket{}
	for _, t := range n.nearby {
		if t.FindInvalidFieldSum(n.fields) == 0 {
			validTickets = append(validTickets, t)
		}
	}
	// now we have the valid tickets
	// we need to find a "column" in the ticket that
	// matches for all tickets, for each field.
	// let's make a map of fields to columns
	fmap := map[string]map[int]bool{}
	columns := len(validTickets[0])
	output := 1 // we multiply
	for i := 0; i < columns; i++ {
		for field, rfn := range n.fields {
			if _, exists := fmap[field]; !exists {
				fmap[field] = map[int]bool{}
			}
			// try to validate every ticket to this field.
			// this could give duplicates.
			// yes it can, so we need to find all possibilities for each column
			// and then reduce it by find those with only one possibility and
			// so on...
			valid := true
			for _, t := range validTickets {
				if !rfn(t[i]) {
					// cannot be this column
					valid = false
					break
				}
			}
			if valid {
				// we found a match for this field
				fmap[field][i] = true
			}
		}
	}

	finalMap := reduceOptions(fmap)

	for field, col := range finalMap {
		if strings.HasPrefix(field, "departure ") {
			output *= n.mine[col]
		}
	}

	return fmt.Sprintf("%d", output)
}

func reduceOptions(m map[string]map[int]bool) map[string]int {
	final := map[string]int{}
	cols := []int{}

	for {
		change := false
		// iterate.
		for field, poss := range m {
			if _, ok := final[field]; ok {
				// got this one
				continue
			}
			// remove all the cols from the poss
			for _, v := range cols {
				delete(poss, v)
			}
			if len(poss) == 1 {
				// only one possibility left!
				for col := range poss {
					final[field] = col
					cols = append(cols, col)
					change = true
					break
				}
			}
		}
		if !change {
			// we must have done as much as we could
			return final
		}
	}
}
