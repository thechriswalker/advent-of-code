package main

import (
	"fmt"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2023, 13, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	return solveX(input, 0)
}

// Implement Solution to Problem 2
func solve2(input string) string {
	return solveX(input, 1)
}

func solveX(input string, targetSmudges int) string {
	gs := parseGrids(input)

	sum := 0

	for _, g := range gs {

		// find a horizontal reflection line
		x := g.FindSmudgedReflection(targetSmudges)
		if x > 0 {
			// fmt.Printf("Grid has horizontal reflection line at: %d\n", x)
			// fmt.Println("-----------------------------------------")
			// fmt.Println(g)
			// horizontal
			sum += 100 * x
		}
		// find a vertical reflection line
		t := g.Transpose()
		y := t.FindSmudgedReflection(targetSmudges)
		if y > 0 {
			// fmt.Printf("Grid has vertical reflection line at: %d\n", x)
			// fmt.Println("-----------------------------------------")
			// fmt.Println(g)
			// vertical
			sum += y
		}
		if x == -1 && y == -1 {
			// fmt.Println("No Reflection in grid!")
			// fmt.Println("-----------------------------------------")
			// fmt.Println(g)
			panic("No reflection in grid!")
		}
	}
	return fmt.Sprint(sum)
}

type Grid []string

func (g Grid) String() string {
	sb := strings.Builder{}

	sb.WriteString("   ")
	for i := range g[0] {
		fmt.Fprintf(&sb, "%d", (i+1)/10)
	}
	sb.WriteString("\n   ")
	for i := range g[0] {
		fmt.Fprintf(&sb, "%d", (i+1)%10)
	}
	sb.WriteByte('\n')
	for i, l := range g {
		fmt.Fprintf(&sb, "%02d ", i+1)
		sb.WriteString(l)
		sb.WriteByte('\n')
	}
	return sb.String()
}

func (g Grid) Transpose() Grid {
	//fmt.Printf("Transposing Grid:\n%v", g)
	t := make(Grid, len(g[0]))
	s := make([]byte, len(g))
	for i := 0; i < len(t); i++ {
		// first string of t is made up of the first chars of
		// each string in g
		// the i'th string of T is made up of the i'th chars
		// of each string in g
		for j := 0; j < len(g); j++ {
			s[j] = g[j][i]
		}
		t[i] = string(s)
	}
	//fmt.Printf("Transposed:\n%v", t)
	return t
}

// finds a reflection in the grid.
// horizontally only
func (g Grid) FindReflection() int {
	return g.FindSmudgedReflection(0)
}

// similar to before, but instead of
// stopping a found reflection, we will track the
// changes required to make it a good reflection.
// 0 changes will be the original solution
// 1 change will be the new solution
// >1 change is a fail.
func (g Grid) FindSmudgedReflection(targetSmudges int) int {
	for i := 0; i < len(g)-1; i++ {
		if g.countDifferences(i, targetSmudges) {
			return i + 1
		}
	}
	return -1
}

func (g Grid) countDifferences(i, targetSmudges int) bool {
	j, k := i+1, i
	n := 0
	for {
		if j == len(g) || k == -1 {
			return n == targetSmudges
		}
		n += diff(g[j], g[k])
		if n > targetSmudges {
			return false
		}
		j++
		k--
	}
}

func diff(s1, s2 string) int {
	d := 0
	for i := range s1 {
		if s1[i] != s2[i] {
			d++
		}
	}
	return d
}

func parseGrids(input string) []Grid {
	grids := []Grid{}
	curr := Grid{}

	aoc.MapLines(input, func(line string) error {
		if line == "" {
			//	fmt.Printf("parsed grid:\n%v", curr)
			grids = append(grids, curr)
			curr = Grid{}
		} else {
			curr = append(curr, line)
		}
		return nil
	})
	if len(curr) > 0 {
		grids = append(grids, curr)
	}
	return grids
}
