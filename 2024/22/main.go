package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 22, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {

	sum := 0

	ints := aoc.ToIntSlice(input, '\n')
	for range 2000 {
		for i := range ints {
			ints[i] = prng(ints[i])
		}
	}
	for _, i := range ints {
		sum += i
	}

	return fmt.Sprint(sum)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// we need to find the sequence of 4 changes, that will result in the highest sum of numbers of when those changes occur in the various buyer sequences.
	// so foor each buyer, let's create a map of all the possible "4 change"  sequences, and the find the one amongst all of them that has the highest sum of numbers.

	// this is super naive and works out EVERYTHING, and takes 4 seconds to run.
	// so I changed it to sum the sequences as they are generated, and only keep the highest sum.
	// then we just need to find the max value in the map.
	// that brought it down to <400ms

	sequenceCache := map[[4]int]int{}
	secrets := aoc.ToIntSlice(strings.TrimSpace(input), '\n')
	// cache of current sequence starting mod 4
	sequence := [4]int{}
	for _, secret := range secrets {
		prev := 0
		seen := map[[4]int]struct{}{}
		curr := secret
		var next int
		for j := 0; j < 2000; j++ {
			next = prng(curr)
			price := next % 10
			diff := price - prev
			sequence[j%4] = diff
			if j > 2 {
				seq := [4]int{sequence[(j-3)%4], sequence[(j-2)%4], sequence[(j-1)%4], sequence[j%4]}
				if _, ok := seen[seq]; !ok {
					sequenceCache[seq] = sequenceCache[seq] + price
					seen[seq] = struct{}{}
				}
			}
			curr = next
			prev = price
		}
	}

	// now we have all the sequences, let's find the one with the highest sum
	max := 0
	//bestSeq := [4]int{}
	for _, sum := range sequenceCache {
		if sum > max {
			max = sum
		}
	}
	//fmt.Println("best sequence", bestSeq)
	// 13461553007 too high
	return fmt.Sprint(max)
}

func prng(curr int) (next int) {
	next = mixAndPrune(curr, curr*64)
	next = mixAndPrune(next, next/32)
	next = mixAndPrune(next, next*2048)
	return
}

func mixAndPrune(secret, b int) int {
	return (secret ^ b) % 16777216
}
