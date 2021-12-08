package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 15, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return rateCookie(input, 0)
}
func rateCookie(input string, calorieTarget int) string {
	// this "looks" like a gradient ascent type problem
	// i.e. pick a start point and tweak while values go up.
	// but we only have 4 ingredients and 100 slots
	// we might be able to just iterate... we can drop bad cookies
	// on the way (as soon as any property is <= 0)
	ingredients := [][5]int{}
	var x string
	aoc.MapLines(input, func(line string) error {
		i := [5]int{}
		fmt.Sscanf(line, `%s capacity %d, durability %d, flavor %d, texture %d, calories %d`, &x, &(i[0]), &(i[1]), &(i[2]), &(i[3]), &(i[4]))
		ingredients = append(ingredients, i)
		return nil
	})
	fmt.Println("Found", len(ingredients), "ingredients", ingredients)
	max := 0
	scores := [5]int{}
	for d := range distributions() {
		scores[0] = 0
		scores[1] = 0
		scores[2] = 0
		scores[3] = 0
		scores[4] = 0 // calories

		// if d[0] == 44 && d[1] == 56 {
		// 	fmt.Println("MagicDistribution:", d)
		// }
		for i, n := range d {
			scores[0] += n * ingredients[i][0]
			scores[1] += n * ingredients[i][1]
			scores[2] += n * ingredients[i][2]
			scores[3] += n * ingredients[i][3]
			scores[4] += n * ingredients[i][4]
		}
		// if d[0] == 44 && d[1] == 56 {
		// 	fmt.Println(scores)
		// }
		if calorieTarget != 0 && scores[4] != calorieTarget {
			continue
		}
		if scores[0] <= 0 || scores[1] <= 0 || scores[2] <= 0 || scores[3] <= 0 {
			continue
		}
		product := scores[0] * scores[1] * scores[2] * scores[3]

		if product > max {
			max = product
		}
		// if d[0] == 44 && d[1] == 56 {
		// 	fmt.Println("Magic Score:", product)
		// }
	}

	return fmt.Sprintf("%d", max)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return rateCookie(input, 500)
}

// all possible way splits of 100
func distributions() <-chan [4]int {
	ch := make(chan [4]int)
	go func() {
		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				for c := 0; c < 100; c++ {
					for d := 0; d < 100; d++ {
						if a+b+c+d == 100 {
							ch <- [4]int{a, b, c, d}
						}
					}
				}
			}
		}
		close(ch)
	}()
	return ch
}
