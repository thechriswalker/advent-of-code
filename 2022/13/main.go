package main

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 13, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	pairs := []*ListOrValue{}

	aoc.MapLines(input, func(line string) error {
		if line != "" {
			//		fmt.Println("Parsing: ", line)
			parsed := ParseListOrValue(line)
			if parsed.String() != line {
				fmt.Println("INPUT: ", line)
				fmt.Println("PARSE: ", parsed.String())
				panic("BAD PARSE!")
			}
			pairs = append(pairs, parsed)

		}
		return nil
	})

	sum := 0

	for i := 0; i < len(pairs); i += 2 {
		index := (i / 2) + 1
		a := pairs[i]
		b := pairs[i+1]

		// fmt.Println("Index:", (i/2)+1)
		// fmt.Println("A: ", a)
		// fmt.Println("B: ", b)

		c := a.Compare(b)
		//fmt.Printf("INDEX: %d VALUE: %d\n", index, c)
		if c == -1 {
			sum = sum + index
		}
	}

	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	pairs := []*ListOrValue{}

	aoc.MapLines(input, func(line string) error {
		if line != "" {
			//		fmt.Println("Parsing: ", line)
			parsed := ParseListOrValue(line)
			if parsed.String() != line {
				fmt.Println("INPUT: ", line)
				fmt.Println("PARSE: ", parsed.String())
				panic("BAD PARSE!")
			}
			pairs = append(pairs, parsed)

		}
		return nil
	})
	d1, d2 := ParseListOrValue("[[2]]"), ParseListOrValue("[[6]]")

	pairs = append(pairs, d1, d2)

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Compare(pairs[j]) == -1
	})

	// find the indixes of the packets.
	product := 1
	for i, p := range pairs {
		if p == d1 {
			product *= i + 1
		}
		if p == d2 {
			product *= i + 1
		}
	}

	return fmt.Sprint(product)
}

// go is massively unsuited to this either/or typing...

type ListOrValue struct {
	IsValue bool
	Value   int
	List    []*ListOrValue
}

func (a *ListOrValue) String() string {
	if a.IsValue {
		return strconv.Itoa(a.Value)
	}
	sb := strings.Builder{}
	sb.WriteByte('[')
	for i, l := range a.List {
		if i != 0 {
			sb.WriteByte(',')
		}
		sb.WriteString(l.String())
	}
	sb.WriteByte(']')

	return sb.String()
}

func (a *ListOrValue) PrettyPrint() string {
	var sb strings.Builder
	a.prettyPrint(0, &sb)

	return sb.String()
}

func indent(depth int, w io.Writer) {
	for i := 0; i < depth; i++ {
		w.Write([]byte(". "))
	}
}

func (a *ListOrValue) prettyPrint(depth int, sb *strings.Builder) {
	indent(depth, sb)
	if a.IsValue {
		sb.WriteString(strconv.Itoa(a.Value))
		return
	}
	sb.WriteByte('[')
	for i, l := range a.List {
		if i != 0 {
			sb.WriteByte(',')
		}
		sb.WriteByte('\n')
		l.prettyPrint(depth+1, sb)
	}
	sb.WriteByte('\n')
	indent(depth, sb)
	sb.WriteByte(']')
}

// a -1, 0, 1 value
func (a *ListOrValue) Compare(b *ListOrValue) int {
	return a.compare(b, 0)
}

func (a *ListOrValue) compare(b *ListOrValue, depth int) int {
	if a.IsValue && b.IsValue {
		// indent(depth, os.Stdout)
		// fmt.Println("Compare:", a.Value, "vs", b.Value)
		if a.Value == b.Value {
			//fmt.Println("Equal Order:", a.Value, "=", b.Value)
			return 0
		}
		if a.Value < b.Value {
			// indent(depth, os.Stdout)
			// fmt.Println("Correct Order:", a.Value, "<", b.Value)
			return -1
		}
		// indent(depth, os.Stdout)
		// fmt.Println("Inorrect Order:", a.Value, ">", b.Value)
		return 1
	}
	// at least on is a list.
	var left, right []*ListOrValue
	if a.IsValue {
		left = []*ListOrValue{a}
	} else {
		left = a.List
	}
	if b.IsValue {
		right = []*ListOrValue{b}
	} else {
		right = b.List
	}
	// indent(depth, os.Stdout)
	// fmt.Println("Compare Lists:", left, "vs", right)
	// now compare them element by element.
	for i, l := range left {
		if i >= len(right) {
			// right ran out first.
			//	indent(depth, os.Stdout)
			//	fmt.Println("Inorrect Order: left list longer, len(left):", len(left), "len(right)", len(right))
			return 1
		}
		c := l.compare(right[i], depth+1)
		if c != 0 {
			return c
		}
	}
	// we got to the end of left.
	// if there were right-hand values remaining then we are in order.
	if len(right) > len(left) {
		// indent(depth, os.Stdout)
		// fmt.Println("Correct Order: right list longer, len(left):", len(left), "len(right)", len(right))
		return -1
	}
	//fmt.Println("Equal Order: lists equivalent, len(left):", len(left), "len(right)", len(right))
	return 0
}

func ParseListOrValue(line string) *ListOrValue {
	l, _ := parseListOrValue(line, 0)
	return l
}

func parseListOrValue(line string, offset int) (*ListOrValue, int) {
	//fmt.Println("parseListOrValue:", line[offset:])
	if line[offset] == '[' {
		return parseList(line, offset+1)
	}
	if line[offset] >= '0' && line[offset] <= '9' {
		// just a number
		// but we don't know how many digits.
		return parseValue(line, offset)
	}
	panic("Bad format")
}

func parseValue(line string, offset int) (*ListOrValue, int) {
	//fmt.Println("parseValue:", line[offset:])
	n := offset
	for {
		if line[n] >= '0' && line[n] <= '9' {
			n++
			continue
		}
		break
	}
	v, _ := strconv.Atoi(line[offset:n])
	return &ListOrValue{IsValue: true, Value: v}, n
}

func parseList(line string, offset int) (*ListOrValue, int) {
	//fmt.Println("parseList:", line[offset:])

	// offset should point past the opening to a value. '['
	list := []*ListOrValue{}
	var v *ListOrValue
	for {
		if line[offset] == ']' {
			// we are done.
			return &ListOrValue{
				List: list,
			}, offset + 1
		}
		// otherwise parse another value, from offset
		v, offset = parseListOrValue(line, offset)
		list = append(list, v)
		if line[offset] == ',' {
			// there is another value.
			offset++
		}
	}
}
