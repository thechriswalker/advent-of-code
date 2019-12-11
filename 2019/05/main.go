package main

import (
	"fmt"

	"../../aoc"
	"../intcode"
)

func main() {
	aoc.Run(2019, 5, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	p := intcode.New(input)
	p.EnqueueInput(1)
	done := p.RunAsync()
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
	p.EnqueueInput(5)
	p.RunAsync()
	return fmt.Sprintf("%d", p.GetOutput())
}
