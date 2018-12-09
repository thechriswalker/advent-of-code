package main

import (
	"fmt"

	"../../aoc"
)

func main() {
	aoc.Run(2018, 9, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	var players, finalMarble int
	fmt.Sscanf(input, "%d players; last marble is worth %d points", &players, &finalMarble)
	return fmt.Sprintf("%d", game(players, finalMarble))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	var players, finalMarble int
	fmt.Sscanf(input, "%d players; last marble is worth %d points", &players, &finalMarble)
	return fmt.Sprintf("%d", game(players, finalMarble*100))
}

type Marble struct {
	prev  *Marble
	next  *Marble
	value int
}

func game(players, finalMarble int) int {
	scores := make([]int, players)
	var highScore int
	// make a linked list of marbles.
	current := &Marble{value: 0}
	current.prev = current
	current.next = current
	for i := 1; i <= finalMarble; i++ {
		if i%23 == 0 {
			// score
			scorer := (i - 1) % players

			scores[scorer] += i // they get marble i
			// go back 7 marbles
			current = current.prev
			current = current.prev
			current = current.prev
			current = current.prev
			current = current.prev
			current = current.prev
			current = current.prev
			// they get that value
			//	fmt.Printf("Marble %d + marble %d by elf %d\n", i, current.value, scorer)
			scores[scorer] += current.value // and the one 7 marbles back
			// keep track of the high score
			if highScore < scores[scorer] {
				highScore = scores[scorer]
			}
			// and that marble is removed
			current.prev.next = current.next
			current.next.prev = current.prev
			// and the current is set to the next marble
			current = current.next
		} else {
			// add a marble 1 places clockwise
			// move one forward
			current = current.next
			// add a marble
			next := &Marble{value: i, prev: current, next: current.next}
			current.next.prev = next
			current.next = next
			current = next
		}
	}
	return highScore
}
