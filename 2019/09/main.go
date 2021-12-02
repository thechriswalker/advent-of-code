package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/2019/intcode"
	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 9, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return runForOneOutput(input, 1)
}

func runForOneOutput(code string, inputs ...int64) string {
	p := intcode.New(code)
	//p.Debug = true
	done := p.RunAsync()
	for _, in := range inputs {
		p.NextInput(in)
	}
	out := []int64{}
	for {
		select {
		case x := <-p.Output:
			out = append(out, x)
			if len(out) > 1 {
				return fmt.Sprintf("Got multiple outputs: %v", out)
			}
		case <-done:
			return fmt.Sprintf("%d", out[len(out)-1])
		}
	}
}

// Implement Solution to Problem 2
func solve2(input string) string {
	p := intcode.New(input)
	done := p.RunAsync()
	p.NextInput(2)
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
