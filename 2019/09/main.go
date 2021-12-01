package main

import (
	"fmt"

	"../intcode"
	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 9, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	p := intcode.New(input)
	p.EnqueueInput(1)
	done := p.RunAsync()
	out := []int64{}
	for {
		select {
		case x := <-p.Output:
			out = append(out, x)
		case <-done:
			if len(out) > 1 {
				return fmt.Sprintf("Got multiple outputs: %v", out)
			}
			return fmt.Sprintf("%d", out[len(out)-1])
		}
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	p := intcode.New(input)
	p.EnqueueInput(2)
	done := p.RunAsync()
	out := []int64{}
	for {
		select {
		case x := <-p.Output:
			out = append(out, x)
		case <-done:
			if len(out) > 1 {
				return fmt.Sprintf("Got multiple outputs: %v", out)
			}
			return fmt.Sprintf("%d", out[len(out)-1])
		}
	}
}
