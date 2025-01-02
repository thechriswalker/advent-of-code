package main

import (
	"fmt"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2024, 18, solve1, solve2)
}

var (
	GridSize   = 71
	Part1Bytes = 1024
)

// Implement Solution to Problem 1
func solve1(input string) string {
	g := aoc.NewFixedByteGrid(GridSize, GridSize, '.', nil)
	fallen := 0
	aoc.MapLines(input, func(line string) error {
		if fallen == Part1Bytes {
			return nil
		}
		var p aoc.V2
		fmt.Sscanf(line, "%d,%d", &p.X, &p.Y)
		g.Setv(p, '#')
		fallen++
		return nil
	})

	start := aoc.Vec2(0, 0)
	finish := aoc.Vec2(GridSize-1, GridSize-1)

	shortestPath := aoc.ShortestPathLength(g, start, finish, '#')

	return fmt.Sprint(shortestPath)
}

// Implement Solution to Problem 2
func solve2(input string) string {

	// this method is super slow, as we are recalculating the shortest path for each new potential blocker
	// it was 1.5seconds until I skipped the first "Part1Bytes" checks.
	// I reckon we could speed it up by keeping track of the shortest path and only recalculating the new path if the new blocker is in the way
	// but my helper function for finding the shortest path, doesn't return the path, just the length - much more effort to keep the path as well.
	// perhaps I should make a new helper.

	// so I did that and it is too slow, mostly because there are too many paths when the memory space is empty.
	// this is not a problem when we can uses a fast cache for visited nodes, but we can't do that here.
	// so the initial implementation is the "best"
	return solve2withPathLengthOnly(input)
}

func solve2withPathLengthOnly(input string) string {
	g := aoc.NewFixedByteGrid(GridSize, GridSize, '.', nil)
	var blocker aoc.V2
	processed := 0
	found := false
	start := aoc.Vec2(0, 0)
	finish := aoc.Vec2(GridSize-1, GridSize-1)
	aoc.MapLines(input, func(line string) error {
		if found {
			return nil
		}
		fmt.Sscanf(line, "%d,%d", &blocker.X, &blocker.Y)
		g.Setv(blocker, '#')
		// we know that this many will be fine
		if processed < Part1Bytes {
			processed++
			return nil
		}
		if aoc.ShortestPathLength(g, start, finish, '#') < 0 {
			found = true
		}
		return nil
	})
	return fmt.Sprintf("%d,%d", blocker.X, blocker.Y)
}

// this works for the test case, but not on a 71x71 grid...
func solve2withPathFinding(input string) string {
	g := aoc.NewFixedByteGrid(GridSize, GridSize, '.', nil)
	var blocker aoc.V2
	processed := 0
	found := false
	start := aoc.Vec2(0, 0)
	finish := aoc.Vec2(GridSize-1, GridSize-1)
	var currentPathPoints map[aoc.V2]struct{}

	createPathPoints := func(paths [][]aoc.V2) map[aoc.V2]struct{} {
		newPoints := make(map[aoc.V2]struct{})
		for _, path := range paths {
			for _, point := range path {
				newPoints[point] = struct{}{}
			}
		}
		return newPoints
	}

	aoc.MapLines(input, func(line string) error {
		if found {
			return nil
		}
		fmt.Sscanf(line, "%d,%d", &blocker.X, &blocker.Y)
		g.Setv(blocker, '#')
		// we know that this many will be fine
		if processed < Part1Bytes {
			processed++
			return nil
		}
		processed++
		if currentPathPoints == nil {
			fmt.Println("get shortest paths on byte", processed)
			paths := aoc.GetShortestPaths(g, start, finish, '#')
			fmt.Println("found", len(paths), "paths")
			if len(paths) == 0 {
				found = true
				return nil
			}
			currentPathPoints = createPathPoints(paths)
			return nil
		}
		if _, ok := currentPathPoints[blocker]; ok {
			// recalculate the path
			fmt.Println("get shortest paths on byte", processed)
			paths := aoc.GetShortestPaths(g, start, finish, '#')
			fmt.Println("found", len(paths), "paths")
			if len(paths) == 0 {
				found = true
				return nil
			}
			currentPathPoints = createPathPoints(paths)
		}
		return nil
	})

	return fmt.Sprintf("%d,%d", blocker.X, blocker.Y)
}
