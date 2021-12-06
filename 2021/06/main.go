package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2021, 6, solve1, solve2)
}

type Shoal struct {
	Fish []*Fish
}

type Fish struct {
	Timer int
}

func (f *Fish) Tick() bool {
	if f.Timer == 0 {
		f.Timer = 6
		return true
	}
	f.Timer--
	return false
}

func (s *Shoal) Tick() {
	newFish := []*Fish{}
	for i := range s.Fish {
		if s.Fish[i].Tick() {
			newFish = append(newFish, &Fish{Timer: 8})
		}
	}
	s.Fish = append(s.Fish, newFish...)
}
func (s *Shoal) State() string {
	sb := &strings.Builder{}
	fmt.Fprintf(sb, "%d", s.Fish[0].Timer)
	for i := 1; i < len(s.Fish); i++ {
		fmt.Fprintf(sb, ",%d", s.Fish[i].Timer)
	}
	fmt.Fprintf(sb, " (%d fish)", len(s.Fish))
	return sb.String()
}

// Implement Solution to Problem 1
func solve1(input string) string {
	// not 353099
	return calcNumberFish(input, 80)
}

func calcNumberFish(input string, ticks int) string {
	shoal := &Shoal{Fish: []*Fish{}}
	for _, n := range aoc.ToIntSlice(input, ',') {
		//fmt.Println("adding fish: ", n)
		shoal.Fish = append(shoal.Fish, &Fish{Timer: n})
	}
	//fmt.Printf("Initial state: %s\n", shoal.State())
	for i := 0; i < ticks; i++ {
		shoal.Tick()
		//fmt.Printf("After %2d days: %s\n", i+1, shoal.State())
	}

	return fmt.Sprintf("%d", len(shoal.Fish))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	// ha! that won't work!
	//return calcNumberFish(input, 256)
	shoal := &Shoal2{Fish: map[int]int{}}
	for _, n := range aoc.ToIntSlice(input, ',') {
		//fmt.Println("adding fish: ", n)
		shoal.Fish[n]++
	}
	for i := 0; i < 256; i++ {
		shoal.Tick()
		//fmt.Printf("After %2d days: %s\n", i+1, shoal.State())
	}

	return fmt.Sprintf("%d", shoal.Len())
}

// ok so first solution is no good.
// how about we group fish into "clusters" based on their timers.

type Shoal2 struct {
	Fish map[int]int
}

func (s *Shoal2) Tick() {
	// capture the number that  will spawn.
	// decrement everything else,
	n0 := s.Fish[0]
	n1 := s.Fish[1]
	n2 := s.Fish[2]
	n3 := s.Fish[3]
	n4 := s.Fish[4]
	n5 := s.Fish[5]
	n6 := s.Fish[6]
	n7 := s.Fish[7]
	n8 := s.Fish[8]

	s.Fish[8] = n0
	s.Fish[7] = n8
	s.Fish[6] = n7 + n0 // 0s spawn a 6
	s.Fish[5] = n6
	s.Fish[4] = n5
	s.Fish[3] = n4
	s.Fish[2] = n3
	s.Fish[1] = n2
	s.Fish[0] = n1
}

func (s *Shoal2) Len() int {
	sum := 0
	for i := 0; i < 9; i++ {
		sum += s.Fish[i]
	}
	return sum
}
