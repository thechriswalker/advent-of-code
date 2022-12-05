package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 5, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	stacks, instructions := parseStacks(input)

	for _, ins := range instructions {
		//fmt.Println(ins)
		move(stacks, ins.N, ins.Src, ins.Dst)
	}

	// now read off the tops of all the stacks.
	sb := strings.Builder{}
	for _, s := range stacks {
		if len(s) > 0 {
			sb.WriteByte(s[len(s)-1])
		}
	}

	return sb.String()
}

// Implement Solution to Problem 2
func solve2(input string) string {
	stacks, instructions := parseStacks(input)

	for _, ins := range instructions {
		//fmt.Println(ins)
		move2(stacks, ins.N, ins.Src, ins.Dst)
	}

	// now read off the tops of all the stacks.
	sb := strings.Builder{}
	for _, s := range stacks {
		if len(s) > 0 {
			sb.WriteByte(s[len(s)-1])
		}
	}

	return sb.String()
}

type Stack []byte
type Instruction struct {
	N, Src, Dst int
}

func (i Instruction) String() string {
	return fmt.Sprintf("INS: move %d from %d to %d", i.N, i.Src+1, i.Dst+1)
}

func pop(s Stack) (Stack, byte) {
	if len(s) == 0 {
		panic("Pop from empty stack")
	}
	return s[:len(s)-1], s[len(s)-1]
}

func push(s Stack, b byte) Stack {
	return append(s, b)
}

func move(stacks []Stack, n, src, dst int) {
	// pop the top
	var b byte
	for i := 0; i < n; i++ {
		stacks[src], b = pop(stacks[src])
		stacks[dst] = push(stacks[dst], b)
	}
}

func move2(stacks []Stack, n, src, dst int) {
	// we won't use pop/push any more.
	ssrc := stacks[src]
	move := ssrc[len(ssrc)-n:]
	stacks[src] = ssrc[:len(ssrc)-n]
	stacks[dst] = append(stacks[dst], move...)
}

func parseStacks(input string) ([]Stack, []Instruction) {
	split := strings.Index(input, "\n\n")

	// the stack defs are in lines and will be easier to just reverse...
	sdef, idef := input[:split], input[split+2:]

	// the stacks will be more easily done backwards...
	stacklines := strings.Split(sdef, "\n")
	// we work through them backwards. assuming a max of nine stacks.
	stacks := make([]Stack, 9)
	// the first (well, last) line will be just numbers, so ignore.
	for i := len(stacklines) - 2; i >= 0; i-- {
		for j := 0; j < 9; j++ {
			// the crate is at index
			idx := 1 + (j * 4)
			if idx < len(stacklines[i]) {
				b := stacklines[i][idx]
				if b >= 'A' && b <= 'Z' {
					stacks[j] = push(stacks[j], b)
				}
			}
		}
	}
	instructions := []Instruction{}
	// the instructions must be done the right way around.
	var n, src, dst int
	for _, ins := range strings.Split(idef, "\n") {
		//fmt.Println(ins)
		_, err := fmt.Sscanf(ins, "move %d from %d to %d", &n, &src, &dst)
		if err == nil {
			instructions = append(instructions, Instruction{
				N:   n,
				Src: src - 1,
				Dst: dst - 1,
			})
		}
	}
	return stacks, instructions
}
