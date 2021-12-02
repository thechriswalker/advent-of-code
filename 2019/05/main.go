package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/2019/intcode"
	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2019, 5, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	p := intcode.New(input)
	done := p.RunAsync()
	p.Input <- func() int64 { return 1 }
	var lastOutput int64
	cont := true
	for cont {
		select {
		case <-done:
			cont = false
			break
		case o := <-p.Output:
			lastOutput = o
		}
	}
	return fmt.Sprintf("%d", lastOutput)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	p := intcode.New(input)
	p.RunAsync()
	p.Input <- func() int64 { return 5 }
	return fmt.Sprintf("%d", p.GetOutput())
}
