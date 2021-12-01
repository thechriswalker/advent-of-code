package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2020, 15, solve1, solve2)
}

func getNumbers(in string) []int {
	p := strings.Split(in, ",")
	stack := []int{}
	for _, v := range p {
		n, err := strconv.Atoi(strings.TrimSpace(v))
		if err == nil {
			stack = append(stack, n)
		}
	}
	return stack
}

func solveTo(input string, turns int) int {
	// key is the number, value the last turn spoken
	game := map[int]int{}
	setup := getNumbers(input)
	// put all but the last number in. we "commit" the number after the turn
	for i := 0; i < len(setup)-1; i++ {
		game[setup[i]] = i + 1
		//fmt.Printf("TURN:%d, GAME:%v\n", i+1, game)
	}
	last := setup[len(setup)-1]
	next := 0
	turn := len(setup) + 1
	limit := turns
	for turn <= limit {
		when, seen := game[last]
		//fmt.Printf("TURN:%d LAST:%d SEEN:%v GAME:%v\n", turn, last, seen, game)
		if !seen {
			// this number has never come up before.
			next = 0
		} else {
			next = turn - when - 1
		}
		//fmt.Printf("TURN:%d LAST:%d NEXT:%d GAME:%v\n", turn, last, next, game)
		setup = append(setup, next)
		//fmt.Println("GAME:", setup)
		// now commit
		game[last] = turn - 1
		last = next
		turn++
	}
	return last
}

// Implement Solution to Problem 1
func solve1(input string) string {

	return fmt.Sprintf("%d", solveTo(input, 2020))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return fmt.Sprintf("%d", solveTo(input, 30000000))
}
