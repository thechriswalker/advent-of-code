package main

import (
	"fmt"
	"math"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 9, solve1, solve2)
}

type City struct {
	ID        int
	Name      string
	Distances map[int]int
}

type CityList struct {
	list []*City
}

func (cl *CityList) GetNamed(name string) *City {
	for _, c := range cl.list {
		if c.Name == name {
			return c
		}
	}
	// nope Add the city.
	c := &City{
		ID:        len(cl.list),
		Name:      name,
		Distances: map[int]int{},
	}
	cl.list = append(cl.list, c)
	return c
}

type Path struct {
	Distance int
	Path     string
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// we only have 8 locations in our input. so a N! implementation
	// is feasible. but we can discard a lot of routes as soon as they
	// exceed a given value, also using a breadth-first approach allows
	// us to "cache" a bunch of calculations
	cities := &CityList{list: []*City{}}
	var from, to string
	var dist int
	aoc.MapLines(input, func(line string) error {
		fmt.Sscanf(line, "%s to %s = %d", &from, &to, &dist)
		f := cities.GetNamed(from)
		t := cities.GetNamed(to)
		f.Distances[t.ID] = dist
		t.Distances[f.ID] = dist
		return nil
	})
	min := math.MaxInt64
	var recur func(dist Path, rem []*City, prev *City)
	recur = func(dist Path, rem []*City, prev *City) {
		if len(rem) == 0 {
			if min > dist.Distance {
				min = dist.Distance
				//	fmt.Println("NewMin:", dist.Distance, "Path:", dist.Path)
			}
			return
		}
		// other iterate over each
		for i := range rem {
			next := rem[i]
			nextPath := dist
			nextRem := make([]*City, len(rem)-1)
			copy(nextRem, rem[:i])
			copy(nextRem[i:], rem[i+1:])
			if prev != nil {
				nextPath.Distance += prev.Distances[next.ID]
				if nextPath.Distance > min {
					// no point continuing here...
					return
				}
				nextPath.Path += " -> " + next.Name
			} else {
				nextPath.Path += next.Name
			}
			recur(nextPath, nextRem, next)
		}
	}
	recur(Path{}, cities.list, nil)
	// min should have the value!
	return fmt.Sprintf("%d", min)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// same as part one but with a max check intstead of min
	cities := &CityList{list: []*City{}}
	var from, to string
	var dist int
	aoc.MapLines(input, func(line string) error {
		fmt.Sscanf(line, "%s to %s = %d", &from, &to, &dist)
		f := cities.GetNamed(from)
		t := cities.GetNamed(to)
		f.Distances[t.ID] = dist
		t.Distances[f.ID] = dist
		return nil
	})
	max := 0
	var recur func(dist Path, rem []*City, prev *City)
	recur = func(dist Path, rem []*City, prev *City) {
		if len(rem) == 0 {
			if max < dist.Distance {
				max = dist.Distance
				//	fmt.Println("NewMin:", dist.Distance, "Path:", dist.Path)
			}
			return
		}
		// other iterate over each
		for i := range rem {
			next := rem[i]
			nextPath := dist
			nextRem := make([]*City, len(rem)-1)
			copy(nextRem, rem[:i])
			copy(nextRem[i:], rem[i+1:])
			if prev != nil {
				nextPath.Distance += prev.Distances[next.ID]
				// this doesn't work in MAX
				// if nextPath.Distance > min {
				// 	// no point continuing here...
				// 	return
				// }
				nextPath.Path += " -> " + next.Name
			} else {
				nextPath.Path += next.Name
			}
			recur(nextPath, nextRem, next)
		}
	}
	recur(Path{}, cities.list, nil)
	// min should have the value!
	return fmt.Sprintf("%d", max)
}
