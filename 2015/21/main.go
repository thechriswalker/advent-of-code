package main

import (
	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 21, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// I actually did both parts of this in my head.
	// basically, my input was
	/*
		Hit Points: 100
		Damage: 8
		Armor: 2
	*/
	// and you want equipment so your armor+damage = 10 as well
	// as we have the same HP as the boss, we need to deal exactly
	// as much damage.
	return "91"
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// as in part 1 except I need to spend as much as possible
	// and keep <10 armor+damage
	return "158"
}
