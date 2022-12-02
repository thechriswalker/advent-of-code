package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2022, 2, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {

	d := byte('X' - 'A')
	shapeScoreDiff := byte('A' - 1)

	score := 0

	aoc.MapLines(input, func(line string) error {
		a, b := line[0], line[2]-d

		// add your choice to score.
		score += int(b - shapeScoreDiff)

		// did we draw?
		if a == b {
			score += 3
		} else {
			// did we win?
			switch b {
			case 'A': // rock
				// beats scissors
				if a == 'C' {
					score += 6
				}
			case 'B': // paper
				// beats rock
				if a == 'A' {
					score += 6
				}
			case 'C': // scissors
				// beats paper
				if a == 'B' {
					score += 6
				}
			}
		}
		return nil
	})

	return fmt.Sprint(score)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	score := 0

	// maps other players rune to the score for your rune.
	// so to win we want what will beat the key,
	wins := map[byte]int{
		'A': 2,
		'B': 3,
		'C': 1,
	}

	draws := map[byte]int{
		'A': 1,
		'B': 2,
		'C': 3,
	}

	losses := map[byte]int{
		'A': 3,
		'B': 1,
		'C': 2,
	}

	aoc.MapLines(input, func(line string) error {
		a, b := line[0], line[2]

		switch b {
		case 'X':
			// loss,
			score += losses[a]
		case 'Y':
			// draw
			score += 3 + draws[a]
		case 'Z':
			// win
			score += 6 + wins[a]
		}
		return nil
	})

	return fmt.Sprint(score)
}
