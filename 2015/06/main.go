package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2015, 6, solve1, solve2)
}

// stride is 1000
type MillionGrid map[int]int

const Stride = 1000

func (g MillionGrid) On(x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			g[x*1000+y] = 1
		}
	}
}
func (g MillionGrid) Off(x1, y1, x2, y2 int) {
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			delete(g, x*1000+y)
		}
	}
}
func (g MillionGrid) Toggle(x1, y1, x2, y2 int) {
	var z int
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			z = x*1000 + y
			if g[z] == 1 {
				delete(g, z)
			} else {
				g[z] = 1
			}
		}
	}
}

func (g MillionGrid) IncreaseBrightness(x1, y1, x2, y2, inc int) {
	var z int
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			z = x*1000 + y
			g[z] = g[z] + inc
		}
	}
}

func (g MillionGrid) DecreaseBrightness(x1, y1, x2, y2 int) {
	var z, v int
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			z = x*1000 + y
			v = g[z]
			if v < 2 {
				delete(g, z)
			} else {
				g[z] = g[z] - 1
			}
		}
	}
}

// Implement Solution to Problem 1
func solve1(input string) string {
	g := MillionGrid{}
	var x1, x2, y1, y2 int
	var action string
	aoc.MapLines(input, func(line string) error {
		switch {
		case strings.HasPrefix(line, "toggle "):
			action = "toggle"
			line = line[len("toggle "):]
		case strings.HasPrefix(line, "turn on "):
			action = "on"
			line = line[len("turn on "):]
		case strings.HasPrefix(line, "turn off "):
			action = "off"
			line = line[len("turn off "):]
		}
		_, err := fmt.Sscanf(line, "%d,%d through %d,%d", &x1, &y1, &x2, &y2)
		if err == nil {
			switch action {
			case "toggle":
				g.Toggle(x1, y1, x2, y2)
			case "on":
				g.On(x1, y1, x2, y2)
			case "off":
				g.Off(x1, y1, x2, y2)
			}
		} else {
			log.Println("error parsing line:", err)
		}
		return nil
	})
	return fmt.Sprintf("%d", len(g))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	g := MillionGrid{}
	var x1, x2, y1, y2 int
	var action string
	aoc.MapLines(input, func(line string) error {
		switch {
		case strings.HasPrefix(line, "toggle "):
			action = "toggle"
			line = line[len("toggle "):]
		case strings.HasPrefix(line, "turn on "):
			action = "on"
			line = line[len("turn on "):]
		case strings.HasPrefix(line, "turn off "):
			action = "off"
			line = line[len("turn off "):]
		}
		_, err := fmt.Sscanf(line, "%d,%d through %d,%d", &x1, &y1, &x2, &y2)
		if err == nil {
			switch action {
			case "toggle":
				g.IncreaseBrightness(x1, y1, x2, y2, 2)
			case "on":
				g.IncreaseBrightness(x1, y1, x2, y2, 1)
			case "off":
				g.DecreaseBrightness(x1, y1, x2, y2)
			}
		} else {
			log.Println("error parsing line:", err)
		}
		return nil
	})
	sum := 0
	for _, b := range g {
		sum += b
	}
	return fmt.Sprintf("%d", sum)
}
