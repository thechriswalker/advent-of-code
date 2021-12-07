package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 14, solve1, solve2)
}

type Reindeer struct {
	Name    string
	Speed   int
	Stamina int
	Rest    int
}

func (r *Reindeer) Travel(seconds int) int {
	// reindeers can travel "speed" in every full
	// Stamina+rest.
	step := r.Speed * r.Stamina
	period := r.Stamina + r.Rest
	leftover := seconds % period
	numPeriods := seconds / period
	//fmt.Println(r.Name, "can travel", step, " over ", period, ". We have", numPeriods, "full periods and", leftover, "seconds after")
	dist := step * numPeriods
	if leftover > r.Stamina {
		dist += step
	} else {
		dist += leftover * r.Speed
	}
	return dist
}

const format = `%s can fly %d km/s for %d seconds, but then must rest for %d seconds.`

// Implement Solution to Problem 1
func solve1(input string) string {
	return maxDistance(input, 2503)
}
func maxDistance(input string, seconds int) string {
	reindeer := []*Reindeer{}
	aoc.MapLines(input, func(line string) error {
		r := Reindeer{}
		fmt.Sscanf(line, format, &(r.Name), &(r.Speed), &(r.Stamina), &(r.Rest))
		reindeer = append(reindeer, &r)
		return nil
	})

	max := 0
	for _, r := range reindeer {
		n := r.Travel(seconds)
		//fmt.Println(r.Name, "travelled:", n)
		if n > max {
			max = n
		}
	}

	return fmt.Sprintf("%d", max)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return maxPoints(input, 2503)
}

func maxPoints(input string, seconds int) string {
	reindeer := []*Reindeer{}
	aoc.MapLines(input, func(line string) error {
		r := Reindeer{}
		fmt.Sscanf(line, format, &(r.Name), &(r.Speed), &(r.Stamina), &(r.Rest))
		reindeer = append(reindeer, &r)
		return nil
	})

	points := map[*Reindeer]int{}
	positions := map[*Reindeer]int{}

	for i := 1; i <= seconds; i++ {
		// calculate the travel for each reindeer and
		// with a local "max" and then give all reindeers with that max a
		// point.
		var max = 0
		for _, r := range reindeer {
			n := r.Travel(i)
			//fmt.Println(r.Name, "travelled:", n, "at time:", i)
			if n > max {
				max = n
			}
			positions[r] = n
		}
		for _, r := range reindeer {
			if positions[r] == max {
				points[r]++
			}
		}
	}
	var max = 0
	for _, r := range reindeer {
		if points[r] > max {
			max = points[r]
		}
	}
	return fmt.Sprintf("%d", max)
}
