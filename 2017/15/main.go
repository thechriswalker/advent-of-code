package main

import (
	"fmt"
	"math"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2017, 15, solve1, solve2)
}

var (
	PairsToConsiderPart1 = 40_000_000
	PairsToConsiderPart2 = 5_000_000
	FactorA              = uint64(16807)
	FactorB              = uint64(48271)
	MaxInt32             = uint64(math.MaxInt32)
)

// Implement Solution to Problem 1
func solve1(input string) string {
	var genA, genB uint64
	fmt.Sscanf(input, `Generator A starts with %d
		Generator B starts with %d`, &genA, &genB)

	var matches int
	for i := 0; i < PairsToConsiderPart1; i++ {
		// surely there is a better way to do modulu 2^31
		genA = (genA * FactorA) % MaxInt32
		genB = (genB * FactorB) % MaxInt32

		if genA&0xFFFF == genB&0xFFFF {
			matches++
		}
	}

	return fmt.Sprint(matches)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	var genA, genB uint64
	fmt.Sscanf(input, `Generator A starts with %d
		Generator B starts with %d`, &genA, &genB)

	// Channels : 2.5 seconds
	// gA := generatorWithChannel(genA, FactorA, 4)
	// gB := generatorWithChannel(genB, FactorB, 8)

	// var matches int
	// for i := 0; i < PairsToConsiderPart2; i++ {
	// 	if <-gA&0xFFFF == <-gB&0xFFFF {
	// 		matches++
	// 	}
	// }

	// Functions : ~400ms
	gA := generatorWithFunc(genA, FactorA, 4)
	gB := generatorWithFunc(genB, FactorB, 8)
	var matches int
	for i := 0; i < PairsToConsiderPart2; i++ {
		if gA()&0xFFFF == gB()&0xFFFF {
			matches++
		}
	}

	return fmt.Sprint(matches)
}

// channel based generator took 2.5 seconds to run
func GeneratorWithChannel(seed, factor uint64, multiple uint64) chan uint64 {
	gen := seed
	out := make(chan uint64)
	go func() {
		for {
			gen = (gen * factor) % MaxInt32
			if gen%multiple == 0 {
				out <- gen
			}
		}
	}()
	return out
}

// function based generator took ~400ms to run
func generatorWithFunc(seed, factor uint64, multiple uint64) func() uint64 {
	gen := seed
	return func() uint64 {
		for {
			gen = (gen * factor) % MaxInt32
			if gen%multiple == 0 {
				return gen
			}
		}
	}
}
