package main

import (
	"sort"
	"strings"

	"github.com/thechriswalker/advent-of-code/aoc"
)

func main() {
	aoc.Run(2016, 6, solve1, solve2)
}

// Implement Solution to Problem 1
func solve1(input string) string {
	cols := createLists(input)
	s := strings.Builder{}
	for _, c := range cols {
		s.WriteByte(byte(c.MostCommon()))
	}
	return s.String()
}

// Implement Solution to Problem 2
func solve2(input string) string {
	cols := createLists(input)
	s := strings.Builder{}
	for _, c := range cols {
		s.WriteByte(byte(c.LeastCommon()))
	}
	return s.String()
}

func createLists(input string) []*Runes {
	out := []*Runes{}
	firstRow := true
	index := 0
	for _, c := range input {
		switch c {
		case '\n':
			index = 0
			firstRow = false
		default:
			if firstRow {
				out = append(out, NewRunes())
			}
			out[index].Add(c)
			index++
		}
	}
	return out
}

func NewRunes() *Runes {
	return &Runes{
		List: []rune{},
		Map:  map[rune]int{},
	}
}

type Runes struct {
	List   []rune
	Map    map[rune]int
	sorted bool
}

func (r *Runes) Len() int      { return len(r.List) }
func (r *Runes) Swap(i, j int) { r.List[i], r.List[j] = r.List[j], r.List[i] }
func (r *Runes) Less(i, j int) bool {
	ri, rj := r.List[i], r.List[j]
	ci, cj := r.Map[ri], r.Map[rj]

	return ci > cj
}

// A room is real (not a decoy) if the checksum is the five most common letters in the encrypted
// name, in order, with ties broken by alphabetization
func (r *Runes) Add(c rune) {
	n, ok := r.Map[c]
	r.Map[c] = n + 1
	if !ok {
		r.List = append(r.List, c)
	}
	r.sorted = false
}

func (r *Runes) LeastCommon() rune {
	if !r.sorted {
		sort.Sort(r)
		r.sorted = true
	}
	return r.List[len(r.List)-1]
}
func (r *Runes) MostCommon() rune {
	if !r.sorted {
		sort.Sort(r)
		r.sorted = true
	}
	return r.List[0]
}
