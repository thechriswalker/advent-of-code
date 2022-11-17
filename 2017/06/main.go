package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 6, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	i, _ := solveboth(input)
	return fmt.Sprint(i)
}

func solveboth(input string) (iteration, looplength int) {
	memory := aoc.ToUint8Slice(input, '	')
	// let's cache the states the super easy and super inefficient way.
	cache := map[string]int{
		string(memory): 0,
	}
	i := 0
	for {
		//fmt.Println("Iteration:", i, "memory:", memory)
		i++
		memory := redistribute(memory)
		s := string(memory)
		if seen, ok := cache[s]; ok {
			// already exists
			// last seen at iteration "seen"
			looplength = i - seen
			iteration = i
			return
		}
		cache[s] = i
	}
}

func redistribute(memory []uint8) []uint8 {
	highest := uint8(0)
	idx := len(memory)
	for i := len(memory) - 1; i >= 0; i-- {
		if memory[i] >= highest {
			highest = memory[i]
			idx = i
		}
	}
	// clear memory at idx
	memory[idx] = 0
	// now iterate adding memory back.
	for i := 0; i < int(highest); i++ {
		memory[(idx+i+1)%len(memory)]++
	}
	return memory
}

// Implement Solution to Problem 2
func solve2(input string) string {
	_, i := solveboth(input)
	return fmt.Sprint(i)
}
