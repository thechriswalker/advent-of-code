package main

import (
	"fmt"
	"io"
	"strings"

	"../../aoc"
)

func main() {
	aoc.Run(2018, 8, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	r := strings.NewReader(input)
	return fmt.Sprintf("%d", readNodeMetadata(r))
}

// Implement Solution to Problem 2
func solve2(input string) string {
	r := strings.NewReader(input)
	return fmt.Sprintf("%d", getNodeValue(r))
}

func readNodeMetadata(r io.Reader) int {
	var meta, children, sum, v int
	_, err := fmt.Fscanf(r, "%d %d", &children, &meta)
	if err != nil {
		panic(err)
	}
	// do children first
	for i := 0; i < children; i++ {
		sum += readNodeMetadata(r)
	}
	// now add Our meta
	for i := 0; i < meta; i++ {
		_, err := fmt.Fscanf(r, "%d", &v)
		if err != nil {
			panic(err)
		}
		sum += v
	}
	return sum
}

func getNodeValue(r io.Reader) int {
	var meta, children, v int
	_, err := fmt.Fscanf(r, "%d %d", &children, &meta)
	if err != nil {
		panic(err)
	}
	// do children first
	childValues := make([]int, children)
	for i := 0; i < children; i++ {
		childValues[i] = getNodeValue(r)
	}
	// now find our meta
	var metaSum, childSum int
	for i := 0; i < meta; i++ {
		_, err := fmt.Fscanf(r, "%d", &v)
		if err != nil {
			panic(err)
		}
		metaSum += v
		if v > 0 && v <= children {
			childSum += childValues[v-1]
		}
	}
	if children == 0 {
		// no children, return sum of meta
		return metaSum
	}
	// return values of indexed child nodes (1 based)
	return childSum
}
